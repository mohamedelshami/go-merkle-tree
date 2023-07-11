# Go Merkle Tree

This project provides a simple implementation of a Merkle Tree data structure in Go.

## Overview

A Merkle tree also known a hash tree, is a binary tree in which every leaf node is labelled with the hash of a data block, and every non-leaf node is labelled with the cryptographic hash of the labels of its child nodes. This data structure allows efficient and secure verification of the contents of large data structures.

## Install

Simply clone the repository using:

```bash
git clone https://github.com/mohamedelshami/go-merkle-tree.git
```

Or run Go get command

```bash
go get https://github.com/mohamedelshami/go-merkle-tree.git
```

## Run Tests

```bash
 cd /path/go-merkle-tree

 go test 
```

## Usage

This project is a Go package that can be imported into other Go projects. To use this package, you would generally do the following:

1. Import the package.
2. Create a new Merkle Tree.
3. Generate a proof for a leaf.
4. Verify a proof for a leaf.
5. Update a leaf node.

## Documentation

The main public functions of the MerkleTree struct are:

1. `NewMerkleTree(data [][]byte) *MerkleTree` - This function creates a new Merkle Tree from the provided data.

2. `GenerateProof(leaf *Node) [][]byte` - This function generates a proof of inclusion for the given leaf node.

3. `VerifyProof(leaf *Node, proof [][]byte) bool` - This function verifies the provided proof for the given leaf node.

4. `UpdateLeaf(leafIndex int, data []byte)` - This function updates the value of a leaf node and recalculates the necessary hashes.

## License

This project is licensed under the MIT License.