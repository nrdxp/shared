#!/usr/bin/env bash

cmd build-go,bg Build Go
build-go () {
	install-go

	printf "Building go/%s..." "${BUILD_NAME}"
	try "cd ${DIR}/go && CGO_ENABLED=0 GOOS=${BUILD_TARGET_OS} GOARCH=${BUILD_TARGET_ARCH} go build -tags ${GO_BUILD_TAGS} -v -ldflags '-X candid/lib/cli.BuildDate=${BUILD_DATE} -X candid/lib/cli.BuildVersion=${BUILD_VERSION} ${GO_BUILD_VARS} -w' -o ${DIR}/${BUILD_NAME} ."
}
bg () {
	build-go
}

cmd build-hugo,bh Build Hugo
build-hugo () {
	install-hugo

	printf "Building hugo..."
	try "cd ${DIR}/hugo; hugo -e prd --gc --minify"
}
bh () {
	build-hugo
}

cmd build-yaml8n,by Build YAML8n translations
build-yaml8n () {
	for i in "${DIR}"/yaml8n/*; do
		# shellcheck disable=SC2086
		${CR} run -it --rm ${CR_USER} -e "GOOGLE_APPLICATION_CREDENTIALS=/.gcp" -v "${DIR}/.gcp:/.gcp" --pull always ${CR_VOLUME} "${CR_REGISTRY}/candiddev/yaml8n:latest" translate "/work/yaml8n/$(basename "${i}")"
	done
}
by () {
	build-yaml8n
}
