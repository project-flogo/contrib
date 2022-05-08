// Package json2xml activity for converting JSON object into XML
package json2xml

import (
	"encoding/xml"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"strings"
)

type Activity struct{}

// init Flogo activity
func init() {
	_ = activity.Register(&Activity{})
}

// Default xmlHeader
var xmlHeader = []byte(xml.Header)

// Precalculated size of default root tags <doc></doc> from MXJ package
var startDefaultRootTagSize = len(mxj.DefaultRootTag) + 2
var endDefaultRootTagSize = len(mxj.DefaultRootTag) + 3

var metadata = activity.ToMetadata(&Input{}, &Output{})

// Metadata for the Activity
func (ac *Activity) Metadata() *activity.Metadata {
	return metadata
}

func (ac *Activity) Eval(ctx activity.Context) (bool, error) {
	ctx.Logger().Debugf("Activity [%s] JSON2XML", ctx.Name())

	input := &Input{}
	var err = ctx.GetInputObject(input)
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Activity [%s] Can't get input JSON object - %s", ctx.Name(), err.Error()), "JSON2XML-01", nil)
	}

	var out []byte
	out, err = convert(input.Json, input.XmlRootTag, ctx.Logger())
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Activity [%s] Error converting JSON object [%T] to XML - %s", ctx.Name(), input.Json, err.Error()), "JSON2XML-02", nil)
	}

	err = ctx.SetOutputObject(&Output{Xml: out})
	if err != nil {
		return false, activity.NewError(fmt.Sprintf("Activity [%s] Can't set output XML object - %s", ctx.Name(), err.Error()), "JSON2XML-03", nil)
	}

	ctx.Logger().Debugf("Activity [%s] JSON2XML completed", ctx.Name())
	return true, nil
}

func convert(json map[string]interface{}, xmlRootTag string, log log.Logger) ([]byte, error) {
	mxj.XMLEscapeChars(true)
	var data []byte
	var err error
	if len(strings.TrimSpace(xmlRootTag)) == 0 {
		data, err = mxj.AnyXml(json)
		if json != nil && err == nil {
			// remove default root tags added by mxj package
			data = data[startDefaultRootTagSize : len(data)-endDefaultRootTagSize]
		}
	} else {
		data, err = mxj.AnyXml(json, xmlRootTag)
	}

	var buf []byte
	buf = append(buf, xmlHeader...)
	buf = append(buf, data...)

	log.Debugf("Converted JSON object [%T] to XML object [%T] with size of %d", json, buf, len(buf))
	return buf, err
}
