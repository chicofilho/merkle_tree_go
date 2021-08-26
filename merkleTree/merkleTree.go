package merkleTree

import (
	"math"
)

/*
* Merkle Tree representation
* It carries only information about the root and leaf nodes
*
 */

type MerkleTree struct {
	Root     *Node
	Leafs    []*Node
	LeafsMap map[string]*Node
}

func (merkleTree *MerkleTree) populateMap() {
	if len(merkleTree.LeafsMap) > 0 {
		return
	}

	merkleTree.LeafsMap = make(map[string]*Node)
	for _, node := range merkleTree.Leafs {
		merkleTree.LeafsMap[node.getHashString()] = node
	}
}

// helper method to create the tree based of bytes only
func (merkleTree *MerkleTree) CreateTreeBytes(bytes [][]byte) {
	// parsing bytes to nodes
	nodes := make([]*Node, len(bytes), len(bytes))
	for i, val := range bytes {
		nodes[i] = &Node{Content: val}
	}
	// generating the Merkle tree
	merkleTree.Leafs = nodes
	merkleTree.CreateTree()
}

func (merkleTree *MerkleTree) CreateTree() {
	merkleTree.populateMap()
	// handling odd entries with blank value
	if len(merkleTree.Leafs)%2 == 1 {
		merkleTree.Leafs = append(merkleTree.Leafs, &Node{})
	}

	leafs := merkleTree.Leafs
	levelUp := make([]*Node, 0, len(leafs)/2)

	i := 0
	for len(leafs) > 1 {
		if i < len(leafs) { // generating level up
			newNode := Node{}
			newNode.addLeafs(leafs[i], leafs[i+1])
			levelUp = append(levelUp, &newNode)
			i = i + 2
		} else { // reset steps
			i = 0
			leafs = levelUp
			if len(leafs) > 1 && len(leafs)%2 == 1 { //odd entries correction
				leafs = append(leafs, &Node{})
			}
			levelUp = make([]*Node, 0, len(leafs)/2)
		}
	}

	merkleTree.Root = leafs[0] //assigning root: tree is ready
}

// a method for generating the MerkleTree recursively
func (merkleTree *MerkleTree) CreateTreeRecursive(nodes ...*Node) {
	if len(nodes) == 0 { // if no parameter is passed (using optional parameters)
		nodes = merkleTree.Leafs
	}

	merkleTree.populateMap()
	nodeStack := []*Node{}
	if len(nodes) == 1 { // recurstion stop condition
		merkleTree.Root = nodes[0] //assigning root: tree is ready
		return
	}
	if len(nodes)%2 == 1 { // add blank value if nodes are odd
		nodes = append(nodes, &Node{})
	}

	// generating the upper level
	for i := 0; i < len(nodes); i = i + 2 {
		newNode := Node{}
		newNode.addLeafs(nodes[i], nodes[i+1])
		nodeStack = append(nodeStack, &newNode)
	}

	// recursion with upper level generated
	merkleTree.CreateTreeRecursive(nodeStack...)
}

// helper type to create a merkle path
type MerkleStep struct {
	hash        Hash
	isLeftChild bool
}

// helper type to calulate a merkle proof
type MerklePath []MerkleStep

// Merkle proof is a reduce calculation of the merkle path
func (merklePath MerklePath) MerkleProof(hash Hash) Hash {
	proof := hash
	for _, val := range merklePath {
		var appended []byte
		if val.isLeftChild {
			appended = append(val.hash[:], proof[:]...)
		} else {
			appended = append(proof[:], val.hash[:]...)
		}
		proof = calculateHashFromBytes(appended)
	}
	return proof
}

/* Returns an array of the path with the orientation of the node
* and the hash. The goal is this array to be reduced
 */
func (merkleTree *MerkleTree) GetMerklePath(hash Hash) MerklePath {
	node, ok := merkleTree.LeafsMap[hash.ToString()]
	if !ok {
		return MerklePath{}
	}

	size := math.Logb(float64(len(merkleTree.Leafs)))
	merklePath := make(MerklePath, 0, int(size))
	for !node.isRoot() {
		sibling := node.getSibling()
		merkleStep := MerkleStep{hash: sibling.GetHash(), isLeftChild: sibling.isLeftChild()}
		merklePath = append(merklePath, merkleStep)
		node = node.Parent
	}

	return merklePath
}
