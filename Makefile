# RUN LINTER
lint:
	gofmt -w .
	golangci-lint run
	go vet ./...

# GENERATE CONTRACT API FROM ABI
abi:
	abigen --abi=contract/nirgp/abi/NIR_NFT_Genesis_Pass.abi --pkg=nirgp --out=contract/nirgp/nirgp.go

# RUN APP
run:
	go run main.go

.PHONY: lint, abi, run
