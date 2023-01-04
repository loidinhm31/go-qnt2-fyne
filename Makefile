BINARY_NAME=QNM.app
APP_NAME=QN2Management
VERSION=1.0.0
APP_ID=QNM1.0

## build: build binary and package app
build:
	rm -rf ${BINARY_NAME}
	rm -f fyne-md
	fyne package -appVersion ${VERSION} -appID ${APP_ID} -name ${APP_NAME} -release

## run: builds and runs the application
run:
	go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${BINARY_NAME}
	@echo "Cleaned!"

## coordinator: runs all tests
test:
	go test -v ./...