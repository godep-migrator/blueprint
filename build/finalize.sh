#!/bin/bash --
set -e -u -o pipefail

BASEDIR="/opt/science"
TARGETDIR="${BASEDIR}/blueprint"
CONFDIR="${TARGETDIR}/config"

UPSTARTDIR="/etc/init"
PKGDIR="/tmp/pkg"
NGINXETC="/etc/nginx"

NGINXDIR="${BASEDIR}/nginx"
NGINXCONFDIR="${CONFDIR}/nginx"

mv ${PKGDIR}/deploy ${TARGETDIR}
chmod +x ${TARGETDIR}/bin/*

echo "setting up nginx"
apt-get install -y nginx
rm ${NGINXETC}/sites-*/default
ln -s ${CONFDIR}/nginx.conf ${NGINXETC}/sites-enabled/blueprint
mkdir -p ${NGINXDIR}/{logs,html}
mv ${TARGETDIR}/data/* ${NGINXDIR}/html/

# Setup upstart
mv ${CONFDIR}/upstart.conf ${UPSTARTDIR}/blueprint.conf
