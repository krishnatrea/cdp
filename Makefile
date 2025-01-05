all: build  

down: 
	docker-compose down

build: 
	docker-compose up --build  