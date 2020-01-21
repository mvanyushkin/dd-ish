package dd

import (
	"errors"
	"github.com/mvanyushkin/dd-ish/dd/settings"
	"io"
	"os"
)

type CopySession struct {
	sourceFile *os.File
	targetFile *os.File
	offset     uint64
	limit      uint64
}

func DoCopy(config settings.Settings, moveProgressCallback func(float32)) error {
	cs := CopySession{
		offset: config.Offset,
		limit:  config.Limit,
	}
	defer cs.Close()
	e := cs.OpenSourceAndTarget(config)
	if e != nil {
		return e
	}

	return cs.DoCopyInternal(moveProgressCallback)
}

func (c *CopySession) DoCopyInternal(moveProgressCallback func(float32)) error {
	targetSize, err := c.PrepareCopying()
	if err != nil {
		return err
	}

	var bufferSize uint64 = 1024
	if c.limit < bufferSize {
		bufferSize = c.limit
	}

	var currentPosition int64 = 0
	bytesToRead := c.limit
	buffer := make([]byte, bufferSize)
	for {
		readCount, e := c.sourceFile.Read(buffer)
		if e == io.EOF || bytesToRead <= 0 {
			break
		}

		bytesToRead -= uint64(readCount)
		_, e = c.targetFile.Write(buffer[0:readCount])
		if e != nil {
			return e
		}

		currentPosition += int64(readCount)
		progress := float64(currentPosition) / float64(targetSize) * 100
		moveProgressCallback(float32(progress))
	}

	return nil
}

func (c *CopySession) Close() {
	if c.sourceFile != nil {
		c.sourceFile.Close()
	}

	if c.targetFile != nil {
		c.targetFile.Close()
	}
}

func (c *CopySession) OpenSourceAndTarget(config settings.Settings) error {
	src, e := os.Open(config.SourcePath)
	if e != nil {
		return e
	}

	c.sourceFile = src
	os.Remove(config.TargetPath)
	dst, e := os.Create(config.TargetPath)
	if e != nil {
		return e
	}

	c.targetFile = dst
	return nil
}

func (c *CopySession) PrepareCopying() (uint64, error) {
	sourceStat, e := c.sourceFile.Stat()
	if e != nil {
		return 0, e
	}

	if sourceStat.Size() <= int64(c.offset) {
		return 0, errors.New("invalid offset")
	}

	var _, seekError = c.sourceFile.Seek(int64(c.offset), io.SeekStart)
	if seekError != nil {
		return 0, seekError
	}

	fi, e := c.sourceFile.Stat()
	if e != nil {
		return 0, e
	}

	if c.limit != 0 {
		return c.limit, nil
	} else {
		return uint64(fi.Size()) - c.offset, nil
	}
}
