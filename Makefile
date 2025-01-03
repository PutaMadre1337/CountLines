.PHONY: build clean

release:
	go build -o ./bin/tammy ./cmd/main.go

build:
	go build -o /usr/bin/tammy ./cmd/main.go

clean:
	rm -f $(INSTALL_DIR)/tammy
