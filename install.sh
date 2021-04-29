#!/usr/bin/env bash

CONFIG_DIR="$HOME/.arss"
DEFAULT_CLIENT_DIR="$CONFIG_DIR/clients/default"

function info ()
{
    echo "info: $1"
}

function success ()
{
    echo "done!"
    echo ""
}

function build_client () 
{
    info "building client" &&
    prefix="--prefix client" &&
    npm $prefix install &&
    npm $prefix run build &&
    success
}

function build_server () 
{
    info "building server" &&
    go build arss && success
}

function install_client () 
{
    info "installing client" &&
    mkdir -p $DEFAULT_CLIENT_DIR &&
    cp -R client/public/* $DEFAULT_CLIENT_DIR &&
    info "installed client at $DEFAULT_CLIENT_DIR" &&
    success
}

function install_server () 
{
    info "installing server" &&
    mv ./arss $1/arss &&
    info "installed server at $1" &&
    success
}

function build () 
{
    build_client &&
    build_server
}

function install () 
{
    only_client=false
    only_server=false

    for arg in $@; do
        case $arg in
            -p|--path)
            server_path=$2
            shift
            shift
            ;;
            -c|--only-client)
            only_client=true
            shift
            ;;
            -s|--only-server)
            only_server=true
            shift
            ;;
        esac
    done
    
    if [ $only_client = true ]; then
        install_client 
    fi

    if [ $only_server = true ]; then
        install_server $server_path
    fi
    
    if [ $only_client = false ] && [ $only_server = false ]; then
        install_client &&
        install_server $server_path
    fi
}

build && install $@
