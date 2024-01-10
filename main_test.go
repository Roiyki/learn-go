package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Inc(t *testing.T) {
	require.Equal(t, 5, inc(4))
}
