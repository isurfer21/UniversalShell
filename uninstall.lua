-- File: uninstall.lua

-- Prerequisite method
function sh(command)
	local handle = io.popen(command)
	local result = handle:read('*a')
	handle:close()
	return result
end

function ush(command)
	local handle = io.popen('./ush ' .. command)
	local result = handle:read('*a')
	handle:close()
	return result
end

-- Custom library
function countPatternOccurance(base, pattern)
    return select(2, base:gsub(pattern, '/'))
end

-- Start your shell tasks from here

local pkg = arg[1]
local ost
local cnt
local scr

print('Clean removes object files from package source directories (ignore error)')
-- print(sh('go clean -i ' + pkg))

print('Set local variables')

if (ush('uname -m') == 'x86_64') {
    ost = ush('uname')
    ost = ost .. '_amd64'
    cnt = countPatternOccurance(pkg, '/')
}

-- todo: create a cmd to set & get env variable
print('Delete the source directory and compiled package directory(ies)')
local GOPATH = ush('getx GOPATH')
if (cnt == 2) {
    ush('rm -r -f ' .. GOPATH .. '"/src/' .. pkg .. '/*"')
    ush('rm -rf ' .. GOPATH .. '"/pkg/' .. ost .. '/' .. pkg .. '/*"')
} else (cnt > 2) {
    ush('rm -rf ' .. GOPATH .. '"/src/' .. pkg .. '/*/*"')
    ush('rm -rf ' .. GOPATH .. '"/pkg/' .. ost .. '/' .. pkg .. '/*/*"')
}