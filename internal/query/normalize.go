package query

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/model"
)

var archAliases = map[string]model.Arch{
	"x64":     model.ArchX64,
	"x86_64":  model.ArchX64,
	"x8664":   model.ArchX64,
	"amd64":   model.ArchX64,
	"x86":     model.ArchX86,
	"i386":    model.ArchX86,
	"i686":    model.ArchX86,
	"386":     model.ArchX86,
	"ia32":    model.ArchX86,
	"arm64":   model.ArchArm64,
	"aarch64": model.ArchArm64,
	"armv8":   model.ArchArm64,
	"armv8a":  model.ArchArm64,
	"arm":     model.ArchArm,
	"armv7":   model.ArchArm,
	"armv6":   model.ArchArm,
	"armhf":   model.ArchArm,
	"arm32":   model.ArchArm,
}

func ArchAliasTable() map[model.Arch][]string {
	result := map[model.Arch][]string{
		model.ArchArm:   {},
		model.ArchArm64: {},
		model.ArchX64:   {},
		model.ArchX86:   {},
	}
	for alias, arch := range archAliases {
		result[arch] = append(result[arch], alias)
	}
	for arch, aliases := range result {
		result[arch] = uniqueSorted(aliases)
	}
	return result
}

func uniqueSorted(input []string) []string {
	seen := map[string]struct{}{}
	out := make([]string, 0, len(input))
	for _, value := range input {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}

func NormalizeArch(input string) (model.Arch, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("architecture is required")
	}
	lower := strings.ToLower(trimmed)
	clean := strings.NewReplacer("_", "", "-", "", " ", "").Replace(lower)
	if arch, ok := archAliases[clean]; ok {
		return arch, nil
	}
	return "", fmt.Errorf("unknown architecture: %s", input)
}

func NormalizeSyscallName(input string, caseSensitive bool) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("syscall name is required")
	}
	lower := strings.ToLower(trimmed)
	prefixes := []string{"__sys_", "sys_"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(lower, prefix) {
			trimmed = trimmed[len(prefix):]
			lower = lower[len(prefix):]
			break
		}
	}
	if caseSensitive {
		return trimmed, nil
	}
	return strings.ToLower(trimmed), nil
}
