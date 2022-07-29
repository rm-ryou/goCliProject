NAME := json_sorter

check_result :=	@read -p "view result? (y/n): " ans; \
				if [ "$$ans" = y ]; then \
					cat ./dist/result.json; \
				fi

all: $(NAME)		## Creation of executable file

$(NAME): main.go
	go build -o $@ $<

clean:		## Delete executable file
	$(RM) $(NAME)

quick: $(NAME)		## Performing an Quick Sort
	@./$(NAME) -type quick -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

merge: $(NAME)		## Performing an Merge Sort
	@./$(NAME) -type merge -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

insertion: $(NAME)		## Performing an Insertion Sort
	@./$(NAME) -type insertion -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

go_test:		## launch go test -v
	@cd ./cmd; go test -v

docker_build:		## Create ubuntu18.04 image
	@docker build -t test_ubunt_18.04 .

docker_run:		## Container execution
	@docker run -it --name "test_env" -v $(PWD):/root --rm test_ubunt_18.04 /bin/bash

help:		## Show Help
	@grep -E '^[a-zA-Z_]+.*## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[34m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: all clean docker_build docker_run quick merge insertion go_test help