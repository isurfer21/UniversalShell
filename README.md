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
 - Implemented in pure Go language

### Features
 - [x] archive
 - [x] basename
 - [x] cat
 - [ ] chmod
 - [x] clear
 - [x] cp
 - [x] dir
 - [x] dirname
 - [ ] dirs
 - [x] echo
 - [ ] error
 - [x] exec
 - [x] export
 - [ ] find
 - [x] getenv
 - [ ] grep
 - [x] head
 - [x] help
 - [x] homedir
 - [x] id
 - [ ] ln
 - [x] ls
 - [x] mkdir
 - [x] mv
 - [ ] popd
 - [x] printenv
 - [ ] pushd
 - [x] pwd
 - [x] rm
 - [ ] sed
 - [ ] set
 - [x] setenv
 - [ ] sort
 - [ ] tail
 - [x] tar
 - [x] tmpdir
 - [x] touch
 - [ ] uniq
 - [x] uname
 - [x] unsetenv
 - [x] unrar
 - [x] unzip
 - [x] which
 - [x] whoami
 - [x] xz
 - [x] zip

### References
 - Bash, i.e., Linux Terminal
 - Batch, i.e., Windows Command Prompt
 - [shelljs](https://github.com/shelljs/shelljs)
 - [shx](https://github.com/shelljs/shx)

### Known issues / limitations

##### On Windows
 - Terminal error message shows unformatted characters instead of coloured string. So I would recommend to use light-weight *ConEmu* app in place of default *command-prompt*.

### Setup global config file for ```ush```
By default **ush** expects it's config file at HOME directory; otherwise on every command execution, it will show an error message like this
```
$ ush
Config File ".ush" Not Found in "[/Users/abhishekkumar]"
```

So, you can create it's config file using following command
```
$ ush touch ~/.ush.yml
```

It would create a blank file named ```.ush.yml``` at *HOME* directory. If the path would be correct, then it won't show the error message again.

Also, you can choose the format of this config file by changing the extension to any other extension. The supported formats are "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl".

Please refer [Wiki: Home directory](https://en.wikipedia.org/wiki/Home_directory) to know more about it, based on your operating system.

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
drwxr-xr-x	-	Jul 03, 2018 15:45	.git/
-rw-r--r--	24.0B	Jul 02, 2018 11:40	.gitignore
drwxr-xr-x	-	Jul 03, 2018 15:45	cmd/
-rw-r--r--	1.0KB	Jun 29, 2018 11:50	LICENSE
-rw-r--r--	4.6KB	Jul 03, 2018 15:59	README.md
-rwxr-xr-x	9.6MB	Jul 03, 2018 15:53	ush
-rw-r--r--	170.0B	Jul 02, 2018 11:40	ush.go

```

### For end-users
You can generate an executable for your system via following commands.

#### Installation using ```go get``` 

##### On MacOS / Linux 
```
$ go get -d -u github.com/isurfer21/UniversalShell
$ cd $GOPATH/src/github.com/isurfer21/UniversalShell/
$ go build -o $GOPATH/bin/ush ush.go
```

#### On Windows
```
> go get -d -u github.com/isurfer21/UniversalShell
> cd %GOPATH%\src\github.com\isurfer21\UniversalShell\
> go build -o %GOPATH%\bin\ush.exe ush.go
```

#### Uninstalling, if installed via ```go get```

##### On MacOS / Linux 
```
$ cd $GOPATH/src/github.com/isurfer21/UniversalShell/
$ lua ./uninstall.lua
```

##### On Windows 
```
> cd %GOPATH%\src\github.com\isurfer21\UniversalShell\
> lua uninstall.lua
```

#### Manual installation using Git CLI

##### On MacOS / Linux 
```
$ git clone https://github.com/isurfer21/UniversalShell.git
$ cd UniversalShell
$ go build ush.go
$ ls
LICENSE		README.md	cmd		ush		ush.go
```
After that, put **ush** anywhere on your system and export it's path in *.bashrc* or *.bash_profile*.

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

##### Tip / Advice
The *ldflags* can be used to reduce the binary file size
```
$ go build -ldflags="-s -w" ush.go
```

### For developers
While development, you can compile & execute the source-code using following commands.

##### On MacOS / Linux 
```
$ git clone https://github.com/isurfer21/UniversalShell.git
$ cd UniversalShell
$ go get ./...
$ go run ush.go -h
```
The flag *-h* in the last line can be replaced by any available commands.

##### On Windows
```
> git clone https://github.com/isurfer21/UniversalShell.git
> cd UniversalShell
> go get ./...
> go run ush.go -h
```
The flag *-h* in the last line can be replaced by any available commands.

### Cross-platform shell scripting with Ush using Lua
[Lua](http://www.lua.org/) is a powerful, efficient, lightweight, embeddable scripting language. It supports procedural programming, object-oriented programming, functional programming, data-driven programming, and data description.

Lua is **cross-platform**, since the interpreter is written in ANSI C, and has a relatively simple C API.

So we can use Lua in place of Bash/Batch script. Although you can use any other language of your choice but considering the size of Lua interpreter (less than 1 MB), it looks like the best choice.

Pre-compiled Lua libraries and executables can be downloaded from [LuaBinaries](http://luabinaries.sourceforge.net/download.html). 

Place these files from where it can be globally accessible via terminal.

Now to write a shell script, create a *task.lua* file. Append the below method at the top; after which you can write your shell task.
```lua
-- Prerequisite method
function shell(command)
	local handle = io.popen(command)
	local result = handle:read("*a")
	handle:close()
	return result
end

-- Start you shell task from here
print(shell('ush ls -l'))
```

We can modify the above code for Ush only, as given below
```lua
-- Prerequisite method
function ush(command)
	local handle = io.popen('ush ' .. command)
	local result = handle:read("*a")
	handle:close()
	return result
end

-- Start you shell task from here
print(ush('ls -l'))
```

Alternatively, you can use [luash](https://github.com/zserge/luash) or [lit-sh](https://github.com/james2doyle/lit-sh) as well.



