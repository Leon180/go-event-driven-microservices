.PHONY: install-tools
install-tools:
	@./scripts/install-tools.sh


.PHONY: run-read-restaurants-service
run-read-restaurants-service:
	@./scripts/run.sh  read_restaurants

.PHONY: run-write-restaurants-service
run-write-restaurants-service:
	@./scripts/run.sh  write_restaurants


.PHONY: docker-compose-up
docker-compose-up:
	@docker-compose -f deployments/docker-compose/docker-compose.yaml up --build -d

.PHONY: docker-compose-down
docker-compose-down:
	@docker-compose -f deployments/docker-compose/docker-compose.yaml down


.PHONY: build
build:
	@./scripts/build.sh  pkg
	@./scripts/build.sh  accounts
	@./scripts/build.sh  cards
	@./scripts/build.sh  customers
	@./scripts/build.sh  loans

.PHONY: update-dependencies
update-dependencies:
	@./scripts/update-dependencies.sh  pkg
	@./scripts/update-dependencies.sh  accounts
	@./scripts/update-dependencies.sh  cards
	@./scripts/update-dependencies.sh  customers
	@./scripts/update-dependencies.sh  loans

.PHONY: install-dependencies
install-dependencies:
	@./scripts/install-dependencies.sh  pkg
	@./scripts/install-dependencies.sh  accounts
	@./scripts/install-dependencies.sh  cards
	@./scripts/install-dependencies.sh  customers
	@./scripts/install-dependencies.sh  loans

.PHONY: format
format:
	@./scripts/format.sh read_restaurants
	@./scripts/format.sh write_restaurants
	@./scripts/format.sh pkg

.PHONY: lint
lint:
	@./scripts/lint.sh accounts
	@./scripts/lint.sh cards
	@./scripts/lint.sh customers
	@./scripts/lint.sh loans
	@./scripts/lint.sh pkg
