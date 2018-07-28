package standard

import "testing"

func Test_InSlice(t *testing.T) {
	if StringInSlice("test string", []string{}) {
		t.Fail()
	}
	if !StringInSlice("test string", []string{"test string", "other string"}) {
		t.Fail()
	}
}
