package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func flippingMatrix(matrix [][]int32) int32 {
	var result int32
	m := matrix
	n := len(m)

	for i := range n / 2 {
		for j := range n / 2 {
			var max int32

			mirrors := []int32{}
			mirrors = append(mirrors, m[i][j], m[i][n-j-1], m[n-i-1][j], m[n-i-1][n-j-1])

			sort.Slice(mirrors, func(i, j int) bool {
				return mirrors[i] < mirrors[j]
			})

			max = mirrors[len(mirrors)-1]
			result += max
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var matrix [][]int32
		for i := 0; i < 2*int(n); i++ {
			matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var matrixRow []int32
			for _, matrixRowItem := range matrixRowTemp {
				matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
				checkError(err)
				matrixItem := int32(matrixItemTemp)
				matrixRow = append(matrixRow, matrixItem)
			}

			if len(matrixRow) != 2*int(n) {
				panic("Bad input")
			}

			matrix = append(matrix, matrixRow)
		}

		result := flippingMatrix(matrix)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
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
