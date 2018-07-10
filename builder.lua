--[[
File: builder.lua

Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
]]

require 'ushlib'

-- Start your shell tasks from here

function version()
  print([[
Builder   version 1.0
Copyright (c) 2018 Abhishek Kumar. All rights reserved.
  ]])
end

function build()
  print("Build for current machine only")
  print(sh('go build ush.go'))
  print("Done!")
end

function xbuild()
  print("Cross build for MacOSX, Windows & Linux as 32bit & 64bit")
  print(sh('gox -osarch="darwin/386" -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print(sh('gox -osarch="darwin/amd64" -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print(sh('gox -osarch="windows/386" -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print(sh('gox -osarch="windows/amd64" -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print(sh('gox -osarch="linux/386" -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print(sh('gox -osarch="linux/amd64" -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print("Done!")
end

function xbuilds()
  print("Cross build for all OS and Arch")
  print(sh('gox -output="./bin/ush_{{.OS}}_{{.Arch}}"'))
  print("Done!")
end

function release()
  print("Create distributable packages")
  print("Done!")
end

function genico()
  print("Generating appicons from master icon")
  -- TODO: To resolve 'Library not loaded' issue while executing ImageMagick commands
  sh('convert ./img/original.png -resize 512x512 ./img/appicon.icns')
  sh('convert ./img/original.png -resize 128x128 ./img/appicon.ico')
  print("Done!")
end

function help()
  print([[
Options:
  version       to see the current version of the app
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