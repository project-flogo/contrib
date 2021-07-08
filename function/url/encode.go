package url

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnEncode{})
}

type fnEncode struct {
}

// Name returns the name of the function
func (fnEncode) Name() string {
	return "encode"
}

// Sig returns the function signature
func (fnEncode) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

// Eval executes the function
func (fnEncode) Eval(params ...interface{}) (interface{}, error) {
	rawString, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce [%+v] to string: %s", params[0], err.Error())
	}
	u, err := url.Parse(rawString)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse url: %s", err.Error())
	}
	var encoded []string
	if len(u.Scheme) > 0 {
		encoded = append(encoded, u.Scheme+"://")
	}
	if u.User != nil {
		encoded = append(encoded, u.User.String())
	}
	if len(u.Host) > 0 {
		encoded = append(encoded, u.Host)
	}
	baseURL, err := url.Parse(strings.Join(encoded, ""))
	if err != nil {
		return nil, fmt.Errorf("Unable to parse base url [%s]: %s", baseURL, err.Error())
	}
	baseURL.Path += u.Path
	if len(u.Fragment) > 0 {
		baseURL.Fragment += u.Fragment
	}
	if baseURL.Query() != nil {
		baseURL.RawQuery = u.Query().Encode()
	}
	return baseURL.String(), nil
}
