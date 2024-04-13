package parsers

import (
	"fmt"

	"github.com/Frank-Mayer/vied/project/lang"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
)

var parsers map[lang.Lang]*sitter.Parser = nil

func Parser(l lang.Lang) (*sitter.Parser, error) {
	if parsers == nil {
		parsers = make(map[lang.Lang]*sitter.Parser)
	}

	if p, ok := parsers[l]; ok {
		return p, nil
	}

	tsLang := getLang(l)
	if tsLang == nil {
		return nil, fmt.Errorf("unknown language %d", l)
	}

	p := sitter.NewParser()
	p.SetLanguage(tsLang)
	parsers[l] = p

	return p, nil
}

func getLang(l lang.Lang) *sitter.Language {
	switch l {
	case lang.Golang:
		return golang.GetLanguage()
	}
	return nil
}
