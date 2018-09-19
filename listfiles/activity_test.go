package listfiles

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	directory := "test"
	tc.SetInput("directory", directory)
	tc.SetInput("recursive", false)
	done, _ := act.Eval(tc)

	assert.True(t, done)

	//check result attr
	result := tc.GetOutput("filenames")
	assert.Equal(t, 4, len(result.([]string)))
	assert.Contains(t, result, filepath.Join(directory, "a"))
	assert.Contains(t, result, filepath.Join(directory, "b"))
	assert.Contains(t, result, filepath.Join(directory, "c"))
	assert.Contains(t, result, filepath.Join(directory, "test2"))

}

func TestEvalRecursive(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	directory := "test"
	tc.SetInput("directory", directory)
	tc.SetInput("recursive", true)
	done, _ := act.Eval(tc)

	assert.True(t, done)

	//check result attr
	result := tc.GetOutput("filenames")
	assert.Equal(t, 7, len(result.([]string)))
	assert.Contains(t, result, filepath.Join(directory, "a"))
	assert.Contains(t, result, filepath.Join(directory, "b"))
	assert.Contains(t, result, filepath.Join(directory, "c"))
	assert.Contains(t, result, filepath.Join(directory, "test2", "d"))
	assert.Contains(t, result, filepath.Join(directory, "test2", "e"))
	assert.Contains(t, result, filepath.Join(directory, "test2", "f"))
	assert.Contains(t, result, filepath.Join(directory, "test2"))

}
