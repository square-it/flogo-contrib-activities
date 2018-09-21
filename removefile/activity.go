package removefile

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
)

var log = logger.GetLogger("activity-removefile")

// RemoveFileActivity is a stub for your Activity implementation
type RemoveFileActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &RemoveFileActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *RemoveFileActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *RemoveFileActivity) Eval(context activity.Context) (done bool, err error) {

	path := context.GetInput("path").(string)
	recursive := context.GetInput("recursive").(bool)
	done = true
	err = nil

	log.Debugf("Remove file %s - recursively : %t", path, recursive)

	if !recursive {
		err = os.Remove(path)
	} else {
		err = os.RemoveAll(path)
	}

	if err != nil {
		done = false
	}

	return done, err
}
