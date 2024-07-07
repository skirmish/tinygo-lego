package lego

import (
	"errors"

	"tinygo.org/x/bluetooth"
)

type Hub struct {
	device *bluetooth.Device
	svc    *bluetooth.DeviceService
	chr    *bluetooth.DeviceCharacteristic

	buf []byte
}

var (
	// BLE service
	// 00001623-1212-EFDE-1623-785FEABCD123
	hubService        = bluetooth.NewUUID([16]byte{0x00, 0x00, 0x16, 0x23, 0x12, 0x12, 0xef, 0xde, 0x16, 0x23, 0x78, 0x5f, 0xea, 0xbc, 0xd1, 0x23})
	hubCharacteristic = bluetooth.NewUUID([16]byte{0x00, 0x00, 0x16, 0x24, 0x12, 0x12, 0xef, 0xde, 0x16, 0x23, 0x78, 0x5f, 0xea, 0xbc, 0xd1, 0x23})
)

// NewHub creates a new LEGO hub.
func NewHub(dev *bluetooth.Device) *Hub {
	h := &Hub{
		device: dev,
		buf:    make([]byte, 255),
	}

	return h
}

func (h *Hub) Start() (err error) {
	srvcs, err := r.device.DiscoverServices([]bluetooth.UUID{
		hubService,
	})
	if err != nil || len(srvcs) == 0 {
		return errors.New("could not find services")
	}

	h.svc = &srvcs[0]
	println("found LEGO hub service", h.srv.UUID().String())

	chars, err := h.svc.DiscoverCharacteristics([]bluetooth.UUID{
		hubCharacteristic,
	})
	if err != nil || len(chars) == 0 {
		return errors.New("could not find LEGO hub characteristic")
	}

	h.chr = &chars[0]

	return
}
