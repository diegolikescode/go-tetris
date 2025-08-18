package tetris

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tee(t *testing.T) {
	tee := Tee(10)

	assert.Equal(t, "", tee)
}
