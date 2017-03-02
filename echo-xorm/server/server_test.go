package server

import (
	"database/sql"
	"fmt"
	"net"
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"gopkg.in/testfixtures.v2"
)

type Suite struct {
}

const fixturesPath = "testdata"

func waitReachable(hostport string, maxWait time.Duration) error {
	done := time.Now().Add(maxWait)
	for time.Now().Before(done) {
		c, err := net.Dial("tcp", hostport)
		if err == nil {
			return c.Close()
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("cannot connect %v for %v", hostport, maxWait)
}

func TestSuite(t *testing.T) {

	const (
		baseURL = "http://localhost:11110"
		restURL = baseURL + "/rest"
	)

	err := prepareTestDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// create test server
	server, err := NewServer()
	if err != nil {
		log.Fatal("cannot create server")
	}
	go server.Run()
	err = waitReachable(":11110", 3*time.Second)
	if err != nil {
		log.Fatal("server doesn't respond")
	}

	// create and setup resty client
	rc := resty.New()
	rc.SetHeader("Content-Type", "application/json")

	// test runners
	// auth
	rc.SetHostURL(baseURL)
	authSuite := new(authTestSuite)
	authSuite.helloTest(rc, t)
	authSuite.postTest(rc, t)
	// should be called once after auth tests and before other tests
	authSuite.setAuthToken(rc, t)
	// reminders
	rc.SetHostURL(restURL)
	reminderSuite := new(reminderTestSuite)
	reminderSuite.tableNameTest(t)
	reminderSuite.postTest(rc, t)
	reminderSuite.getAllTest(rc, t)
	reminderSuite.getTest(rc, t)
	reminderSuite.putTest(rc, t)
	reminderSuite.deleteTest(rc, t)
	// users
	userSuite := new(userTestSuite)
	userSuite.tableNameTest(t)
	userSuite.postTest(rc, t)
	userSuite.getAllTest(rc, t)
	userSuite.getTest(rc, t)
	userSuite.putTest(rc, t)
	userSuite.deleteTest(rc, t)

}

func prepareTestDatabase() error {
	db, err := sql.Open("sqlite3", "/tmp/echo-xorm-test.sqlite.db")
	if err != nil {
		return err
	}
	err = testfixtures.LoadFixtures(fixturesPath, db, &testfixtures.SQLiteHelper{})
	return err
}
