TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'tools')
HOSTNAME=github.com
NAMESPACE=awlsring
NAME=headscale
BINARY=terraform-provider-${NAME}

default: gen install terradocs

clean:
	rm ${BINARY}

terradocs:
	go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	go generate

gen:
	swagger generate client -f models/headscale.23.0.json -A headscale -t ./internal/gen
	go mod tidy

install:
	go install .

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m