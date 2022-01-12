#!/usr/bin/sh

download=true
do_start=true

show_help() {
    # Shows help dialog
    echo -e "Downloads data and starts server
Usage: $0 [ARG1] [ARG2] ...
Options:
	--no-download\t: Does not download the data
	--update\t: Just downloads the data
	--help\t\t: Shows this dialog"
}

for i in $@
do
    case $i in 
	"--no-download")
	    download=false
	    ;;
	"--update")
	    do_start=false
	    ;;
	"--help" | *)
	    show_help
	    exit
	    ;;
    esac
done

if [[ $download = true ]]
then
    . ./download.sh
fi

if [[ $do_start = true ]]
then
    echo ":: Starting..."
    go run server.go
fi
