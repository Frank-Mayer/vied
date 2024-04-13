package main

import (
	"github.com/Frank-Mayer/vied/project"
	"github.com/Frank-Mayer/vied/ui"
	"github.com/charmbracelet/log"
)

func main() {
	p, err := project.Init(".")
	if err != nil {
		log.Fatal("failed to load project", "err", err)
	}
	defer p.Close()

	ui.Init(p)
}
