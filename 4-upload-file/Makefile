all: help

.PHONY:
help:
	@echo "swagger example doc"
	@echo ""
	@echo "generate-index            generate swagger documentation on index.yaml"
	@echo "serve                     Creates a server to render the content of index.yaml"
	@echo ""

.PHONY:
generate-index:
	@swagger generate spec -o ./docs/index.yaml --scan-models

.PHONY:
serve:
	@swagger-ui-watcher docs/index.yaml -p 8081