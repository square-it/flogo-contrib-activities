package command

import (
	"bytes"
	"context"
	"errors"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var log = logger.GetLogger("activity-command")

// CommandActivity is a stub for your Activity implementation
type CommandActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CommandActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *CommandActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *CommandActivity) Eval(activityContext activity.Context) (done bool, err error) {

	// check input command
	command, ok := activityContext.GetInput("command").(string)

	if !ok {
		s := "The input command is not a string."
		log.Error(s)
		return false, errors.New(s)
	}

	if command == "" {
		s := "The input command is required."
		log.Error(s)
		return false, errors.New(s)
	}

	log.Debugf("Input command : %s", command)

	//check input arguments
	arguments, ok, err := a.checkAndGetStringArrays(activityContext, "arguments")

	if !ok {
		return false, err
	}

	log.Debugf("Input arguments : %s", arguments)

	//check input directory
	directory, ok := activityContext.GetInput("directory").(string)

	if !ok {
		s := "The input directory is not a string."
		log.Error(s)
		return false, errors.New(s)
	}

	log.Debugf("Input directory : %s", directory)

	// check input useCurrentEnvironment
	useCurrentEnvironment, ok := activityContext.GetInput("useCurrentEnvironment").(bool)

	if !ok {
		s := "The input useCurrentEnvironment is not a boolean."
		log.Error(s)
		return false, errors.New(s)
	}

	log.Debugf("Input useCurrentEnvironment : %v", useCurrentEnvironment)

	//check input environment
	var environment []string

	if useCurrentEnvironment {
		environment = os.Environ()
	} else {
		environment = make([]string, 0)
	}

	extendsEnvironment, ok, err := a.checkAndGetStringArrays(activityContext, "environment")

	if !ok {
		return false, err
	}

	environment = append(environment, extendsEnvironment...)

	log.Debugf("Input environment : %s", environment)

	// check input timeout
	timeout, ok := activityContext.GetInput("timeout").(int)

	if !ok {
		s := "The input timeout is not a integer."
		log.Error(s)
		return false, errors.New(s)
	}

	log.Debugf("Input timeout : %v", timeout)

	// check input wait
	wait, ok := activityContext.GetInput("wait").(bool)

	if !ok {
		s := "The input wait is not a boolean."
		log.Error(s)
		return false, errors.New(s)
	}

	log.Debugf("Input wait : %v", wait)

	// check command
	path, err := exec.LookPath(command)

	if err != nil {
		log.Errorf("Command %s not found : %v", command, err)
		return false, err
	}

	log.Debugf("Path of the command : %v", path)

	var cmd *exec.Cmd
	var cmdContext context.Context
	var cancel context.CancelFunc

	if timeout != 0 {
		cmdContext, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		cmd = exec.CommandContext(cmdContext, path, arguments...)
	} else {
		cmd = exec.Command(path, arguments...)
	}

	cmd.Dir = directory
	cmd.Env = environment

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if wait {
		log.Debugf("Run the command and waits for it to complete.")
		err = cmd.Run()

		done, err = a.complete(activityContext, buf, err)

		if !done {
			return done, err
		}

	} else {
		log.Debugf("Run the command but does not wait for it to complete.")
		err = cmd.Start()

		if err != nil {
			return false, err
		}

		activityContext.SetOutput("success", true)
		activityContext.SetOutput("exitCode", 0)
		activityContext.SetOutput("output", "")

		if timeout != 0 {
			go func() {
				defer cancel()
				err = cmd.Wait()
				a.complete(activityContext, buf, err)
			}()
		} else {
			go func() {
				err = cmd.Wait()
				a.complete(activityContext, buf, err)
			}()
		}

	}

	return true, nil
}

func (a *CommandActivity) complete(activityContext activity.Context, buf bytes.Buffer, errCmd error) (done bool, err error) {

	activityContext.SetOutput("output", string(buf.Bytes()))

	if errCmd != nil {

		if exitError, ok := errCmd.(*exec.ExitError); ok {
			activityContext.SetOutput("success", false)
			exitCode := -100

			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			}
			activityContext.SetOutput("exitCode", exitCode)
		} else {
			return false, errCmd
		}
	} else {
		activityContext.SetOutput("success", true)
		activityContext.SetOutput("exitCode", 0)
	}

	log.Debugf("Command success : %v", activityContext.GetOutput("success"))
	log.Debugf("Command exitCode : %v", activityContext.GetOutput("exitCode"))
	log.Debugf("Command output : %s", activityContext.GetOutput("output"))

	return true, nil
}

func (a *CommandActivity) checkAndGetStringArrays(activityContext activity.Context, input string) ([]string, bool, error) {
	inputArrays := activityContext.GetInput(input)
	var arrays = make([]string, 0)

	if inputArrays != nil {
		switch arraysValue := inputArrays.(type) {
		case []interface{}:
			for _, item := range arraysValue {
				s, ok := item.(string)
				if ok {
					arrays = append(arrays, s)
				}
			}
		case []string:
			arrays = arraysValue
		default:
			e := "The input " + input + " is not a array."
			log.Error(e)
			return arrays, false, errors.New(e)
		}
	}

	return arrays, true, nil
}
