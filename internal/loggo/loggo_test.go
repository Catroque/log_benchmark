package main

import (
    "testing"
	"github.com/Catroque/loggo/pkg/logger"
)

func BenchmarkLog_Debug2Info_Simple(b *testing.B) {

	log := logger.New().SetLevel(logger.InfoLevel)

    for i := 0; i < b.N; i++ {
	    log.Debug().Msg("Test simple debug level log.")
    }
}

func BenchmarkLog_Debug2Info_WithFields(b *testing.B) {

	log := logger.New().SetLevel(logger.InfoLevel)

    for i := 0; i < b.N; i++ {
	    log.Debug().Fields(map[string]interface{} {
			"foo1": "bar1",
			"foo2": "bar2",
		}).Msg("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_Simple(b *testing.B) {

	log := logger.New().SetLevel(logger.InfoLevel)

    for i := 0; i < b.N; i++ {
	    log.Info().Msg("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_WithFields(b *testing.B) {

	log := logger.New().SetLevel(logger.InfoLevel)

    for i := 0; i < b.N; i++ {
	    log.Info().Fields(map[string]interface{} {
			"foo1": "bar1",
			"foo2": "bar2",
		}).Msg("Test simple debug level log.")
    }
}

var (
	glog = logger.New().SetLevel(logger.InfoLevel)
)

func BenchmarkLog_Debug2Info_G_Simple(b *testing.B) {

    for i := 0; i < b.N; i++ {
	    glog.Debug().Msg("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_G_Simple(b *testing.B) {

    for i := 0; i < b.N; i++ {
	    glog.Info().Msg("Test simple debug level log.")
    }
}
