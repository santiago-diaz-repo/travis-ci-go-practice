package travis_ci_go_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sum(t *testing.T) {
	want := 11
	subject := Calculator{}
	got := subject.Sum(5, 6)

	assert.Equal(t, want, got, "they should not be equal")
}

func Test_Sustract(t *testing.T) {
	want := -1
	subject := Calculator{}
	got := subject.Subtract(5, 6)

	assert.Equal(t, want, got, "they should not be equal")
}
