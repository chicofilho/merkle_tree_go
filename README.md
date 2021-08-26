# Merkle Tree Implementation in Go

This project has an implementation of the MerkleTree algorithm using go. 

## Running the project

Go is a dependency of this project. Make sure you have it installed in your machine

The file main.go shows the usage of the merkleTree package built here. The file describes an array of bytes and use them to build the Merkle Tree.

To run this project go to the root folder and run:
```
go run .
```
It will compile and run the main function of the `main.go` file. You should see an output of a tree and some details about the Merkle Proof of a given path

If you want to run the tests just go to the root of the folder and run:
```
go test ./... -cover
```

## Package main
The main pacakge uses the Merkle Tree package. There's a printer file with printing methods and the main file. The main file builds a tree and prints information about it and its operations results for this scenario

## Package MerkleTree
The package is divided into three main files. `merkleTree.go` has all the definitions of a merkle tree creation and its operations, but doesn't describe any node level operation. `node.go` on the other hand has all the helper functions to work with a node. `hash.go` has the definition of a type hash to help with manipulation and presentation. 

All files have their unit tests covering basic and more complex scenarios. The project overal coverage is: `98.9%`

Specific to this implementation. The merkle tree carries a hashset with all leafs using their hash as the key to access the nodes data. That way, by having a specific hash it takes O(1) to find the node in the leafs and the merkle path takes the regular O(log(n)) to build and proof.

