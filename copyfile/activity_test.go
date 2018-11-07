package copyfile

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

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

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	done, err := act.Eval(tc)

	//check result attr
	assert.False(t, done)
	assert.True(t, err != nil)
}

func TestEvalCopyFile(t *testing.T) {
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

	source := filepath.Join(dir, "tmpfile")
	destination := filepath.Join(dir, "tmpfile2")
	tc.SetInput("source", source)
	tc.SetInput("destination", destination)

	done, err := act.Eval(tc)

	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, existsFile(destination))
}

func prepareTests(t *testing.T) (dir string) {
	dir = filepath.Join(os.TempDir(), "flogo-tmpdir-tests", t.Name())
	dirs := filepath.Join(dir, "tmpdir")
	err := os.MkdirAll(dirs, 0777)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
	file := filepath.Join(dir, "tmpfile")
	f, err := os.Create(file)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
	f.Close()
	file = filepath.Join(dirs, "tmpfile")
	f, err = os.Create(file)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
	f.Close()

	return dir
}

func existsFile(path string) (exists bool) {
	if _, err := os.Stat(path); err == nil {
		exists = true
	} else {
		exists = false
	}
	return exists
}