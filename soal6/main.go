package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingSort function below.
func countingSort(arr []int32) []int32 {
	max := int32(-1)
	for _, n := range arr {
		if n > max {
			max = n
		}
	}

	fmt.Println(max)

	length := int32(100)

	if max+1 < 100 {
		length = 100
	} else {
		length = max + 1
	}
	out := make([]int32, length)

	for _, n := range arr {
		out[n] += 1
		// fmt.Println("n=",n,"out[n]=",out[n])
	}

	//   for i,n:= range out {
	//         fmt.Println("i=",i,"out[i]=",n)
	//     }

	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 10240*10240)

	stdout, err := os.Create("./output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	fmt.Println("====================================================")
	fmt.Println("arrTemp length", len(arrTemp))
	fmt.Println(arrTemp)
	fmt.Println("====================================================")

	_ = n
	_ = writer

	// var arr []int32

	// for i := 0; i < int(n); i++ {
	// 	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	// 	checkError(err)
	// 	arrItem := int32(arrItemTemp)
	// 	arr = append(arr, arrItem)
	// }

	// result := countingSort(arr)

	// for i, resultItem := range result {
	// 	fmt.Fprintf(writer, "%d", resultItem)

	// 	if i != len(result)-1 {
	// 		fmt.Fprintf(writer, " ")
	// 	}
	// }

	// fmt.Fprintf(writer, "\n")

	// writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
