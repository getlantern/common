PROTOS = $(shell find . -name '*.proto')
SOURCES = $(shell echo "$(PROTOS)" | sed 's/\.proto/\.pb\.go/g')

%.pb.go: %.proto
	protoc --go_out=. --go_opt=paths=source_relative $<

all: $(SOURCES)

.PHONY: all
