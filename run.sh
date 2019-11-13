#!/bin/sh
#
# About: Build OpenDiablo 2 automatically
# Author: liberodark
# License: GNU GPLv3

version="0.0.4"

echo "OpenDiablo 2 Build Script $version"

#=================================================
# RETRIEVE ARGUMENTS FROM THE MANIFEST AND VAR
#=================================================

distribution=$(cat /etc/*release | grep "PRETTY_NAME" | sed 's/PRETTY_NAME=//g' | sed 's/["]//g' | awk '{print $1}')


go_install(){
  # Check OS & go

  if ! command -v go; then

  	echo "Install Go for OpenDiablo 2 ($distribution)? y/n"
	read -r choice
	[ "$choice" != y ] && [ "$choice" != Y ] && exit

    if [[ "$distribution" = CentOS || "$distribution" = CentOS || "$distribution" = Red\ Hat || "$distribution" = Suse || "$distribution" = Oracle ]]; then
      wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
	  sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz
	  rm go*.linux-amd64.tar.gz
	  export PATH=$PATH:/usr/local/go/bin
	  sudo dnf install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel
      
    elif [[ "$distribution" = Fedora ]]; then
      wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
	  sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz
	  rm go*.linux-amd64.tar.gz
	  export PATH=$PATH:/usr/local/go/bin
	  sudo dnf install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel
    
    elif [[ "$distribution" = Debian || "$distribution" = Ubuntu || "$distribution" = Deepin ]]; then
      apt-get update
      apt-get install -y make autoconf automake gcc libc6 libmcrypt-dev libssl-dev openssl packagekit --force-yes
      
    elif [[ "$distribution" = Gentoo ]]; then
      sudo emerge --ask n go
      
    elif [[ "$distribution" = Manjaro || "$distribution" = Arch\ Linux ]]; then
      sudo pacman -S go --noconfirm

    fi
fi
}

# Build
echo "Check Go"
go_install
echo "Build OpenDiablo 2"
go get
go build
echo "Build finished. Please edit config.json before running OpenDiablo2"
