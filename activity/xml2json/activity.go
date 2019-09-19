package xml2json

import (
	"encoding/json"
	"strings"

	xj "github.com/basgys/goxml2json"
	"github.com/project-flogo/core/activity"
)

// Activity is an activity that converts XML data into JSON object.
// inputs: XML data
// outputs: JSON object
type Activity struct {
}

func init() {
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	context.Logger().Debug("Executing XML2JSON activity")

	input := &Input{}
	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}
	xmlData := input.XmlData

	output := &Output{}

	xml := strings.NewReader(xmlData)

	jsonData, err := xj.Convert(xml, xj.WithTypeConverter(xj.Float, xj.Bool, xj.Int, xj.String, xj.Null))
	if err != nil {
		context.Logger().Error(err)
		return false, activity.NewError("Failed to convert XML data", "", nil)
	}

	err = json.Unmarshal(jsonData.Bytes(), &output.JsonObject)
	if err != nil {
		context.Logger().Error(err)
		return false, activity.NewError("Failed to parse JSON data", "", nil)
	}

	err = context.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	context.Logger().Debug("XML2JSON activity completed")
	return true, nil
}
