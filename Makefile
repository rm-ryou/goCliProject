NAME := json_sorter

check_result :=	@read -p "view result? (y/N): " ans; \
				if [ "$$ans" = y ]; then \
					cat ./dist/result.json; \
				fi

all: $(NAME)	## Creation of executable file

$(NAME): main.go
	go build -o $@ $<

clean:			## Delete executable file
	$(RM) $(NAME)

quick: $(NAME)		## Performing an Quick Sort
	@./$(NAME) -type quick -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

merge: $(NAME)		## Performing an Merge Sort
	@./$(NAME) -type merge -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

insertion: $(NAME)	## Performing an Insertion Sort
	@./$(NAME) -type insertion -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

help:	## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[31m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: all clean quick merge insertion help