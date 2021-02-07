all: cli

.PHONY: cli
cli: goparsebgp
	go install -i ./cmd/goparsebgp

goparsebgp: build/lib/libparsebgp.a
	go build -i ./

build/lib/libparsebgp.a:
	./scripts/build-libparsebgp.sh
