package internal

import (
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
