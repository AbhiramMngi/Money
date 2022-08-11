package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMoney(t *testing.T) {
	t.Run("check if the Money object is created", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(20))
	})
}
