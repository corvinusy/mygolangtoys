package restxorm

import (
	"strconv"
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

func (r *Reminder) tableNameTest(t *testing.T) {
	Convey("test Table Name", t, func() {
		r = new(Reminder)
		result := r.TableName()
		So(result, ShouldEqual, "reminders")
	})
}

func (r *Reminder) postTest(rc *resty.Client, t *testing.T) {
	Convey("POST reminders", t, func() {
		_, err := rc.R().SetResult(r).
			SetBody(`{"message":"bla-bla-bla"}`).
			Post("/reminders")

		So(err, ShouldBeNil)
		So(r.Message, ShouldEqual, "bla-bla-bla")
	})

	Convey("POST reminder with bad payload", t, func() {
		rs, err := rc.R().SetBody(`this is not a json`).
			Post("/reminders")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

}

func (r *Reminder) getAllTest(rc *resty.Client, t *testing.T) {
	Convey("GET all reminders", t, func() {
		var reminders []Reminder
		_, err := rc.R().SetResult(&reminders).Get("/reminders")
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)
	})

}

func (r *Reminder) getTest(rc *resty.Client, t *testing.T) {
	Convey("GET reminder", t, func() {
		reminders, err := r.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		_, err = rc.R().SetResult(r).
			Get("/reminders/" + strconv.Itoa(int(reminders[0].ID)))

		So(err, ShouldBeNil)
		So(r.Message, ShouldNotBeBlank)
	})

	Convey("GET reminder with bad id", t, func() {
		rs, err := rc.R().Get("/reminders/badid")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

	Convey("GET non-existent reminder", t, func() {
		rs, err := rc.R().Get("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "404")
	})
}

func (r *Reminder) putTest(rc *resty.Client, t *testing.T) {
	Convey("PUT reminder", t, func() {
		reminders, err := r.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		_, err = rc.R().SetBody(`{"message":"ololo"}`).
			SetResult(r).
			Put("/reminders/" + strconv.Itoa(int(reminders[0].ID)))

		So(err, ShouldBeNil)
		So(r.Message, ShouldEqual, "ololo")
	})

	Convey("PUT reminder with bad id", t, func() {
		rs, err := rc.R().
			SetBody(`{"message":"ololo"}`).
			Put("/reminders/badid")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

	Convey("PUT non-existent reminder", t, func() {
		rs, err := rc.R().
			SetBody(`{"message":"ololo"}`).
			Put("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "404")
	})

	Convey("PUT reminder with bad payload", t, func() {
		reminders, err := r.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		rs, err := rc.R().
			SetBody(`this is not a json`).
			Put("/reminders/" + strconv.Itoa(int(reminders[0].ID)))

		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

}

func (r *Reminder) deleteTest(rc *resty.Client, t *testing.T) {
	Convey("DELETE reminder", t, func() {
		reminders, err := r.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldNotBeEmpty)

		for i := range reminders {
			_, err := rc.R().Delete("/reminders/" + strconv.Itoa(int(reminders[i].ID)))
			So(err, ShouldBeNil)
		}
		reminders, err = r.extractReminders(rc)
		So(err, ShouldBeNil)
		So(reminders, ShouldBeEmpty)
	})

	Convey("Delete reminder with bad id", t, func() {
		rs, err := rc.R().Delete("/reminders/badid")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

	Convey("Delete non-existent reminder", t, func() {
		rs, err := rc.R().Delete("/reminders/1005009999")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "404")
	})

}

func (r *Reminder) extractReminders(rc *resty.Client) ([]Reminder, error) {
	reminders := []Reminder{}
	_, err := rc.R().SetResult(&reminders).Get("/reminders")

	if err != nil {
		return nil, err
	}
	return reminders, nil
}
