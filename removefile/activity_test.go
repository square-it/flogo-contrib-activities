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

func TestEvalRemoveFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	dir := prepareTests(t)
	path := filepath.Join(dir, "tmpfile")
	tc.SetInput("path", path)

	done, err := act.Eval(tc)

	assert.True(t, done)
	assert.True(t, err == nil)
	assert.False(t, existsFile(path), "File %s must be removed", path)
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir")))
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir", "tmpfile")))
}

func TestEvalRemoveFileRecursively(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	dir := prepareTests(t)
	path := filepath.Join(dir, "tmpfile")

	tc.SetInput("path", path)
	tc.SetInput("recursive", true)

	done, err := act.Eval(tc)

	assert.True(t, done)
	assert.True(t, err == nil)
	assert.False(t, existsFile(path), "File %s must be removed", path)
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir")))
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir", "tmpfile")))
}

func TestEvalRemoveUnknownFile(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	dir := prepareTests(t)
	unknwonPath := filepath.Join(dir, "tmpfile2")
	path := filepath.Join(dir, "tmpfile")
	tc.SetInput("path", unknwonPath)

	done, err := act.Eval(tc)

	assert.False(t, done)
	assert.True(t, err != nil)
	assert.True(t, existsFile(path))
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir")))
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir", "tmpfile")))
}

func TestEvalRemoveUnknownDirecory(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	dir := prepareTests(t)
	unknwonPath := filepath.Join(dir, "tmpdir2")
	path := filepath.Join(dir, "tmpfile")
	tc.SetInput("path", unknwonPath)

	done, err := act.Eval(tc)

	assert.False(t, done)
	assert.True(t, err != nil)
	assert.True(t, existsFile(path))
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir")))
	assert.True(t, existsFile(filepath.Join(dir, "tmpdir", "tmpfile")))
}

func TestEvalRemoveDirecoryRecursively(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	dir := prepareTests(t)
	path := filepath.Join(dir, "tmpdir")

	tc.SetInput("path", path)
	tc.SetInput("recursive", true)

	done, err := act.Eval(tc)

	assert.True(t, done)
	assert.True(t, err == nil)
	assert.False(t, existsFile(path))
	assert.False(t, existsFile(filepath.Join(path, "tmpfile")))
}

func TestEvalRemoveDirectory(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	dir := prepareTests(t)
	path := filepath.Join(dir, "tmpdir")
	file := filepath.Join(path, "tmpfile")

	tc.SetInput("path", file)

	done, err := act.Eval(tc)

	assert.True(t, done)
	assert.True(t, err == nil)

	tc.SetInput("path", path)

	done2, err2 := act.Eval(tc)

	assert.True(t, done2)
	assert.True(t, err2 == nil)

	assert.False(t, existsFile(file))
	assert.False(t, existsFile(path))
}

func existsFile(path string) (exists bool) {
	if _, err := os.Stat(path); err == nil {
		exists = true
	} else {
		exists = false
	}
	return exists
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
