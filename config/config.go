package config

import "time"

// DnsConfig struct to use in env package

type DnsConfig struct {
	EnableTCP       bool          `env:"STUB_ENABLE_TCP" envDefault:"true"`
	EnableUDP       bool          `env:"STUB_ENABLE_UDP" envDefault:"true"`
	UpstreamTimeout time.Duration `env:"STUB_UPSTREAM_TIMEOUT" envDefault:"3500ms"`
	UpstreamServer  string        `env:"STUB_UPSTREAM_SERVER" envDefault:"1.1.1.1"`
	UpstreamPort    string        `env:"STUB_UPSTREAM_PORT" envDefault:"853"`
}
