package skiplist

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

const (
	defaultMaxLevel = 12
	minimumMaxLevel = 0
)

type node[K constraints.Ordered, V any] struct {
	key  K
	val  V
	next []*node[K, V]
}

func newNode[K constraints.Ordered, V any](key K, val V, level int) *node[K, V] {
	return &node[K, V]{
		key:  key,
		val:  val,
		next: make([]*node[K, V], level+1),
	}
}

type SkipList[K constraints.Ordered, V any] struct {
	head     *node[K, V]
	level    int
	len      int
	maxLevel int
	rand     *rand.Rand
}

func New[K constraints.Ordered, V any](maxLevel ...int) *SkipList[K, V] {
	maxLvl := defaultMaxLevel
	if len(maxLevel) > 0 {
		maxLvl = maxLevel[0]
	}
	if maxLvl < minimumMaxLevel {
		maxLvl = minimumMaxLevel
	}
	return &SkipList[K, V]{
		head:     newNode[K, V](*new(K), *new(V), maxLvl),
		level:    0,
		len:      0,
		maxLevel: maxLvl,
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (l *SkipList[K, V]) Get(key K) (V, bool) {
	n := l.head
	for i := l.level; i >= 0; i-- {
		for n.next[i] != nil && n.next[i].key < key {
			n = n.next[i]
		}
	}
	n = n.next[0]
	if n == nil || n.key != key {
		return *new(V), false
	}
	return n.val, true
}

func (l *SkipList[K, V]) randLevel() int {
	level := 0
	for level < l.maxLevel && l.rand.Uint32()%2 == 0 {
		level++
	}
	return level
}

func (l *SkipList[K, V]) Insert(key K, val V) {
	nodeLevel := l.randLevel()
	if l.level < nodeLevel {
		l.level = nodeLevel
	}
	levelTraceNodes := make([]*node[K, V], l.level+1)
	n := l.head
	for i := l.level; i >= 0; i-- {
		for n.next[i] != nil && n.next[i].key < key {
			n = n.next[i]
		}
		levelTraceNodes[i] = n
	}
	n = n.next[0]
	if n != nil && n.key == key {
		n.val = val
		return
	}
	n = newNode(key, val, nodeLevel)
	for i := 0; i <= nodeLevel; i++ {
		n.next[i] = levelTraceNodes[i].next[i]
		levelTraceNodes[i].next[i] = n
	}
	l.len++
}

func (l *SkipList[K, V]) Delete(key K) bool {
	levelTraceNodes := make([]*node[K, V], l.level+1)
	n := l.head
	for i := l.level; i >= 0; i-- {
		for n.next[i] != nil && n.next[i].key < key {
			n = n.next[i]
		}
		levelTraceNodes[i] = n
	}
	n = n.next[0]
	if n == nil || n.key != key {
		return false
	}
	for i, next := range n.next {
		levelTraceNodes[i].next[i] = next
	}
	l.len--
	return true
}

func (l *SkipList[K, V]) Length() int {
	return l.len
}
