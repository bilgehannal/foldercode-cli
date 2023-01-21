echo 'Downloading fcode executable file...'
sudo curl https://github.com/bilgehannal/foldercode-cli/releases/download/v0.0.1/fcode -o /usr/local/bin/fcode
echo 'Checking permissions...'
sudo chmod +x /usr/local/bin/fcode
echo 'fcode is installed.'
echo 'Run: fcode --help'