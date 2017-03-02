package server

import (
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

type userTestSuite struct{}

func (s *userTestSuite) tableNameTest(t *testing.T) {
	Convey("test Table Name", t, func() {
		var r User
		result := r.TableName()
		So(result, ShouldEqual, "users")
	})
}

func (s *userTestSuite) postTest(rc *resty.Client, t *testing.T) {
	Convey("POST users", t, func() {
		var result User
		resp, err := rc.R().SetResult(&result).
			SetBody(`{"login":"test_login01", "password":"test_password01"}`).
			Post("/users")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "201")
		So(result.Login, ShouldEqual, "test_login01")
	})

	Convey("POST user with bad payload", t, func() {
		resp, err := rc.R().SetBody(`this is not a json`).
			Post("/users")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "400")
	})
}

func (s *userTestSuite) getAllTest(rc *resty.Client, t *testing.T) {
	Convey("GET all users", t, func() {
		var users []User
		resp, err := rc.R().SetResult(&users).Get("/users")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(len(users), ShouldBeGreaterThanOrEqualTo, 3)
	})
}

func (s *userTestSuite) getTest(rc *resty.Client, t *testing.T) {
	Convey("GET user", t, func() {
		var result User

		resp, err := rc.R().SetResult(&result).
			Get("/users/2")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(result.Login, ShouldNotBeBlank)
	})

	Convey("GET user with bad id", t, func() {
		response, err := rc.R().Get("/users/badid")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

	Convey("GET non-existent user", t, func() {
		response, err := rc.R().Get("/users/1005009999")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "204")
	})
}

func (s *userTestSuite) putTest(rc *resty.Client, t *testing.T) {
	Convey("PUT user", t, func() {
		var result User
		resp, err := rc.R().SetBody(`{"login":"test_login_update01"}`).
			SetResult(&result).
			Put("/users/2")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(result.Login, ShouldEqual, "test_login_update01")
	})

	Convey("PUT user with bad id", t, func() {
		response, err := rc.R().
			SetBody(`{"message":"ololo"}`).
			Put("/users/bad_id")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

	Convey("PUT non-existent user", t, func() {
		response, err := rc.R().
			SetBody(`{"message":"ololo"}`).
			Put("/users/1005009999")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "204")
	})

	Convey("PUT user with bad payload", t, func() {
		resp, err := rc.R().
			SetBody(`this is not a json`).
			Put("/users/2")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "400")
	})

}

func (s *userTestSuite) deleteTest(rc *resty.Client, t *testing.T) {
	Convey("DELETE user", t, func() {
		resp, err := rc.R().Delete("/users/3")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "204")

		_, err = rc.R().Get("/users/3")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "204")
	})

	Convey("Delete user with bad id", t, func() {
		resp, err := rc.R().Delete("/users/badid")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "400")
	})

	Convey("Delete non-existent user", t, func() {
		resp, err := rc.R().Delete("/users/1005009999")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "204")
	})
}
