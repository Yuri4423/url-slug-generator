package url_slug_generator

import (
	"testing"
)

func TestGenerateBasic(t *testing.T) {
	got := Generate("Hello World")
	want := "hello-world"
	if got != want {
		t.Errorf("Generate(%q) = %q, want %q", "Hello World", got, want)
	}
}

func TestGenerateMixedCaseAndDigits(t *testing.T) {
	got := Generate("Go 1.21 Release")
	want := "go-1-21-release"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateNonASCIIDropped(t *testing.T) {
	got := Generate("Café — Mününchen")
	want := "caf-m-n-nchen"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateLeadingTrailingPunct(t *testing.T) {
	got := Generate("!!!Hello World!!!")
	want := "hello-world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateEmpty(t *testing.T) {
	got := Generate("")
	want := ""
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateAllDisallowed(t *testing.T) {
	got := Generate("!!!???")
	want := ""
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSluggerUnderscoreSeparator(t *testing.T) {
	sl := New("_")
	got := sl.Generate("Hello World")
	want := "hello_world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSluggerMultiCharSeparator(t *testing.T) {
	sl := New("--")
	got := sl.Generate("Hello World")
	want := "hello--world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSluggerEmptySeparatorDefaultsToDash(t *testing.T) {
	sl := New("")
	got := sl.Generate("Hello World")
	want := "hello-world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateTabsAndNewlines(t *testing.T) {
	got := Generate("hello\tworld\nagain")
	want := "hello-world-again"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateDigitsOnly(t *testing.T) {
	got := Generate("20240115")
	want := "20240115"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGenerateAdjacentPunctCollapses(t *testing.T) {
	got := Generate("hello   world")
	want := "hello-world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
