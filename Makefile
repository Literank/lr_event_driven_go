TARGET_BIN=./bin/
# Binary names
BINARY_WEB=lrbooks_web
BINARY_TREND=lrbooks_trend
BINARY_REC=lrbooks_rec

build-web:
	@echo "Building $(BINARY_WEB) for Linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${TARGET_BIN}$(BINARY_WEB) service/web/main.go

build-trend:
	@echo "Building $(BINARY_TREND) for Linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${TARGET_BIN}$(BINARY_TREND) service/trend/main.go

build-rec:
	@echo "Building $(BINARY_REC) for Linux..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${TARGET_BIN}$(BINARY_REC) service/recommendation/main.go

build: build-web build-trend build-rec