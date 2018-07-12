--[[
File: ushlib.lua

Copyright © 2018 Abhishek Kumar <isurfer21@gmail.com>
This work is licensed under the 'MIT License'.
]]

UniversalShell = {

-- PREREQUISITES 

  -- Execute (Sh)
  sh = function (command)
  	local handle = io.popen(command)
  	local result = handle:read('*a')
  	handle:close()
  	return result
  end,

  -- Execute (Ush)
  ush = function (command)
  	local handle = io.popen('ush ' .. command)
  	local result = handle:read('*a')
  	handle:close()
  	return trim(result)
  end,

-- DERIVED

  -- Format > Execute (Sh) > Print
  shFmtExecPrnt = function (...)
    return UniversalShell.printOutput(UniversalShell.sh(string.format(...)))
  end,

  -- Format > Execute (Sh)
  shFmtExec = function (...)
    return UniversalShell.sh(string.format(...))
  end,

  -- Format > Execute (Ush) > Print
  ushFmtExecPrnt = function (...)
    return UniversalShell.printOutput(UniversalShell.ush(string.format(...)))
  end,

  -- Format > Execute (Ush)
  ushFmtExec = function (...)
    return UniversalShell.ush(string.format(...))
  end,

  -- Print, if output is not empty 
  printOutput = function (str)
    local output = str
    if output ~= '' then 
      print(output)
    end
    return output
  end
  
}

-- UTILITIES
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

function contains(table, val)
   for i=1,#table do
      if table[i] == val then 
         return true
      end
   end
   return false
end
