#!/usr/bin/env bash

cmd build-go,bg Build Go
build-go () {
	install-go

	printf "Building go/%s..." "${BUILD_NAME}"
	try "cd ${DIR}/go && CGO_ENABLED=0 GOOS=${BUILD_TARGET_OS} GOARCH=${BUILD_TARGET_ARCH} go build -tags ${GO_BUILD_TAGS} -v -ldflags '-X candid/lib/cli.BuildDate=${BUILD_DATE} -X candid/lib/cli.BuildVersion=${BUILD_VERSION} ${GO_BUILD_VARS} -w' -o ${DIR}/${BUILD_NAME} ./${APP_NAME}"
}
bg () {
	build-go
}
