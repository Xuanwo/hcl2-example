package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Config struct {
	Listen    string            `hcl:"listen"`
	Upstreams []*Upstream       `hcl:"upstream,block"`
	Rules     map[string]string `hcl:"rules"`
}

type Upstream struct {
	Name    string   `hcl:",label"`
	Type    string   `hcl:"type"`
	Addr    string   `hcl:"addr"`
	Options hcl.Body `hcl:",remain"`
}

// Parse will parse file content into valid config.
func Parse(src []byte, filename string) (c *Config, err error) {
	var diags hcl.Diagnostics

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, fmt.Errorf("config parse: %w", diags)
	}

	c = &Config{}

	diags = gohcl.DecodeBody(file.Body, nil, c)
	if diags.HasErrors() {
		return nil, fmt.Errorf("config parse: %w", diags)
	}

	return c, nil
}
