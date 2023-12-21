package git

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	dir := t.TempDir()
}

func TestFoo(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want string
	}{
		{name: "hoge test", lhs: "foo", rhs: "", want: "foo"},
		{name: "hoge test", lhs: "foo", rhs: "bar", want: "foobar"},
	}

	for _, tc := range tests {
		got := Foo(tc.lhs, tc.rhs)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

