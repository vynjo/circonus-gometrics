// Copyright 2016 Circonus, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func callServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, r.Method)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestNew(t *testing.T) {
	t.Log("invalid config [nil]")
	{
		expectedError := errors.New("Invalid API configuration (nil)")
		_, err := New(nil)
		if err == nil {
			t.Error("Expected an error")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected an '%#v' error, got '%#v'", expectedError, err)
		}
	}

	t.Log("invalid config [blank]")
	{
		expectedError := errors.New("API Token is required")
		ac := &Config{}
		_, err := New(ac)
		if err == nil {
			t.Error("Expected an error")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected an '%#v' error, got '%#v'", expectedError, err)
		}
	}

	t.Log("API Token, no API App, no API URL")
	{
		ac := &Config{
			TokenKey: "abc123",
		}
		_, err := New(ac)
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	}

	t.Log("API Token, API App, no API URL")
	{
		ac := &Config{
			TokenKey: "abc123",
			TokenApp: "someapp",
		}
		_, err := New(ac)
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	}

	t.Log("API Token, API App, API URL [host]")
	{
		ac := &Config{
			TokenKey: "abc123",
			TokenApp: "someapp",
			URL:      "something.somewhere.com",
		}
		_, err := New(ac)
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	}

	t.Log("API Token, API App, API URL [trailing '/']")
	{
		ac := &Config{
			TokenKey: "abc123",
			TokenApp: "someapp",
			URL:      "something.somewhere.com/somepath/",
		}
		_, err := New(ac)
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	}

	t.Log("API Token, API App, API URL [w/o trailing '/']")
	{
		ac := &Config{
			TokenKey: "abc123",
			TokenApp: "someapp",
			URL:      "something.somewhere.com/somepath",
		}
		_, err := New(ac)
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	}

	t.Log("API Token, API App, API URL [invalid]")
	{
		expectedError := errors.New("parse http://something.somewhere.com\\somepath$: invalid character \"\\\\\" in host name")
		ac := &Config{
			TokenKey: "abc123",
			TokenApp: "someapp",
			URL:      "http://something.somewhere.com\\somepath$",
		}
		_, err := New(ac)
		if err == nil {
			t.Error("Expected an error")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected an '%#v' error, got '%#v'", expectedError, err)
		}
	}

	t.Log("Debug true, no log.Logger")
	{
		ac := &Config{
			TokenKey: "abc123",
			Debug:    true,
		}
		_, err := New(ac)
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
	}
}

func TestApiCall(t *testing.T) {
	server := callServer()
	defer server.Close()

	ac := &Config{
		TokenKey: "foo",
		TokenApp: "bar",
		URL:      server.URL,
	}

	apih, err := NewAPI(ac)
	if err != nil {
		t.Errorf("Expected no error, got '%+v'", err)
	}

	t.Log("invalid URL path")
	{
		_, err := apih.apiCall("GET", "", nil)
		expectedError := errors.New("Invalid URL path")
		if err == nil {
			t.Errorf("Expected error")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected %+v go '%+v'", expectedError, err)
		}
	}

	t.Log("URL path fixup, prefix '/'")
	{
		call := "GET"
		resp, err := apih.apiCall(call, "nothing", nil)
		if err != nil {
			t.Errorf("Expected no error, got '%+v'", resp)
		}
		expected := fmt.Sprintf("%s\n", call)
		if string(resp) != expected {
			t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
		}
	}

	t.Log("URL path fixup, remove '/v2' prefix")
	{
		call := "GET"
		resp, err := apih.apiCall(call, "/v2/nothing", nil)
		if err != nil {
			t.Errorf("Expected no error, got '%+v'", resp)
		}
		expected := fmt.Sprintf("%s\n", call)
		if string(resp) != expected {
			t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
		}
	}

	calls := []string{"GET", "PUT", "POST", "DELETE"}
	for _, call := range calls {
		t.Logf("Testing %s call", call)
		resp, err := apih.apiCall(call, "/", nil)
		if err != nil {
			t.Errorf("Expected no error, got '%+v'", resp)
		}

		expected := fmt.Sprintf("%s\n", call)
		if string(resp) != expected {
			t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
		}
	}

}

func TestApiGet(t *testing.T) {
	server := callServer()
	defer server.Close()

	ac := &Config{
		TokenKey: "foo",
		TokenApp: "bar",
		URL:      server.URL,
	}

	client, err := NewClient(ac)
	if err != nil {
		t.Errorf("Expected no error, got '%+v'", err)
	}

	resp, err := client.Get("/")

	if err != nil {
		t.Errorf("Expected no error, got '%+v'", resp)
	}

	expected := "GET\n"
	if string(resp) != expected {
		t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
	}

}

func TestApiPut(t *testing.T) {
	server := callServer()
	defer server.Close()

	ac := &Config{
		TokenKey: "foo",
		TokenApp: "bar",
		URL:      server.URL,
	}

	client, err := NewClient(ac)
	if err != nil {
		t.Errorf("Expected no error, got '%+v'", err)
	}

	resp, err := client.Put("/", nil)

	if err != nil {
		t.Errorf("Expected no error, got '%+v'", resp)
	}

	expected := "PUT\n"
	if string(resp) != expected {
		t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
	}

}

func TestApiPost(t *testing.T) {
	server := callServer()
	defer server.Close()

	ac := &Config{
		TokenKey: "foo",
		TokenApp: "bar",
		URL:      server.URL,
	}

	client, err := NewClient(ac)
	if err != nil {
		t.Errorf("Expected no error, got '%+v'", err)
	}

	resp, err := client.Post("/", nil)

	if err != nil {
		t.Errorf("Expected no error, got '%+v'", resp)
	}

	expected := "POST\n"
	if string(resp) != expected {
		t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
	}

}

func TestApiDelete(t *testing.T) {
	server := callServer()
	defer server.Close()

	ac := &Config{
		TokenKey: "foo",
		TokenApp: "bar",
		URL:      server.URL,
	}

	client, err := NewClient(ac)
	if err != nil {
		t.Errorf("Expected no error, got '%+v'", err)
	}

	resp, err := client.Delete("/")

	if err != nil {
		t.Errorf("Expected no error, got '%+v'", resp)
	}

	expected := "DELETE\n"
	if string(resp) != expected {
		t.Errorf("Expected\n'%s'\ngot\n'%s'\n", expected, resp)
	}

}
