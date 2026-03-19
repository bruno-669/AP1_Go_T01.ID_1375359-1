package main

import (
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		testname string
		input    string
		want     string
		wanErr   bool
		errMsg   string
	}{
		{
			testname: "base test",
			input:    "1 2 3 4 5\n2 7 9 3 4 9 0\n",
			want:     "2 3 4",
			wanErr:   false,
		},
		{
			testname: "base test only one",
			input:    "1 1 1 1 1\n1 2\n",
			want:     "1",
			wanErr:   false,
		},
		{
			testname: "Invalid input slices one test",
			input:    "1 2 3 t 5\n2 7 9 3 4 9 0\n",
			want:     "",
			wanErr:   true,
			errMsg:   "invalid input",
		},
		{
			testname: "Invalid input slices two test",
			input:    "1 2 3 5\n2 7 9 3 4 s 0\n",
			want:     "",
			errMsg:   "invalid input",
			wanErr:   true,
		},
		{
			testname: "Empty input slices one test",
			input:    "\n2 7 9 3 4 9 0\n",
			want:     "",
			wanErr:   true,
			errMsg:   "Empty input",
		},
		{
			testname: "Empty input slices one test",
			input:    "1 2 3 4\n\n",
			want:     "",
			wanErr:   true,
			errMsg:   "Empty input",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.testname, func(t *testing.T) {
			out, err := SlicesHandler(strings.NewReader(testcase.input))
			if (err != nil) != testcase.wanErr {
				t.Errorf("SlicesHandler() error = %v, wantErr %v", err, testcase.wanErr)
			}
			if testcase.wanErr && testcase.errMsg != "" && !strings.Contains(err.Error(), testcase.errMsg) {
				t.Errorf("SlicesHandler() error = %q, want containing %q", err.Error(), testcase.errMsg)
			}
			if out != testcase.want {
				t.Errorf("SlicesHandler() = %q, want %q", out, testcase.want)
			}
		})
	}

}
