# UniversalShell
Shell commands that works similar everywhere

### Sample

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

### Build

```
$ cd UniversalShell

$ pwd
~/UniversalShell

$ go build -o ./bin/ush ./src/ush.go

$ ls ./bin
ush
```

### Development

```
$ cd UniversalShell

$ pwd
~/UniversalShell

$ go run ./src/ush.go -h
```