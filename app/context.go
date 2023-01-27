package app

import (
	"auditor/env"
	"auditor/response"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

// Context app context
type Context struct {
	Environment *env.Environment
	Results     *response.Results
	RedisClient *redis.Client
	RedisLock   *redislock.Client
}

// NewContext new application context
func NewContext(e *env.Environment, r *response.Results) *Context {
	// debug := e.Production == false
	return &Context{
		Environment: e,
		Results:     r,
	}
}
