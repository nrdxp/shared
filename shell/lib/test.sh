#!/usr/bin/env bash

cmd test-go Test Go code
test-go () {
	install-go
	run-postgresql-start

	sleep 5
	(cd "${DIR}/go" && ${EXEC_GO} test ./... -coverprofile coverage.out -p=1)
	(cd "${DIR}/go" && ${EXEC_GO} tool cover -func=coverage.out)
}

cmd test-web Test Web code
test-web() {
	install-node

	${EXEC_NPM} run test
}

