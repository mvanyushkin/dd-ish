package dd

import (
	"github.com/mvanyushkin/dd-ish/dd/settings"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestSmoke(t *testing.T) {
	testSourceFileName := "test_source"

	wd, _ := os.Getwd()
	os.Remove(testSourceFileName)
	source, _ := os.Create(testSourceFileName)
	source.Write(make([]byte, 1024*16))
	source.Close()

	sourcePath := path.Join(wd, testSourceFileName)
	targetPath := path.Join(wd, "test_dst")
	e := DoCopy(settings.Settings{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Offset:     0,
	}, func(f float32) {})

	assert.Nil(t, e)

	sourceStat, e := os.Stat(sourcePath)
	targetStat, e := os.Stat(targetPath)

	assert.Equal(t, sourceStat.Size(), targetStat.Size())
	os.Remove(sourcePath)
	os.Remove(targetPath)
}

func TestWhenSourceDoesntExist(t *testing.T) {
	testSourceFileName := "doesnt_exist"
	wd, _ := os.Getwd()
	sourcePath := path.Join(wd, testSourceFileName)
	targetPath := path.Join(wd, "test_dst")
	e := DoCopy(settings.Settings{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Offset:     0,
	}, func(f float32) {})

	assert.NotNil(t, e)
}

func TestWhenTargetDirIsInvalid(t *testing.T) {
	testSourceFileName := "doesnt_exist"
	wd, _ := os.Getwd()
	sourcePath := path.Join(wd, testSourceFileName)
	targetPath := path.Join("zalupa", "test_dst")
	e := DoCopy(settings.Settings{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Offset:     0,
	}, func(f float32) {})

	assert.NotNil(t, e)
}

func TestWhenOffsetHasIncorrectValue(t *testing.T) {
	testSourceFileName := "doesnt_exist"
	wd, _ := os.Getwd()
	sourcePath := path.Join(wd, testSourceFileName)
	targetPath := path.Join("zalupa", "test_dst")
	e := DoCopy(settings.Settings{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Offset:     0,
	}, func(f float32) {})

	assert.NotNil(t, e)
}

func TestWhenOffsetGreaterThanSourceFileSize(t *testing.T) {
	testSourceFileName := "test_source"

	wd, _ := os.Getwd()
	os.Remove(testSourceFileName)
	source, _ := os.Create(testSourceFileName)
	source.Write(make([]byte, 1024*16))
	source.Close()

	sourcePath := path.Join(wd, testSourceFileName)
	targetPath := path.Join(wd, "test_dst")
	e := DoCopy(settings.Settings{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Offset:     1024*16 + 1,
	}, func(f float32) {})

	assert.NotNil(t, e)
	os.Remove(sourcePath)
	os.Remove(targetPath)
}

func TestWhenOffsetEqualToSourceSourceFileSize(t *testing.T) {
	testSourceFileName := "test_source"

	wd, _ := os.Getwd()
	os.Remove(testSourceFileName)
	source, _ := os.Create(testSourceFileName)
	source.Write(make([]byte, 1024*16))
	source.Close()

	sourcePath := path.Join(wd, testSourceFileName)
	targetPath := path.Join(wd, "test_dst")
	e := DoCopy(settings.Settings{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Offset:     1024 * 16,
	}, func(f float32) {})

	assert.NotNil(t, e)
	os.Remove(sourcePath)
	os.Remove(targetPath)
}
