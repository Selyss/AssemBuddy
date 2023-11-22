package assembuddy

import (
	"reflect"
	"testing"
)

func TestGetSyscallData(t *testing.T) {
	opts := CLIOptions{
		Syscall:     "read",
		Arch:        "x64",
		ListArch:    false,
		PrettyPrint: false,
	}

	want := []Syscall{}

	syscall := Syscall{
		Arch:        "x64",
		Name:        "read",
		ReturnValue: "0x00",
		Arg0:        "unsigned int fd",
		Arg1:        "char *buf",
		Arg2:        "size_t count",
		Arg3:        "",
		Arg4:        "",
		Arg5:        "",
	}

	want = append(want, syscall)

	got, err := GetSyscallData(&opts)
	if !(reflect.DeepEqual(want, got)) || err != nil {
		t.Errorf("Expected opts to parse correctly %v", err)
	}
}
