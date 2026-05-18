package consistence_hash

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"sort"
)

type Node struct {
	Key string
}

type VirtualNode struct {
	Hash uint32
	Node Node
}

type HashRing struct {
	hasher   func(string) uint32
	replicas int
	keys     []uint32
	nodes    map[uint32]Node
}

func NewHashRing() *HashRing {
	r := &HashRing{
		hasher: func(s string) uint32 {
			sum := md5.Sum([]byte(s))
			return binary.BigEndian.Uint32(sum[:4])
		},
		replicas: 10,
		nodes:    make(map[uint32]Node),
	}
	return r
}
func (r *HashRing) AddNode(node Node) {
	if node.Key == "" {
		return
	}

	for i := 0; i < r.replicas; i++ {
		virtualKey := fmt.Sprintf("%s#%d", node.Key, i)
		hash := r.hasher(virtualKey)
		r.keys = append(r.keys, hash)
		r.nodes[hash] = node
	}

	sort.Slice(r.keys, func(i, j int) bool {
		return r.keys[i] < r.keys[j]
	})
}

func (r *HashRing) GetNode(key string) (Node, bool) {
	if len(r.keys) == 0 {
		return Node{}, false
	}

	idx, ok := r.locate(r.hasher(key))
	if !ok {
		return Node{}, false
	}
	node, ok := r.nodes[r.keys[idx]]
	return node, ok
}

func (r *HashRing) VirtualNodes() []VirtualNode {
	result := make([]VirtualNode, 0, len(r.keys))
	for _, hash := range r.keys {
		result = append(result, VirtualNode{
			Hash: hash,
			Node: r.nodes[hash],
		})
	}
	return result
}

func (r *HashRing) LocateKey(key string) (uint32, VirtualNode, bool) {
	if len(r.keys) == 0 {
		return 0, VirtualNode{}, false
	}

	hash := r.hasher(key)
	idx, ok := r.locate(hash)
	if !ok {
		return 0, VirtualNode{}, false
	}

	virtualHash := r.keys[idx]
	return hash, VirtualNode{
		Hash: virtualHash,
		Node: r.nodes[virtualHash],
	}, true
}

func (r *HashRing) locate(hash uint32) (int, bool) {
	if len(r.keys) == 0 {
		return 0, false
	}

	idx := sort.Search(len(r.keys), func(i int) bool {
		return r.keys[i] >= hash
	})
	if idx == len(r.keys) {
		idx = 0
	}

	return idx, true
}
