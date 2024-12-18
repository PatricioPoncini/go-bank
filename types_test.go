package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	_, err := NewAccount("john", "doe", "fake_password")
	assert.Nil(t, err)
}
