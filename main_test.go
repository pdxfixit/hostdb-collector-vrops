package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	loadConfig()

	os.Exit(m.Run())

}

func TestHttpRequest(t *testing.T) {

	testString := "Hello World."

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, testString)
		if err != nil {
			t.Errorf("%v", err)
		}
	}))
	defer ts.Close()

	header := map[string]string{
		"test": "yes",
	}

	responseBytes, err := httpRequest("GET", ts.URL, nil, header)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, responseBytes, []byte(testString))

}

func TestRequestToStruct(t *testing.T) {

	// setup fake http server for test
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "{\"test\":true}")
		if err != nil {
			t.Error(err.Error())
		}
	}))
	defer ts.Close()

	testStruct := struct {
		Test bool `json:"test"`
	}{}

	if err := requestToStruct(ts.URL, &testStruct); err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, testStruct.Test, true)

}
