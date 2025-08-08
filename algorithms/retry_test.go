package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetryAlwaysFails(t *testing.T) {
	var counter int
	err := Retry(func() error {
		counter++
		return fmt.Errorf("some error")
	}, 2)

	assert.Error(t, err, "expected error")
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, counter, 2)
}
