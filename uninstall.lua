--[[
File: uninstall.lua

Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
]]

require 'ushlib'

-- Custom library
function parent(str)
    local p = split(str,'[\\/]+')
    return string.format('%s/%s', p[1], p[2])
end

-- Start your shell tasks from here

local pkg = 'github.com/isurfer21/UniversalShell'

print('Clean removes object files from package source directories (ignore error)')
print(sh('go clean -i ' .. pkg))

print('Delete the source directory and compiled package directory(ies)')
local GOPATH = ush('getenv GOPATH')

if (ush(string.format('ls -e "%s/src/%s/"', GOPATH, pkg)) == 'true') then
	print('Deleting package source')
	print(ush(string.format('rm -v -r -f "%s/src/%s/"', GOPATH, pkg)))
	print('Deleting package source container, if empty')
	print(ush(string.format('rm -v "%s/src/%s"', GOPATH, parent(pkg))))
end 

if (ush('uname -m') == 'x86_64') then
    local ost = ush('uname') .. '_amd64'
	if (ush(string.format('ls -e "%s/pkg/%s/%s/"', GOPATH, ost, pkg)) == 'true') then
		print('Deleting package objects')
		print(ush(string.format('rm -v -r -f "%s/pkg/%s/%s/"', GOPATH, ost, pkg)))
		print('Deleting package objects container, if empty')
		print(ush(string.format('rm -v "%s/pkg/%s/%s"', GOPATH, ost, parent(pkg))))
	end 
end

print('Delete package binary')
local ext = ''
if (ush('uname') == 'windows') then
	ext = '.exe'
end
print(ush(string.format('rm -v -f "%s/bin/ush%s"', GOPATH, ext)))
