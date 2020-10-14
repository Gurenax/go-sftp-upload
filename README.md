# Go SFTP Upload
Uploads files to your SFTP server using SCP and SSHPASS

<br/>

## Install sshpass (Required)
### Mac OS
```
brew install esolitos/ipa/sshpass
```
### Ubuntu
```
apt-get install sshpass
```

<br/><br/>

## Setup Environment Variables
Your `.env` file should contain values for
```
PASSWORD=
SRC=
USERNAME=
SERVER=
DEST=
```

<br/><br/>

## Running the App
```
go build
./go-sftp-upload
```

