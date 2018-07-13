--[[
File: builder.lua

Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
]]

require 'UniversalShell'

-- CORE

S = UniversalShell.shFmtExecPrnt
s = UniversalShell.shFmtExec
U = UniversalShell.ushFmtExecPrnt
u = UniversalShell.ushFmtExec

-- TASKS

function version()
  print([[
Builder   version 1.0
Copyright (c) 2018 Abhishek Kumar. All rights reserved.
  ]])
end

function setup()
  print("Setup all the go dependencies at GOPATH")
  S('go get github.com/spf13/cobra')
  S('go get github.com/spf13/viper')
  S('go get github.com/mitchellh/go-homedir')
  S('go get github.com/inhies/go-bytesize')
  S('go get github.com/mholt/archiver')
  print("Done!")
end

function build()
  print("Build for current machine only")
  S('go build ush.go')
  print("Done!")
end

function xbuild()
  print("Cross build for MacOSX, Windows & Linux as 32bit & 64bit")
  S('gox -osarch="darwin/386" -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  S('gox -osarch="darwin/amd64" -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  S('gox -osarch="windows/386" -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  S('gox -osarch="windows/amd64" -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  S('gox -osarch="linux/386" -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  S('gox -osarch="linux/amd64" -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  print("Done!")
end

function xbuilds()
  print("Cross build for all OS and Arch")
  S('gox -output="./bin/ush_{{.OS}}_{{.Arch}}"')
  print("Done!")
end

function release()
  print("# Create distributable packages")

  local rootDir = u('pwd')
  local srcDir = rootDir
  local binDir = rootDir..'/bin'
  local pubDir = rootDir..'/pub'

  if u('ls -e "%s"', binDir) ~= 'true' then
    U('mkdir "%s"', binDir)
  end

  if u('ls -e "%s"', pubDir) ~= 'true' then
    U('mkdir "%s"', pubDir)
  end

  print("Copying icons")
  U('cp "%s/img/appicon.ico" "%s/appicon.ico"', srcDir, binDir)
  U('cp "%s/img/appicon.icns" "%s/appicon.icns"', srcDir, binDir)
  
  print("Copying doc files")
  U('cp "%s/LICENSE" "%s/LICENSE"', rootDir, binDir)
  U('cp "%s/README.md" "%s/README.md"', rootDir, binDir)

  print("# Create compressed binary distributable files")

  if contains(arg, '-win') then 
    print('Publishing release for Windows')
    if u('ls -e "%s/Ush_windows_x86-64.zip"', pubDir) == 'true' then
      print('Delete last release build for Windows')
      U('rm "%s/Ush_windows_x86-64.zip"', pubDir)
    end 
    local tmpDir = pubDir.."/Ush_windows_x86-64"
    U('mkdir "%s"', tmpDir)
    U('cp "%s/LICENSE" "%s/LICENSE"', binDir, tmpDir)
    U('cp "%s/README.md" "%s/README.md"', binDir, tmpDir)
    U('cp "%s/appicon.ico" "%s/appicon.ico"', binDir, tmpDir)
    U('mkdir "%s/bin"', tmpDir)
    U('cp "%s/ush_windows_386.exe" "%s/bin/ush_windows_32bit.exe"', binDir, tmpDir)
    U('cp "%s/ush_windows_amd64.exe" "%s/bin/ush_windows_64bit.exe"', binDir, tmpDir)
    U('zip "%s/Ush_windows_x86-64.zip" "%s/"', pubDir, tmpDir)
    U('rm -r -f "%s"', tmpDir)
    if contains(arg, '-rmb') then 
      print("Delete all builds for Windows")
      U('rm -r -f "%s/ush_windows_386.exe"', binDir)
      U('rm -r -f "%s/ush_windows_amd64.exe"', binDir)
    end
    print("Published")
  end

  if contains(arg, '-mac') then 
    print('Publishing release for MacOSX')
    if u('ls -e "%s/Ush_macos_x86-64.gz"', pubDir) == 'true' then
      print('Delete last release build for MacOSX')
      U('rm "%s/Ush_macos_x86-64.gz"', pubDir)
    end 
    local tmpDir = pubDir.."/Ush_macos_x86-64"
    U('mkdir "%s"', tmpDir)
    U('cp "%s/LICENSE" "%s/LICENSE"', binDir, tmpDir)
    U('cp "%s/README.md" "%s/README.md"', binDir, tmpDir)
    U('cp "%s/appicon.ico" "%s/appicon.ico"', binDir, tmpDir)
    U('mkdir "%s/bin"', tmpDir)
    U('cp "%s/ush_darwin_386" "%s/bin/ush_darwin_32bit"', binDir, tmpDir)
    U('cp "%s/ush_darwin_amd64" "%s/bin/ush_darwin_64bit"', binDir, tmpDir)
    U('tar -c -z -f "%s/Ush_macos_x86-64.gz" "%s/"', pubDir, tmpDir)
    U('rm -r -f "%s"', tmpDir)
    if contains(arg, '-rmb') then 
      print("Delete all builds for MacOSX")
      U('rm -r -f "%s/ush_darwin_386"', binDir)
      U('rm -r -f "%s/ush_darwin_amd64"', binDir)
    end
    print("Published")
  end

  if contains(arg, '-nix') then 
    print('Publishing release for Linux')
    if u('ls -e "%s/Ush_linux_x86-64.bz2"', pubDir) == 'true' then
      print('Delete last release build for Linux')
      U('rm "%s/Ush_linux_x86-64.bz2"', pubDir)
    end 
    local tmpDir = pubDir.."/Ush_linux_x86-64"
    U('mkdir "%s"', tmpDir)
    U('cp "%s/LICENSE" "%s/LICENSE"', binDir, tmpDir)
    U('cp "%s/README.md" "%s/README.md"', binDir, tmpDir)
    U('cp "%s/appicon.ico" "%s/appicon.ico"', binDir, tmpDir)
    U('mkdir "%s/bin"', tmpDir)
    U('cp "%s/ush_linux_386" "%s/bin/ush_linux_32bit"', binDir, tmpDir)
    U('cp "%s/ush_linux_amd64" "%s/bin/ush_linux_64bit"', binDir, tmpDir)
    if u('ls -e "%s/ush_linux_arm"', binDir) == 'true' then
      U('cp "%s/ush_linux_arm" "%s/bin/ush_linux_arm"', binDir, tmpDir)
    end
    U('tar -c -j -f "%s/Ush_linux_x86-64.bz2" "%s/"', pubDir, tmpDir)
    U('rm -r -f "%s"', tmpDir)
    if contains(arg, '-rmb') then 
      print("Delete all builds for Linux")
      U('rm -r -f "%s/ush_linux_386"', binDir)
      U('rm -r -f "%s/ush_linux_amd64"', binDir)
      if u('ls -e "%s/ush_linux_arm"', binDir) == 'true' then
        U('rm -r -f "%s/ush_linux_arm"', binDir)
      end
    end
    print("Published")
  end

  print("Done!")
end

function genico()
  print("Generating appicons from master icon")
  -- TODO: To resolve 'Library not loaded' issue while executing ImageMagick commands
  local rootDir = u('pwd')
  local imgDir = rootDir..'/img'
  S('convert "%s/original.png" -resize 512x512 "%s/appicon.icns"', imgDir, imgDir)
  S('convert "%s/original.png" -resize 128x128 "%s/appicon.ico"', imgDir, imgDir)
  print("Done!")
end

function help()
  print([[
Options:
  version       to see the current version of the app
  setup         to install all the go dependencies 
  build         to build a binary executable for current OS & Arch only
  xbuild        to cross build binary executables for MacOSX, Windows, Linux as 32 & 64 bit
  xbuilds       to cross build binary executables for all OS & Arch
  release       to create compressed binary distributable files
  genico        to generate application icons in varied sizes & formats
  help          to see the menu of command line options
  ]])
end

function main(command)
  if command ~= nil then 
    action = {
      ['version'] = version,
      ['build'] = build,
      ['xbuild'] = xbuild,
      ['xbuilds'] = xbuilds,
      ['release'] = release,
      ['genico'] = genico,
      ['help'] = help,
    }
    action[command]()
  else 
    print('Error: Command is missing!')
  end 
end

main(arg[1])