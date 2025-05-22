Update_File=./script/update.sh

update_lib:
	@if [ ! -x "$(Update_File)" ]; then \
		sudo chmod +x $(Update_File) ;\
	fi
	zsh $(Update_File)
	#sh ./script/update-lib.sh


.PHONY: update_lib