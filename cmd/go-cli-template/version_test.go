package main

import (
	"testing"

	"github.com/wreulicke/snap"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	s := snap.New()
	version(s, false)
	s.Assert(t)
}
