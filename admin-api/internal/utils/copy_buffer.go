package utils

import (
	"bytes"
	"io"
	"strings"
)

type CopyBuffer struct {
	buf *bytes.Buffer
	err error
}

func NewCopyBuffer(b *bytes.Buffer) *CopyBuffer {
	return &CopyBuffer{buf: b}
}

func (cb *CopyBuffer) Next() bool {
	return cb.buf.Len() > 0
}

func (cb *CopyBuffer) Values() ([]interface{}, error) {
	line, err := cb.buf.ReadString('\n')
	if err != nil && err != io.EOF {
		cb.err = err
		return nil, err
	}
	line = strings.TrimSuffix(line, "\n")

	values := strings.Split(line, "\t")
	returnValues := make([]interface{}, len(values))
	for i, v := range values {
		returnValues[i] = v
	}

	return returnValues, nil
}

func (cb *CopyBuffer) Err() error {
	return cb.err
}
