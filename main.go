package main

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/mvanyushkin/dd-ish/dd"
	"github.com/mvanyushkin/dd-ish/dd/settings"
	"os"
)

func main() {
	config := settings.Instance()
	bar := pb.StartNew(100)

	fmt.Printf("Copy from %v\n to %v\n offset %v\n", config.SourcePath, config.TargetPath, config.Offset)

	e := dd.DoCopy(*config, func(progress float32) {
		bar.SetCurrent(int64(progress))
	})

	if e != nil {
		fmt.Printf(e.Error())
		os.Exit(-1)
	}

	bar.Finish()
	os.Exit(0)
}
