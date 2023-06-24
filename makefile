all: build_linux build_win

build_linux:
	@echo "build_linux"
	@go build -o bin/gopro_files_renamer

build_win:
	@echo "build_win"
	@env GOOS=windows GOARCH=amd64 go build -o bin/gopro_files_renamer.exe

test: build_linux
	@echo "run test"
	@rm -rf ./tmp/
	@cp -R ./test_data ./tmp/
	@echo "before -----------------------------------------------------------------------------------------------------"
	@ls tmp
	@cp bin/gopro_files_renamer tmp/gopro_files_renamer
	@cd tmp && ./gopro_files_renamer
	@echo "after ------------------------------------------------------------------------------------------------------"
	@ls tmp
	@rm -rf ./tmp/
