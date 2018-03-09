package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 多叉树的一个节点
type PathTreeNode struct {
	Content   string
	Score     string
	BranchMap map[string]*PathTreeNode
}

// 一个简单的多叉树而已
type PathTree struct {
	RootNode *PathTreeNode
}

func NewPathTree() *PathTree {
	return new(PathTree)
}

func (pt *PathTree) AddPath(contentpath string, scorepath string) {
	contents := strings.Split(contentpath, "_")
	scores := strings.Split(scorepath, "_")
	// fmt.Println("contents:", contents, "scores:", scores)
	if len(contents) != len(scores) || len(contents) == 0 {
		errstr := fmt.Sprintf("Addpath: {%v} {%v}", contentpath, scorepath)
		panic(errstr)
	}
	if pt.RootNode == nil {
		pt.RootNode = new(PathTreeNode)
		pt.RootNode.Content = contents[0]
		pt.RootNode.Score = scores[0]
		pt.RootNode.BranchMap = make(map[string]*PathTreeNode)
		b, _ := json.Marshal(pt)
		fmt.Println("pt:", string(b))
	}
	if pt.RootNode.Content != contents[0] {
		errstr := fmt.Sprintf("Addpath: {%v [0]} != {%v}", contentpath, pt.RootNode.Content)
		panic(errstr)
	}
	contents = contents[1:]
	scores = scores[1:]
	// fmt.Println("1:啥效果", contents, scores)
	curNode := pt.RootNode
	b, _ := json.Marshal(curNode)
	fmt.Println("curNode:", string(b))

	for {
		fmt.Println("len(contents):", len(contents))
		if len(contents) == 0 {
			break
		}
		if curNode.BranchMap == nil {
			curNode.BranchMap = make(map[string]*PathTreeNode)
		}
		curNewContent := contents[0]
		fmt.Println("contents[0]:", contents[0], "curNode.BranchMap[curNewContent]:", curNode.BranchMap[curNewContent])
		if _node, ok := curNode.BranchMap[curNewContent]; ok {
			// this is an old path
			b, _ := json.Marshal(curNode)
			fmt.Println("curNode111111111", string(b))
			curNode = _node
			c, _ := json.Marshal(curNode)
			fmt.Println("curNode22222222222", string(c))
			contents = contents[1:]
			scores = scores[1:]
			continue
		} else {
			// new path
			_newNode := new(PathTreeNode)
			_newNode.Content = curNewContent
			_newNode.Score = scores[0]
			_newNode.BranchMap = make(map[string]*PathTreeNode)
			// b, _ := json.Marshal(_newNode)
			// fmt.Println("_newNode:", string(b))
			curNode.BranchMap[curNewContent] = _newNode
		}
		b, _ := json.Marshal(pt)
		fmt.Println("pttttttttttt:", string(b))
	}
}

var path = []string{"1_1_1_1", "1_2_1_1", "1_2_2_1"}
var score = "0_0_0_1"

func main() {
	pt := NewPathTree()
	for _, onepath := range path {
		pt.AddPath(onepath, score)
	}
	b, _ := json.Marshal(pt)
	fmt.Println(string(b))
}
