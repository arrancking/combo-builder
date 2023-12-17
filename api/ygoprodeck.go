package api

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type ygoProAPI struct {
	URL string
}

type response struct {
	Cards []card `json:"data"`
}

func (a ygoProAPI) GetCard(name string) (*card, error) {
	resp, err := http.Get(
		strings.Join(
			[]string{a.URL, "?name=", name},
			"",
		),
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respStruct response
	err = json.Unmarshal(body, &respStruct)
	if err != nil {
		return nil, err
	}

	if len(respStruct.Cards) != 1 {
		slog.Error(
			"Unexpected response from YGOProdeck",
			"response length", len(respStruct.Cards),
			"response", respStruct,
		)
		return nil, errors.New("unexpected response from YGOProdeck")
	}

	return &respStruct.Cards[0], nil
}
