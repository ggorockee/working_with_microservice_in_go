REPOSITORY=microservice
LINUX_VERSION=alpine3.20
DOCKER_HOST=ggorockee

BROKER_BINARY=brokerApp




up_build: build_broker push_broker
	@echo "build push done!"

build_broker:
	@echo "[INFO] building broker binary"
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api/
	@echo "[INFO] build_broker done\n"

push_broker:
	@echo "[INFO] Change docker tag"
	cd ./broker-service && docker build -t ${DOCKER_HOST}/${REPOSITORY}:broker-${LINUX_VERSION} -f broker-service.dockerfile .
	@echo "[INFO] push ggorockee docker hub"
	docker push ${DOCKER_HOST}/${REPOSITORY}:broker-${LINUX_VERSION}
	@echo "[INFO] push_broker done\n"
