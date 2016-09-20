package server

import (
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

type reminderTestSuite struct{}

func (s *reminderTestSuite) tableNameTest(t *testing.T) {
	Convey("test Table Name", t, func() {
		var r Reminder
		result := r.TableName()
		So(result, ShouldEqual, "reminders")
	})
}

func (s *reminderTestSuite) postTest(rc *resty.Client, t *testing.T) {
	Convey("POST reminders", t, func() {
		var result Reminder
		resp, err := rc.R().SetResult(&result).
			SetBody(`{"message":"bla-bla-bla"}`).
			Post("/reminders")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "201")
		So(result.Message, ShouldEqual, "bla-bla-bla")
	})

	Convey("POST reminder with bad payload", t, func() {
		resp, err := rc.R().SetBody(`this is not a json`).
			Post("/reminders")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "400")
	})
}

func (s *reminderTestSuite) getAllTest(rc *resty.Client, t *testing.T) {
	Convey("GET all reminders", t, func() {
		var reminders []Reminder
		resp, err := rc.R().SetResult(&reminders).Get("/reminders")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(reminders, ShouldNotBeEmpty)
	})
}

func (s *reminderTestSuite) getTest(rc *resty.Client, t *testing.T) {
	Convey("GET reminder", t, func() {
		var result Reminder
		resp, err := rc.R().SetResult(&result).
			Get("/reminders/2")

		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "200")
		So(result.Message, ShouldEqual, "This is message 2!")
	})

	Convey("GET reminder with bad id", t, func() {
		response, err := rc.R().Get("/reminders/badid")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

	Convey("GET non-existent reminder", t, func() {
		response, err := rc.R().Get("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "204")
	})
}

func (s *reminderTestSuite) putTest(rc *resty.Client, t *testing.T) {
	Convey("PUT reminder", t, func() {
		var result Reminder

		_, err := rc.R().SetBody(`{"message":"ololo"}`).
			SetResult(&result).
			Put("/reminders/3")

		So(err, ShouldBeNil)
		So(result.Message, ShouldEqual, "ololo")
	})

	Convey("PUT reminder with bad id", t, func() {
		response, err := rc.R().
			SetBody(`{"message":"ololo"}`).
			Put("/reminders/bad_id")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

	Convey("PUT non-existent reminder", t, func() {
		response, err := rc.R().
			SetBody(`{"message":"ololo"}`).
			Put("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "204")
	})

	Convey("PUT reminder with bad payload", t, func() {
		response, err := rc.R().
			SetBody(`this is not a json`).
			Put("/reminders/1")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

}

func (s *reminderTestSuite) deleteTest(rc *resty.Client, t *testing.T) {
	Convey("DELETE reminder", t, func() {
		resp, err := rc.R().Delete("/reminders/3")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "204")

		_, err = rc.R().Get("/reminders/3")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "204")
	})

	Convey("Delete reminder with bad id", t, func() {
		resp, err := rc.R().Delete("/reminders/badid")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "400")
	})

	Convey("Delete non-existent reminder", t, func() {
		resp, err := rc.R().Delete("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(resp.Status(), ShouldStartWith, "204")
	})
}
