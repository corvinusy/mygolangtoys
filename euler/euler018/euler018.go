package main

import (
    "fmt"
	"strings"
	"strconv"
)

type Node struct {
//	leftRoot *Node
//	rightRoot *Node
	leftLeaf *Node
	rightLeaf *Node
	row int
	col int
	value int
	weight int
}


type Tree struct {
	nodes [15][15]*Node
	root *Node
}

func (t *Tree) insertNode(node *Node) {

	t.nodes[node.row][node.col] = node

	fmt.Printf("\n%d %d", node.row, node.col);

	if node.row < 1 {
		return
	}

	if node.col < node.row {
		t.nodes[node.row - 1][node.col].leftLeaf = node
		fmt.Printf("left") 
	}

	if node.col > 0 {
		t.nodes[node.row - 1][node.col - 1].rightLeaf = node
		fmt.Printf("right")		
	}

}

func (n *Node) initNode(value int, row int, col int) {
//	n.leftRoot = nil
//	n.rightRoot = nil
	n.leftLeaf = nil
	n.rightLeaf = nil
	n.row = row
	n.col = col
	n.value = value
	n.weight = 0
}



func main() {

	var source string = 
`75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`


	var strs []string
	strs = strings.SplitAfter(source,"\n")

	var strnums []string

	nums := make ([][]int, 15)
	for i, _ := range nums {
		nums[i] = make([]int, 15)
	}

	tree := new(Tree)
	rootNode := new(Node)
	rootNode.initNode(75, 0, 0)
	tree.insertNode(rootNode)
	tree.root = rootNode

	var tmp int64
	
	for i := 1; i < len(strs); i++ {
		strnums = strings.Fields(strs[i])
		for j, sn := range strnums  {
			tmp, _ = strconv.ParseInt(sn, 10, 0)
			nums[i][j] = int(tmp)
			node := new(Node)
			node.initNode(nums[i][j], i, j)
			tree.insertNode(node)
		}
	}
	
	//lets calc weights

	fmt.Println(tmp)
}


