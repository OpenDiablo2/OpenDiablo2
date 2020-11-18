#!/bin/bash
#
# About: Build OpenDiablo 2 automatically
# Author: liberodark
# License: GNU GPLv3

version="0.0.8"
go_version="1.13.4"
echo "OpenDiablo 2 Build Script $version"

#=================================================
# RETRIEVE ARGUMENTS FROM THE MANIFEST AND VAR
#=================================================
export PATH=$PATH:/usr/local/go/bin

distribution=$(cat /etc/*release | grep "PRETTY_NAME" | sed 's/PRETTY_NAME=//g' | sed 's/["]//g' | awk '{print $1}')

go_install() {
	# Check OS & go

	if ! command -v go >/dev/null 2>&1; then

		echo "Install Go for OpenDiablo 2 ($distribution)? y/n"
		read -r choice
		[ "$choice" != y ] && [ "$choice" != Y ] && exit

		if [ "$distribution" = "CentOS" ] || [ "$distribution" = "Red\ Hat" ] || [ "$distribution" = "Oracle" ]; then
			echo "Downloading Go"
			wget https://dl.google.com/go/go"$go_version".linux-amd64.tar.gz >/dev/null 2>&1
			echo "Install Go"
			sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz >/dev/null 2>&1
			echo "Clean unneeded files"
			rm go*.linux-amd64.tar.gz

		elif [ "$distribution" = "Fedora" ]; then
			echo "Downloading Go"
			wget https://dl.google.com/go/go"$go_version".linux-amd64.tar.gz >/dev/null 2>&1
			echo "Install Go"
			sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz >/dev/null 2>&1
			echo "Clean unneeded files"
			rm go*.linux-amd64.tar.gz

		elif [ "$distribution" = "Debian" ] || [ "$distribution" = "Ubuntu" ] || [ "$distribution" = "Deepin" ]; then
			echo "Downloading Go"
			wget https://dl.google.com/go/go"$go_version".linux-amd64.tar.gz >/dev/null 2>&1
			echo "Install Go"
			sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz >/dev/null 2>&1
			echo "Clean unneeded files"
			rm go*.linux-amd64.tar.gz

		elif [ "$distribution" = "Gentoo" ]; then
			sudo emerge --ask n go

		elif [ "$distribution" = "Manjaro" ] || [ "$distribution" = "Arch\ Linux" ]; then
			sudo pacman -S go --noconfirm

		elif [ "$distribution" = "OpenSUSE" ] || [ "$distribution" = "SUSE" ]; then
			echo "Downloading Go"
			wget https://dl.google.com/go/go"$go_version".linux-amd64.tar.gz >/dev/null 2>&1
			echo "Install Go"
			sudo tar -C /usr/local -xzf go*.linux-amd64.tar.gz >/dev/null 2>&1
			echo "Clean unneeded files"
			rm go*.linux-amd64.tar.gz

		fi
	fi
}

dep_install() {
	if [ "$distribution" = "CentOS" ] || [ "$distribution" = "Red\ Hat" ] || [ "$distribution" = "Oracle" ]; then
		sudo yum install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel alsa-lib-devel libXi-devel >/dev/null 2>&1

	elif [ "$distribution" = "Fedora" ]; then
		sudo dnf install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel alsa-lib-devel libXi-devel >/dev/null 2>&1

	elif [ "$distribution" = "Debian" ] || [ "$distribution" = "Ubuntu" ] || [ "$distribution" = "Deepin" ]; then
		sudo apt-get install -y libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libgl1-mesa-dev libsdl2-dev libasound2-dev >/dev/null 2>&1

	elif [ "$distribution" = "Gentoo" ]; then
		sudo emerge --ask n libXcursor libXrandr libXinerama libXi libGLw libglvnd libsdl2 alsa-lib >/dev/null 2>&1

	elif [ "$distribution" = "Manjaro" ] || [ "$distribution" = "Arch\ Linux" ]; then
		mesa_detect_arch=$(pacman -Q | grep mesa)

		if [ -z "$mesa_detect_arch" ]; then
			sudo pacman -S libxcursor libxrandr libxinerama libxi mesa libglvnd sdl2 sdl2_mixer sdl2_net alsa-lib --noconfirm >/dev/null 2>&1
		else
			sudo pacman -S libxcursor libxrandr libxinerama libxi libglvnd sdl2 sdl2_mixer sdl2_net alsa-lib --noconfirm >/dev/null 2>&1
		fi

	elif [ "$distribution" = "OpenSUSE" ] || [ "$distribution" = "SUSE" ]; then
		sudo zypper install -y libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel Mesa-libGL-devel alsa-lib-devel libXi-devel >/dev/null 2>&1

	elif [[ "$OSTYPE" == "darwin"* ]]; then
		# are there dependencies required? did I just have all of them already?
		echo "Mac OS detected, no dependency installation necessary..."

	fi
}

# Build
echo "Check Go"
go_install

echo "Install libraries"
if [ ! -e "$HOME/.config/OpenDiablo2" ]; then
	mkdir -p $HOME/.config/OpenDiablo2
fi

if [ -e "$HOME/.config/OpenDiablo2/.libs" ]; then
	echo "libraries is installed"
else
	echo "OK" >"$HOME/.config/OpenDiablo2/.libs"
	dep_install
fi

echo "Build OpenDiablo 2"
go get -d
go build

echo "Build finished. Running OpenDiablo2 will generate a config.json file."
echo "If there are subsequent errors, please inspect and edit the config.json file. See doc/index.html for more details"
