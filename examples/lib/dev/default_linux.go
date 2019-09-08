package dev

import (
	"github.com/apersenius/ble"
	"github.com/apersenius/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
