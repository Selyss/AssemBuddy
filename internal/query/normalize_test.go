package query

import "testing"

func TestNormalizeArch(t *testing.T) {
	cases := map[string]string{
		"x86_64": "x64",
		"AMD64":  "x64",
		"x86":    "x86",
		"i386":   "x86",
		"arm64":  "arm64",
		"aarch64": "arm64",
		"armv7":  "arm",
		"arm":    "arm",
	}
	for input, expected := range cases {
		arch, err := NormalizeArch(input)
		if err != nil {
			t.Fatalf("unexpected error for %s: %v", input, err)
		}
		if string(arch) != expected {
			t.Fatalf("expected %s for %s, got %s", expected, input, arch)
		}
	}
}

func TestNormalizeSyscallName(t *testing.T) {
	name, err := NormalizeSyscallName(" SYS_read ", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "read" {
		t.Fatalf("expected read, got %s", name)
	}

	caseSensitive, err := NormalizeSyscallName("Sys_Open", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if caseSensitive != "Open" {
		t.Fatalf("expected Open, got %s", caseSensitive)
	}
}
