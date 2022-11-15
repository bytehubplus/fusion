package common

import "testing"

func TestNew(t *testing.T) {
	conf := NewConfig("a", "f", "p")
	t.Log(conf.FileName)
}

func TestGetNode(t *testing.T) {
	conf := NewConfig("a", "f", "p")
	conf.GetNodes()
}
