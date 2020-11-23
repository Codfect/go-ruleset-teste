package opts

import "testing"

func TestDependsAA(t *testing.T) {
	s := NewRuleSet()
	s.AddDep("a", "a")
	if !s.IsCoherent() {
		t.Error("s.IsCoherent failed")
	}
}

func TestDependsAB_BA(t *testing.T) {
	s := NewRuleSet()
	s.AddDep("a", "b")
	s.AddDep("b", "a")
	if !s.IsCoherent() {
		t.Error("s.IsCoherent failed")
	}
}

func TestExclusiveAB(t *testing.T) {
	s := NewRuleSet()
	s.AddDep("a", "b")
	s.AddConflict("a", "b")
	if s.IsCoherent() {
		t.Error("s.IsCoherent failed")
	}
}

func TestExclusiveAB_BC(t *testing.T) {
	s := NewRuleSet()
	s.AddDep("a", "b")
	s.AddDep("b", "c")
	s.AddConflict("a", "c")
	if s.IsCoherent() {
		t.Error("s.IsCoherent failed")
	}
}

func TestDeepDeps(t *testing.T) {
	s := NewRuleSet()
	s.AddDep("a", "b")
	s.AddDep("b", "c")
	s.AddDep("c", "d")
	s.AddDep("d", "e")
	s.AddDep("a", "f")
	s.AddConflict("e", "f")

	if s.IsCoherent() {
		t.Error("IsCoherent failed")
	}
}