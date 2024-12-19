package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

func ToIntE(i interface{}) (int, error) {
	return cast.ToIntE(i)
}

func ToInt(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

func ToInt8E(i interface{}) (int8, error) {
	return cast.ToInt8E(i)
}

func ToInt8(i interface{}) int8 {
	v, _ := ToInt8E(i)
	return v
}

func ToInt16E(i interface{}) (int16, error) {
	return cast.ToInt16E(i)
}

func ToInt16(i interface{}) int16 {
	v, _ := ToInt16E(i)
	return v
}

func ToInt32E(i interface{}) (int32, error) {
	return cast.ToInt32E(i)
}

func ToInt32(i interface{}) int32 {
	v, _ := ToInt32E(i)
	return v
}

func ToInt64E(i interface{}) (int64, error) {
	return cast.ToInt64E(i)
}

func ToInt64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

func ToBoolE(i interface{}) (bool, error) {
	return cast.ToBoolE(i)
}

func ToBool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

func ToBoolSliceE(i interface{}) ([]bool, error) {
	return cast.ToBoolSliceE(i)
}

func ToBoolSlice(i interface{}) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

func ToStringE(i interface{}) (string, error) {
	return cast.ToStringE(i)
}

func ToString(i interface{}) string {
	v, _ := ToStringE(i)
	return v
}

func ToFloat32E(i interface{}) (float32, error) {
	return cast.ToFloat32E(i)
}

func ToFloat32(i interface{}) float32 {
	v, _ := ToFloat32E(i)
	return v
}

func ToFloat64E(i interface{}) (float64, error) {
	return cast.ToFloat64E(i)
}

func ToFloat64(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

func ToStringMapE(i interface{}) (map[string]interface{}, error) {
	return cast.ToStringMapE(i)
}

func ToStringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

func ToStringSliceE(i interface{}) ([]string, error) {
	return cast.ToStringSliceE(i)
}

func ToStringSlice(i interface{}) []string {
	v, _ := ToStringSliceE(i)
	return v
}

func ToDurationE(i interface{}) (d time.Duration, err error) {
	return cast.ToDurationE(i)
}

func ToDuration(i interface{}) (d time.Duration) {
	v, _ := ToDurationE(i)
	return v
}

func ToMapFromStructE(input interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &result,
	})
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ToMapFromStruct(input interface{}) map[string]interface{} {
	v, _ := ToMapFromStructE(input)
	return v
}

func ToStructFromMapE(input interface{}, out interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  out,
	})
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func ToStructFromMapEWeakWithTag(input interface{}, out interface{}, tag string) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          tag,
		Result:           out,
		WeaklyTypedInput: true,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func ToStructFromStringE(input interface{}, out interface{}) error {
	var data []byte
	switch s := input.(type) {
	case string:
		data = []byte(s)
	case []byte:
		data = s
	default:
		return fmt.Errorf("input:%T. need:string or []byte", input)
	}

	return json.Unmarshal(data, out)
}

func ToStringFromStructE(input interface{}) (string, error) {
	dat, err := json.Marshal(input)
	return string(dat), err
}

func ToStringFromStructWithoutEscape(input interface{}) string {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(input)
	return string(buffer.Bytes())
}

func ToStringFromStruct(input interface{}) string {
	v, _ := ToStringFromStructE(input)
	return v
}

func ToStringFromMapE(input interface{}) (string, error) {
	dat, err := json.Marshal(input)
	return string(dat), err
}

func ToStringFromMap(input interface{}) string {
	v, _ := ToStringFromMapE(input)
	return v
}
