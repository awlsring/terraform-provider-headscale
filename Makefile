TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'tools')
HOSTNAME=github.com
NAMESPACE=awlsring
NAME=headscale
BINARY=terraform-provider-${NAME}
OS_ARCH=darwin_arm64

default: install terradocs

clean:
	rm ${BINARY}

terradocs:
	go generate

gen:
	swagger generate client -f ${SWAGGER_DOC} -A headscale -t ./internal/gen
	go mod tidy

install: install-tools
	go install .

install-tools:
	go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@latest

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m