
# Color pallete
COLOUR_GREEN=\033[0;32m
COLOUR_RED=\033[0;31m
COLOUR_BLUE=\033[0;34m
COLOUR_END=\033[0m


start_backend_air:
	@echo "$(COLOUR_GREEN)--- Starting Backend Server with Air ---$(COLOUR_END)"
	@air -build.cmd "go build -o ./tmp/main ./cmd/api"

start_frontend_air:
	@echo "$(COLOUR_GREEN)--- Starting Frontend Server with Air ---$(COLOUR_END)"
	@air -build.cmd "go build -o ./tmp/main ./cmd/web"

run_backend:
	@echo "$(COLOUR_GREEN)--- Starting Backend Server ---$(COLOUR_END)"
	@go run ./cmd/api

run_frontend:
	@echo "$(COLOUR_GREEN)--- Starting Frontend Server ---$(COLOUR_END)"
	@go run ./cmd/web
