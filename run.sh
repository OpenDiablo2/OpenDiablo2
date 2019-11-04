#!/bin/bash
#
# About: Build OpenDiablo 2 automatically
# Author: liberodark
# License: GNU GPLv3

version="0.0.1"

echo "OpenDiablo 2 Build Script $version"


#=================================================
# RETRIEVE ARGUMENTS FROM THE MANIFEST AND VAR
#=================================================

distribution=$(cat /etc/*release | grep "PRETTY_NAME" | sed 's/PRETTY_NAME=//g' | sed 's/["]//g' | awk '{print $1}')

compile_od2(){
	  echo "Build OpenDiablo 2"
      go get
	  go build
	  chmod +x OpenDiablo2
	  echo "Build finished. Please edit config.json before running OpenDiablo2"
      }

go_run(){
echo "Install Go for OpenDiablo 2 ($distribution)"

  # Check OS & go

  if ! command -v go &> /dev/null; then

    if [[ "$distribution" = CentOS || "$distribution" = CentOS || "$distribution" = Red\ Hat || "$distribution" = Suse || "$distribution" = Oracle ]]; then
      sudo yum install -y go &> /dev/null

      compile_od2 || exit
      
    elif [[ "$distribution" = Fedora ]]; then
      sudo dnf install -y go &> /dev/null
    
      compile_od2 || exit
    
    elif [[ "$distribution" = Debian || "$distribution" = Ubuntu || "$distribution" = Deepin ]]; then
      sudo apt-get update &> /dev/null
      sudo apt-get install -y go --force-yes &> /dev/null
    
      compile_od2 || exit
      
    elif [[ "$distribution" = Manjaro || "$distribution" = Arch\ Linux ]]; then
      sudo pacman -S go --noconfirm &> /dev/null
    
      compile_od2 || exit

    fi
fi
}

go_run
