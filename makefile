Update_File=./script/update.sh
IMAGE_NAME=gateway
Build_File=./script/build.sh

update_lib:
	@if [ ! -x "$(Update_File)" ]; then \
		sudo chmod +x $(Update_File) ;\
	fi
	zsh $(Update_File)
	#sh ./script/update-lib.sh

buildBinary:
	@if [ ! -x "$(Update_File)" ]; then \
    		sudo chmod +x $(Build_File) ;\
    	fi
	./script/build.sh cmd/main.go $(IMAGE_NAME)
build: buildBinary
	docker build --progress=plain -t $(IMAGE_NAME) .
run: build
	docker run --rm -p 8585:8585 $(IMAGE_NAME)

.PHONY: update_lib,build,run,buildBinary