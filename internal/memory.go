package internal

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Free        uint64
	Used        uint64
	Total       uint64
	UsedPercent float64
}

func NewMemInfo() (MemInfo, error) {
	mem, err := mem.VirtualMemory()
	if err != nil {
		return MemInfo{}, errors.Wrap(err, "Failed to load memory info")
	}

	return MemInfo{
		Free:        mem.Free,
		Used:        mem.Used,
		Total:       mem.Total,
		UsedPercent: mem.UsedPercent,
	}, nil
}
