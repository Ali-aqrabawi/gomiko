package nokia

import (
	"github.com/asadarafat/gomiko/pkg/connections"
	"github.com/asadarafat/gomiko/pkg/driver"
	"github.com/asadarafat/gomiko/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\n")
	return &SROSDevice{
		Prompt: "",
		Driver: devDriver,
	}, nil
}
