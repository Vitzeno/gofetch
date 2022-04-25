package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func main() {
	mem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load memory info: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("mem.Free: %v\n", mem.Free)
	fmt.Printf("mem.Used: %v\n", mem.Used)
	fmt.Printf("mem.Total: %v\n", mem.Total)
	fmt.Printf("mem.UsedPercent: %v\n", mem.UsedPercent)

	cpusInfo, err := cpu.Info()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load CPU info: %v \n", err)
		os.Exit(1)
	}

	for _, cpu := range cpusInfo {
		fmt.Printf("cpu.ModelName: %v\n", cpu.ModelName)
		fmt.Printf("cpu.Cores: %v\n", cpu.Cores)
		fmt.Printf("cpu.Mhz: %v\n", cpu.Mhz)
	}

	diskInfo, err := disk.Usage("/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load disk info: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("diskInfo.Free: %v\n", diskInfo.Free)
	fmt.Printf("diskInfo.Used: %v\n", diskInfo.Used)
	fmt.Printf("diskInfo.Total: %v\n", diskInfo.Total)
	fmt.Printf("diskInfo.UsedPercent: %v\n", diskInfo.UsedPercent)

	hostInfo, err := host.Info()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load host info: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("hostInfo.OS: %v\n", hostInfo.OS)
	fmt.Printf("hostInfo.KernelVersion: %v\n", hostInfo.KernelVersion)
	fmt.Printf("hostInfo.Platform: %v\n", hostInfo.Platform)
	fmt.Printf("hostInfo.Uptime: %v\n", hostInfo.Uptime)

	processes, err := process.Pids()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load process info: %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("num of processes: %v\n", len(processes))

}
