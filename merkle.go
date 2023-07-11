// Copyright 2023 Mohamed Elshami
// Licensed under the MIT License
package main

import (
	"bytes"
	"crypto/sha256"
)

// Node represents a node in the Merkle tree.
type Node struct {
	Hash   []byte
	Parent *Node
	Left   *Node
	Right  *Node
}

// getSibling returns the sibling of a node if it exists, otherwise it returns nil.
func (n *Node) getSibling() *Node {
	if n.Parent == nil {
		return nil
	}

	if n.Parent.Left == n {
		return n.Parent.Right
	} else {
		return n.Parent.Left
	}
}

// MerkleTree represents a Merkle tree structure.
type MerkleTree struct {
	Root   *Node
	Leaves []*Node
}

// NewMerkleTree constructs a new MerkleTree from the given data.
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []*Node

	// If no data is provided, return an empty tree with 0 hash.
	if len(data) == 0 {
		return &MerkleTree{Root: &Node{Hash: computeHash(nil, nil)}, Leaves: nodes}
	}

	for _, value := range data {
		node := &Node{Hash: computeHash(value, nil)}
		nodes = append(nodes, node)
	}

	var leaves []*Node = append([]*Node{}, nodes...)

	for len(nodes) > 1 {
		var levelNodes []*Node

		for i := 0; i < len(nodes); i += 2 {
			var node *Node
			if i+1 < len(nodes) {
				node = &Node{Hash: computeHash(nodes[i].Hash, nodes[i+1].Hash), Left: nodes[i], Right: nodes[i+1]}
				nodes[i].Parent = node
				nodes[i+1].Parent = node
			} else {
				node = nodes[i] // For odd nodes, the last node is a leaf node at this level. Other implementation adds 0 has instead
			}
			levelNodes = append(levelNodes, node)
		}

		nodes = levelNodes
	}

	return &MerkleTree{Root: nodes[0], Leaves: leaves}
}

// GenerateProof generates a proof of inclusion for the given leaf.
func (m *MerkleTree) GenerateProof(leaf *Node) [][]byte {
	proof := [][]byte{}
	currentNode := leaf

	for currentNode != m.Root {
		sibling := currentNode.getSibling()

		if sibling != nil {
			proof = append(proof, sibling.Hash)
		}

		currentNode = currentNode.Parent
	}

	return proof
}

// VerifyProof verifies a proof of inclusion for the given leaf and proof.
func (m *MerkleTree) VerifyProof(leaf *Node, proof [][]byte) bool {
	hash := leaf.Hash

	for _, p := range proof {
		hash = computeHash(hash, p)
	}

	return bytes.Equal(hash, m.Root.Hash)
}

// UpdateLeaf updates the hash of the given leaf and all its ancestors.
func (m *MerkleTree) UpdateLeaf(leaf *Node, data []byte) {
	leaf.Hash = computeHash(data, nil)
	currentNode := leaf

	for currentNode != m.Root {
		sibling := currentNode.getSibling()
		var siblingHash []byte
		if sibling != nil {
			siblingHash = sibling.Hash
		}

		currentNode.Parent.Hash = computeHash(currentNode.Hash, siblingHash)
		currentNode = currentNode.Parent
	}
}

// computeHash computes and returns the SHA256 hash of the concatenation of two data slices.
func computeHash(a, b []byte) []byte {
	hashArray := sha256.Sum256(append(a, b...))
	return hashArray[:]
}
