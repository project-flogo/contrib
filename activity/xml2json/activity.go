package xml2json

import (
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
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
	output := &Output{}

	xmlData, err := coerce.ToBytes(input.XmlData)
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Failed to convert input data to bytes: %s", err.Error()), "", nil)
	}
	s, err := mxj.NewMapXml(xmlData)
	if err != nil {
		context.Logger().Error(err)
		return false, activity.NewError(fmt.Sprintf("Failed to convert XML data: %s", err.Error()), "", nil)
	}
	output.JsonObject = s.Old()
	err = context.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	context.Logger().Debug("XML2JSON activity completed")
	return true, nil
}
