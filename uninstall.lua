-- File: uninstall.lua

-- Prerequisite method
function sh(command)
	local handle = io.popen(command)
	local result = handle:read('*a')
	handle:close()
	return result
end

function ush(command)
	local handle = io.popen('ush ' .. command)
	local result = handle:read('*a')
	handle:close()
	return trim(result)
end

-- Custom library
function trim (str)
  if str == '' then
    return str
  else  
    local startPos = 1
    local endPos   = #str

    while (startPos < endPos and str:byte(startPos) <= 32) do
      startPos = startPos + 1
    end

    if startPos >= endPos then
      return ''
    else
      while (endPos > 0 and str:byte(endPos) <= 32) do
        endPos = endPos - 1
      end

      return str:sub(startPos, endPos)
    end
  end
end

function split(str, pat)
   local t = {}  -- NOTE: use {n = 0} in Lua-5.0
   local fpat = "(.-)" .. pat
   local last_end = 1
   local s, e, cap = str:find(fpat, 1)
   while s do
      if s ~= 1 or cap ~= "" then
         table.insert(t,cap)
      end
      last_end = e+1
      s, e, cap = str:find(fpat, last_end)
   end
   if last_end <= #str then
      cap = str:sub(last_end)
      table.insert(t, cap)
   end
   return t
end

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
