package log

import (
	"errors"
	"testing"
)

func TestNewFmtIOLog(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		err      error
		expected interface{}
	}{
		{"Testing With Error", "", errors.New("Error"), nil},
		{"Testing Without Error", "Created a parking lot with 6 slots", nil, nil},
		{"Testing Without Error and Blank Value", "", nil, nil},
	}
	logObj := NewFmtIOLog()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			output := logObj.Log(test.response, test.err)
			if test.expected != output {
				t.Errorf("OutPut is not matched with expected")
			}
		})
	}
}
