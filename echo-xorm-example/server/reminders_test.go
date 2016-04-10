package server

import (
	"strconv"
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

func (s *Suite) tableNameTest(t *testing.T) {
	Convey("test Table Name", t, func() {
		var r Reminder
		result := r.TableName()
		So(result, ShouldEqual, "reminders")
	})
}

func (s *Suite) helloTest(rc *resty.Client, t *testing.T) {
	Convey("POST reminders", t, func() {
		response, err := rc.R().Get("/")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "200")
	})
}

func (s *Suite) postTest(rc *resty.Client, t *testing.T) {
	Convey("POST reminders", t, func() {
		var result Reminder
		_, err := rc.R().SetResult(&result).
			SetBody(`{"message":"bla-bla-bla"}`).
			Post("/reminders")

		So(err, ShouldBeNil)
		So(result.Message, ShouldEqual, "bla-bla-bla")
	})

	Convey("POST reminder with bad payload", t, func() {
		response, err := rc.R().SetBody(`this is not a json`).
			Post("/reminders")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})
}

func (s *Suite) getAllTest(rc *resty.Client, t *testing.T) {
	Convey("GET all reminders", t, func() {
		var reminders []Reminder
		_, err := rc.R().SetResult(&reminders).Get("/reminders")
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)
	})
}

func (s *Suite) getTest(rc *resty.Client, t *testing.T) {
	Convey("GET reminder", t, func() {
		var result Reminder
		reminders, err := s.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		_, err = rc.R().SetResult(&result).
			Get("/reminders/" + strconv.Itoa(int(reminders[0].ID)))

		So(err, ShouldBeNil)
		So(result.Message, ShouldNotBeBlank)
	})

	Convey("GET reminder with bad id", t, func() {
		response, err := rc.R().Get("/reminders/badid")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

	Convey("GET non-existent reminder", t, func() {
		response, err := rc.R().Get("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "404")
	})
}

func (s *Suite) putTest(rc *resty.Client, t *testing.T) {
	Convey("PUT reminder", t, func() {
		var result Reminder
		reminders, err := s.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		_, err = rc.R().SetBody(`{"message":"ololo"}`).
			SetResult(&result).
			Put("/reminders/" + strconv.Itoa(int(reminders[0].ID)))

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
		reminders, err := s.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		response, err := rc.R().
			SetBody(`this is not a json`).
			Put("/reminders/" + strconv.Itoa(int(reminders[0].ID)))

		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

}

func (s *Suite) deleteTest(rc *resty.Client, t *testing.T) {
	Convey("DELETE reminder", t, func() {
		reminders, err := s.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		for i := range reminders {
			_, err = rc.R().Delete("/reminders/" + strconv.Itoa(int(reminders[i].ID)))
			So(err, ShouldBeNil)
		}
		reminders, err = s.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldBeEmpty)
	})

	Convey("Delete reminder with bad id", t, func() {
		response, err := rc.R().Delete("/reminders/badid")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "400")
	})

	Convey("Delete non-existent reminder", t, func() {
		response, err := rc.R().Delete("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(response.Status(), ShouldStartWith, "204")
	})

}

func (s *Suite) extractReminders(rc *resty.Client) ([]Reminder, error) {
	reminders := []Reminder{}
	_, err := rc.R().SetResult(&reminders).Get("/reminders")

	if err != nil {
		return nil, err
	}
	return reminders, nil
}
