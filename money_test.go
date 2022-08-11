package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMoney(t *testing.T) {
	t.Run("check if the money object is created", func(t *testing.T) {
		assert.IsType(t, Money{}, NewMoney(20, 50))
	})
	t.Run("check for panic when paise is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			NewMoney(100, -20)
		})
	})
	t.Run("check for panic when rupees is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			NewMoney(-100, 20)
		})
	})
	t.Run("check for no panic when paise is non-negative", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewMoney(100, 20)
		})
	})
}

func TestAddMoney(t *testing.T) {
	t.Run("check if the sum is not negative", func(t *testing.T) {
		actual := NewMoney(0, 10)
		actual.AddMoney(NewMoney(100, 10))
		assert.Less(t, 0, actual.paise)
	})
	t.Run("check the sum is correct if otherMoney paise is 0", func(t *testing.T) {
		actual := NewMoney(0, 10)
		actual.AddMoney(NewMoney(0, 0))
		assert.Equal(t, 10, actual.paise)
	})
	t.Run("check the sum is correct if paise is 0", func(t *testing.T) {
		actual := NewMoney(0, 0)
		actual.AddMoney(NewMoney(10, 10))
		assert.Equal(t, 1010, actual.paise)
	})
	t.Run("check if the sum is correct", func(t *testing.T) {
		actual := NewMoney(100, 10)
		actual.AddMoney(NewMoney(10, 100))
		assert.Equal(t, NewMoney(110, 110).paise, actual.paise)
	})
	t.Run("check if the sum is correct when both objects are swapped", func(t *testing.T) {
		actual := NewMoney(100, 10)
		actual.AddMoney(NewMoney(10, 100))
		expected := NewMoney(10, 100)
		expected.AddMoney(NewMoney(100, 10))
		assert.Equal(t, expected.paise, actual.paise)
	})
}

func TestSubtractMoney(t *testing.T) {
	t.Run("check if the difference is not negative", func(t *testing.T) {
		actual := NewMoney(100, 10)
		actual.SubtractMoney(NewMoney(5, 12))
		assert.Less(t, 0, actual.paise)
	})
	t.Run("check the difference when the money to be subtracted is 0", func(t *testing.T) {
		actual := NewMoney(10, 22)
		actual.SubtractMoney(NewMoney(0, 0))
		assert.Equal(t, 1022, actual.paise)
	})
	t.Run("check for panic when the money is 0", func(t *testing.T) {
		actual := NewMoney(0, 0)
		assert.Panics(t, func() {
			actual.SubtractMoney(NewMoney(10, 0))
		})
	})
	t.Run("check for panic when paise drops below 0", func(t *testing.T) {
		actual := NewMoney(100, 10)
		assert.Panics(t, func() {
			actual.SubtractMoney(NewMoney(130, 14))
		})
	})
	t.Run("check if the difference is correct", func(t *testing.T) {
		actual := NewMoney(10, 22)
		actual.SubtractMoney(NewMoney(5, 10))
		assert.Equal(t, NewMoney(5, 12).paise, actual.paise)
	})
}
