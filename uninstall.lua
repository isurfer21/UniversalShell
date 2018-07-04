-- File: uninstall.lua

-- Prerequisite method
function ush(command)
	local handle = io.popen('ush ' .. command)
	local result = handle:read("*a")
	handle:close()
	return result
end

-- Start your shell tasks from here
print(ush('ls -l'))