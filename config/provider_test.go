package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileProvider(t *testing.T) {
	Convey("successful config read", t, func() {
		clear()
		os.Setenv("CONFIG_FILE", "../config.example.yml")
		v := viper.New()
		err := fileProvider(v)
		Convey("error should be nil", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("environment variable is missing", t, func() {
		clear()
		v := viper.New()
		err := fileProvider(v)
		Convey("error should not be nil", func() {
			So(err.Error(), ShouldEqual, "unable to process environment variable: env: required environment variable \"CONFIG_FILE\" is not set")
		})
	})

	Convey("file is missing", t, func() {
		clear()
		os.Setenv("CONFIG_FILE", "./config.example.yml")
		v := viper.New()
		err := fileProvider(v)
		Convey("error should not be nil", func() {
			So(err.Error(), ShouldEqual, "unable to read config file: open ./config.example.yml: no such file or directory")
		})
	})
}
