#!/bin/sh

REPO="${1:-stable}"
URL=https://dl.bintray.com/emccode/dvdcli/$REPO/latest
OS=$(uname -s)
ARCH=$(uname -m)
SUDO=$(which sudo)
BIN_DIR=/usr/bin
BIN_FILE=$BIN_DIR/dvdcli

function sudo() {
    if [[ $(id -u) -eq 0 ]]; then $@; else $SUDO $@; fi
}

function is_coreos() {
    grep -q DISTRIB_ID=CoreOS /etc/lsb-release &> /dev/null
    if [[ $? -eq 0 ]]; then echo 0; else echo 1; fi
}

IS_COREOS=$(is_coreos)

# how to detect the linux distro was taken from http://bit.ly/1JkNwWx
if [[ -e /etc/redhat-release || \
      -e /etc/redhat-version ]]; then

    #echo "installing rpm"
    sudo rpm -ih --quiet $URL/dvdcli-latest-$ARCH.rpm > /dev/null

elif [[ $ARCH = x86_64 && \
       (-e /etc/debian-release || \
        -e /etc/debian-version || \
        -e /etc/lsb-release) &&
        $IS_COREOS -eq 1 ]]; then

    #echo "installing deb"
    curl -sSLO $URL/dvdcli-latest-$ARCH.deb && \
        sudo dpkg -i dvdcli-latest-$ARCH.deb && \
        rm -f dvdcli-latest-$ARCH.deb

else
    if [[ $IS_COREOS -eq 0 ]]; then
        BIN_DIR=/opt/bin
        BIN_FILE=$BIN_DIR/dvdcli
    elif [[ $OS = Darwin ]]; then
        BIN_DIR=/usr/local/bin
        BIN_FILE=$BIN_DIR/dvdcli
    fi

    sudo mkdir -p $BIN_DIR && \
      curl -sSLO $URL/dvdcli-$OS-$ARCH.tar.gz && \
      sudo tar xzf dvdcli-$OS-$ARCH.tar.gz -C $BIN_DIR && \
      rm -f dvdcli-$OS-$ARCH.tar.gz && \
      sudo chmod 0755 $BIN_FILE && \
      sudo chown 0 $BIN_FILE && \
      sudo chgrp 0 $BIN_FILE
fi

echo
echo "dvdcli has been installed to $BIN_FILE"
echo
$BIN_FILE version
echo
