package gpiopwm

import (
	"errors"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/stianeikeland/go-rpio"
	"os"
)

const (
	cycleLengthInput = "cycleLength"
	dutyLengthInput  = "dutyLength"
	frequencyInput   = "pwmFrequency"
	pinNumberInput   = "pinNumber"
)

type Activity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Activity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return a.metadata
}

// Enable PWM, then set Frequency and DutyCycle on the given GPIO pin
func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	pinNumber, ok := context.GetInput(pinNumberInput).(int)
	if !ok {
		logger.Errorf("Input value for %s is not a valid int", pinNumberInput)
		return false, errors.New(fmt.Sprintf("Input value for %s is not a valid int", pinNumberInput))
	}

	frequency, ok := context.GetInput(frequencyInput).(int)
	if !ok {
		logger.Errorf("Input value for %s is not a valid int", frequencyInput)
		return false, errors.New(fmt.Sprintf("Input value for %s is not a valid int", frequencyInput))
	}

	dutyLength, ok := context.GetInput(dutyLengthInput).(uint32)
	if !ok {
		logger.Errorf("Input value for %s is not a valid uint32", dutyLengthInput)
		return false, errors.New(fmt.Sprintf("Input value for %s is not a valid uint32", dutyLengthInput))
	}

	cycleLength, ok := context.GetInput(cycleLengthInput).(uint32)
	if !ok {
		logger.Errorf("Input value for %s is not a valid uint32", cycleLengthInput)
		return false, errors.New(fmt.Sprintf("Input value for %s is not a valid uint32", cycleLengthInput))
	}

	if os.Getegid() != 0 {
		logger.Error("PWM control requires to run as root")
		return false, errors.New("PWM control requires to run as root")
	}

	rpioErr := rpio.Open()
	defer rpio.Close()

	if rpioErr != nil {
		logger.Error("Failed to open RPIO", rpioErr.Error())
		return false, errors.New(fmt.Sprintf("Failed to open RPIO", rpioErr.Error()))
	}

	pin := rpio.Pin(pinNumber)
	pin.Pwm()
	pin.Freq(frequency)
	pin.DutyCycle(dutyLength, cycleLength)

	return true, nil
}
