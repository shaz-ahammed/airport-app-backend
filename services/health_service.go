package services

import (
	"airport-app-backend/middleware"
	"airport-app-backend/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
	"os"
	"runtime"
	"sync"
)

var doOnce sync.Once

var gomaxprocs int
var numCpu int
var pid int
var pPid int
var operatingSystem string
var arch string

var memStats runtime.MemStats

type IHealthRepository interface {
	GetAppHealth(c context.Context, ctx *gin.Context) models.AppHealth
}

func (repo *ServiceRepository) GetAppHealth(c context.Context, ctx *gin.Context) models.AppHealth {
	_, span := trace.StartSpan(c, "get_app_health")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)

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

	appHealth := models.AppHealth{
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

func getMemStats() models.MemoryStats {
	log.Debug().Msg("Getting memory stats")

	runtime.ReadMemStats(&memStats)

	memStats := models.MemoryStats{
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
