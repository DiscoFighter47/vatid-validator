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

	testData := []struct {
		des     string
		provide string
		file    string
		panic   string
	}{
		{
			des:     "read config successful",
			provide: "file",
			file:    "../config.example.yml",
		},
		{
			des:   "config provider environment missing",
			panic: "log.fatal called: [environment variable missing: [CONFIG_PROVIDER]]",
		},
		{
			des:     "invalid config provider",
			provide: "invalid",
			panic:   "log.fatal called: [invalid config provider: invalid]",
		},
		{
			des:     "file is missing",
			provide: "file",
			file:    "./config.example.yml",
			panic:   "log.fatal called: [unable to read config: unable to read config file: open ./config.example.yml: no such file or directory]",
		},
		{
			des:     "invalid file",
			provide: "file",
			file:    "../test/config.invalid.yml",
			panic:   "log.fatal called: [unable to unmarshal config: 1 error(s) decoding:\n\n* 'App.Port' expected type 'int64', got unconvertible type '[]interface {}', value: '[8080]']",
		},
		{
			des:     "empty file",
			provide: "file",
			file:    "../test/config.empty.yml",
			panic:   "log.fatal called: [invalid config: App.Port: less than min, Euvies.Timeout: zero value]",
		},
	}

	for _, td := range testData {
		Convey(td.des, t, func() {
			clear()
			os.Setenv("CONFIG_PROVIDER", td.provide)
			os.Setenv("CONFIG_FILE", td.file)
			if td.panic == "" {
				Convey("get config should not panic", func() {
					So(f, ShouldNotPanic)
				})
			} else {
				Convey("get config should panic", func() {
					So(f, ShouldPanicWith, td.panic)
				})
			}
		})
	}
}

func clear() {
	once = sync.Once{}
	cnf = Config{}
	os.Clearenv()
}
