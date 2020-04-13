package main

import (
	"io/ioutil"
	"testing"
)

func TestParseNormal(t *testing.T) {
	filename := "testdata/1.hcl"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	c, err := Parse(content, filename)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", c)
}

func TestParseMissingField(t *testing.T) {
	filename := "testdata/2.hcl"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	c, err := Parse(content, filename)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", c)
}
