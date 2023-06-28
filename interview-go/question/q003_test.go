package question

import "testing"

func Test_reverString(t *testing.T) {
	s := "avdad"
	rs, ok := reverString(s)
	if ok {
		t.Logf("string %s rever is %s", s, rs)
	}
}
