package stringutil

import "testing"

func TestReverse(t *testing.T){
	cases := []struct {
		in, want string
	}{
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"a", "a"},
		{"中文", "文中"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want{
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
