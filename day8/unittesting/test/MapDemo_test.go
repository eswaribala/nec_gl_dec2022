package test

import (
	assert2 "github.com/stretchr/testify/assert"
	"necdec2022/day8/unittesting/src"
	"testing"
)

func TestCheckUserExistOrNot(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(true, src.CheckUserId(int32(45)))

}
