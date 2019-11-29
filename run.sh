#!/bin/bash
#
# About: Build OpenDiablo 2 automatically
# Author: liberodark
# License: GNU GPLv3

version="0.0.7"

echo "OpenDiablo 2 Build Script $version"

#=================================================
# RETRIEVE ARGUMENTS FROM THE MANIFEST AND VAR
#=================================================
export PATH=$PATH:/usr/local/go/bin

distribution=$(cat /etc/*release | grep "PRETTY_NAME" | sed 's/PRETTY_NAME=//g' | sed 's/["]//g' | awk '{print $1}')


go_install(){
  # Check OS & go

  if ! command -v go > /dev/null 2>&1; then

    echo "Install Go for OpenDiablo 2 ($distribution)? y/n"
	    read -r choice
	    [ "$choice" != y ] && [ "$choice" != Y ] && exit

    if [ "$distribution" = "CentOS" ] || [ "$distribution" = "Red\ Hat" ] || [ "$distribution" = "Oracle" ]; then
      echo "Downloading Go"
      	wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Install Go"
	  	sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Clean unless files"
	  	rm go*.linux-amd64.tar.gz
      echo "Install libraries"
	    sudo yum install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel alsa-lib-devel libXi-devel > /dev/null 2>&1
      
    elif [ "$distribution" = "Fedora" ]; then
      echo "Downloading Go"
      	wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Install Go"
	    sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Clean unless files"
	    rm go*.linux-amd64.tar.gz
      echo "Install libraries"
	    sudo dnf install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel alsa-lib-devel libXi-devel > /dev/null 2>&1
    
    elif [ "$distribution" = "Debian" ] || [ "$distribution" = "Ubuntu" ] || [ "$distribution" = "Deepin" ]; then
      echo "Downloading Go"
      	wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Install Go"
	    sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Clean unless files"
	    rm go*.linux-amd64.tar.gz
      echo "Install libraries"
	    sudo apt-get install -y libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libsdl2-dev libasound2-dev > /dev/null 2>&1
      
    elif [ "$distribution" = "Gentoo" ]; then
      sudo emerge --ask n go libXcursor libXrandr libXinerama libXi libGLw libglvnd libsdl2 alsa-lib
      
    elif [ "$distribution" = "Manjaro" ] || [ "$distribution" = "Arch\ Linux" ]; then
      sudo pacman -S go libxcursor libxrandr libxinerama libxi mesa libglvnd sdl2 sdl2_mixer sdl2_net alsa-lib --noconfirm
	  
	elif [ "$distribution" = "OpenSUSE" ] || [ "$distribution" = "SUSE" ]; then
	  echo "Downloading Go"
      	wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Install Go"
	    sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz > /dev/null 2>&1
      echo "Clean unless files"
	    rm go*.linux-amd64.tar.gz
      echo "Install libraries"
      sudo zypper install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel Mesa-libGL-devel alsa-lib-devel libXi-devel > /dev/null 2>&1

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
