package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	setupHandles()
	code := m.Run()
	os.Exit(code)
}

func MockingStoreAnArticle(w http.ResponseWriter, r *http.Request) {

	v := Context{
		map[string]interface{}{"name": strings.Split(r.URL.Path, "/")[2]},
		w,
		r,
	}
	s.router.handlers["PUT"]["/articles/:name"](&v)
}

func TestMockingStoreAnArticle(t *testing.T) {

	req, err := http.NewRequest("PUT", "/articles/One", strings.NewReader("test_data_123"))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	MockingStoreAnArticle(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := "test_data_123"
	if val, ok := w["One"]; ok {
		if val != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", val, expected)
		}
	}
}

func MockingReadAnArticle(w http.ResponseWriter, r *http.Request) {

	v := Context{
		map[string]interface{}{"name": strings.Split(r.URL.Path, "/")[2]},
		w,
		r,
	}
	s.router.handlers["GET"]["/articles/:name"](&v)
}

func TestMockingReadAnArticle(t *testing.T) {

	req, err := http.NewRequest("Get", "/articles/One", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	MockingReadAnArticle(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "test_data_123"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func MockingListArticles(w http.ResponseWriter, r *http.Request) {

	v := Context{
		nil,
		w,
		r,
	}
	s.router.handlers["GET"]["/articles"](&v)
}

func TestMockingListArticles(t *testing.T) {

	req, err := http.NewRequest("PUT", "/articles/Two", strings.NewReader("test_data_456"))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	MockingStoreAnArticle(rr, req)

	req, err = http.NewRequest("GET", "/articles", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()

	MockingListArticles(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected1 := "[\"One\",\"Two\"]\n"
	expected2 := "[\"Two\",\"One\"]\n"
	if (rr.Body.String() != expected1) && (rr.Body.String() != expected2) {
		t.Errorf("handler returned unexpected body: got %v want %v or %v",
			rr.Body.String(), expected1, expected2)
	}
}
