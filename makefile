IMAGE=coding:latest
PROJECT=golang-coding

sh:
	@docker-compose build sh
	@docker-compose run --rm sh

test:
	@docker-compose build test
	@docker-compose run --rm test
