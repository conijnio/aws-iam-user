package config

import (
	"github.com/go-ini/ini"
	osUser "os/user"
	"path"
)

type IConfig interface {
	UpdateKeys(accessKey string, secretAccessKey string) error
	Rollback(error error) error
	AccessKeyId() string
}

type Config struct {
	profile           string
	cfg               *ini.File
	section           *ini.Section
	accessKeyId       string
	secretAccessKeyId string
	file              IFile
}

func New(profile string) *Config {
	return &Config{profile: profile, file: &File{}}
}

func (c *Config) resolveConfigPath() string {
	usr, _ := osUser.Current()
	return path.Join(usr.HomeDir, ".aws", "credentials")
}

func (c *Config) readCurrentKeys() {
	if c.section == nil {
		c.cfg, _ = c.file.readConfig(c.resolveConfigPath())
		c.section = c.cfg.Section(c.profile)
	}

	c.accessKeyId = c.section.Key("aws_access_key_id").String()
	c.secretAccessKeyId = c.section.Key("aws_secret_access_key").String()
}

func (c *Config) AccessKeyId() string {
	return c.accessKeyId
}

func (c *Config) UpdateKeys(accessKey string, secretAccessKey string) error {
	c.readCurrentKeys()
	c.section.Key("aws_access_key_id").SetValue(accessKey)
	c.section.Key("aws_secret_access_key").SetValue(secretAccessKey)

	return c.file.writeConfig(c.resolveConfigPath(), c.cfg)
}

func (c *Config) Rollback(err error) error {
	c.section.Key("aws_access_key_id").SetValue(c.accessKeyId)
	c.section.Key("aws_secret_access_key").SetValue(c.secretAccessKeyId)

	c.file.writeConfig(c.resolveConfigPath(), c.cfg)
	return err
}
