package controller

import "testing"
import (
	"github.com/hyperledger/fabric/consensus"
	"github.com/stretchr/testify/assert"
)

func TestNewConsenter(t *testing.T) {
	var stack consensus.Stack
	son := NewConsenter(stack)
	assert.NotNil(t, son)
}
