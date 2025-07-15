TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'tools')
HOSTNAME=github.com
NAMESPACE=awlsring
NAME=headscale
BINARY=terraform-provider-${NAME}
HEADSCALE_VERSION=25.0

default: gen install terradocs

clean:
	rm ${BINARY}

terradocs:
	go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go generate

gen:
	swagger generate client -f models/headscale.$(HEADSCALE_VERSION).json -A headscale -t ./internal/gen
	go mod tidy

install:
	go install .

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	scripts/run_tests.sh