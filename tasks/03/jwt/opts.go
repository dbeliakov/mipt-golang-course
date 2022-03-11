package jwt

import "time"

type Option func(*config)

func WithSignMethod(m SignMethod) Option {
	return func(c *config) {
		c.SignMethod = m
	}
}

func WithKey(k []byte) Option {
	return func(c *config) {
		c.Key = k
	}
}

func WithTTL(ttl time.Duration) Option {
	return func(c *config) {
		c.TTL = &ttl
	}
}

func WithExpires(t time.Time) Option {
	return func(c *config) {
		c.Expires = &t
	}
}

type config struct {
	SignMethod SignMethod
	Key        []byte
	TTL        *time.Duration
	Expires    *time.Time
}
