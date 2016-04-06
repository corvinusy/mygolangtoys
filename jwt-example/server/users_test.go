package server

import (
	"strconv"
	"testing"

	"github.com/go-resty/resty"
	. "github.com/smartystreets/goconvey/convey"
)

func (u *User) tableNameTest(t *testing.T) {
	Convey("test Table Name", t, func() {
		u = new(User)
		result := u.TableName()
		So(result, ShouldEqual, "users")
	})
}

func (u *User) postTest(client *resty.Client, t *testing.T) {
	Convey("POST users", t, func() {
		_, err := client.R().SetResult(u).SetBody(`{"message":"bla-bla-bla"}`).Post("/users")
		So(err, ShouldBeNil)
		So(u.Message, ShouldEqual, "bla-bla-bla")
	})

	Convey("POST user with bad payload", t, func() {
		rs, err := client.R().SetBody(`this is not a json`).Post("/users")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

}

func (u *User) getAllTest(client *resty.Client, t *testing.T) {
	Convey("GET all users", t, func() {
		var users []User
		_, err := client.R().SetResult(&users).Get("/users")
		So(err, ShouldBeNil)
		So(users, ShouldNotBeEmpty)
	})

}

func (u *User) getTest(client *resty.Client, t *testing.T) {
	Convey("GET /user/{id}", t, func() {
		users, err := u.extractUsers(client)
		So(err, ShouldBeNil)
		So(users, ShouldNotBeEmpty)

		_, err = client.R().SetResult(u).
			Get("/users/" + strconv.Itoa(int(users[0].ID)))

		So(err, ShouldBeNil)
		So(u.Message, ShouldNotBeBlank)
	})

	Convey("GET /user/{id} with bad id", t, func() {
		rs, err := client.R().Get("/users/badid")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

	Convey("GET /user/{id} non-existent id", t, func() {
		rs, err := client.R().Get("/users/1005009999")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "404")
	})
}

func (u *User) putTest(client *resty.Client, t *testing.T) {
	Convey("PUT /user/{id}", t, func() {
		users, err := u.extractUsers(client)
		So(err, ShouldBeNil)
		So(users, ShouldNotBeEmpty)

		_, err = client.R().SetBody(`{"message":"ololo"}`).
			SetResult(u).
			Put("/users/" + strconv.Itoa(int(users[0].ID)))

		So(err, ShouldBeNil)
		So(u.Message, ShouldEqual, "ololo")
	})

	Convey("PUT /user/{id} with bad id", t, func() {
		rs, err := client.R().
			SetBody(`{"message":"ololo"}`).
			Put("/users/badid")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

	Convey("PUT /user/{id} for non-existent id", t, func() {
		rs, err := client.R().
			SetBody(`{"message":"ololo"}`).
			Put("/users/1005009999")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "404")
	})

	Convey("PUT /user/{id} with bad payload", t, func() {
		users, err := u.extractUsers(client)
		So(err, ShouldBeNil)
		So(users, ShouldNotBeEmpty)

		rs, err := client.R().
			SetBody(`this is not a json`).
			Put("/users/" + strconv.Itoa(int(users[0].ID)))

		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

}

func (u *User) deleteTest(client *resty.Client, t *testing.T) {
	Convey("DELETE user", t, func() {
		users, err := u.extractUsers(client)
		So(err, ShouldBeNil)
		So(users, ShouldNotBeEmpty)

		for i := range users {
			_, err = client.R().Delete("/users/" + strconv.Itoa(int(users[i].ID)))
			So(err, ShouldBeNil)
		}
		users, err = u.extractUsers(client)
		So(err, ShouldBeNil)
		So(users, ShouldBeEmpty)
	})

	Convey("Delete user with bad id", t, func() {
		rs, err := client.R().Delete("/users/badid")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "400")
	})

	Convey("Delete non-existent user", t, func() {
		rs, err := client.R().Delete("/users/1005009999")
		So(err, ShouldBeNil)
		So(rs.Status(), ShouldStartWith, "404")
	})

}

func (u *User) extractUsers(client *resty.Client) ([]User, error) {
	var (
		users []User
		err   error
	)
	_, err = client.R().SetResult(&users).Get("/users")
	if err != nil {
		return nil, err
	}
	return users, nil
}
