package config

import "github.com/go-ini/ini"

type IFile interface {
	readConfig(filename string) (*ini.File, error)
	writeConfig(filename string, cfg *ini.File) error
}

type File struct{}

func (f *File) readConfig(filename string) (*ini.File, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (f *File) writeConfig(filename string, cfg *ini.File) error {
	return cfg.SaveTo(filename)
}
