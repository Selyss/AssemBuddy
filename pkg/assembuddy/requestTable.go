package assembuddy

import (
	"encoding/json"
	"io"
	"net/http"
)

type SystemCall struct {
	Arch        string `json:"arch"`
	Name        string `json:"name"`
	ReturnValue string `json:"return"`
	Arg0        string `json:"arg0"`
	Arg1        string `json:"arg1"`
	Arg2        string `json:"arg2"`
	Arg3        string `json:"arg3"`
	Arg4        string `json:"arg4"`
	Arg5        string `json:"arg5"`
	Nr          int    `json:"nr"`
}

func fetchData(endpointURL string) ([]SystemCall, error) {
	response, err := http.Get(endpointURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var systemCalls []SystemCall
	err = json.Unmarshal(body, &systemCalls)
	if err != nil {
		return nil, err
	}

	return systemCalls, nil
}

func GetFromArchitecture(arch string) ([]string, error) {
	return nil, nil
}
