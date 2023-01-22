echo 'Downloading and installing fcode executable file...'

BASE_RELEASES_URL="https://github.com/bilgehannal/foldercode-cli/releases/download"
VERSION="$1"
DISTRO="$2"
ARCH="$3"
TAR_NAME="foldercode-cli_${VERSION}_${DISTRO}_${ARCH}.tar.gz"
wget "$BASE_RELEASES_URL/v${VERSION}/${TAR_NAME}"
tar -xzvf $TAR_NAME fcode
chmod +x fcode
sudo mv fcode /usr/local/bin/fcode
rm $TAR_NAME

echo 'fcode is installed.'
echo 'Run: fcode --help'