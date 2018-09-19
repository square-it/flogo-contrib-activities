package listfiles

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os"
	"path/filepath"
)

var log = logger.GetLogger("activity-listfiles")

// ListFilesActivity is a stub for your Activity implementation
type ListFilesActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ListFilesActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ListFilesActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ListFilesActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	directory := context.GetInput("directory").(string)
	recursive := context.GetInput("recursive").(bool)

	// Use the log object to log the greeting
	log.Debugf("listfiles of %s - recursively : %t", directory, recursive)

	// Set the result as part of the context
	context.SetOutput("filenames", list(true, directory, recursive))

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}

func list(isRoot bool, directory string, recursive bool) (filenames []string) {

	filenames = make([]string, 0)

	fileInfo, err := os.Stat(directory)

	if err == nil {

		if !isRoot {
			filenames = append(filenames, directory)
			log.Debugf("filename : %s", directory)
		}

		if (isRoot || recursive) && fileInfo.IsDir() {
			d, err := os.Open(directory)

			if err == nil {
				defer d.Close()

				names, err := d.Readdirnames(-1)
				if err == nil {

					for _, name := range names {
						filenames = append(filenames, list(false, filepath.Join(directory, name), recursive)...)
					}
				}
			}
		}
	}

	return filenames
}
