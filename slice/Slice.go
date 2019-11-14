package main

import "fmt"

func main() {
	const elementSize = 1000

	srcData := make([]int, elementSize)

	for i := 0; i < elementSize; i++ {
		srcData[i] = i + 1
	}

	refData := srcData

	copyData := make([]int, elementSize)

	copy(copyData, srcData)

	srcData[0] = 1001

	fmt.Println(refData[0])
	fmt.Println(srcData[0], srcData[len(srcData)-1])
	copy(copyData, srcData[1:5])

	for i := 0; i < 5; i++ {
		fmt.Println(copyData[i])
	}
}
