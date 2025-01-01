all: build run 

run: 
	docker run cdp 

build: 
	docker build -t cdp .