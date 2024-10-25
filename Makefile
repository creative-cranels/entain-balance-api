GOPATH = $(shell go env | egrep GOPATH | egrep -o '/[^"]+')

.PHONY: all
all:            ## Start dev server
	@command -v air || (echo "air is not installed. Please run $$ make devsetup"; exit 1)
	@$(MAKE) docs
	air
