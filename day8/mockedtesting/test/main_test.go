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
	// specify our return value. Since the code in divByRand
	// passes 10 into randomInt, we pass 10 in as the argument
	// to go with randomInt, and specify that we want the
	// method to return 6.
	m.On("GetOTP", 10000).Return(rand.Int31n(10))

	// now run divByRand and assert that we got back the
	// return value we expected, just like in a Go test that
	// doesn't use Testify Mock.
	quotient := src.GenerateOTP(10000, m)

	// check that randomInt was called with the number 10;
	// if not then the test fails
	m.AssertCalled(t, "GetOTP", 10000)
	assert.True(quotient > 0)
}
