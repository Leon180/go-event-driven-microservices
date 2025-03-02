.PHONY: install-tools
install-tools:
	@./scripts/install-tools.sh

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
	@./scripts/format.sh accounts
	@./scripts/format.sh cards
	@./scripts/format.sh customers
	@./scripts/format.sh loans
	@./scripts/format.sh pkg

.PHONY: lint
lint:
	@./scripts/lint.sh accounts
	@./scripts/lint.sh cards
	@./scripts/lint.sh customers
	@./scripts/lint.sh loans
	@./scripts/lint.sh pkg
