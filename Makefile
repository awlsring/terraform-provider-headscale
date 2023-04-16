TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=github.com
NAMESPACE=awlsring
NAME=headscale
BINARY=terraform-provider-${NAME}
VERSION=0.1.1
OS_ARCH=darwin_arm64

default: install

clean:
	rm ${BINARY}

gen:
	swagger generate client -f ${SWAGGER_DOC} -A headscale -t ./internal

install:
	go install .

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m