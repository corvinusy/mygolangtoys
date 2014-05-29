package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const (
	SIZE   = 40
	GREY   = 0
	ORANGE = -1
	RED    = -2
)

type Node struct {
	num      int
	color    int
	edges    [40]int
	distance int
}

func main() {

	execPath, _ := exec.LookPath(os.Args[0])

	fname := path.Dir(execPath) + "/network.txt"

	tree := getTreeFromFile(fname)

	//check data
	/*
		for _, t := range tree {
			fmt.Println(t)
		}
	*/
	mst := make([]Node, 0)
	mst = append(mst, tree[0])

	near := list.New()

	updateNear(near, mst, tree)

	for len(mst) < SIZE {
		mst = append(mst, getBest(near))
		updateNear(near, mst, tree)
		/*
			fmt.Printf("%d:[", len(mst))
			for _, m := range mst {
				fmt.Printf(",%d", m.num)
			}
			fmt.Println("]")
		*/
	}

	minDistance := 0

	for _, m := range mst {
		minDistance += m.distance
	}

	allDistance := 0

	for _, t := range tree {
		for j := 0; j < t.num; j++ {
			allDistance += t.edges[j]
		}
	}

	fmt.Println(allDistance - minDistance)

	return
}

/*----------------------------------------------------------------------------*/
func updateNear(near *list.List, mst, tree []Node) {

	for _, m := range mst {
		for i, v := range m.edges {
			if v != 0 && !isAlwaysInMST(mst, i) {
				found := false
				// search list and update found
				for e := near.Front(); e != nil; e = e.Next() {
					if i == e.Value.(Node).num {
						found = true
						if e.Value.(Node).distance > v {
							near.Remove(e)
							tree[i].distance = v
							near.PushBack(tree[i])
						}
						break
					}
				}
				if !found {
					tree[i].distance = v
					near.PushBack(tree[i])
				}

			}
		}
	}

	return

}

/*----------------------------------------------------------------------------*/
func isAlwaysInMST(mst []Node, n int) bool {

	for _, m := range mst {
		if m.num == n {
			return true
		}
	}

	return false
}

/*----------------------------------------------------------------------------*/
func getBest(near *list.List) Node {

	result := near.Front()

	for e := near.Front(); e != nil; e = e.Next() {
		if e.Value.(Node).distance < result.Value.(Node).distance {
			result = e
		}
	}

	return near.Remove(result).(Node)
}

/*----------------------------------------------------------------------------*/
func getTreeFromFile(fname string) []Node {

	var (
		strs    []string
		strnums []string
		tmp     int64
	)

	//read file into source

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("FILE NOT READ")
	}

	source := strings.Trim(string(content), "\n")

	//create result placeholder
	tree := make([]Node, SIZE)

	strs = strings.Split(source, "\n")

	for i, s := range strs {
		strnums = strings.Split(s, ",")
		tree[i].num = i
		for j, sn := range strnums {
			if sn != "-" {
				tmp, _ = strconv.ParseInt(sn, 10, 0)
				tree[i].edges[j] = int(tmp)
			}
		}
	}

	return tree
}
