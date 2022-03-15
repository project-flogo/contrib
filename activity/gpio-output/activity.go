package gpio_output

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/log"
	"github.com/stianeikeland/go-rpio/v4"
)

const (
	actionTurnOn  = "TurnOn"
	actionTurnOff = "TurnOff"
	actionToggle  = "Toggle"
)

const (
	jsonGpioPin = "gpioPin"
	jsonAction  = "action"
)

type Input struct {
	GpioPin int    `md:"GPIOPin, required"`
	Action  string `md:"Action, required, allowed(TurnOn, TurnOff, Toggle)"`
}

type Activity struct {
}

func init() {
	_ = activity.Register(&Activity{})
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		jsonGpioPin: i.GpioPin,
		jsonAction:  i.Action,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {
	var err error
	i.GpioPin, err = coerce.ToInt(values[jsonGpioPin])
	if err != nil {
		return err
	}

	i.Action, err = coerce.ToString(values[jsonAction])

	return err
}

func (a *Activity) Metadata() *activity.Metadata {
	return activity.ToMetadata(&Input{})
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	settings := &Input{}
	ctx.GetInputObject(settings)

	pinNumber := settings.GpioPin
	pin := rpio.Pin(pinNumber)
	action := settings.Action

	log.RootLogger().Debug("Action %s, pin %d", action, pinNumber)

	openError := rpio.Open()
	if openError != nil {
		log.RootLogger().Error("Failed to open pin")
		return false, openError
	}
	pin.Output()

	if action == actionTurnOn {
		pin.Write(rpio.High)
	} else if action == actionTurnOff {
		pin.Write(rpio.Low)
	} else if action == actionToggle {
		pin.Toggle()
	} else {
		log.RootLogger().Error("Unknown action %s", action)
		return false, nil
	}

	return true, nil
}
