package log

import (
	"fmt"
)

type fmtIO struct{}

func NewFmtIOLog() *fmtIO {
	return &fmtIO{}
}

func (s *fmtIO) Log(log string, err error) error {
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(log)
	return nil
}
