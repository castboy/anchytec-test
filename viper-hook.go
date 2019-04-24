package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"reflect"
)

//self-define hookFunc to parse `float64` type  to `decimal.Decimal` type.

func floatToDecimal() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Float64 {
			return data, nil
		}
		if t != reflect.TypeOf(decimal.Decimal{}) {
			return data, nil
		}

		// Convert it by parsing
		dc := decimal.NewFromFloat(data.(float64))

		return dc, nil
	}
}

func decodeHookWithTag(hook mapstructure.DecodeHookFunc, tagName string) viper.DecoderConfigOption {
	return func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = hook
		c.TagName = tagName
	}
}

func main() {
	iViper := viper.New()

	decodeHookFunc := floatToDecimal()
	decoderConfigOption := decodeHookWithTag(decodeHookFunc, "json")

	iViper.Set("salary", float64((100)))

	var dc decimal.Decimal

	err := iViper.UnmarshalKey("salary", &dc, decoderConfigOption)
	fmt.Println(err, dc)

	return
}
