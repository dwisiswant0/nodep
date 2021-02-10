package main

import (
	"errors"
	"fmt"
)

func (opt *options) validate() error {
	if opt.reg == "" {
		return errors.New("no registry given")
	}

	switch opt.reg {
	case "npm", "pip", "gem":
	default:
		return fmt.Errorf("invalid registry for '%s'", opt.reg)
	}

	if opt.dep == "" {
		return errors.New("no dependencies given")
	}

	if list, err := readFile(opt.dep); err == nil {
		opt.list = list
	} else {
		opt.list = []string{opt.dep}
	}

	return nil
}
