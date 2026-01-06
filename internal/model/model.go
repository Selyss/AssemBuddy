package model

import "time"

type Arch string

const (
	ArchArm   Arch = "arm"
	ArchArm64 Arch = "arm64"
	ArchX64   Arch = "x64"
	ArchX86   Arch = "x86"
)

var CanonicalArchOrder = []Arch{ArchArm, ArchArm64, ArchX64, ArchX86}

var CanonicalArchSet = map[Arch]struct{}{
	ArchArm:   {},
	ArchArm64: {},
	ArchX64:   {},
	ArchX86:   {},
}

type SyscallRecord struct {
	Arch       string `json:"arch"`
	Number     int    `json:"nr"`
	Name       string `json:"name"`
	References string `json:"references,omitempty"`
	Return     string `json:"return,omitempty"`
	Arg0       string `json:"arg0,omitempty"`
	Arg1       string `json:"arg1,omitempty"`
	Arg2       string `json:"arg2,omitempty"`
	Arg3       string `json:"arg3,omitempty"`
	Arg4       string `json:"arg4,omitempty"`
	Arg5       string `json:"arg5,omitempty"`
}

type Meta struct {
	SchemaVersion      string            `json:"schema_version"`
	DatasetVersion     string            `json:"dataset_version"`
	GeneratedAt        time.Time         `json:"generated_at"`
	Architectures      []string          `json:"architectures"`
	RecordCountByArch  map[string]int    `json:"record_count_by_arch"`
	SourceNotes        string            `json:"source_notes"`
}
