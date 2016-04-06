/*
package main

import (
	"fmt"
	"hash/fnv"
	"sort"
)

func Hash(x int) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(fmt.Sprintf("%d", x)))
	return hasher.Sum64()
}

const ITEMS = 1000000
const NODES = 100
const VIRTUAL_NODES_PER_NODE = 1000

type nodeSlice []uint64

func (ns nodeSlice) Len() int               { return len(ns) }
func (ns nodeSlice) Swap(a int, b int)      { ns[a], ns[b] = ns[b], ns[a] }
func (ns nodeSlice) Less(a int, b int) bool { return ns[a] < ns[b] }

func main() {
	ring := make([]uint64, NODES*VIRTUAL_NODES_PER_NODE)
	hashesToNode := make(map[uint64]int, NODES*VIRTUAL_NODES_PER_NODE)
	for p, n := 0, 0; n < NODES; n++ {
		for v := 0; v < VIRTUAL_NODES_PER_NODE; v++ {
			h := Hash(n*1000000 + v)
			ring[p] = h
			p++
			hashesToNode[h] = n
		}
	}
	sort.Sort(nodeSlice(ring))

	countPerNode := make([]int, NODES)
	for i := 0; i < ITEMS; i++ {
		h := Hash(i)
		x := sort.Search(len(ring), func(x int) bool { return ring[x] >= h })
		if x >= len(ring) {
			x = 0
		}
		countPerNode[hashesToNode[ring[x]]]++
	}
	min := ITEMS
	max := 0
	for n := 0; n < NODES; n++ {
		if countPerNode[n] < min {
			min = countPerNode[n]
		}
		if countPerNode[n] > max {
			max = countPerNode[n]
		}
	}
	t := ITEMS / NODES
	fmt.Printf("%d to %d assigments per node, target was %d.\n", min, max, t)
	fmt.Printf("That's %.02f%% under and %.02f%% over.\n",
		float64(t-min)/float64(t)*100, float64(max-t)/float64(t)*100)

	ring2 := make([]uint64, (NODES+1)*VIRTUAL_NODES_PER_NODE)
	copy(ring2, ring)
	hashesToNode2 := make(map[uint64]int, (NODES+1)*VIRTUAL_NODES_PER_NODE)
	for k, v := range hashesToNode {
		hashesToNode2[k] = v
	}
	for p, v := NODES*VIRTUAL_NODES_PER_NODE, 0; v < VIRTUAL_NODES_PER_NODE; v++ {
		h := Hash(NODES*1000000 + v)
		ring2[p] = h
		p++
		hashesToNode2[h] = NODES
	}
	sort.Sort(nodeSlice(ring2))

	moved := 0
	for i := 0; i < ITEMS; i++ {
		h := Hash(i)
		x := sort.Search(len(ring), func(x int) bool { return ring[x] >= h })
		if x >= len(ring) {
			x = 0
		}
		x2 := sort.Search(len(ring2), func(x int) bool { return ring2[x] >= h })
		if x2 >= len(ring2) {
			x2 = 0
		}
		if hashesToNode[ring[x]] != hashesToNode2[ring2[x2]] {
			moved++
		}
	}
	fmt.Printf("%d items moved, %.02f%%.\n",
		moved, float64(moved)/float64(ITEMS)*100)
}
*/
package main

import (
	"fmt"
	"hash/fnv"

	"github.com/gholt/ring"
)

func Hash(x int) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(fmt.Sprintf("%d", x)))
	return hasher.Sum64()
}

const ITEMS = 1000000
const NODES = 100

func main() {
	nodeIDsToNode := make(map[uint64]int)
	b := ring.NewBuilder(64)
	for n := 0; n < NODES; n++ {
		bn, _ := b.AddNode(true, 1, nil, nil, "", nil)
		nodeIDsToNode[bn.ID()] = n
	}
	ring := b.Ring()

	nodeIDsToNode2 := make(map[uint64]int, NODES+1)
	for k, v := range nodeIDsToNode {
		nodeIDsToNode2[k] = v
	}
	b.PretendElapsed(b.MoveWait() + 1)
	bn, _ := b.AddNode(true, 1, nil, nil, "", nil)
	nodeIDsToNode2[bn.ID()] = NODES
	ring2 := b.Ring()

	countPerNode := make([]int, NODES)
	for i := 0; i < ITEMS; i++ {
		h := Hash(i)
		x := ring.ResponsibleNodes(uint32(h >> (64 - ring.PartitionBitCount())))[0].ID()
		countPerNode[nodeIDsToNode[x]]++
	}
	min := ITEMS
	max := 0
	for n := 0; n < NODES; n++ {
		if countPerNode[n] < min {
			min = countPerNode[n]
		}
		if countPerNode[n] > max {
			max = countPerNode[n]
		}
	}
	t := ITEMS / NODES
	fmt.Printf("%d to %d assigments per node, target was %d.\n", min, max, t)
	fmt.Printf("That's %.02f%% under and %.02f%% over.\n",
		float64(t-min)/float64(t)*100, float64(max-t)/float64(t)*100)

	moved := 0
	for i := 0; i < ITEMS; i++ {
		h := Hash(i)
		x := ring.ResponsibleNodes(uint32(h >> (64 - ring.PartitionBitCount())))[0].ID()
		x2 := ring2.ResponsibleNodes(uint32(h >> (64 - ring2.PartitionBitCount())))[0].ID()
		if nodeIDsToNode[x] != nodeIDsToNode2[x2] {
			moved++
		}
	}
	fmt.Printf("%d items moved, %.02f%%.\n",
		moved, float64(moved)/float64(ITEMS)*100)
}
