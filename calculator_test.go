package main1

import "testing"

func Test_Add(t *testing.T) {
	a, b, c := 2, 2, 4

	result := add(a, b)
	if result != c {
		t.Fail()
	} else {
		t.Logf("success")
	}
}

//func TestSubtract(t *testing.T) {
//	t.Parallel()
//	var want int = 2
//	got := Subtract(4, 2)
//	if want != got {
//		t.Fail()
//	} else {
//		t.Logf("success")
//	}
//}
