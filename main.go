// Copyright 2023 Mohamed Elshami
// Licensed under the MIT License
package main

import "fmt"

func main() {
	leaves := [][]byte{
		[]byte("A"),
		[]byte("B"),
		[]byte("C"),
	}

	tree := NewMerkleTree(leaves)

	fmt.Println("Root hash:", tree.Root.Hash)

	proof := tree.GenerateProof(tree.Leaves[0])
	fmt.Println("Proof:", proof)

	ok := tree.VerifyProof(tree.Leaves[0], proof)
	fmt.Println("Proof verified:", ok)
}
