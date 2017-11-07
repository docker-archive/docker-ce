package testutil

import (
	"io"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ErrorContains checks that the error is not nil, and contains the expected
// substring.
func ErrorContains(t require.TestingT, err error, expectedError string, msgAndArgs ...interface{}) {
	require.Error(t, err, msgAndArgs...)
	assert.Contains(t, err.Error(), expectedError, msgAndArgs...)
}

// DevZero acts like /dev/zero but in an OS-independent fashion.
var DevZero io.Reader = devZero{}

type devZero struct{}

func (d devZero) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = '\x00'
	}
	return len(p), nil
}
