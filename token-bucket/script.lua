-- KEYS[1] = bucket key
-- ARGV[1] = capacity
-- ARGV[2] = refill_rate (token per ms)
-- ARGV[3] = now (ms)
-- ARGV[4] = cost

-- get values from golang
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local cost = tonumber(ARGV[4])

local data = redis.call("HMGET", KEYS[1], "tokens", "last_refill")

local tokens = tonumber(data[1])
local last_refill = tonumber(data[2])

if tokens == nil then
	tokens = capacity
	last_refill = now
end

local delta = math.max(0, now - last_refill)
local new_tokens = math.min(capacity, tokens + delta * refill_rate)

local allowed = 0
if new_tokens >= cost then
	new_tokens = new_tokens - cost
	allowed = 1
end

redis.call("HMSET", KEYS[1], "tokens", new_tokens, "last_refill", now)

-- TTL to auto cleanup key
local ttl = math.ceil(capacity / refill_rate)
redis.call("PEXPIRE", KEYS[1], ttl)

return { allowed, new_tokens }
