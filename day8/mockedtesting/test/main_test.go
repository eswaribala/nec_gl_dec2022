package test

import (
	assert2 "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"necdec2022/day8/mockedtesting/src"
	"testing"
)

type mockOTP struct{ mock.Mock }

func newMockOTP() *mockOTP { return &mockOTP{} }

var m mockOTP

func (m *mockOTP) GetOTP(seed int) int {
	args := m.Called(seed)
	return args.Int(0)
}
func TestGenearteOTP(t *testing.T) {

	assert := assert2.New(t)
	// get our mockRand
	m := newMockOTP()

	m.On("GetOTP", 10000).Return(int(rand.Int31n(10)))

	quotient := src.GenerateOTP(10000, m)

	// check that GetOTP was called with the number 10000;
	// if not then the test fails
	m.AssertCalled(t, "GetOTP", 10000)
	assert.True(quotient > 0)
}
