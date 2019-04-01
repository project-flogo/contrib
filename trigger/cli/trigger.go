package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/trigger"
)

var triggerMd = trigger.NewMetadata(&HandlerSettings{}, &Output{}, &Reply{})

func init() {
	trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	singleton = &Trigger{config: config, commands: make(map[string]*handlerCmd)}
	return singleton, nil
}

var singleton *Trigger

// Trigger CLI trigger struct
type Trigger struct {
	config   *trigger.Config
	logger   log.Logger
	settings *Settings
	commands map[string]*handlerCmd
}

type handlerCmd struct {
	handler  trigger.Handler
	settings *HandlerSettings
	flagSet  *flag.FlagSet
	cmdFlags map[string]interface{}
}

// Metadata implements trigger.Trigger.Metadata
func (t *Trigger) Metadata() *trigger.Metadata {
	return triggerMd
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = ctx.Logger()

	if len(ctx.GetHandlers()) == 0 {
		return fmt.Errorf("no commands found for cli trigger '%s'", t.config.Id)
	}

	s := &Settings{}
	err := metadata.MapToStruct(t.config.Settings, s, true)
	if err != nil {
		return err
	}

	t.settings = s

	unamedHandler := false

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		handlerName := handler.Name()
		if handlerName == "" {
			if unamedHandler {
				return fmt.Errorf("at most one handler can be unamed in the cli trigger")
			} else {
				unamedHandler = true
				handlerName = "default"
			}
		}

		if _, exists := t.commands[handlerName]; exists {
			return fmt.Errorf("cannot have duplicate handler names in the cli trigger")
		}
		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err

		}

		hCmd := &handlerCmd{handler: handler, settings: s, cmdFlags: make(map[string]interface{})}

		// Subcommands
		cmd := flag.NewFlagSet(handlerName, flag.ContinueOnError)
		hCmd.flagSet = cmd

		for _, desc := range s.FlagDesc {
			descParts := strings.Split(desc.(string), "||")

			name := strings.TrimSpace(descParts[0])
			value := strings.TrimSpace(descParts[1])
			usage := strings.TrimSpace(descParts[2])

			if strings.EqualFold(value, "true") || strings.EqualFold(value, "false") {
				tmpVal := strings.ToLower(value)
				boolVal, _ := strconv.ParseBool(tmpVal)
				boolPtr := cmd.Bool(name, boolVal, usage)
				hCmd.cmdFlags[name] = boolPtr
			} else {
				strPtr := cmd.String(name, value, usage)
				hCmd.cmdFlags[name] = strPtr
			}
		}

		log.RootLogger().Tracef("Adding command %s", handlerName)
		t.commands[handlerName] = hCmd

	}

	return nil
}

func (t *Trigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	return nil
}

func Invoke() (string, error) {

	logger := trigger.GetLogger(support.GetRef(singleton))

	lvl := os.Getenv("FLOGO_LOG_LEVEL")
	if lvl == "" {
		log.SetLogLevel(log.RootLogger(), log.ErrorLevel)
		log.SetLogLevel(logger, log.ErrorLevel)
	} else {
		log.SetLogLevel(logger, log.ToLogLevel(lvl))
	}

	cliPath, _ := os.Executable()
	cliName := filepath.Base(cliPath)
	
	//Todo?
	if singleton.settings.SingleCmd {

	}

	var cmdName string

	if len(os.Args) == 1 {

		if singleton.settings.SingleCmd {
			//cli is a single command, assumes one handler
			var hCmd *handlerCmd
			for _, cmd := range singleton.commands {
				hCmd = cmd
				break
			}

			if hCmd == nil {
				fmt.Fprintf(os.Stderr, "Error: cli improperly configured, needs at least one handler\n")
				os.Exit(1)
			}

			flags, args := getFlagsAndArgs(hCmd, cliName, false)
			return singleton.Invoke(hCmd.handler, flags, args)
		}

		help(cliName, singleton, false)
		os.Exit(0)
	} else {
		cmdName = os.Args[1]
	}

	if strings.EqualFold(cmdName, "help") {
		if len(os.Args) == 2 {
			help(cliName, singleton, false)
			os.Exit(0)
		}

		subCmd := os.Args[2]

		handlerCmd, exists := singleton.commands[subCmd]
		if !exists {
			fmt.Fprintf(os.Stderr, "Error: unknown command %#q\n", cmdName)
			help(cliName, singleton, true)
			os.Exit(1)
		}

		helpCmd(cliName, handlerCmd, false)
		os.Exit(0)
	}

	if strings.EqualFold(cmdName, "version") {

		fmt.Fprintf(os.Stdout, "%s version %s\n", cliName, engine.GetAppVersion())
		os.Exit(0)
	}

	hCmd, exists := singleton.commands[cmdName]
	if !exists {
		fmt.Fprintf(os.Stderr, "Error: unknown command %#q\n", cmdName)
		help(cliName, singleton, true)
		os.Exit(1)
	}

	flags, args := getFlagsAndArgs(hCmd, cliName, true)

	return singleton.Invoke(hCmd.handler, flags, args)
}

func (t *Trigger) Invoke(handler trigger.Handler, flags map[string]interface{}, args []string) (string, error) {

	t.logger.Debugf("invoking handler '%s'", handler)

	data := map[string]interface{}{
		"args":  args,
		"flags": flags,
	}

	results, err := handler.Handle(context.Background(), data)

	if err != nil {
		t.logger.Debugf("error: %s", err.Error())
		return "", err
	}

	replyData := results["data"]
	stringData, _ := coerce.ToString(replyData)

	return stringData, nil

	//if replyData != nil {
	//	data, err := json.Marshal(replyData)
	//	if err != nil {
	//		return "", err
	//	}
	//	return string(data), nil
	//}
}

func help(cliName string, t *Trigger, isErr bool) {
	printMainUsage(cliName, t, isErr)
}

func helpCmd(cliName string, hc *handlerCmd, isErr bool) {
	printCmdUsage(cliName, hc, isErr)
}
