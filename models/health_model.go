package models

type AppHealth struct {
	Goroutines      int         `json:"goroutines"`
	CGoCalls        int64       `json:"cgoCalls"`
	Gomaxprocs      int         `json:"gomaxprocs"`
	NumCpu          int         `json:"numCpu"`
	Pid             int         `json:"pid"`
	PPid            int         `json:"ppid"`
	OperatingSystem string      `json:"os"`
	Arch            string      `json:"arch"`
	MemStats        MemoryStats `json:"memStats"`
  Status          string      `json:"status"`
}

type MemoryStats struct {
	TotalMemObtainedFromSysMb uint64 `json:"totalMemObtainedFromSysMb"`
	TotalMemAllocatedMb       uint64 `json:"totalMemAllocatedMb"`
	MemAllocatedMb            uint64 `json:"memAllocatedMb"`
	LastGcEpoch               uint64 `json:"lastGcEpoch"`
	NumGc                     uint32 `json:"numGc"`
}
