#!/usr/bin/env bash

cmd lint-go,lg Lint Go Code
lint-go () {
	install-go
	install-golangci-lint

	printf "Linting Go..."
	try "(${EXEC_GOLANGCILINT} --verbose --timeout=5m run && ${EXEC_GOVULNCHECK} -tags release ./...)"
}
lg () {
	lint-go
}

cmd lint-shell,ls Lint Shell code
lint-shell () {
	install-shellcheck

	printf "Linting Shell..."
	try "${EXEC_SHELLCHECK} -e SC2153 -x ${DIR}/m ${DIR}/shell/lib/*"
}
ls () {
	lint-shell
}

cmd lint-web,lw Lint Web code
lint-web () {
	install-node

	printf "Linting Web..."
	try "${EXEC_NPM} run lint"
}
lw () {
	lint-web
}

cmd lint-yaml8n,ly Lint YAML8n translations
lint-yaml8n() {
	install-yaml8n

	for i in "${DIR}"/yaml8n/*; do
		try "${EXEC_YAML8N} -x input_path=${i} validate"
	done
}
ly () {
	lint-yaml8n
}
