// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package webapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsPluginConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementPluginConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsPluginConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookPluginConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementPluginConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_PluginConfig(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookPluginConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_PluginConfig(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_PluginConfig(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_PluginConfig(val, result))
}

func testDecodeRaw_PluginConfig(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_PluginConfig(vStringSlice, result))
}

func TestPluginConfig_GetPFlagSet(t *testing.T) {
	val := PluginConfig{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestPluginConfig_SetFlags(t *testing.T) {
	actual := PluginConfig{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_readRateLimiter.qps", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("readRateLimiter.qps", testValue)
			if vInt, err := cmdFlags.GetFloat64("readRateLimiter.qps"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.ReadRateLimiter.QPS)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_readRateLimiter.burst", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("readRateLimiter.burst", testValue)
			if vInt, err := cmdFlags.GetInt("readRateLimiter.burst"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.ReadRateLimiter.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_writeRateLimiter.qps", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("writeRateLimiter.qps", testValue)
			if vInt, err := cmdFlags.GetFloat64("writeRateLimiter.qps"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.WriteRateLimiter.QPS)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_writeRateLimiter.burst", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("writeRateLimiter.burst", testValue)
			if vInt, err := cmdFlags.GetInt("writeRateLimiter.burst"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.WriteRateLimiter.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.size", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("caching.size", testValue)
			if vInt, err := cmdFlags.GetInt("caching.size"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.Caching.Size)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.resyncInterval", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := DefaultPluginConfig.Caching.ResyncInterval.String()

			cmdFlags.Set("caching.resyncInterval", testValue)
			if vString, err := cmdFlags.GetString("caching.resyncInterval"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vString), &actual.Caching.ResyncInterval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.workers", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("caching.workers", testValue)
			if vInt, err := cmdFlags.GetInt("caching.workers"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.Caching.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.maxSystemFailures", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("caching.maxSystemFailures", testValue)
			if vInt, err := cmdFlags.GetInt("caching.maxSystemFailures"); err == nil {
				testDecodeJson_PluginConfig(t, fmt.Sprintf("%v", vInt), &actual.Caching.MaxSystemFailures)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
