package settings

import (
	"flag"
	"sync"
)

type Settings struct {
	SourcePath string
	TargetPath string
	Offset     uint64
	Limit      uint64
}

var once sync.Once

var settings Settings

func Instance() *Settings {
	once.Do(func() {
		settings = Settings{
			SourcePath: "",
			TargetPath: "",
			Offset:     0,
		}

		flag.StringVar(&settings.SourcePath, "src", "", "the source file, must be final")
		flag.StringVar(&settings.TargetPath, "dst", "", "the destination file")
		flag.Uint64Var(&settings.Offset, "offset", 0, "source")
		flag.Uint64Var(&settings.Limit, "limit", 0, "source")
		flag.Parse()
	})
	return &settings
}
