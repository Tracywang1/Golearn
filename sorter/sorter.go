package main

import (
	"algorithms/bubblesort"
	"algorithms/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "inflie", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted value")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open input file", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A tpp long line ,seems upexpected.")
			return
		}
		str := string(line)
		fmt.Println("Line number is", str)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Faile to ceate the output file", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		fmt.Println("Sorted values:", value, "\n")
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	fmt.Println("---------Begin------------")
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile = ", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	fmt.Println(len(values))
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported")
		}
		t2 := time.Now()

		fmt.Println("The Sorting process cossts", t2.Sub(t1), "to complete")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
