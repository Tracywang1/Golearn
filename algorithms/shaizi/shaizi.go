package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Player struct {
	Name       string
	Score      int
	CurBlockId int
}

type WorldMap struct {
	Id      int
	Name    string
	Rewards int
}

func Newmap() *WorldMap {
	return new(WorldMap)
}

func Shaizi() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func Steps(rand []int) int {
	var step int
	for _, v := range rand {
		step = step + v
	}
	if step >= 20 {
		return step - 20
	}
	return step
}

func read(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func main() {
	file := "map.csv"
	fmt.Println(read(file))
	// rand := make([]int, 10)
	// for i := 0; i < 10; i++ {
	// 	rand[i] = Shaizi()
	// 	fmt.Println(rand[i], Steps(rand))
	// }
}
