package common

import "github.com/charmbracelet/log"

type dirs struct {
	Config string
	Data   string
}

var Dirs dirs

func init() {
	var err error
	Dirs, err = getGlobalDirs()
	if err != nil {
		log.Fatal("failed to init global directories", "err", err)
	}
}
