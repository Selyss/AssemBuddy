package render

import (
	"encoding/json"

	"github.com/Selyss/AssemBuddy/internal/model"
)

func RenderJSONRecord(record model.SyscallRecord) (string, error) {
	payload, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return "", err
	}
	return string(payload), nil
}

func RenderJSONRecords(records []model.SyscallRecord) (string, error) {
	payload, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return "", err
	}
	return string(payload), nil
}
