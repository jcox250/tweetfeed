ENV_FILE = $(PWD)/.env
ifneq ("$(wildcard $(ENV_FILE))", "")
    include $(ENV_FILE)
    export $(shell sed 's/=.*//' $(ENV_FILE))
endif


build:
	docker-compose build

run:
	docker-compose build
	docker-compose up

