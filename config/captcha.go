package config

import "time"

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}

type Attribute struct {
	Kind          string        `json:"kind" yaml:"kind" default:"letter"`
	Length        int           `json:"length" yaml:"length" default:"4"`
	Width         int           `json:"width" yaml:"width" default:"80"`
	Height        int           `json:"height" yaml:"height" default:"40"`
	Expire        time.Duration `json:"expire" yaml:"expire" default:"40"`
	CollectNumber int           `json:"collectNum" yaml:"collectNum" default:"10000"`
}
