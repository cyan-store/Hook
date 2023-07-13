APP = "hook" # App name
PORT = 3015
STRIPE = stripe # https://stripe.com/docs/stripe-cli
UPX = upx # https://upx.github.io/ (optional)

format:
	@printf "Formating $(APP)-\n"
	@go fmt ./...
	@go mod tidy

## Production
.PHONY: build
build:
	@printf "Building $(APP)-\n"
	@go build -o ./bin/$(APP) ./main.go
	@printf "Do you want to minify $(APP) with upx? (y/N) "; \
		read answer_minify; \
		if [ $$answer_minify == "y" ]; then $(UPX) ./bin/$(APP); fi

## Development
.PHONY: dev
dev:
	@printf "Running $(APP)-\n"
	@go run *.go

.PHONY: listen
listen:
	@printf "Stripe webhook 127.0.0.1:$(PORT)"
	@$(STRIPE) listen --forward-to 127.0.0.1:$(PORT)/checkout
