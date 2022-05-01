package internal

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Free        uint64
	Used        uint64
	Total       uint64
	UsedPercent float64
}

func NewDiskInfo(path string) (DiskInfo, error) {
	diskInfo, err := disk.Usage(path)
	if err != nil {
		return DiskInfo{}, errors.Wrap(err, "Failed to load disk info")
	}

	return DiskInfo{
		Free:        diskInfo.Free,
		Used:        diskInfo.Used,
		Total:       diskInfo.Total,
		UsedPercent: diskInfo.UsedPercent,
	}, nil
}

func (d DiskInfo) String() string {
	return fmt.Sprintf("Free: %v GB\nUsed: %v GB\nTotal: %v GB\nUsedPercent: %.0f%%\n", d.Free/1024/1024/1024, d.Used/1024/1024/1024, d.Total/1024/1024/1024, d.UsedPercent)
}
