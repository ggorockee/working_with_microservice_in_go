REPOSITORY=microservice
LINUX_VERSION=alpine3.20
DOCKER_HOST=ggorockee

BROKER_BINARY=brokerApp
MAIL_BINARY=mailServiceApp
LOGGER_BINARY=loggerServiceApp
AUTH_BINARY=authApp
FRONT_BINARY=frontApp

TAG=k8s

###### VERSION ######
BROKER_VERSION=5
MAIL_VERSION=1
LOGGER_VERSION=1
AUTH_VERSION=2
FRONT_VERSION=11
########################




push_front_on_hub: build_front push_front

push_mail_on_hub: build_mailer push_mailer

push_logger_on_hub: build_logger push_logger

push_broker_on_hub: build_broker push_broker

push_auth_on_hub: build_auth push_auth

build_broker:
	@echo "[INFO] building broker binary"
	cd ./broker-service && go get github.com/ggorockee/working_with_microservice_in_go/broker-service/cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api/
	@echo "[INFO] build_broker done\n"

push_broker:
	@echo "[INFO] Change docker tag"
	cd ./broker-service && docker build -t ${DOCKER_HOST}/${REPOSITORY}:broker-${LINUX_VERSION}-${TAG}-${BROKER_VERSION} -f broker-service.dockerfile .
	@echo "[INFO] push ggorockee docker hub"
	docker push ${DOCKER_HOST}/${REPOSITORY}:broker-${LINUX_VERSION}-${TAG}-${BROKER_VERSION}
	@echo "\n[OUTPUT]: ${DOCKER_HOST}/${REPOSITORY}:broker-${LINUX_VERSION}-${TAG}-${BROKER_VERSION}\n"
	@echo "[INFO] push_broker done\n"

build_mailer:
	@echo "[INFO] building mail binary"
	cd ./mail-service && go get mail-service/cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/api/
	@echo "[INFO] build mail done\n"

push_mailer:
	@echo "[INFO] Change docker tag"
	cd ./mail-service && docker build -t ${DOCKER_HOST}/${REPOSITORY}:mailer-${LINUX_VERSION}-${TAG}-${MAIL_VERSION} -f mail-service.dockerfile .
	@echo "[INFO] push ggorockee docker hub"
	docker push ${DOCKER_HOST}/${REPOSITORY}:mailer-${LINUX_VERSION}-${TAG}-${MAIL_VERSION}
	@echo "\n[OUTPUT]: ${DOCKER_HOST}/${REPOSITORY}:mailer-${LINUX_VERSION}-${TAG}-${MAIL_VERSION}\n"
	@echo "[INFO] push_broker done\n"

build_logger:
	@echo "[INFO] building logger binary"
	cd ./logger-service && go get log-service/cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api/
	@echo "[INFO] build logger done\n"

push_logger:
	@echo "[INFO] Change docker tag"
	cd ./logger-service && docker build -t ${DOCKER_HOST}/${REPOSITORY}:logger-${LINUX_VERSION}-${TAG}-${LOGGER_VERSION} -f logger-service.dockerfile .
	@echo "[INFO] push ggorockee docker hub"
	docker push ${DOCKER_HOST}/${REPOSITORY}:logger-${LINUX_VERSION}-${TAG}-${LOGGER_VERSION}
	@echo "\n[OUTPUT]: ${DOCKER_HOST}/${REPOSITORY}:logger-${LINUX_VERSION}-${TAG}-${LOGGER_VERSION}\n"
	@echo "[INFO] push_broker done\n"

build_auth:
	@echo "[INFO] building auth binary"
	cd ./authentication-service && go get authentication-service/cmd/api && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api/
	@echo ""
	@echo "[INFO] build auth done"
	@echo ""

push_auth:
	@echo "[INFO] Change docker tag"
	cd ./authentication-service && docker build -t ${DOCKER_HOST}/${REPOSITORY}:auth-${LINUX_VERSION}-${TAG}-${AUTH_VERSION} -f authentication-service.dockerfile .
	@echo "[INFO] push ggorockee docker hub"
	docker push ${DOCKER_HOST}/${REPOSITORY}:auth-${LINUX_VERSION}-${TAG}-${AUTH_VERSION}
	@echo ""
	@echo "[OUTPUT]: ${DOCKER_HOST}/${REPOSITORY}:auth-${LINUX_VERSION}-${TAG}-${AUTH_VERSION}"
	@echo ""
	@echo "[INFO] push_broker done"
	@echo ""


build_front:
	@echo "[INFO] building front binary"
	cd ./front-end && go get front-end/cmd/web && env GOOS=linux CGO_ENABLED=0 go build -o ${FRONT_BINARY} ./cmd/web/
	@echo ""
	@echo "[INFO] build front done"
	@echo ""

push_front:
	@echo "[INFO] Change docker tag"
	cd ./front-end && docker build -t ${DOCKER_HOST}/${REPOSITORY}:front-${LINUX_VERSION}-${TAG}-${FRONT_VERSION} -f frontend.dockerfile .
	@echo "[INFO] push ggorockee docker hub"
	docker push ${DOCKER_HOST}/${REPOSITORY}:front-${LINUX_VERSION}-${TAG}-${FRONT_VERSION}
	@echo ""
	@echo "[OUTPUT]: ${DOCKER_HOST}/${REPOSITORY}:front-${LINUX_VERSION}-${TAG}-${FRONT_VERSION}"
	@echo ""
	@echo "[INFO] push front done"
	@echo ""

