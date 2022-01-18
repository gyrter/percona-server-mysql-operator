#!/bin/bash

set -e
set -o xtrace

XTRABACKUP_USER=${XTRABACKUP_USER:-xtrabackup}
XTRABACKUP_PASSWORD=$(</etc/mysql/mysql-users-secret/$XTRABACKUP_USER)

main() {
	until /var/lib/mysql/healthcheck replication; do
		echo "waiting for the replication to become active"
		sleep 5
	done

	xtrabackup --backup=1 --user=${XTRABACKUP_USER} --password=${XTRABACKUP_PASSWORD}

	pkill -e -15 mysqld
}

main
