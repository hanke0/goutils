package assert_test

import (
	"errors"
	"testing"

	"github.com/hanke0/goutils/assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, 1, 1)
	assert.NotEqual(t, 1, 2)
}

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
	assert.NotNil(t, 1)
}

func TestWantError(t *testing.T) {
	assert.WantError(t, true, errors.New("123"))
	assert.WantError(t, false, nil)
}

func TestTrue(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
}
