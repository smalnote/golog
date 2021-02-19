package redis

import "context"

// Config redist connection config
type Config struct {
	Hostname string
	Port     int
	Password string
}

type configContextKey struct{}

// Context return a context with value of config
func (c *Config) Context(ctx context.Context) context.Context {
	return context.WithValue(ctx, configContextKey{}, c)
}

// FromContext get config from context
func FromContext(ctx context.Context) *Config {
	if cfg, ok := ctx.Value(configContextKey{}).(*Config); ok {
		return cfg
	}
	return nil
}
