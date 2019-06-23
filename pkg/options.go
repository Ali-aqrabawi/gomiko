package gomiko

import (
	"github.com/Ali-aqrabawi/gomiko/pkg/types"
)

type OptionDevice func(interface{}) error

func OptionCiscoSecret(secret string) func(device interface{}) error {
	return func(device interface{}) error {
		device.(types.CiscoDevice).SetSecret(secret)
		return nil
	}
}
