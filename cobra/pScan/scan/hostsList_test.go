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
