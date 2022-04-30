package internal

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/mem"
)

type MemInfo struct {
	Free        uint64
	Used        uint64
	Total       uint64
	UsedPercent float64
	PageTables  uint64
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
		PageTables:  mem.PageTables,
	}, nil
}

func (m MemInfo) String() string {
	return fmt.Sprintf("Free: %v GB\nUsed: %v GB\nTotal: %v GB\nUsedPercent: %.0f%%\nPagedTables: %v\n", m.Free/1024/1024/1024, m.Used/1024/1024/1024, m.Total/1024/1024/1024, m.UsedPercent, m.PageTables)
}
