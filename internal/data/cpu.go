package internal

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	ModelName string
	Threads   int32
	Mhz       float64
	Usage     float64
}

func NewCPUInfo() ([]CPUInfo, error) {
	var cpuInfo []CPUInfo
	cpus, err := cpu.Info()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load CPU info")
	}

	usage, err := cpu.Percent(0, false)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load CPU usage")
	}

	for _, cpu := range cpus {
		cpuInfo = append(cpuInfo, CPUInfo{
			ModelName: cpu.ModelName,
			Threads:   cpu.Cores,
			Mhz:       cpu.Mhz,
			Usage:     usage[0],
		})
	}

	return cpuInfo, nil
}

func (c CPUInfo) String() string {
	return fmt.Sprintf("ModelName: %v\nThreads: %v\nMhz: %v\nUsage: %.0f%% \n", c.ModelName, c.Threads, c.Mhz, c.Usage)
}
