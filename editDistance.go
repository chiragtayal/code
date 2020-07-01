package main

import (
	"fmt"
	"math"
)

/*
Levenstien Distance: https://en.wikipedia.org/wiki/Levenshtein_distance

edit distance is a way of quantifying how dissimilar two strings (e.g., words) are to one another by counting the
minimum number of operations required to transform one string into the other.

There are three operations permitted on a word: replace, delete, insert.
For example, the edit distance between "a" and "b" is 1, the edit distance between "abc" and "def" is 3.

Example:
A = "benyam"
B = "ephrem"

Count number of operations we need to transform A to B using 3 operations (replace, insert, delete)
i:= rows , j:= cols
|_Replace__|___Insert___|
|__Delete__|_CurrentPos_|
 */

func main() {
	B := "zoologicoarchaeologist"
	A := "zoogeologist"

	fmt.Println(editDistance(A,B))
}

func editDistance(A, B string) int {

	a := len(A)
	b := len(B)

	m := make([][]int, a+1)
	for i := range m{
		m[i] = make([]int, b+1)
	}

	for i := 0; i <= a; i++ {
		m[i][0] = i
	}
	for j := 0; j <= b; j++ {
		m[0][j] = j
	}
	printMatrix(m)

	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			printMatrix(m)
			if A[i] == B[j] {
				m[i+1][j+1] = m[i][j]
			} else {
//				r := m[i][j] + 1   // m[i-1+1][j-1+1] because size of m is [a+1][b+1]
//				i := m[i][j+1] + 1 // m[i-1+1][j+1]
//				d := m[i+1][j] + 1 // m[i+1][j-1+1]

				m[i+1][j+1] = min([]int{m[i][j], m[i][j+1], m[i+1][j]}) + 1

				fmt.Println("Val: ", m[i+1][j+1])
			}
		}
	}
	printMatrix(m)
	return m[a][b]
}

func min(ar []int) int {
	if len(ar) == 0 {
		return -1
	}

	min := math.MaxInt64
	for _, e := range ar {
		if e < min {
			min = e
		}
	}
	return min
}

func printMatrix(m [][]int) {
	for i := range m {
		fmt.Println(m[i])
	}
	fmt.Println()
}