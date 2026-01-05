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
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Arch      string `json:"arch"`
	Signature string `json:"signature,omitempty"`
	ABI       string `json:"abi,omitempty"`
	Notes     string `json:"notes,omitempty"`
	Instr     string `json:"instr,omitempty"`
	Since     string `json:"since,omitempty"`
}

type Meta struct {
	SchemaVersion      string            `json:"schema_version"`
	DatasetVersion     string            `json:"dataset_version"`
	GeneratedAt        time.Time         `json:"generated_at"`
	Architectures      []string          `json:"architectures"`
	RecordCountByArch  map[string]int    `json:"record_count_by_arch"`
	SourceNotes        string            `json:"source_notes"`
}
