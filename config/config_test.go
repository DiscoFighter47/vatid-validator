package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"bou.ke/monkey"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetConfig(t *testing.T) {
	logFatalPatch := func(v ...interface{}) {
		panic(fmt.Sprintf("log.fatal called: %v", v))
	}
	patch := monkey.Patch(log.Fatalln, logFatalPatch)
	defer patch.Unpatch()

	f := func() {
		GetConfig()
	}

	Convey("read config successfully", t, func() {
		clear()
		os.Setenv("config_provider", "file")
		os.Setenv("config_file", "../config.example.yml")
		Convey("get config should not panic", func() {
			So(f, ShouldNotPanic)
		})
	})

	Convey("config provider environment missing", t, func() {
		clear()
		Convey("get config should panic", func() {
			So(f, ShouldPanicWith, "log.fatal called: [environment variable missing: [config_provider]]")
		})
	})

	Convey("invalid config provider", t, func() {
		clear()
		os.Setenv("config_provider", "invalid")
		Convey("get config should panic", func() {
			So(f, ShouldPanicWith, "log.fatal called: [invalid config provider: invalid]")
		})
	})

	Convey("file is missing", t, func() {
		clear()
		os.Setenv("config_provider", "file")
		os.Setenv("config_file", "./config.example.yml")
		Convey("get config should panic", func() {
			So(f, ShouldPanicWith, "log.fatal called: [unable to read config: unable to read config file: open ./config.example.yml: no such file or directory]")
		})
	})

	Convey("invalid file", t, func() {
		clear()
		os.Setenv("config_provider", "file")
		os.Setenv("config_file", "../test/config.invalid.yml")
		Convey("get config should panic", func() {
			So(f, ShouldPanicWith, "log.fatal called: [unable to unmarshal config: 1 error(s) decoding:\n\n* 'App.Port' expected type 'int64', got unconvertible type '[]interface {}', value: '[8080]']")
		})
	})

	Convey("empty file", t, func() {
		clear()
		os.Setenv("config_provider", "file")
		os.Setenv("config_file", "../test/config.empty.yml")
		Convey("get config should panic", func() {
			So(f, ShouldPanicWith, "log.fatal called: [invalid config: App.Port: less than min, Euvies.Timeout: zero value]")
		})
	})
}

func clear() {
	once = sync.Once{}
	cnf = Config{}
	os.Clearenv()
}
