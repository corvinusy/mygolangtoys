package restxorm

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/go-resty/resty"
)

const baseURL = "http://localhost:8090"

func waitReachable(hostport string, maxWait time.Duration) error {
	done := time.Now().Add(maxWait)
	for time.Now().Before(done) {
		c, err := net.Dial("tcp", hostport)
		if err == nil {
			c.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("cannot connect %v for %v", hostport, maxWait)
}

func TestSuite(t *testing.T) {

	// create test server
	server, err := NewServer()
	if err != nil {
		log.Fatal("cannot create server")
	}
	go server.Run()
	waitReachable("localhost:8090", 5*time.Second)

	// create and setup resty client
	rc := resty.New()
	rc.SetHeader("Content-Type", "application/json")
	rc.SetHostURL(baseURL)

	// suite runners
	r := new(Reminder)
	r.tableNameTest(t)
	r.postTest(rc, t)
	r.getAllTest(rc, t)
	r.getTest(rc, t)
	r.putTest(rc, t)
	r.deleteTest(rc, t)
}
