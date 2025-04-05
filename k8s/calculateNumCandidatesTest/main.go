/*
kube-schedulerのpreemptionロジックのうちcalculateNumCandidates()という関数で、
Node数が多い時に一定数に絞り込む処理をしている。
この処理の意味がいまいちわからないため検証。
pkg/scheduler/framework/plugins/defaultpreemption/default_preemption.go
*/

package main

import "fmt"

const MinCandidateNodesPercentage int32 = 10
const MinCandidateNodesAbsolute int32 = 100

func main() {
	nodeNums := []int32{90, 99, 100, 101, 110, 200, 220, 500, 900, 999, 1000, 1100, 1500, 2000, 3000}
	for _, numNode := range nodeNums {
		fmt.Printf("numNode: %v\n", numNode)
		calculateNumCandidates(numNode)
		fmt.Println("--------------------------------------------------")
	}
}

func calculateNumCandidates(numNodes int32) {

	// var numNodes int32 = 900
	// var n int32

	// デバッグ
	// fmt.Printf("numNodes: %v\n", numNodes)
	// fmt.Printf("MinCandidateNodesPercentage: %v\n", MinCandidateNodesPercentage)
	// fmt.Printf("MinCandidateNodesAbsolute: %v\n", MinCandidateNodesAbsolute)

	n := (numNodes * MinCandidateNodesPercentage) / 100

	// デバッグ
	fmt.Printf("n_1: %v\n", n)

	if n < MinCandidateNodesAbsolute {
		n = MinCandidateNodesAbsolute
	}

	// デバッグ
	fmt.Printf("n_2: %v\n", n)

	if n > numNodes {
		n = numNodes
	}

	// デバッグ
	fmt.Printf("n_final: %v\n", n)
}
