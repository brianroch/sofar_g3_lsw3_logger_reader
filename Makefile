IMG_VERSION="0.16.0"
IMG_NAME="ghcr.io/brianroch/sofar-exporter"
TAG=${IMG_NAME}":"${IMG_VERSION}
TAG_ARM=${IMG_NAME}":"${IMG_VERSION}"-arm"
DOCKERFILE="Dockerfile"
DOCKERFILE_ARM=Dockerfile-arm

build:
	go build -o sofar

build-arm:
	env GOOS=linux GOARCH=arm64 go build -o sofar-arm

docker-build:
	docker build -t ${TAG} -f ${DOCKERFILE} .

docker-push:
	docker push ${TAG}

docker-build-and-push: docker-build docker-push

docker-build-arm:
	docker build --platform linux/arm64 -t ${TAG_ARM} -f ${DOCKERFILE_ARM} .

docker-push-arm:
	docker push ${TAG_ARM}

docker-build-and-push-arm: docker-build-arm docker-push-arm
