.PHONY: keypair migrate-create migrate-up migrate-down migrate-force

PWD = $(shell pwd)/wordboard
ACCTPATH = $(PWD)/account
MPATH = $(ACCTPATH)/migrations

# Default number of migrations to execute up or down
N = 1

# Set address and port for Database
ADDRESS = 47.103.204.39
PORT = 5432
USER = postgres
PASSWORD = 12345

# Create keypair should be in your file below
create-keypair:
	@echo "Creating an rsa 256 key pair"
    # TODO: the relative path seems not work well
	openssl genpkey -algorithm RSA -out $(ACCTPATH)\rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)\rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)\rsa_public_$(ENV).pem

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

migrate-up:
	migrate -path $(MPATH) -database postgres://$(USER):$(PASSWORD)@$(ADDRESS):$(PORT)/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -path $(MPATH) -database postgres://$(USER):$(PASSWORD)@$(ADDRESS):$(PORT)/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -path $(MPATH) -database postgres://$(USER):$(PASSWORD)@$(ADDRESS):$(PORT)/postgres?sslmode=disable force $(VERSION)