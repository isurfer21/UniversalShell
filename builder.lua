--[[
File: builder.lua

Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
]]

require 'ushlib'

-- Common utils
function shfp(...)
  local output = sh(string.format(...))
  if output ~= '' then 
    print(output)
  end
  return output
end

function ushfp(...)
  local output = ush(string.format(...))
  if output ~= '' then 
    print(output)
  end
  return output
end

function ushf(...)
  return ush(string.format(...))
end

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

  local rootDir = ush('pwd')
  local srcDir = rootDir
  local binDir = rootDir..'/bin'
  local pubDir = rootDir..'/pub'

  if ushf('ls -e "%s"', binDir) ~= 'true' then
    ushfp('mkdir "%s"', binDir)
  end

  if ushf('ls -e "%s"', pubDir) ~= 'true' then
    ushfp('mkdir "%s"', pubDir)
  end

  print("- Copying icons")
  ushfp('cp "%s/img/appicon.ico" "%s/appicon.ico"', srcDir, binDir)
  ushfp('cp "%s/img/appicon.icns" "%s/appicon.icns"', srcDir, binDir)
  
  print("- Copying doc files")
  ushfp('cp "%s/LICENSE" "%s/LICENSE"', rootDir, binDir)
  ushfp('cp "%s/README.md" "%s/README.md"', rootDir, binDir)

  print("Create compressed binary distributable files")

  if contains(arg, '-win') then 
    print('  Publishing release for Windows')
    ushfp('rm "%s/Ush_windows_x86-64.zip"', pubDir)
    local tmpDir = pubDir.."/Ush_windows_x86-64"
    ushfp('mkdir "%s"', tmpDir)
    ushfp('cp "%s/LICENSE" "%s/LICENSE"', binDir, tmpDir)
    ushfp('cp "%s/README.md" "%s/README.md"', binDir, tmpDir)
    ushfp('cp "%s/appicon.ico" "%s/appicon.ico"', binDir, tmpDir)
    ushfp('mkdir "%s/bin"', tmpDir)
    ushfp('cp "%s/Ush_windows_386.exe" "%s/bin/Ush_windows_386.exe"', binDir, tmpDir)
    ushfp('cp "%s/Ush_windows_amd64.exe" "%s/bin/Ush_windows_amd64.exe"', binDir, tmpDir)
    shfp('zip -r -X "%s/Ush_windows_x86-64.zip" .', pubDir)
    -- ushfp('rm -r -f "%s"', tmpDir)
    if contains(arg, '-rmb') then 
      print("- Delete all builds for Windows")
      ushfp('rm -r -f "%s/Ush_windows_386.exe"', binDir)
      ushfp('rm -r -f "%s/Ush_windows_amd64.exe"', binDir)
    end
    print("  Published.")
  end

  if contains(arg, '-mac') then 
    print('  Publishing release for MacOSX')
  end

  if contains(arg, '-nix') then 
    print('  Publishing release for Linux')
  end

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