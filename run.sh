#!/bin/sh
#
# About: Build OpenDiablo 2 automatically
# Author: liberodark
# License: GNU GPLv3
set -eu

version="0.0.2"

echo "OpenDiablo 2 Build Script $version"

#=================================================
# RETRIEVE ARGUMENTS FROM THE MANIFEST AND VAR
#=================================================

distro=$(sed -n 's#PRETTY_NAME="\([^"]*\)"#\1#p' /etc/os-release)

# Check OS & go

if ! command -v go >/dev/null
then
	echo "Install Go for OpenDiablo 2 ($distro)? y/n"
	read -r choice
	[ "$choice" != y ] && [ "$choice" != Y ] && exit

	case "$(echo "$distro" | tr '[:upper:]' '[:lower:]')" in
		centos|red\ hat|suse|oracle)
			sudo yum install -y go >/dev/null 2>&1
			;;
		fedora)
			sudo dnf install -y go >/dev/null 2>&1
			;;
		debian|ubuntu|deepin)
			sudo apt-get update >/dev/null 2>&1
			sudo apt-get install -y go --force-yes >/dev/null 2>&1
			;;
		manjaro|arch\ linux)
			sudo pacman -S go --noconfirm >/dev/null 2>&1
			;;
		gentoo/linux)
			sudo emerge --ask n go /dev/null 2>&1
			;;
		*)
			echo "$distro: unsupported distribution, please install Go by yourself"
			exit 1
			;;
	esac
fi

# Build
echo "Build OpenDiablo 2"
go get
go build
chmod u+x OpenDiablo2
echo "Build finished. Please edit config.json before running OpenDiablo2"
