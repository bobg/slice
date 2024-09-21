package slice

import "testing"

func TestSubslice(t *testing.T) {
	var (
		s   = From("a", "b", "c", "d", "e")
		got = s.Subslice(1, 3)
	)
	if l := got.Len(); l != 2 {
		t.Errorf("got %d, want 2", l)
	}
	if v := got.At(0); v != "b" {
		t.Errorf(`got %q, want "b"`, v)
	}
	if v := got.At(1); v != "c" {
		t.Errorf(`got %q, want "c"`, v)
	}
}

func TestCopy(t *testing.T) {
	var (
		from = From("a", "b", "c", "d", "e")
		to   = Make[string](3, 3)
		n    = from.Copy(to)
	)
	if n != 3 {
		t.Errorf("got %d, want 3", n)
	}
	if l := to.Len(); l != 3 {
		t.Errorf("got %d, want 3", l)
	}
	if v := to.At(0); v != "a" {
		t.Errorf(`got %q, want "a"`, v)
	}
	if v := to.At(1); v != "b" {
		t.Errorf(`got %q, want "b"`, v)
	}
	if v := to.At(2); v != "c" {
		t.Errorf(`got %q, want "c"`, v)
	}
}

func TestAppend(t *testing.T) {
	var s *Slice[string]
	if l := s.Len(); l != 0 {
		t.Errorf("got %d, want 0", l)
	}
	s = s.Append("a", "b", "c")
	if l := s.Len(); l != 3 {
		t.Errorf("got %d, want 3", l)
	}
	if v := s.At(0); v != "a" {
		t.Errorf(`got %q, want "a"`, v)
	}
	if v := s.At(1); v != "b" {
		t.Errorf(`got %q, want "b"`, v)
	}
	if v := s.At(2); v != "c" {
		t.Errorf(`got %q, want "c"`, v)
	}
}
