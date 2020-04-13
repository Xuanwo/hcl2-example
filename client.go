package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type client struct {
	cfg *Upstream

	// DoT related config
	TLSServerName string `hcl:"tls_server_name,optional"`
}

func NewClient(cfg *Upstream) (*client, error) {
	c := &client{cfg: cfg}

	var diags hcl.Diagnostics
	diags = gohcl.DecodeBody(cfg.Options, nil, c)
	if diags.HasErrors() {
		return nil, fmt.Errorf("new domain list: %w", diags)
	}

	return c, nil
}
