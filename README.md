# Ush - Universal Shell 
Copyright (c) 2018 [Abhishek Kumar](https://github.com/isurfer21).
It is licensed under the **MIT License**.

> Shell commands that works similar everywhere

##### Have you ever written a script in bash for Mac or Linux and then tried to port the same for Windows? 

Well, if you did then you might have noticed the sheer frustation because the bash commands or scripts are incompatible on Windows command-line. So we can't use them straight forward.

##### Then, what options do we have?

Either we can use Unix shells like Msys2, Cygwin, Babun, etc. or install Ubuntu which is an experimental feature available only on Windows 10. Alternatively, if your system have high RAM then you can use VirtualBox, Virtual Machine, Docker, etc. as well.

##### What challenges do we face with them?

Out of the available options my personal preference is Msys2 because it can generate windows compatible executables. Next would be Ubuntu because it can access Windows file system just like it's own. Rest of them provides their own sandboxed environment which is less suitable.

But one thing is common in all of them that they are so heavy. Pretty basic installation will consume more than 300 MB of your storage space. So just for running a script, it would be too much.

Thus, I have started this project out of frustation of not having any light-weight tool that runs the same sets of commands on all platforms.

### Objective
To create a cross-platform application as a portable executable that could run built-in shell commands in a same way on native terminal of all platforms including Windows, MacOS, Linux.

### Check-list
 - Cross-platform application
 - Portable executable program
 - CLI based tool
 - Built-in shell commands
 - Compatible to Windows, MacOS, Linux
 - Required no external dependency

### Known issues / limitations

##### On Windows
 - Terminal error message shows unformatted characters instead of coloured string. So I would recommend to use light-weight *ConEmu* app in place of default *command-prompt*.

### Sample usage

##### Command: pwd
```
$ ush pwd
/Users/abhishekkumar/Documents/UniversalShell
```

##### Command: ls
```
$ ush ls
.git  .gitignore  bin  LICENSE  README.md  src

$ ush ls -y
.git
.gitignore
bin
LICENSE
README.md
src

$ ush ls -l
drwxr-xr-x	544.0B	Jun 30, 2018 02:56	Dir	.git
-rw-r--r--	4.0B	Jun 29, 2018 10:15		.gitignore
drwxr-xr-x	102.0B	Jun 30, 2018 01:56	Dir	bin
-rw-r--r--	1.0KB	Jun 28, 2018 22:16		LICENSE
-rw-r--r--	275.0B	Jun 30, 2018 01:53		README.md
drwxr-xr-x	102.0B	Jun 29, 2018 22:25	Dir	src

```

### For end-users
You can generate an executable for your system via following commands.

##### On MacOS / Linux 
```
$ git clone https://github.com/isurfer21/UniversalShell.git
$ cd UniversalShell
$ go build ./ush.go
$ ls
LICENSE		README.md	cmd		ush		ush.go
```
After that, put **ush** anywhere on your system and export it's path in *.bashrc* or *.bash_profile*.

##### On MacOS / Linux via Go package manager
```
go get -d github.com/isurfer21/UniversalShell
cd $GOPATH/src/github.com/isurfer21/UniversalShell/
go build -o $GOPATH/bin/ush ./ush.go
```

##### On Windows
```
> git clone https://github.com/isurfer21/UniversalShell.git
> cd UniversalShell
> go build ush.go
> dir /b
LICENSE
README.md
cmd
ush.exe
ush.go
```
After that, put **ush.exe** anywhere on your system and set it's path in *environment variable*.

### For developers
While development, you can compile & execute the source-code using following commands.

##### On MacOS / Linux 
```
$ git clone https://github.com/isurfer21/UniversalShell.git
$ cd UniversalShell
$ go run ./ush.go -h
```
The flag *-h* in the last line can be replaced by any available commands.

##### On Windows
```
$ git clone https://github.com/isurfer21/UniversalShell.git
$ cd UniversalShell
$ go run ush.go -h
```
The flag *-h* in the last line can be replaced by any available commands.
