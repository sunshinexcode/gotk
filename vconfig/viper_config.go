package vconfig

import "github.com/spf13/viper"

type (
	DecoderConfigOption = viper.DecoderConfigOption
)

func Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error {
	return viper.Unmarshal(rawVal, opts...)
}
