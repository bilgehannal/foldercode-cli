compile:
	go build github.com/bilgehannal/foldercode-cli/cmd/fcode && echo 'Compiling Succesfull...'

unit_test:
	go test ./... -coverprofile=coverage.out -coverpkg=./...

open_test_result:
	go test ./... -coverprofile=coverage.out -coverpkg=./... && go tool cover -html=coverage.out