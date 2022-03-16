[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

## _EZPasswordSpray_

EZPasswordSpray is a spraying tool that takes a user and a password file and checks if there are valid Outlook credentials in it.

## Quick Start
There are 2 ways of getting started:

*Case 1 - Download the executable:*\
Download the latest release from [Downloads]. After extracting the folder you will have the executable to start the program and you are ready to go. 

*Case 2 - Download the source code and compile it:*\
After Downloading the source code from Github you will have to build a new executable file.

Depending on your target hardware and operating system you have to set the go environment variables:
```sh
set GOOS=your OS
set GOARCH=your CPU architecture
```

with the following command you receive the list of possible variables:
```sh
go tool dist list
```

then navigate to the project folder within your terminal and type the following command to build the executable:
```sh
go get github.com/schollz/progressbar
go build
```

After building, you can start the executable by double-clicking. The program needs paramters to start properly.

## Parameter Description

<code>-u</code> The users parameter is a path to the text file with the User Principal Names. In each line should be one User Principal Name.  
<code>-p</code> The passwords parameter is a path to the text file with the passwords. In each line should be one password.

[Downloads]:https://github.com/NSIDE-ATTACK-LOGIC/EZPasswordSpray/releases/tag/V1.0_Windows
