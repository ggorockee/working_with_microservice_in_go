VOL=volume
REPOSITORY=ggorockee/dind
DIND_TAG=27.2.1-alpine3.20-go1.23-crossplatform-v3
NAME=dind

#####################################################
build_image_tag=3

run: daemon exec

daemon:
	@echo "[INFO] docker pull docker in docker"
	mkdir -p ${VOL}
	docker run -idt  --privileged -v "./${VOL}:/data" -w /data --name ${NAME} ${REPOSITORY}:${DIND_TAG}
	@echo "Done!!\n"

exec:
	@echo [INFO] docker pull docker in docker
	docker exec -it ${NAME} sh
	@echo "Done!!\n"


down:
	@echo "[INFO] docker cleaning"
	docker rm -f ${NAME}
	@echo "[INFO] done!"

create_image_and_push:
	@echo "[INFO] Create & Push Docker in Docker Image"
	docker buildx build --push --platform linux/arm64/v8,linux/amd64 --tag ${REPOSITORY}:27.2.1-alpine3.20-go1.23-crossplatform-v${build_image_tag} -f createImageGO.dockerfile .
	@echo "[INFO] Done Docker in Docker Image"