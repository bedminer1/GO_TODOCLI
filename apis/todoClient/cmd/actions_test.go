package cmd

import (
	"bytes"
	"errors"
	"fmt"
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