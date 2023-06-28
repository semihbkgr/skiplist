package skiplist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	l := New[int, string]()
	assert.NotNil(t, l)
	assert.Equal(t, defaultMaxLevel, l.maxLevel)
}

func TestNew_maxLevel(t *testing.T) {
	maxLevel := 18

	l := New[int, string](maxLevel)
	assert.NotNil(t, l)
	assert.Equal(t, maxLevel, l.maxLevel)
}

func TestNew_negativeMaxLevel(t *testing.T) {
	maxLevel := -3

	l := New[int, string](maxLevel)
	assert.NotNil(t, l)
	assert.Equal(t, minimumMaxLevel, l.maxLevel)
}

func TestSkipList_Insert(t *testing.T) {
	l := New[int, string]()

	for i := 0; i < 100; i++ {
		v := fmt.Sprintf("%v-%d", time.Now(), i)
		l.Insert(i, v)
	}
}

func TestSkipList_Insert_duplicatedKeys(t *testing.T) {
	l := New[int, string]()

	for i := 0; i < 100; i++ {
		v := fmt.Sprintf("%v-%d", time.Now(), i)
		l.Insert(i, v)
	}

	kvMap := make(map[int]string)
	for i := 0; i < 100; i++ {
		v := fmt.Sprintf("%v-%d", time.Now(), i)
		l.Insert(i, v)
		kvMap[i] = v
	}

	for i := 0; i < 100; i++ {
		v, ok := l.Get(i)
		assert.True(t, ok)
		assert.Equal(t, v, kvMap[i])
	}

	length := l.Length()
	assert.Equal(t, 100, length)
}

func TestSkipList_Get(t *testing.T) {
	l := New[int, string]()

	for i := 0; i < 100; i++ {
		v, ok := l.Get(i)
		assert.False(t, ok)
		assert.Empty(t, v)
	}

	kvMap := make(map[int]string)
	for i := 0; i < 100; i++ {
		v := fmt.Sprintf("%v-%d", time.Now(), i)
		l.Insert(i, v)
		kvMap[i] = v
	}

	for i := 0; i < 100; i++ {
		v, ok := l.Get(i)
		assert.True(t, ok)
		assert.Equal(t, v, kvMap[i])
	}
}

func TestSkipList_Delete(t *testing.T) {
	l := New[int, string]()

	for i := 0; i < 100; i++ {
		ok := l.Delete(i)
		assert.False(t, ok)
	}

	for i := 0; i < 100; i++ {
		v := fmt.Sprintf("%v-%d", time.Now(), i)
		l.Insert(i, v)
	}

	for i := 0; i < 100; i++ {
		ok := l.Delete(i)
		assert.True(t, ok)
	}
}

func TestSkipList_Length(t *testing.T) {
	l := New[int, string]()

	length := l.Length()
	assert.Zero(t, length)

	for i := 0; i < 100; i++ {
		v := fmt.Sprintf("%v-%d", time.Now(), i)
		l.Insert(i, v)
	}

	length = l.Length()
	assert.Equal(t, 100, length)

	for i := 0; i < 100; i++ {
		l.Delete(i)
	}

	length = l.Length()
	assert.Zero(t, length)
}
