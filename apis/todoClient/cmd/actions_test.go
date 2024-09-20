package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestListAction(t *testing.T) {
	testCases := []struct {
		name     string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		closeServer bool
	}{
		{
			name:     "Results",
			expError: nil,
			expOut:   "-  1  Task 1\n-  2  Task 2\n",
			resp:     testResp["resultsMany"],
		},
		{
			name:     "NoResults",
			expError: ErrNotFound,
			resp:     testResp["noResults"],
		},
		{
			name:        "NoConnection",
			expError:    ErrConnection,
			resp:        testResp["noResults"],
			closeServer: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.resp.Status)
				fmt.Fprintln(w, tc.resp.Body)
			})
			defer cleanup()
			if tc.closeServer {
				cleanup()
			}

			var out bytes.Buffer
			err := listAction(&out, url)

			if tc.expError != nil {
				if err == nil {
					t.Fatal("expected an error")
				}

				if !errors.Is(err, tc.expError) {
					t.Error("wrong error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %q", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("expected %s, got %s", tc.expOut, out.String())
			}
		})
	}
}

func TestViewAction(t *testing.T) {
	testCases := []struct {
		name     string
		expError error
		expOut   string
		resp     struct {
			Status int
			Body   string
		}
		id string
	}{
		{
			name:     "ResultsOne",
			expError: nil,
			expOut:   "Task:         Task 1\nCreated at:   Oct/28 @08:28\nCompleted:    No\n",
			resp:     testResp["resultsOne"],
			id: "1",
		},
		{
			name:     "NotFound",
			expError: ErrNotFound,
			resp:     testResp["notFound"],
			id: "1",
		},
		{
			name:        "InvalidID",
			expError:    ErrNotNumber,
			resp:        testResp["noResults"],
			id: "a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.resp.Status)
				fmt.Fprintln(w, tc.resp.Body)
			})
			defer cleanup()

			var out bytes.Buffer

			err := viewAction(&out, url, tc.id)
			if tc.expError != nil {
				if err == nil {
					t.Fatal("expected error, got no error")
				}

				if !errors.Is(err, tc.expError) {
					t.Error("wrong error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %q", err)
			}

			if tc.expOut != out.String() {
				t.Errorf("wrong output, expected %q, got %q", tc.expOut, out.String())
			}
		})
	}
}

func TestAddAction(t *testing.T) {
	expURLPath := "/todo"
	expMethod := http.MethodPost
	expBody := "{\"task\":\"Task 1\"}\n"
	expContentType := "application/json"
	expOut := "Added task \"Task 1\" to the list.\n"
	args := []string{"Task", "1"}

	url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expURLPath {
			t.Errorf("expected path %q, got %q", expURLPath, r.URL.Path)
		}
		if r.Method != expMethod {
			t.Error("expect POST method")
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		r.Body.Close()

		if string(body) != expBody {
			t.Errorf("Expected body %q, got %q", expBody, string(body))
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != expContentType {
			t.Error("expected content-type to be application/json")
		}
		w.WriteHeader(testResp["created"].Status)
		fmt.Fprintln(w, testResp["created"].Body)
	})
	defer cleanup()

	var out bytes.Buffer
	if err := addAction(&out, url, args); err != nil {
		t.Fatalf("Unexpected error: %q", err)
	}

	if expOut != out.String() {
		t.Error("Unexpected output")
	}
}

func TestCompleteAction(t *testing.T) {
	expURLPath := "/todo/1"
	expMethod := http.MethodPatch
	expQuery := "complete"
	expOut := "Item number 1 marked as completed\n"
	arg := "1"

	url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expURLPath {
			t.Error("unexpected path")
		}

		if r.Method != expMethod {
			t.Error("unexpected method")
		}

		if _, ok := r.URL.Query()[expQuery]; !ok {
			t.Error("expected query 'complete'")
		}

		w.WriteHeader(testResp["noContent"].Status)
		fmt.Fprintln(w, testResp["noContent"].Body)
	})
	defer cleanup()

	var out bytes.Buffer

	if err := completeAction(&out, url, arg); err != nil {
		t.Fatal("unexpected error")
	}

	if expOut != out.String() {
		t.Errorf("unexpected output: %q\n expected: %q", out.String(), expOut)
	}
}