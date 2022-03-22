.PHONY: create-keypair

PWD = $(shell pwd)
ACCTPATH = $(PWD)/account

create-keypair:
	@echo "Creating an rsa 256 key pair"
    # TODO: the relative path seems not work well
    # openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
    # openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

	openssl genpkey -algorithm RSA -out 'D:\# Desktop\wordboard\account\rsa_private_test.pem' -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in 'D:\# Desktop\wordboard\account\rsa_private_test.pem' -pubout -out 'D:\# Desktop\wordboard\account\rsa_public_test.pem'