--[[
File: uninstall.lua

Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
]]

require 'UniversalShell'

-- CORE

S = UniversalShell.shFmtExecPrnt
s = UniversalShell.shFmtExec
U = UniversalShell.ushFmtExecPrnt
u = UniversalShell.ushFmtExec

-- SEMI-CORE

function parent(str)
    local p = split(str,'[\\/]+')
    return string.format('%s/%s', p[1], p[2])
end

-- TASKS

local pkg = 'github.com/isurfer21/UniversalShell'

print('Clean removes object files from package source directories (ignore error)')
S('go clean -i %s', pkg)

print('Delete the source directory and compiled package directory(ies)')
local GOPATH = u('getenv GOPATH')

if u('ls -e "%s/src/%s/"', GOPATH, pkg) == 'true' then
	print('Deleting package source')
	U('rm -v -r -f "%s/src/%s/"', GOPATH, pkg)
	print('Deleting package source container, if empty')
	U('rm -v "%s/src/%s"', GOPATH, parent(pkg))
end 

if u('uname -m') == 'x86_64' then
    local ost = u('uname') .. '_amd64'
	if u('ls -e "%s/pkg/%s/%s/"', GOPATH, ost, pkg) == 'true' then
		print('Deleting package objects')
		U('rm -v -r -f "%s/pkg/%s/%s/"', GOPATH, ost, pkg)
		print('Deleting package objects container, if empty')
		U('rm -v "%s/pkg/%s/%s"', GOPATH, ost, parent(pkg))
	end 
end

print('Delete package binary')
local ext = ''
if u('uname') == 'windows' then
	ext = '.exe'
end
U('rm -v -f "%s/bin/ush%s"', GOPATH, ext)
