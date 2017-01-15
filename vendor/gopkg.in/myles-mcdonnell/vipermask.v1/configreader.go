package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"time"
)

type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
}

type ConfigReader struct {
	vipers []*viper.Viper
}

func NewFromFiles(files ...string) Config {

	arr := make([]*viper.Viper, len(files))

	for index, file := range files {
		arr[index] = buildAndRead(file)
	}

	return NewFromVipers(arr[:])
}

func buildAndRead(file string) *viper.Viper {

	viper := viper.New()
	viper.SetConfigFile(file)
	viper.ReadInConfig()
	return viper
}

func NewFromVipers(vipers []*viper.Viper) Config {
	lv := new(ConfigReader)
	lv.vipers = vipers
	return lv
}

func (v *ConfigReader) get(key string) interface{} {
	for _, v := range v.vipers {
		val := v.Get(key)
		if val != nil {
			return val
		}
	}
	return nil
}

func (v *ConfigReader) GetString(key string) string {

	return cast.ToString(v.get(key))
}

func (v *ConfigReader) GetBool(key string) bool {

	return cast.ToBool(v.get(key))
}

func (v *ConfigReader) GetInt(key string) int {

	return cast.ToInt(v.get(key))
}

func (v *ConfigReader) GetFloat64(key string) float64 {

	return cast.ToFloat64(v.get(key))
}

func (v *ConfigReader) GetTime(key string) time.Time {

	return cast.ToTime(v.get(key))
}

func (v *ConfigReader) GetDuration(key string) time.Duration {

	return cast.ToDuration(v.get(key))
}
