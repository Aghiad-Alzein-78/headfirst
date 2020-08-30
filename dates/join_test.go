package testo4

import "testing"

// func TestTwoElements(t *testing.T) {
// 	list := []string{"apple", "orange"}
// 	if JoinWithCommas(list) != "apple and orange" {
// 		t.Error("The test failed")
// 	}
// }

// func TestThreeElements(t *testing.T) {
// 	list := []string{"apple", "orange", "banana"}
// 	if JoinWithCommas(list) != "apple,orange and banana" {
// 		t.Error("The test failed")
// 	}
// }

// func TestOneElement(t *testing.T) {
// 	list := []string{"apple"}
// 	if JoinWithCommas(list) != "apple" {
// 		t.Error("Wrong Output")
// 	}
// }

type testData struct {
	list []string
	want string
}

func TestData(t *testing.T) {
	testdata := []testData{
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "banana"}, want: "apple,orange and banana"},
	}
}
