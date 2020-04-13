package main

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	content, err := ioutil.ReadFile("testdata/test.hcl")
	if err != nil {
		t.Error(err)
	}

	c, err := Parse(content, "testdata/test.hcl")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", c)
}
