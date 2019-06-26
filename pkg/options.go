package gomiko

import "github.com/Ali-aqrabawi/gomiko/pkg/types"

type DeviceOption func(interface{}) error

func SecretOption(secret string) func(device interface{}) error {
	return func(device interface{}) error {
		device.(types.CiscoDevice).SetSecret(secret)
		return nil
	}
}
