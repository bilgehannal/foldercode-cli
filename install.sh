echo 'Downloading and installing fcode executable file...'

BASE_RELEASES_URL="https://github.com/bilgehannal/foldercode-cli/releases/download"
VERSION=$(wget -q -O - https://api.github.com/repos/bilgehannal/foldercode-cli/releases/latest|grep '"tag_name": '|cut -d 'v' -f 2|cut -d '"' -f 1)
DISTRO="$1"
ARCH="$2"
TAR_NAME="foldercode-cli_${VERSION}_${DISTRO}_${ARCH}.tar.gz"
wget "$BASE_RELEASES_URL/v${VERSION}/${TAR_NAME}"
tar -xzvf $TAR_NAME fcode
chmod +x fcode
sudo mv fcode /usr/local/bin/fcode
rm $TAR_NAME

echo 'fcode is installed.'
echo 'Run: fcode --help'