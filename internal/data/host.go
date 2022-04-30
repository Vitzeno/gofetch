package internal

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	OS            string
	KernelVersion string
	Platform      string
	Uptime        uint64
}

func NewHostInfo() (HostInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return HostInfo{}, errors.Wrap(err, "Failed to load host info")
	}

	return HostInfo{
		OS:            hostInfo.OS,
		KernelVersion: hostInfo.KernelVersion,
		Platform:      hostInfo.Platform,
		Uptime:        hostInfo.Uptime,
	}, nil
}

func (h HostInfo) String() string {
	return fmt.Sprintf("OS: %v\nKernelVersion: %v\nPlatform: %v\nUptime: %v\n", h.OS, h.KernelVersion, h.Platform, h.Uptime)
}
