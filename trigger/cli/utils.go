package cli

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func getFlagsAndArgs(hCmd *handlerCmd, cliName string, subCmd bool) (map[string]interface{}, []string) {
	hCmd.flagSet.SetOutput(ioutil.Discard)

	idx := 1
	if subCmd {
		idx = 2
	}

	err := hCmd.flagSet.Parse(os.Args[idx:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s.\n", err.Error())
		printCmdUsage(cliName, hCmd, true)
		os.Exit(1)
	}

	flags := make(map[string]interface{})

	for key, value := range hCmd.cmdFlags {
		flags[key] = value
	}
	args := hCmd.flagSet.Args()

	return flags, args
}

func printMainUsage(cliName string, trg *Trigger, isErr bool) {

	w := os.Stderr

	if !isErr {
		w = os.Stdout
		fmt.Fprintf(w, "%s\n", trg.settings.Long)
	}

	bw := bufio.NewWriter(w)

	data := &mainUsageData{Name: cliName, Use: trg.settings.Usage, Long: trg.settings.Long}

	for name, cmd := range trg.commands {
		data.Cmds = append(data.Cmds, &mainCmdUsageData{Name: name, Short: cmd.settings.Short})
	}
	data.Cmds = append(data.Cmds, &mainCmdUsageData{Name: "help", Short: "help on command"})
	data.Cmds = append(data.Cmds, &mainCmdUsageData{Name: "version", Short: "prints cli version"})

	RenderTemplate(bw, mainUsageTpl, data)
	bw.Flush()
}

func printCmdUsage(cliName string, cmd *handlerCmd, isErr bool) {

	w := os.Stderr

	if !isErr {
		w = os.Stdout
		fmt.Fprintf(w, "%s\n", cmd.settings.Long)
	}

	bw := bufio.NewWriter(w)

	data := &cmdUsageData{CliName: cliName, Name: cmd.handler.Name(), Use: cmd.settings.Usage, Long: cmd.settings.Long}

	flags := GetFlags(cmd.flagSet)

	for _, flg := range flags {
		n, _ := flag.UnquoteUsage(flg)

		usage := "-" + flg.Name + " " + n
		data.Flags = append(data.Flags, &cmdFlagUsageData{Usage: usage, Short: flg.Usage})
	}

	RenderTemplate(bw, cmdUsageTpl, data)
	bw.Flush()
}

type mainUsageData struct {
	Name string
	Use  string
	Long string
	Cmds []*mainCmdUsageData
}
type mainCmdUsageData struct {
	Name  string
	Short string
}

var mainUsageTpl = `Usage:
    {{.Name}} {{.Use}}

Commands:{{range .Cmds}}
    {{.Name | printf "%-12s"}} {{.Short}}{{end}}
`

type cmdUsageData struct {
	CliName string
	Name    string
	Use     string
	Long    string
	Flags   []*cmdFlagUsageData
}
type cmdFlagUsageData struct {
	Usage string
	Short string
}

var cmdUsageTpl = `Usage:
    {{.CliName}} {{.Name}} {{.Use}}

Flags: {{range .Flags}}
    {{.Usage | printf "%-20s"}} {{.Short}}{{end}}

`

func RenderTemplate(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func GetFlags(fs *flag.FlagSet) []*flag.Flag {

	var flags []*flag.Flag

	fs.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})

	return flags
}
