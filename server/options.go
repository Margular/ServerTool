package server

import (
	"fmt"
	"net"
)

var GlobalOptions = NewOptions()

type Options struct {
	Host            string
	Port            uint
	DownloadBasedir string
	Debug           bool
}

func NewOptions() *Options {
	return &Options{}
}

func (opt *Options) validate() error {
	if err := opt.validateHost(); err != nil {
		return err
	}

	if err := opt.validatePort(); err != nil {
		return err
	}

	return nil
}

func (opt *Options) validateHost() error {
	ip := net.ParseIP(opt.Host)

	if ip == nil {
		return fmt.Errorf("error param host %s", opt.Host)
	}

	return nil
}

func (opt *Options) validatePort() error {
	if opt.Port > 65535 {
		return fmt.Errorf("listen on port %d is impossible", opt.Port)
	}

	return nil
}
