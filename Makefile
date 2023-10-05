
.PHONY: ts-re
ts-re: 
	cd testserver && make re

.PHONY: ts-down
ts-down: 
	cd testserver && make down

.PHONY: run
up:
	go run .