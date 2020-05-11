// +build !linux

package device

import (
	"github.com/k-vpm/wireguard-go/conn"
	"github.com/k-vpm/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
