package makedirectory

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
	"strconv"
)

var log = logger.GetLogger("activity-makedirectory")

// MakeDirectoryActivity is a stub for your Activity implementation
type MakeDirectoryActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MakeDirectoryActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MakeDirectoryActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MakeDirectoryActivity) Eval(context activity.Context) (done bool, err error) {

	path, ok := context.GetInput("path").(string)

	if !ok {
		log.Error("The input path is not a string.")
		return false, nil
	}

	if (path == "") {
		log.Error("The input path is required.")
		return false, nil
	}

	log.Debugf("Input path : %s", path)

	all, ok := context.GetInput("all").(bool)

	if !ok {
		log.Error("The input all is not a boolean.")
		return false, nil
	}

	log.Debugf("Input all : %v", all)

	permissions, ok := context.GetInput("permissions").(string)

	if !ok {
		log.Error("The input permissions is not a string.")
		return false, nil
	}

	if (permissions == "") {
		permissions = a.metadata.Input["permissions"].Value().(string)
	}

	log.Debugf("Input permissions : %s", path)

	fmInt, err := strconv.ParseUint(permissions, 8, 32)

	if err != nil {
		return true, err
	}

	fm := os.FileMode(fmInt)

	if !all {
		err = os.Mkdir(path, fm)
	} else {
		err = os.MkdirAll(path, fm)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
