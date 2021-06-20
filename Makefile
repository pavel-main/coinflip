TOKEN_SALE_REPO := $(GOPATH)/src/github.com/pavel-main/token-sale
SALE_CONTRACT := $(TOKEN_SALE_REPO)/build/contracts/ShopTokenSale.json
TOKEN_CONTRACT := $(TOKEN_SALE_REPO)/build/contracts/ShopToken.json

.PHONY: abigen build ensure install

abigen-sale:
	@cat $(SALE_CONTRACT) | jq -c .abi > token_sale.abi
	@abigen --abi token_sale.abi --pkg contracts --type TokenSale --out contracts/token_sale.go
	@rm -rf token_sale.abi

abigen-token:
	@cat $(TOKEN_CONTRACT) | jq -c .abi > shop_token.abi
	@abigen --abi shop_token.abi --pkg contracts --type ShopToken --out contracts/shop_token.go
	@rm -rf shop_token.abi

abigen: abigen-sale abigen-token

build:
	go build -ldflags="`govvv -flags`" -o ./coinflip

install:
	go get -u github.com/golang/dep/cmd/dep
	go get github.com/ahmetb/govvv

ensure:
	dep ensure -v

