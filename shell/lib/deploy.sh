#!/usr/bin/env bash

# shellcheck disable=SC2317
cmd deploy,d Deploy to environments
deploy () {
	if [ -n "${VAULT_SSH_ROLE}" ]; then
		run-vault-ssh
	fi

	chmod 0600 "${DIR}/id_rsa"

	printf "Deploying %s using %s..." "${APP_NAME}" "${VAULT_SSH_ROLE}"
	# shellcheck disable=SC2016
	try 'for host in ${DEPLOY_HOSTS}; do
		ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i ${DIR}/id_rsa -p ${host#*:} root@${host%:*}
	done'

	run deploy-post
}
d () {
	deploy
}
