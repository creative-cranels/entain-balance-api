.PHONY: setup

default: run

setup:
	@echo "🌀 Setting up the system..."
	docker-compose run app bash setup.sh
	@echo "✔  System is ready"

run:
	docker-compose up