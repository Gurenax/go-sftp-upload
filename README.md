# Go SFTP Upload
Uploads files to your SFTP server using `scp (Secure copy)` and `sshpass (SSH password provider)`

<br/>

## Use Case
Sometimes we have to work on old school pipelines where we can only use `sftp` to update source files instead of the modern `git` pipelines that we are used to. <br/>
Such an example is uploading wordpress files in which some providers do not provide you access to ssh authentication via `rsa keys`.<br/>
My method will provide you with an alternative to simply use a `.env` file where you can store your credentials.

<br/><br/>

## Install sshpass (Required)
`sshpass` will allow us to execute `scp` without being prompted for a password.
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
USERNAME=
PASSWORD=
SERVER=
SRC=
DEST=
```
`SRC` is the local directory/files you wish to upload.<br/>
`DEST` is the remote directory where you wish to upload the local directory/files.<br/>
`USERNAME, PASSWORD and SERVER` should be self-explanatory once you have your sftp credentials.
<br/><br/>

## Build and Run the App
```
go build
./go-sftp-upload
```

