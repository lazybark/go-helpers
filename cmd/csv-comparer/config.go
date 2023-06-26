package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

type Config struct {
	First        string `arg:"-f"`
	Second       string `arg:"-s"`
	SepOne       string `arg:"--sepone"`
	SepTwo       string `arg:"--septwo"`
	KeyCol       string `arg:"--keycol"`
	ColsString   string `arg:"--cols"`
	WriteDiffs   bool   `arg:"--wdiff"`
	DiffPath     string `arg:"--diffpath"`
	WriteDeleted bool   `arg:"--wdel"`
	DelPath      string `arg:"--delpath"`
}

func ParseEnv() (Config, error) {
	var c Config
	err := arg.Parse(&c)
	if err != nil {
		err = fmt.Errorf("[ParseEnv]: %w", err)
	}
	return c, err
}
