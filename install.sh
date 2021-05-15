#!/usr/bin/env bash

CONFIG_DIR="~/.arss"
DEFAULT_CLIENT_DIR="$CONFIG_DIR/clients/default"

function info ()
{
    echo "[info]: $1"
}

function success ()
{
    echo "done!"
    echo ""
}

function error ()
{
    echo "[error]: $1"
}

function error_exit ()
{
    echo ""
    error "INSTALLATION FAILED"
    error "see the log above for information on which step failed"
    echo "For usage details, type ./install.sh --help"
    exit 1
}

function help_text ()
{
    echo "install.sh - Installation script for arss"
    echo ""
    echo "Usage:"
    echo "  ./install.sh [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -p, --path <path/to/installation/directory>"
    echo "      Which directory to install arss in?"
    echo "  -c, --only-client"
    echo "      Only install client"
    echo "  -s, --only-server"
    echo "      Only install server"
    echo "  -h, --help"
    echo "      Prints help information"
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
    mkdir -p "$DEFAULT_CLIENT_DIR" &&
    cp -R client/public/* "$DEFAULT_CLIENT_DIR" &&
    info "installed client at $DEFAULT_CLIENT_DIR" &&
    success
}

function install_server () 
{
    info "installing server" &&
    mv ./arss "$1/arss" &&
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
    if [ "$only_client" = true ]; then
        install_client 
    fi

    if [ "$only_server" = true ]; then
        install_server "$server_path"
    fi
    
    if [ "$only_client" = false ] && [ "$only_server" = false ]; then
        install_client &&
        install_server "$server_path"
    fi
}

function parse_args ()
{
    only_client=false
    only_server=false

    for arg in "$@"; do
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
            -h|--help)
            help_text
            exit 0
            ;;
        esac
    done

    (build && install "$@") || error_exit "error occurred"
}

parse_args "$@"
