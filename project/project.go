package project

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Frank-Mayer/vied/project/outline"
	"github.com/charmbracelet/log"
)

type Project struct {
	Root    string
	Outline *outline.Outline
}

func Init(root string) (*Project, error) {
	projectPath, err := filepath.Abs(root)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to get absolute path of %q", root), err)
	}
	p := &Project{
		Root: projectPath,
	}

	if p.Outline, err = outline.Init(projectPath); err != nil {
		return nil, errors.Join(errors.New("failed to create new filetree object"), err)
	}

	log.Info("project loaded", "root", p.Root)

	return p, nil
}

func (p *Project) Close() error {
	return nil
}

func (p *Project) Reload(file string) {}
