all: cli

.PHONY: cli
cli:
	go install -i ./cmd/goparsebgp
