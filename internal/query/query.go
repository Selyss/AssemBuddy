package query

import (
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/model"
	"github.com/Selyss/AssemBuddy/internal/store"
)

type QueryOptions struct {
	Name          string
	Arch          model.Arch
	AllArch       bool
	Exact         bool
	CaseSensitive bool
}

func QueryByName(s *store.Store, opts QueryOptions) ([]model.SyscallRecord, error) {
	if opts.AllArch {
		return queryAllArch(s, opts)
	}
	if opts.Arch == "" {
		return nil, fmt.Errorf("architecture is required")
	}
	nameKey, err := NormalizeSyscallName(opts.Name, opts.CaseSensitive)
	if err != nil {
		return nil, err
	}

	records := s.ByArch[opts.Arch]
	if opts.Exact {
		if !opts.CaseSensitive {
			if rec, ok := s.ByArchName[opts.Arch][nameKey]; ok {
				return []model.SyscallRecord{rec}, nil
			}
			return nil, nil
		}
		for _, rec := range records {
			if rec.Name == nameKey {
				return []model.SyscallRecord{rec}, nil
			}
		}
		return nil, nil
	}

	matches := make([]model.SyscallRecord, 0)
	for _, rec := range records {
		candidate := rec.Name
		if !opts.CaseSensitive {
			candidate = strings.ToLower(candidate)
		}
		if strings.Contains(candidate, nameKey) {
			matches = append(matches, rec)
		}
	}
	return matches, nil
}

func queryAllArch(s *store.Store, opts QueryOptions) ([]model.SyscallRecord, error) {
	nameKey, err := NormalizeSyscallName(opts.Name, opts.CaseSensitive)
	if err != nil {
		return nil, err
	}
	if opts.Exact && !opts.CaseSensitive {
		return s.AllByName[nameKey], nil
	}

	results := make([]model.SyscallRecord, 0)
	for _, arch := range model.CanonicalArchOrder {
		records := s.ByArch[arch]
		for _, rec := range records {
			candidate := rec.Name
			if !opts.CaseSensitive {
				candidate = strings.ToLower(candidate)
			}
			if opts.Exact {
				if candidate == nameKey {
					results = append(results, rec)
				}
				continue
			}
			if strings.Contains(candidate, nameKey) {
				results = append(results, rec)
			}
		}
	}
	return results, nil
}

func ListArch(s *store.Store, arch model.Arch, filter string, caseSensitive bool) ([]model.SyscallRecord, error) {
	if arch == "" {
		return nil, fmt.Errorf("architecture is required")
	}
	records := s.ByArch[arch]
	if filter == "" {
		return records, nil
	}
	key := filter
	if !caseSensitive {
		key = strings.ToLower(filter)
	}
	matches := make([]model.SyscallRecord, 0)
	for _, rec := range records {
		candidate := rec.Name
		if !caseSensitive {
			candidate = strings.ToLower(candidate)
		}
		if strings.Contains(candidate, key) {
			matches = append(matches, rec)
		}
	}
	return matches, nil
}
