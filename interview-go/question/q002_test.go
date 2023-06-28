package question

import (
	"fmt"
	"testing"
)

func Test_isUniqueString(t *testing.T) {
	s := "asdfaincasdfjasdfadfashgh"
	fmt.Printf("string for %s is unique string is: %#v", s, isUniqueString(s))
}
