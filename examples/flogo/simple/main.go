package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"

	_ "github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/core/support/log"

	_ "github.com/project-flogo/contrib/trigger/rest"
	_ "github.com/project-flogo/rules/ruleaction"
)

var (
	cpuProfile    = flag.String("cpuprofile", "", "Writes CPU profile to the specified file")
	memProfile    = flag.String("memprofile", "", "Writes memory profile to the specified file")
	cfgJson       string
	cfgCompressed bool
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to create CPU profiling file due to error - %s", err.Error()))
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}

	e, err := engine.New(cfg)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}

	code := engine.RunEngine(e)

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to create memory profiling file due to error - %s", err.Error()))
			os.Exit(1)
		}

		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Println(fmt.Sprintf("Failed to write memory profiling data to file due to error - %s", err.Error()))
			os.Exit(1)
		}
		f.Close()
	}

	os.Exit(code)
}
