package internal

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	ModelName string
	Cores     int32
	Mhz       float64
}

func NewCPUInfo() ([]CPUInfo, error) {
	var cpuInfo []CPUInfo
	cpus, err := cpu.Info()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load CPU info")
	}

	for _, cpu := range cpus {
		cpuInfo = append(cpuInfo, CPUInfo{
			ModelName: cpu.ModelName,
			Cores:     cpu.Cores,
			Mhz:       cpu.Mhz,
		})
	}

	return cpuInfo, nil
}
