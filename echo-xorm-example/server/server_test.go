package server

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/go-resty/resty"
)

type Suite struct {
}

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

	const baseURL = "http://localhost:11999"
	// create test server
	server, err := NewServer()
	if err != nil {
		log.Fatal("cannot create server")
	}
	go server.Run()
	waitReachable("localhost:11999", 5*time.Second)

	// create and setup resty client
	rc := resty.New()
	rc.SetHeader("Content-Type", "application/json")
	rc.SetHostURL(baseURL)

	// suite runners
	s := new(Suite)
	s.tableNameTest(t)
	s.helloTest(rc, t)
	s.postTest(rc, t)
	s.getAllTest(rc, t)
	s.getTest(rc, t)
	s.putTest(rc, t)
	s.deleteTest(rc, t)
}
