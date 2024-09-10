VOL=volume
REPOSITORY=ggorockee/dind
DIND_TAG=27.2.1-alpine3.20-go1.23
NAME=dind


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