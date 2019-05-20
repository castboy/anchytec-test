package main

import "fmt"

type Config struct{
	reticulatedSplines bool
	cities int
}

type Terrain struct {
	config Config
}

type Option interface {
	apply(*Config)
}

func NewTerrain(options ...Option) *Terrain {
	var t Terrain

	for _, option := range options {
		option.apply(&t.config)
	}

	return &t
}

// one implemention.
type splines struct{}

func (s *splines) apply(config *Config) {
	config.reticulatedSplines = true
}

func WithReticulatedSplines() Option {
	return new(splines)
}

// another implemention.
type cities struct {
	cities int
}

func (c *cities) apply(config *Config) {
	config.cities = c.cities
}

func WithCities(n int) Option {
	return &cities{
		cities: n,
	}
}

func main() {
	t := NewTerrain(WithReticulatedSplines(), WithCities(9))
	fmt.Println(t)
}