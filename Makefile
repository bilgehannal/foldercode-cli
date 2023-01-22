compile:
	go build github.com/bilgehannal/foldercode-cli/cmd/fcode && echo 'Compiling Succesfull...'

compile_linux:
	GOOS=linux GOARCH=amd64 go build github.com/bilgehannal/foldercode-cli/cmd/fcode && mv fcode fcode.linux && echo 'Compiling Succesfull...'

compile_macosx:
	go build github.com/bilgehannal/foldercode-cli/cmd/fcode && mv fcode fcode.macosx && echo 'Compiling Succesfull...'

unit_test:
	go test ./... -coverprofile=coverage.out -coverpkg=./...

open_test_result:
	go test ./... -coverprofile=coverage.out -coverpkg=./... && go tool cover -html=coverage.out