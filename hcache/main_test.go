package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBlackHandler(t *testing.T) {

	blackHandlerFunc := http.HandlerFunc(blackHandler)
	var etag string

	// first request
	if r, err := http.NewRequest("GET", "", nil); err != nil {
		t.Errorf("%v", err)
	} else {
		recorder := httptest.NewRecorder()
		blackHandlerFunc.ServeHTTP(recorder, r)
		if recorder.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusOK)
		}
		// record etag to test cache
		etag = recorder.Header().Get("Etag")
	}

	// test caching
	if r, err := http.NewRequest("GET", "", nil); err != nil {
		t.Errorf("%v", err)
	} else {
		r.Header.Set("If-None-Match", etag)
		recorder := httptest.NewRecorder()
		blackHandlerFunc.ServeHTTP(recorder, r)
		if recorder.Code != http.StatusNotModified {
			t.Errorf("returned %v. Expected %v.", recorder.Code, http.StatusNotModified)
		}
	}
}
