package store

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Selyss/AssemBuddy/data"
	"github.com/Selyss/AssemBuddy/internal/model"
)

const metaFile = "meta.json"

var archFiles = map[model.Arch]string{
	model.ArchArm:   "syscalls_arm.json",
	model.ArchArm64: "syscalls_arm64.json",
	model.ArchX64:   "syscalls_x64.json",
	model.ArchX86:   "syscalls_x86.json",
}

type Store struct {
	Meta        model.Meta
	ByArch      map[model.Arch][]model.SyscallRecord
	ByArchName  map[model.Arch]map[string]model.SyscallRecord
	AllByName   map[string][]model.SyscallRecord
}

func Load() (*Store, error) {
	metaBytes, err := data.EmbeddedFS.ReadFile(metaFile)
	if err != nil {
		return nil, fmt.Errorf("read meta: %w", err)
	}

	var meta model.Meta
	if err := json.Unmarshal(metaBytes, &meta); err != nil {
		return nil, fmt.Errorf("parse meta: %w", err)
	}

	store := &Store{
		Meta:       meta,
		ByArch:     make(map[model.Arch][]model.SyscallRecord),
		ByArchName: make(map[model.Arch]map[string]model.SyscallRecord),
		AllByName:  make(map[string][]model.SyscallRecord),
	}

	for arch, filename := range archFiles {
		payload, err := data.EmbeddedFS.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", filepath.Base(filename), err)
		}
		var records []model.SyscallRecord
		if err := json.Unmarshal(payload, &records); err != nil {
			return nil, fmt.Errorf("parse %s: %w", filepath.Base(filename), err)
		}
		store.ByArch[arch] = records
		byName := make(map[string]model.SyscallRecord, len(records))
		for _, rec := range records {
			nameKey := strings.ToLower(strings.TrimSpace(rec.Name))
			byName[nameKey] = rec
			store.AllByName[nameKey] = append(store.AllByName[nameKey], rec)
		}
		store.ByArchName[arch] = byName
	}

	for name, items := range store.AllByName {
		store.AllByName[name] = sortByArch(items)
	}

	if err := store.validateMeta(); err != nil {
		return nil, err
	}

	store.sortArchLists()

	return store, nil
}

func (s *Store) validateMeta() error {
	archSet := map[string]struct{}{}
	for _, arch := range s.Meta.Architectures {
		archSet[arch] = struct{}{}
	}
	for _, arch := range model.CanonicalArchOrder {
		if _, ok := archSet[string(arch)]; !ok {
			return fmt.Errorf("meta missing architecture: %s", arch)
		}
	}
	for arch, records := range s.ByArch {
		if s.Meta.RecordCountByArch == nil {
			continue
		}
		if count, ok := s.Meta.RecordCountByArch[string(arch)]; ok && count != len(records) {
			return fmt.Errorf("meta record count mismatch for %s: %d != %d", arch, count, len(records))
		}
	}
	return nil
}

func (s *Store) sortArchLists() {
	for arch, records := range s.ByArch {
		sorted := make([]model.SyscallRecord, len(records))
		copy(sorted, records)
		sort.Slice(sorted, func(i, j int) bool {
			if sorted[i].Number == sorted[j].Number {
				return sorted[i].Name < sorted[j].Name
			}
			return sorted[i].Number < sorted[j].Number
		})
		s.ByArch[arch] = sorted
	}
}

func sortByArch(records []model.SyscallRecord) []model.SyscallRecord {
	archOrder := map[string]int{}
	for idx, arch := range model.CanonicalArchOrder {
		archOrder[string(arch)] = idx
	}
	sorted := make([]model.SyscallRecord, len(records))
	copy(sorted, records)
	sort.Slice(sorted, func(i, j int) bool {
		return archOrder[sorted[i].Arch] < archOrder[sorted[j].Arch]
	})
	return sorted
}
