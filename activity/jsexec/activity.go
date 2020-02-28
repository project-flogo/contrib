package jsexec

import (
	"encoding/json"
	"errors"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/robertkrimen/otto"
)

var activityMetadata = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func init() {
	_ = activity.Register(&Activity{}, New)
}

// Activity is a javascript activity
type Activity struct {
	script string
}

// New creates a new javascript activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := Settings{}
	err := metadata.MapToStruct(ctx.Settings(), &settings, true)
	if err != nil {
		return nil, err
	}

	logger := ctx.Logger()
	logger.Debugf("Setting: %b", settings)

	act := Activity{
		script: settings.Script,
	}

	return &act, nil
}

// Metadata return the metadata for the activity
func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

// Eval executes the activity
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := Input{}
	err = ctx.GetInputObject(&input)
	if err != nil {
		return false, err
	}

	output := Output{}
	vm, err := NewVM(nil)
	if err != nil {
		output.Error = true
		output.ErrorMessage = err.Error()
		return false, err
	}
	//todo is ok to ignore the errors for the setInVM calls?
	_ = vm.setInVM("parameters", input.Parameters)
	_ = vm.setInVM("result", map[string]interface{}{})

	_, err = vm.vm.Run(a.script)
	if err != nil {
		output.Error = true
		output.ErrorMessage = err.Error()
		return false, err
	}
	result, err := vm.getFromVM("result")
	if err != nil {
		output.Error = true
		output.ErrorMessage = err.Error()
		return false, err
	}
	output.Result, err = coerce.ToObject(result)
	if err != nil {
		return false, err
	}

	err = ctx.SetOutputObject(&output)
	if err != nil {
		return false, err
	}

	return true, nil
}

// VM represents a VM object.
type VM struct {
	vm *otto.Otto
}

// NewVM initializes a new VM with defaults.
func NewVM(defaults map[string]interface{}) (vm *VM, err error) {
	vm = &VM{}
	vm.vm = otto.New()
	for k, v := range defaults {
		if v != nil {
			vm.vm.Set(k, v)
		}
	}
	return vm, err
}

// EvaluateToBool evaluates a string condition within the context of the VM.
func (vm *VM) EvaluateToBool(condition string) (truthy bool, err error) {
	if condition == "" {
		return true, nil
	}
	res, err := vm.vm.Object(condition)
	if err != nil {
		return false, err
	}
	truthy, err = res.Value().ToBoolean()
	if err != nil {
		err = errors.New("condition does not evaluate to bool")
		return false, err
	}
	return truthy, err
}

// setInVM sets the object name and value in the VM.
func (vm *VM) setInVM(name string, object interface{}) (err error) {
	var valueJSON json.RawMessage
	var vmObject map[string]interface{}
	valueJSON, err = json.Marshal(object)
	if err != nil {
		return err
	}
	err = json.Unmarshal(valueJSON, &vmObject)
	if err != nil {
		return err
	}
	return vm.vm.Set(name, vmObject)
}

// getFromVM extracts the current object value from the VM.
func (vm *VM) getFromVM(name string) (object interface{}, err error) {
	obj, err := vm.vm.Object(name) //todo is ok to ignore the error?
	if err != nil {
		return object, err
	}

	object, err = obj.Value().Export()
	return object, err
}

// setPrimitiveInVM sets primitive value in VM.
func (vm *VM) setPrimitiveInVM(name string, primitive interface{}) error {
	return vm.vm.Set(name, primitive)
}
