package sleep

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"time"
)

var log = logger.GetLogger("activity-sleep")

// SleepActivity is a stub for your Activity implementation
type SleepActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SleepActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *SleepActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *SleepActivity) Eval(activityContext activity.Context) (done bool, err error) {
	// check input timeToSleepInMs
	durationParameter, ok := activityContext.GetInput("duration").(string)

	if !ok {
		log.Error("The duration parameter is not a string.")
		return false, nil
	}

	duration, err := time.ParseDuration(durationParameter)

	if err != nil {
		log.Errorf("Unable to parse duration '%s'.", durationParameter)
		log.Error("Please read available format at https://golang.org/pkg/time/#ParseDuration.")
		return false, nil
	}

	if duration == 0 {
		log.Debug("No need to sleep.")
		return true, nil
	}

	log.Debugf("Amount of time to sleep: %s", duration)

	time.Sleep(duration)

	log.Debugf("Thread has slept for %d ms", duration)

	return true, nil
}
