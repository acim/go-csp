package seen_test

import (
	"testing"

	"github.com/acim/go-csp/pkg/seen"
)

func TestSeen(t *testing.T) {
	s := seen.CreateMap(2)
	s.SetSeen(1)
	if !s.Seen(1) {
		t.Errorf("got false, want true")
	}
	if s.Seen(2) {
		t.Errorf("got true, want false")
	}
}

func TestAllSeen(t *testing.T) {
	s := seen.CreateMap(2)
	s.SetSeen(1)
	if s.AllSeen() {
		t.Errorf("got true, want false")
	}
	s.SetSeen(2)
	if !s.AllSeen() {
		t.Errorf("got false, want true")
	}
}
