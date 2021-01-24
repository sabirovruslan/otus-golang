package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyCache(t *testing.T) {
	c := NewCache(10)

	_, ok := c.Get("aaa")
	require.False(t, ok)

	_, ok = c.Get("bbb")
	require.False(t, ok)
}

func TestSimpleCache(t *testing.T) {
	c := NewCache(5)

	wasInCache := c.Set("aaa", 100)
	require.False(t, wasInCache)

	wasInCache = c.Set("bbb", 200)
	require.False(t, wasInCache)

	val, ok := c.Get("aaa")
	require.True(t, ok)
	require.Equal(t, 100, val)

	val, ok = c.Get("bbb")
	require.True(t, ok)
	require.Equal(t, 200, val)

	wasInCache = c.Set("aaa", 300)
	require.True(t, wasInCache)

	val, ok = c.Get("aaa")
	require.True(t, ok)
	require.Equal(t, 300, val)

	val, ok = c.Get("ccc")
	require.False(t, ok)
	require.Nil(t, val)
}

func TestLogic(t *testing.T) {
	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}
