NAME := json_sorter

check_result :=	@read -p "view result? (y/N): " ans; \
				if [ "$$ans" = y ]; then \
					cat ./dist/result.json; \
				fi

all: $(NAME)

$(NAME): main.go
	go build -o $@ $<

clean:
	$(RM) $(NAME)

quick:
	@./$(NAME) -type quick -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

merge:
	@./$(NAME) -type merge -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

insertion:
	@./$(NAME) -type insertion -json ./misc/data/user.json -output ./dist/result.json
	$(check_result)

.PHONY: all clean quick merge insertion