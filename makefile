up:
	@echo "Starting the build process..."
	docker compose up --build -d

down:
	@echo "Stopping the build process..."
	docker compose down
