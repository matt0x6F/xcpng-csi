package xolib

type (
	xolib struct {
		Config *Config
	}

	// Xolib is the main struct interface for the xolib library
	Xolib interface {
		Init() error
		Call(*MessageRequest) (*MessageResult, error)
	}
)

// NewXolib initializes the xolib package
func NewXolib(config *Config) (Xolib, error) {
	return &xolib{
		Config: config,
	}, nil
}
