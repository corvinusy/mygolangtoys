package server

import (
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

type authTestSuite struct{}

func (s *authTestSuite) helloTest(rc *resty.Client, t *testing.T) {
	Convey("GET /", t, func() {
		resp, err := rc.R().Get("/")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
	})
}

func (s *authTestSuite) versionTest(rc *resty.Client, t *testing.T) {
	Convey("GET /", t, func() {
		var result versionResponse
		resp, err := rc.R().SetResult(&result).Get("/version")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(result.ServerTime, ShouldNotEqual, 0)
		So(result.Version, ShouldNotBeBlank)
	})
}

func (s *authTestSuite) postTest(rc *resty.Client, t *testing.T) {
	Convey("GET /auth", t, func() {
		var result authResponse
		resp, err := rc.R().
			SetBody(`{"login":"admin", "password":"admin"}`).
			SetResult(&result).
			Post("/auth")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(result.Result, ShouldEqual, "OK")
		So(result.Token, ShouldNotBeBlank)
	})
}

func (s *authTestSuite) setAuthToken(rc *resty.Client, t *testing.T) {
	Convey("GET /auth", t, func() {
		var result authResponse
		_, err := rc.R().
			SetBody(`{"login":"admin", "password":"admin"}`).
			SetResult(&result).
			Post("/auth")

		So(err, ShouldBeNil)
		rc.SetAuthToken(result.Token)
	})
}
