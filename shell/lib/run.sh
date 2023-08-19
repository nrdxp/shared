#!/usr/bin/env bash

cmd run-github Curl GitHub\'s API
run-github () {
	run-vault-secrets-github

	method=""
	if [[ -n ${2+x} ]]; then
		method="-X ${2}"
	fi

	data=()
	if [[ -n ${3+x} ]]; then
		data=(-d "${3}")
	fi

	#shellcheck disable=SC2086
	curl ${method} "${data[@]}" -fsL -H "Accept: application/vnd.github+json" -H "Authorization: Bearer ${GITHUB_TOKEN}" -H "X-GitHub-Api-Version: 2022-11-28" "https://api.github.com${1}"
}

cmd run-github-release-id Get release ID
run-github-release-id () {
	run-github "${GITHUB_PATH}/releases/tags/${BUILD_TAG}" | jq -r .id
}

cmd run-hugo-start Start Hugo
run-hugo-start () {
	install-hugo
	run-network
	run-hugo-stop

	printf "Starting Hugo..."
	try "${CR} run \
		-d \
		${CR_LOGOPTS} \
		${CR_USER} \
		--name candiddev_hugo_${APP_NAME} \
		-p 1313:1313 \
		--restart always \
		-v ${DIR}:/hugo \
		-w /hugo/hugo \
		debian:bullseye \
		../.bin/hugo server --verbose --watch --bind 0.0.0.0 -b /"
}

cmd run-hugo-stop Stop Hugo
run-hugo-stop () {
	printf "Stopping Hugo..."
	try "${CR} rm -f candiddev_hugo_${APP_NAME}"
}

cmd run-network Start network
run-network () {
	printf "Starting network..."
	try "${CR} network create candiddev --subnet 172.31.0.0/24 || true"
}

cmd run-postgresql-start, Run PostgreSQL container
run-postgresql-start () {
	run-network

	printf "Running PostgreSQL container..."
	try "${CR} run \
		-d \
		-e POSTGRES_PASSWORD=postgres \
		-e PGDATA=/pgtmpfs \
		${CR_LOGOPTS} \
		--name candiddev_postgresql \
		--network candiddev \
		-p 127.0.0.1:5432:5432 \
		--restart always \
		--tmpfs=/pgtmpfs \
		-v ${DIR}/containers/initdb.sql:/docker-entrypoint-initdb.d/initdb.sql \
		docker.io/postgres:${VERSION_POSTGRESQL} \
		-c log_statement=all && sleep 2 || true"
}

cmd run-postgresql-backup,rpb Run PostgreSQL backup
run-postgresql-backup () {
	printf "Running PostgreSQL backup..."
	try "${CR} exec -u postgres -i candiddev_postgresql pg_dumpall > ${DIR}/postgresql.sql"
}
rpb () {
	run-postgresql-backup
}

cmd run-postgresql-cli,rpc Run PostgreSQL CLI
run-postgresql-cli () {
	run-postgresql

	# quoting this causes the command to break
	# shellcheck disable=SC2086
	${CR} exec -t ${CR_EXEC_POSTGRESQL}
}
rpc () {
	run-postgresql-cli
}

cmd run-postgresql-restore,rpr Run PostgreSQL restore
run-postgresql-restore () {
	clean-postgresql

	printf "Running PostgreSQL restore..."
	try "${CR} exec ${CR_EXEC_POSTGRESQL} < postgresql.sql"
}
rpr () {
	run-postgresql-restore
}

cmd run-postgresql-stop Stop PostgreSQL container
run-postgresql-stop () {
	printf "Stopping PostgreSQL container..."
	try "${CR} rm -fv candiddev_postgresql || true"
}

cmd run-start Start all containers
run-start () {
	run_cmds run-*-start

	printf "network..."
	try "${CR} network rm candiddev || true"
}

cmd run-stop Stop all running containers
run-stop () {
	run_cmds run-*-stop

	printf "Stopping network..."
	try "${CR} network rm candiddev || true"
}

cmd run-vault-auth-google,rvag Authenticate to Vault using Google
run-vault-auth-google () {
	install-vault

	if ! vault token lookup -format=json | jq .data.policies | grep google_engineering > /dev/null 2>&1; then
		${EXEC_VAULT} login -method=oidc -path=google
		VAULT_TOKEN=$(cat ~/.vault-token 2> /dev/null)
		export VAULT_TOKEN
	fi
}
rvag () {
	run-vault-auth-google
}

cmd run-vault-secrets-github Retreive a GitHub secret from Vault
run-vault-secrets-github () {
	if [[ -z ${GITHUB_TOKEN} ]]; then
		install-vault

		GITHUB_TOKEN=$(${EXEC_VAULT} write -field=token github/token repository_ids="${GITHUB_REPOSITORY_ID}" org_name=candiddev)
	fi
}

cmd run-vault-secrets-kv Retreive a KV secret from Vault
run-vault-secrets-kv () {
	install-vault

	${EXEC_VAULT} read -field="${1}" "${2}"
}

cmd run-yaml8n-start,rys Run YAML8n listeners
run-yaml8n-start () {
	run-network
	run-yaml8n-stop

	for i in "${DIR}"/yaml8n/*; do
		name="$(basename "${i}" | cut -d. -f1)"
		printf "Running YAML8n %s container..." "${name}"

		try "${CR} run \
			-d \
			${CR_LOGOPTS} \
			${CR_USER} \
			--name candiddev_yaml8n_${name} \
			--network candiddev \
			--pull always \
			--restart always \
			-v ${DIR}:/work \
			-w /work \
			${CR_REGISTRY}/candiddev/yaml8n:latest \
			watch /work/yaml8n/${name}.yaml"
	done
}
rys () {
	run-yaml8n-start
}
cmd run-yaml8n-stop Stop YAML8n listeners
run-yaml8n-stop() {
	for i in "${DIR}"/yaml8n/*; do
		${CR} rm -f "candiddev_yaml8n_$(basename "${i}" | cut -d. -f1)" || true
	done
}
