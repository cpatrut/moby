#!/usr/bin/env bash
set -e

# This file is used to auto-generate Dockerfiles for making debs via 'make deb'
#
# usage: ./generate.sh [versions]
#    ie: ./generate.sh
#        to update all Dockerfiles in this directory
#    or: ./generate.sh ubuntu-trusty
#        to only update ubuntu-trusty/Dockerfile
#    or: ./generate.sh ubuntu-newversion
#        to create a new folder and a Dockerfile within it
#
# Note: non-LTS versions are not guaranteed to work.

cd "$(dirname "$(readlink -f "$BASH_SOURCE")")"

versions=( "$@" )
if [ ${#versions[@]} -eq 0 ]; then
	versions=( */ )
fi
versions=( "${versions[@]%/}" )

for version in "${versions[@]}"; do
	echo "${versions[@]}"
	distro="${version%-*}"
	suite="${version##*-}"
	from="aarch64/${distro}:${suite}"

	mkdir -p "$version"
	echo "$version -> FROM $from"
	cat > "$version/Dockerfile" <<-EOF
		#
		# THIS FILE IS AUTOGENERATED; SEE "contrib/builder/deb/aarch64/generate.sh"!
		#

		FROM $from

	EOF

	dockerBuildTags='apparmor pkcs11 selinux'
	runcBuildTags='apparmor selinux'

	# this list is sorted alphabetically; please keep it that way
	packages=(
		apparmor # for apparmor_parser for testing the profile
		bash-completion # for bash-completion debhelper integration
		btrfs-tools # for "btrfs/ioctl.h" (and "version.h" if possible)
		build-essential # "essential for building Debian packages"
		cmake # tini dep
		curl ca-certificates # for downloading Go
		debhelper # for easy ".deb" building
		dh-apparmor # for apparmor debhelper
		dh-systemd # for systemd debhelper integration
		git # for "git commit" info in "docker -v"
		libapparmor-dev # for "sys/apparmor.h"
		libdevmapper-dev # for "libdevmapper.h"
		libltdl-dev # for pkcs11 "ltdl.h"
		pkg-config # for detecting things like libsystemd-journal dynamically
		vim-common # tini dep
	)

	case "$suite" in
		trusty)
			packages+=( libsystemd-journal-dev )
			# aarch64 doesn't have an official downloadable binary for go.
			# And gccgo for trusty only includes Go 1.2 implementation which
			# is too old to build current go source, fortunately trusty has
			# golang-1.6-go package can be used as bootstrap.
			packages+=( golang-1.6-go )
			;;
		jessie)
			packages+=( libsystemd-journal-dev )
			# aarch64 doesn't have an official downloadable binary for go.
			# And gccgo for jessie only includes Go 1.2 implementation which
			# is too old to build current go source, fortunately jessie backports
			# has golang-1.6-go package can be used as bootstrap.
			packages+=( golang-1.6-go libseccomp-dev )

			dockerBuildTags="$dockerBuildTags seccomp"
			runcBuildTags="$runcBuildTags seccomp"
			;;
		stretch|xenial)
			packages+=( libsystemd-dev )
			packages+=( golang-go libseccomp-dev )

			dockerBuildTags="$dockerBuildTags seccomp"
			runcBuildTags="$runcBuildTags seccomp"
			;;
		*)
			echo "Unsupported distro:" $distro:$suite
			rm -fr "$version"
			exit 1
			;;
	esac

	case "$suite" in
		jessie)
			echo 'RUN echo deb http://ftp.debian.org/debian jessie-backports main > /etc/apt/sources.list.d/backports.list' >> "$version/Dockerfile"
			;;
		*)
			;;
	esac

	# update and install packages
	echo "RUN apt-get update && apt-get install -y ${packages[*]} --no-install-recommends && rm -rf /var/lib/apt/lists/*" >> "$version/Dockerfile"
	echo >> "$version/Dockerfile"

	case "$suite" in
		jessie|trusty)
			echo 'RUN update-alternatives --install /usr/bin/go go /usr/lib/go-1.6/bin/go 100' >> "$version/Dockerfile"
			echo >> "$version/Dockerfile"
			;;
		*)
			;;
	esac

	echo "# Install Go" >> "$version/Dockerfile"
	echo "# aarch64 doesn't have official go binaries, so use the version of go installed from" >> "$version/Dockerfile"
	echo "# the image to build go from source." >> "$version/Dockerfile"

	awk '$1 == "ENV" && $2 == "GO_VERSION" { print; exit }' ../../../../Dockerfile.aarch64 >> "$version/Dockerfile"
	echo 'RUN mkdir /usr/src/go && curl -fsSL https://golang.org/dl/go${GO_VERSION}.src.tar.gz | tar -v -C /usr/src/go -xz --strip-components=1 \' >> "$version/Dockerfile"
	echo '	&& cd /usr/src/go/src \' >> "$version/Dockerfile"
	echo '	&& GOOS=linux GOARCH=arm64 GOROOT_BOOTSTRAP="$(go env GOROOT)" ./make.bash' >> "$version/Dockerfile"
	echo >> "$version/Dockerfile"

	echo 'ENV PATH /usr/src/go/bin:$PATH' >> "$version/Dockerfile"
	echo >> "$version/Dockerfile"

	echo "ENV AUTO_GOPATH 1" >> "$version/Dockerfile"
	echo >> "$version/Dockerfile"

	echo "ENV DOCKER_BUILDTAGS $dockerBuildTags" >> "$version/Dockerfile"
	echo "ENV RUNC_BUILDTAGS $runcBuildTags" >> "$version/Dockerfile"
done
