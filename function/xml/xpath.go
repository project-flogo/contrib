package xml

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type fnXPATH struct {
}

func init() {
	function.Register(&fnXPATH{})
}

func (s *fnXPATH) Name() string {
	return "xpath"
}
func (s *fnXPATH) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeBool}, false
}

func (s *fnXPATH) Eval(in ...interface{}) (interface{}, error) {

	//validate the passed first argument (should be an XPATH) is at least a string
	xpath, err := coerce.ToString(in[0])
	if err != nil {
		return nil, fmt.Errorf("xpath function first parameter [%+v] must be string", in[0])
	}
	//validate the passed second argument (should be an XML) is at least a string
	xml, err := coerce.ToString(in[1])
	if err != nil {
		return nil, fmt.Errorf("xpath function second parameter [%+v] must be string", in[1])
	}
	//validate the passed third argument (should be a boolean) is at least true or false
	returnAsXML, err := coerce.ToBool(in[2])
	if err != nil {
		return nil, fmt.Errorf("xpath function thrid parameter [%+v] must be boolean", in[2])
	}
	//ok now try to check the passed XML in the first param is a valid XML structure
	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		//fmt.Println("The passed XML value does not Parse / is not valid")
		panic(err)
	}
	//next lets specify a nodeArray which will store the  result from the passed XPATH against the xml
	var nodeList []*xmlquery.Node
	//now query the passed xpath against the doc to generate a list of Nodes (hopefully more than 0)
	nodeList, err = xmlquery.QueryAll(doc, xpath)
	//check there is no error, panic otherwise
	if err != nil {
		//fmt.Println("The passed XML value does not Parse / is not valid")
		panic(err)
	}
	var parsedStrings []string
		if len(nodeList) == 0 {
			// No matching nodes - add a blank string or handle as needed
			parsedStrings = append(parsedStrings, "") 
		} else {
			// Iterate and process found nodes
			for index := range nodeList {
				parsedStrings = append(parsedStrings, nodeList[index].OutputXML(returnAsXML))
			}
	}

	//now we want to return all the Strings as a single return String
	returnString := strings.Join(parsedStrings, "\n" )

	return returnString, nil
	//return result.InnerText(), nil
}