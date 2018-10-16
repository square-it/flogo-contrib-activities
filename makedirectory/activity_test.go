package removefile

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

	//setup attrs

	path := filepath.Join(dir, "path")

	tc.SetInput("path", path)

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, existsFile(path))
}

func TestEvalPermissions(t *testing.T) {

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

	//setup attrs

	path := filepath.Join(dir, "path")

	tc.SetInput("path", path)
	tc.SetInput("permissions", "0700")

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, existsFile(path))
	assert.True(t, checkPermissions(path,0700))

}

func TestEvalAll(t *testing.T) {

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

	//setup attrs

	path := filepath.Join(dir, "path")
	path2 := filepath.Join(path, "path2")
	path3 := filepath.Join(path2, "path3")

	tc.SetInput("path", path3)
	tc.SetInput("all", true)

	done, err := act.Eval(tc)

	//check result attr
	assert.True(t, done)
	assert.True(t, err == nil)
	assert.True(t, existsFile(path))
	assert.True(t, existsFile(path2))
	assert.True(t, existsFile(path3))
}

func prepareTests(t *testing.T) (dir string) {
	dir = filepath.Join(os.TempDir(), "flogo-tmpdir-tests", t.Name())
	dirs := filepath.Join(dir, "tmpdir")
	err := os.MkdirAll(dirs, 0700)
	if err != nil {
		t.Error(err)
		t.Failed()
		return
	}
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

func checkPermissions(path string, permissions os.FileMode) (ok bool) {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().Perm() == permissions {
			ok = true
		} else {
			ok = false
		}
	} else {
		ok = false
	}
	return ok
}
