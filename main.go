package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type cell struct {
	val  uint8
	next *cell
	prev *cell
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Execute with 1 argument: the brainfuck file path")
		os.Exit(0)
	}
	filename := os.Args[1]
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(filename, "not found")
	}

	i := 0
	currCell := &cell{
		val:  0,
		next: nil,
		prev: nil,
	}
	loopstack := make([]int, 0)
	for i < len(input) {
		switch string(input[i]) {
		case ">":
			if currCell.next == nil {
				newCell := &cell{
					val:  0,
					next: nil,
					prev: currCell,
				}
				currCell.next = newCell
			}
			currCell = currCell.next
			i++
		case "<":
			if currCell.prev == nil {
				newCell := &cell{
					val:  0,
					next: currCell,
					prev: nil,
				}
				currCell.prev = newCell
			}
			currCell = currCell.prev
			i++
		case "+":
			if currCell.val == 255 {
				currCell.val = 0
			} else {
				currCell.val++
			}
			i++
		case "-":
			if currCell.val == 0 {
				currCell.val = 255
			} else {
				currCell.val--
			}
			i++
		case ".":
			fmt.Print(string(currCell.val))
			i++
		case ",":
			var x string
			fmt.Scan(&x)
			val := []byte(x)[0]
			if val < 0 && val > 255 {
				log.Fatalln("Input must be ASCII character")
			} else {
				currCell.val = val
			}
			i++
		case "[":
			if currCell.val == 0 {
				loopstack = append(loopstack, i)
				currLoopstack := i
				for {
					i++
					if string(input[i]) == "[" {
						loopstack = append(loopstack, i)
					} else if string(input[i]) == "]" {
						pop := loopstack[len(loopstack)-1]
						loopstack = append(loopstack[0 : len(loopstack)-1])
						if pop == currLoopstack {
							break
						}
					}
				}
			} else {
				loopstack = append(loopstack, i)
			}
			i++
		case "]":
			if currCell.val != 0 {
				i = loopstack[len(loopstack)-1]
			} else {
				loopstack = append(loopstack[0 : len(loopstack)-1])
			}
			i++
		default:
			i++
		}
	}
	fmt.Println()
	fmt.Println("--------------------\nSuccess")
}
