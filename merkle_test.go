// Copyright 2023 Mohamed Elshami
// Licensed under the MIT License
package main

import (
	"bytes"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	leaves := [][]byte{
		[]byte("A"),
		[]byte("B"),
		[]byte("C"),
		[]byte("D"),
	}

	tree := NewMerkleTree(leaves)

	if len(tree.Leaves) != len(leaves) {
		t.Errorf("Number of leaves incorrect, got: %d, want: %d", len(tree.Leaves), len(leaves))
	}

	// Calculate expected root hash for leaves A, B, C, D
	hashA := computeHash([]byte("A"), nil)
	hashB := computeHash([]byte("B"), nil)
	hashC := computeHash([]byte("C"), nil)
	hashD := computeHash([]byte("D"), nil)
	hashAB := computeHash(hashA, hashB)
	hashCD := computeHash(hashC, hashD)
	expectedRootHash := computeHash(hashAB, hashCD)

	if !bytes.Equal(tree.Root.Hash, expectedRootHash) {
		t.Errorf("Incorrect root hash, got: %x, want: %x", tree.Root.Hash, expectedRootHash)
	}
}

func TestGenerateProof(t *testing.T) {
	leaves := [][]byte{
		[]byte("A"),
		[]byte("B"),
		[]byte("C"),
		[]byte("D"),
	}

	tree := NewMerkleTree(leaves)

	proof := tree.GenerateProof(tree.Leaves[0])

	if len(proof) == 0 {
		t.Errorf("Proof is empty, expected a proof")
	}

	// Calculate expected proof for leaf "A"
	hashB := computeHash([]byte("B"), nil)
	hashC := computeHash([]byte("C"), nil)
	hashD := computeHash([]byte("D"), nil)
	hashCD := computeHash(hashC, hashD)
	expectedProof := [][]byte{hashB, hashCD}

	for i, p := range proof {
		if !bytes.Equal(p, expectedProof[i]) {
			t.Errorf("Incorrect proof, got: %x, want: %x", p, expectedProof[i])
		}
	}
}

func TestVerifyProof(t *testing.T) {
	leaves := [][]byte{
		[]byte("A"),
		[]byte("B"),
		[]byte("C"),
		[]byte("D"),
	}

	tree := NewMerkleTree(leaves)

	proof := tree.GenerateProof(tree.Leaves[0])

	ok := tree.VerifyProof(tree.Leaves[0], proof)

	if !ok {
		t.Errorf("Proof verification failed")
	}
}

func TestUpdateLeaf(t *testing.T) {
	leaves := [][]byte{
		[]byte("A"),
		[]byte("B"),
		[]byte("C"),
		[]byte("D"),
	}

	tree := NewMerkleTree(leaves)

	// keep old root hash for comparison
	oldRootHash := tree.Root.Hash

	tree.UpdateLeaf(tree.Leaves[0], []byte("E")) // update the first leaf ("A") with new data ("E")

	// Check if the root hash changed after the update
	if bytes.Equal(oldRootHash, tree.Root.Hash) {
		t.Errorf("Root hash did not change after leaf update")
	}

	// Check if the leaf's data was correctly updated
	newHash := computeHash([]byte("E"), nil)
	if !bytes.Equal(tree.Leaves[0].Hash, newHash[:]) {
		t.Errorf("Leaf update failed, got: %x, want: %x", tree.Leaves[0].Hash, newHash[:])
	}
}
