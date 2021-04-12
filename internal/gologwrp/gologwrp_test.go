package main

import (
    "testing"
	"github.com/ademirfteo/gologwrp/pkg/logger"
)

func BenchmarkLog_Debug2Info_Simple(b *testing.B) {

	log := logger.New().Level("info").JsonFormat()

    for i := 0; i < b.N; i++ {
		log.Debug("Test simple debug level log.")
    }
}

func BenchmarkLog_Debug2Info_WithFields(b *testing.B) {

	log := logger.New().Level("info").JsonFormat()

    for i := 0; i < b.N; i++ {
		log.WithFields(map[string]interface{} {
			"foo1": "bar1",
			"foo2": "bar2",
		}).Debug("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_Simple(b *testing.B) {

	log := logger.New().Level("info").JsonFormat()

    for i := 0; i < b.N; i++ {
		log.Info("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_WithFields(b *testing.B) {

	log := logger.New().Level("info").JsonFormat()

    for i := 0; i < b.N; i++ {
		log.WithFields(map[string]interface{} {
			"foo1": "bar1",
			"foo2": "bar2",
		}).Info("Test simple debug level log.")
    }
}

var (
	glog = logger.New().Level("info").JsonFormat()
)

func BenchmarkLog_Debug2Info_G_Simple(b *testing.B) {

    for i := 0; i < b.N; i++ {
		glog.Debug("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_G_Simple(b *testing.B) {

    for i := 0; i < b.N; i++ {
		glog.Info("Test simple debug level log.")
    }
}
