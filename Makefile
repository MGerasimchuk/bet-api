good:
	docker-compose down -v --remove-orphans ; docker-compose up
rebuild:
	docker-compose down -v --remove-orphans ; docker-compose up --build
