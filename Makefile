run:
	docker compose up --build backend

up:
	docker-compose up --build

clean-containers:
	docker-compose down -v