package main

import (
	"log/slog"
	"os"

	"github.com/arrancking/combo-builder/api"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	slog.Info("JSON logger set")

	server := api.Init()

	card, err := server.GetCard("Sacred Phoenix of Nephthys")
	if err != nil {
		slog.Error(
			"Error retrieving card",
			"Error", err,
		)
		return
	}

	slog.Info(
		"Card retrieved",
		"Card", card,
	)
}
