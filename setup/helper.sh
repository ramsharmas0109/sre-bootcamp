#! /bin/bash

set -x
#write two function to install docker and docker-compose
#also detect os version as well before installing

detect_os() {
    OS=$(uname -s)
    if [[ $OS == "Darwin" ]]; then
        install_docker_mac
    fi

    if [[ $OS == "Linux" ]];then
        install_docker_linux
    fi
}

install_docker_linux() {
    sudo apt update -y;
    sudo apt install docker.io docker-compose-v2 -y;
}

install_docker_mac() {
    brew install docker;
    brew install docker-compose;
}

detect_os
