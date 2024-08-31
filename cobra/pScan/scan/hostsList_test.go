package scan_test

import (
	"errors"
	"testing"

	"github.com/bedminer1/cobra/pScan/scan"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{
			name: "AddNew",
			host: "host2",
			expectLen: 2,
			expectErr: nil,
		},
		{
			name: "AddExisting",
			host: "host1",
			expectLen: 1,
			expectErr: scan.ErrExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// init hl with 'host1'
			hl := &scan.HostsList{}
			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}

			// test tc
			err := hl.Add(tc.host)
			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("Expected error %q, got %q instead\n ", tc.expectErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %q\n", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}
			if hl.Hosts[1] != tc.host {
				t.Errorf("Expected host %q, got %q instead\n", tc.host, hl.Hosts[1])
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{
			name: "RemoveExisting",
			host: "host1",
			expectLen: 1,
			expectErr: nil,
		},
		{
			name: "RemoveNotFound",
			host: "host3",
			expectLen: 1,
			expectErr: scan.ErrNotExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &scan.HostsList{}
			for _, h := range []string{"host1", "host2"} {
				if err := hl.Add(h); err != nil {
					t.Fatal(err)
				}
			}
			err := hl.Remove(tc.host)
			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil\n")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("Expected error %q, got %q", tc.expectErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %q", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}
			if hl.Hosts[0] == tc.host {
				t.Errorf("Expected host %q to be removed\n", tc.host)
			}

		})
	}
}