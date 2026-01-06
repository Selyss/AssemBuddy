package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Selyss/AssemBuddy/internal/model"
	"github.com/Selyss/AssemBuddy/internal/query"
)

type Options struct {
	OutDir    string
	Overwrite bool
	Arches    []string
	APIBase   string
}

type apiSyscall struct {
	Arch       string `json:"arch"`
	Name       string `json:"name"`
	Number     int    `json:"nr"`
	Arg0       string `json:"arg0"`
	Arg1       string `json:"arg1"`
	Arg2       string `json:"arg2"`
	Arg3       string `json:"arg3"`
	Arg4       string `json:"arg4"`
	Arg5       string `json:"arg5"`
	ReturnType string `json:"return"`
	References string `json:"references"`
}

func GenerateDataset(opts Options) (*model.Meta, error) {
	arches, err := resolveArches(opts.Arches)
	if err != nil {
		return nil, err
	}
	if opts.APIBase == "" {
		opts.APIBase = "https://api.syscall.sh/v1"
	}
	if opts.OutDir == "" {
		opts.OutDir = "data"
	}

	if err := os.MkdirAll(opts.OutDir, 0o755); err != nil {
		return nil, fmt.Errorf("create output dir: %w", err)
	}

	counts := make(map[string]int)
	for _, arch := range arches {
		syscalls, err := fetchArch(opts.APIBase, arch)
		if err != nil {
			return nil, err
		}
		records := transformRecords(syscalls)
		counts[string(arch)] = len(records)

		filename := filepath.Join(opts.OutDir, fmt.Sprintf("syscalls_%s.json", arch))
		if err := writeJSON(filename, records, opts.Overwrite); err != nil {
			return nil, err
		}
	}

	archStrings := make([]string, 0, len(arches))
	for _, arch := range arches {
		archStrings = append(archStrings, string(arch))
	}
	meta := &model.Meta{
		SchemaVersion:     "1",
		DatasetVersion:    time.Now().UTC().Format("2006-01-02"),
		GeneratedAt:       time.Now().UTC(),
		Architectures:     archStrings,
		RecordCountByArch: counts,
		SourceNotes:       "Data sourced from api.syscall.sh",
	}
	if err := writeJSON(filepath.Join(opts.OutDir, "meta.json"), meta, opts.Overwrite); err != nil {
		return nil, err
	}

	return meta, nil
}

func resolveArches(arches []string) ([]model.Arch, error) {
	if len(arches) == 0 {
		return model.CanonicalArchOrder, nil
	}
	result := make([]model.Arch, 0, len(arches))
	for _, arch := range arches {
		parsed, err := query.NormalizeArch(arch)
		if err != nil {
			return nil, err
		}
		result = append(result, parsed)
	}
	return result, nil
}

func fetchArch(apiBase string, arch model.Arch) ([]apiSyscall, error) {
	syscallsURL := fmt.Sprintf("%s/syscalls/%s", strings.TrimRight(apiBase, "/"), arch)

	var syscalls []apiSyscall
	if err := getJSON(syscallsURL, &syscalls); err != nil {
		return nil, err
	}
	return syscalls, nil
}

func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("fetch %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("fetch %s: status %d: %s", url, resp.StatusCode, strings.TrimSpace(string(body)))
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode %s: %w", url, err)
	}
	return nil
}

func transformRecords(syscalls []apiSyscall) []model.SyscallRecord {
	records := make([]model.SyscallRecord, 0, len(syscalls))
	for _, call := range syscalls {
		records = append(records, model.SyscallRecord{
			Arch:       normalizeArch(call.Arch),
			Number:     call.Number,
			Name:       call.Name,
			References: call.References,
			Return:     call.ReturnType,
			Arg0:       call.Arg0,
			Arg1:       call.Arg1,
			Arg2:       call.Arg2,
			Arg3:       call.Arg3,
			Arg4:       call.Arg4,
			Arg5:       call.Arg5,
		})
	}
	return records
}

func normalizeArch(arch string) string {
	parsed, err := query.NormalizeArch(arch)
	if err != nil {
		return arch
	}
	return string(parsed)
}

func writeJSON(path string, payload interface{}, overwrite bool) error {
	if _, err := os.Stat(path); err == nil && !overwrite {
		return fmt.Errorf("file exists: %s", path)
	}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal %s: %w", path, err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
