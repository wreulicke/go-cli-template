package main

import (
	"io"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Parallel()

	version(io.Discard, false)
}
