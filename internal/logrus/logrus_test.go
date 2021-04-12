package main

import (
	"testing"
	log "github.com/sirupsen/logrus"
)

func BenchmarkLog_Debug2Info_Simple(b *testing.B) {

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "02-01-2006T15:04:05.000Z"})

    for i := 0; i < b.N; i++ {
		log.Debug("Test simple debug level log.")
    }
}

func BenchmarkLog_Debug2Info_WithFields(b *testing.B) {

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "02-01-2006T15:04:05.000Z"})

    for i := 0; i < b.N; i++ {
		log.WithFields(log.Fields {
			"foo1": "bar1",
			"foo2": "bar2",
		}).Debug("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_Simple(b *testing.B) {

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "02-01-2006T15:04:05.000Z"})

    for i := 0; i < b.N; i++ {
		log.Info("Test simple debug level log.")
    }
}

func BenchmarkLog_Info2Info_WithFields(b *testing.B) {

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "02-01-2006T15:04:05.000Z"})

    for i := 0; i < b.N; i++ {
		log.WithFields(log.Fields {
			"foo1": "bar1",
			"foo2": "bar2",
		}).Info("Test simple debug level log.")
    }
}

var (
	glog = log.New()
)

func init() {
	glog.SetLevel(log.InfoLevel)
	glog.SetFormatter(&log.JSONFormatter{TimestampFormat: "02-01-2006T15:04:05.000Z"})
}

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
