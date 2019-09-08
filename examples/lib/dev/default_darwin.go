package dev

import (
	"github.com/apersenius/ble"
	"github.com/apersenius/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
