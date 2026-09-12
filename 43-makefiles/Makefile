# Targets: allows to run a specific command
# ex, help: will run "make help"
# PHONY is special declaration to let know that "target" is a command but not a file
# ex: install: @echo "hello", make install -> hello
#  but if install file exists in pwd, then make install, looks for the file. to solve we use .PHONY: install
# @ dont print command
main_package_path = ./
binary_name = export
build_dir = ./bin

.PHONY: help
help:
	@echo "Usage:"
# 	list everything in the Makefile with // and print
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# multi-archirtecture binary build
# The below is the documentation which is read by sed command above
## build: build the application
.PHONY: build
build:
	@mkdir -p ${build_dir}
	GOARCH=amd64 GOOS=darwin go build -o ${build_dir}/${binary_name}-darwin ${main_package_path}
	GOARCH=amd64 GOOS=linux go build -o ${build_dir}/${binary_name}-linux ${main_package_path}
	GOARCH=amd64 GOOS=windows go build -o ${build_dir}/${binary_name}-windows ${main_package_path}

# clean runs only when confirm target succeeds
## clean: remove the build directory and its contents
.PHONY: clean
clean: confirm
	@echo "Cleaning build directory..."
	@rm -rf ${build_dir}
