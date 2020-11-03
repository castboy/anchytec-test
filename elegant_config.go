package main

import "fmt"

type Cfg struct {
	a string
	b int
	c bool
}

func NewCfgA(a string) func(cfg *Cfg) {
	return func(cfg *Cfg) {
		cfg.a = a
	}
}

func NewCfgB(b int) func(cfg *Cfg) {
	return func(cfg *Cfg) {
		cfg.b = b
	}
}

func NewCfgC(c bool) func(cfg *Cfg) {
	return func(cfg *Cfg) {
		cfg.c = c
	}
}

func NewCfg(opts ...func(cfg *Cfg)) *Cfg {
	cfg := &Cfg{}

	for i := range opts {
		opts[i](cfg)
	}

	return cfg
}

func main() {
	a := NewCfgA("wmg")
	b := NewCfgB(30)
	c := NewCfgC(true)

	cfg := NewCfg(a, b, c)
	fmt.Println(cfg)
}
