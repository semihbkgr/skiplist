package skiplist

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"time"
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
	size     int
	maxLevel int
	rand     *rand.Rand
}

func NewSkipList[K constraints.Ordered, V any](maxLevel int) *SkipList[K, V] {
	return &SkipList[K, V]{
		head:     newNode[K, V](*new(K), *new(V), maxLevel),
		level:    0,
		size:     0,
		maxLevel: maxLevel,
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (l *SkipList[K, V]) nextLevel() int {
	level := 0
	for level < l.maxLevel && l.rand.Uint32()%2 == 0 {
		level++
	}
	return level
}

func (l *SkipList[K, V]) Find(key K) (V, bool) {
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

func (l *SkipList[K, V]) Insert(key K, val V) {
	nodeLevel := l.nextLevel()
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
	n = newNode(key, val, nodeLevel)
	for i := 0; i <= nodeLevel; i++ {
		n.next[i] = levelTraceNodes[i].next[i]
		levelTraceNodes[i].next[i] = n
	}
	l.size++
}
