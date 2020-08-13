package server

import (
	"net/http"
	"os"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

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
}

type MemoryStats struct {
	TotalMemObtainedFromSysMb uint64 `json:"totalMemObtainedFromSysMb"`
	TotalMemAllocatedMb       uint64 `json:"totalMemAllocatedMb"`
	MemAllocatedMb            uint64 `json:"memAllocatedMb"`
	LastGcEpoch               uint64 `json:"lastGcEpoch"`
	NumGc                     uint32 `json:"numGc"`
}

var doOnce sync.Once

var gomaxprocs int
var numCpu int
var pid int
var pPid int
var operatingSystem string
var arch string

var memStats runtime.MemStats

func (srv *AppServer) handleHealth(ctx *gin.Context) {
	log.Debug().Msg("Getting application health information")
	appHealth := getAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}

func getAppHealth() AppHealth {
	doOnce.Do(func() {
		log.Debug().Msg("Performing one-time lookup of constant runtime information")

		gomaxprocs = runtime.GOMAXPROCS(-1)
		numCpu = runtime.NumCPU()
		pid = os.Getpid()
		pPid = os.Getppid()
		operatingSystem = runtime.GOOS
		arch = runtime.GOARCH
	})

	log.Debug().Msg("Getting runtime information")

	appHealth := AppHealth{
		Goroutines:      runtime.NumGoroutine(),
		CGoCalls:        runtime.NumCgoCall(),
		Gomaxprocs:      gomaxprocs,
		NumCpu:          numCpu,
		Pid:             pid,
		PPid:            pPid,
		OperatingSystem: operatingSystem,
		Arch:            arch,
		MemStats:        getMemStats(),
	}

	return appHealth
}

func getMemStats() MemoryStats {
	log.Debug().Msg("Getting memory stats")

	runtime.ReadMemStats(&memStats)

	memStats := MemoryStats{
		TotalMemObtainedFromSysMb: bytesToMB(memStats.Sys),
		TotalMemAllocatedMb:       bytesToMB(memStats.TotalAlloc),
		MemAllocatedMb:            bytesToMB(memStats.Alloc),
		LastGcEpoch:               memStats.LastGC,
		NumGc:                     memStats.NumGC,
	}

	return memStats
}

func bytesToMB(bytes uint64) uint64 {
	return bytes / 1024 / 1024
}
