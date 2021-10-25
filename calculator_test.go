package travis_ci_go_practice

import "testing"

func Test_Sum(t *testing.T){
	want := 11
	subject := Calculator{}
	got := subject.Sum(5,6)
	if got != want {
		t.Errorf("wanted: %d, got: %d",want,got)
	}
}


