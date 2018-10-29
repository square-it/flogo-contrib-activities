package command

import (
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func init() {
	log.SetLogLevel(logger.ErrorLevel)
}

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEvalInputCommandRequired(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err == nil)
}

func TestEvalInputCommandWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", make(chan int))

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err == nil)
}

func TestEvalInputArgumentsDefault(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("wait", act.Metadata().Input["wait"].Value())
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", act.Metadata().Input["timeout"].Value())

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") != "")
}

func TestEvalInputArguments(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	arguments := []interface{}{"-l"}
	tc.SetInput("arguments", arguments)
	tc.SetInput("wait", act.Metadata().Input["wait"].Value())
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", act.Metadata().Input["timeout"].Value())

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") != "")
}

func TestEvalInputArgumentsWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("arguments", make([]int, 0))

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err != nil)
}

func TestEvalInputDirectoryWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("directory", make(chan int))

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err == nil)
}

func TestEvalInputUseCurrentEnvironmentWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("useCurrentEnvironment", "not a boolean")

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err == nil)
}

func TestEvalInputEnvironment(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	environment := []interface{}{"KEY=VALUE"}
	tc.SetInput("environment", environment)
	tc.SetInput("wait", act.Metadata().Input["wait"].Value())
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", act.Metadata().Input["timeout"].Value())

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") != "")

	environmentStr := []string{"KEY=VALUE"}
	tc.SetInput("environment", environmentStr)

	done, err = act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") != "")
}

func TestEvalInputEnvironmentWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("environment", make([]int, 0))

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err != nil)
}

func TestEvalInputTimeoutWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("timeout", "not a string")

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err == nil)
}

func TestEvalInputWaitWrongType(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "ls")
	tc.SetInput("wait", "not a boolean")

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err == nil)
}

func TestLookPathNotFound(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("command", "notFound")
	tc.SetInput("wait", act.Metadata().Input["wait"].Value())
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", act.Metadata().Input["timeout"].Value())

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err != nil)
}

func TestRun(t *testing.T) {

	dir := prepareTests(t)

	defer func() {
		os.RemoveAll(dir)

		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	arguments := make([]string, 0)
	arguments = append(arguments, "-1", dir)

	tc.SetInput("command", "ls")
	tc.SetInput("arguments", arguments)
	tc.SetInput("wait", act.Metadata().Input["wait"].Value())
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", act.Metadata().Input["timeout"].Value())

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") == "a\nb\nc\n")
	assert.True(t, tc.GetOutput("success") == true)
	assert.True(t, tc.GetOutput("exitCode") == 0)
}

func TestRunTimeout(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	arguments := make([]string, 0)
	arguments = append(arguments, "2")

	tc.SetInput("command", "sleep")
	tc.SetInput("arguments", arguments)
	tc.SetInput("wait", act.Metadata().Input["wait"].Value())
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", 1)

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") == "")
	assert.True(t, tc.GetOutput("success") == false)
	assert.True(t, tc.GetOutput("exitCode") != 0)
}

func TestRunNoWaiting(t *testing.T) {

	dir := prepareTests(t)

	defer func() {
		os.RemoveAll(dir)

		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	arguments := make([]string, 0)
	arguments = append(arguments, "-c", "sleep 1 && ls -1 "+dir)

	tc.SetInput("command", "sh")
	tc.SetInput("arguments", arguments)
	tc.SetInput("wait", false)
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", act.Metadata().Input["timeout"].Value())

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") == "")
	assert.True(t, tc.GetOutput("success") == true)
	assert.True(t, tc.GetOutput("exitCode") == 0)

	time.Sleep(2 * time.Second)
	assert.True(t, tc.GetOutput("output") == "a\nb\nc\n")
	assert.True(t, tc.GetOutput("success") == true)
	assert.True(t, tc.GetOutput("exitCode") == 0)
}

func TestRunNoWaitingTimeout(t *testing.T) {

	dir := prepareTests(t)

	defer func() {
		os.RemoveAll(dir)

		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	arguments := make([]string, 0)
	arguments = append(arguments, "-c", "sleep 1 && ls -1 "+dir)

	tc.SetInput("command", "sh")
	tc.SetInput("arguments", arguments)
	tc.SetInput("wait", false)
	tc.SetInput("useCurrentEnvironment", act.Metadata().Input["useCurrentEnvironment"].Value())
	tc.SetInput("timeout", 1)

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, tc.GetOutput("output") == "")
	assert.True(t, tc.GetOutput("success") == true)
	assert.True(t, tc.GetOutput("exitCode") == 0)

	time.Sleep(2 * time.Second)
	assert.True(t, tc.GetOutput("output") == "")
	assert.True(t, tc.GetOutput("success") == false)
	assert.True(t, tc.GetOutput("exitCode") != 0)
}

func prepareTests(t *testing.T) (dir string) {
	dir = filepath.Join(os.TempDir(), "flogo-tmpdir-tests", t.Name())

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}

	file := filepath.Join(dir, "a")
	f, err := os.Create(file)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
	f.Close()

	file = filepath.Join(dir, "b")
	f, err = os.Create(file)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
	f.Close()

	file = filepath.Join(dir, "c")
	f, err = os.Create(file)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
	f.Close()

	return dir
}
