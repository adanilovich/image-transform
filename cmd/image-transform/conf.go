package main

import (
	"errors"

	"github.com/asfeather.com/internal/myconf"
)

func newConf(confPath string) (*myconf.MyConf, error) {
	conf, err := myconf.New(confPath)
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, errors.New("Bad file parsing")
	}

	return conf, nil
}
