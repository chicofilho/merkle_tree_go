package merkleTree

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing a simple tree with two leafs only
func TestCreateTree(t *testing.T) {
	bytes := [][]byte{[]byte("first"), []byte("second")}
	merkleTree := MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)
	assert.NotNil(t, merkleTree.Root)
	assert.Equal(t, merkleTree.Root.getHashString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")

	// creating the tree manually for comparison
	newNode := Node{}
	newNode.addLeafsBytes(bytes[0], bytes[1])
	assert.Equal(t, len(merkleTree.Leafs), 2)
	assert.Equal(t, merkleTree.Root.getHashString(), newNode.getHashString())

}

// testing an odd tree
func TestCreateOddTree(t *testing.T) {
	// a tree with three leafs as start
	bytes := [][]byte{[]byte("first"), []byte("second"), []byte("third")}
	merkleTree := MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)
	assert.NotNil(t, merkleTree.Root)
	assert.Equal(t, len(merkleTree.Leafs), 4)
	assert.Equal(t, merkleTree.Root.getHashString(), "17a3c8b6a398e426c6356adb7f1dbc8d91db23a1")

	// a tree with thirteen leafs as start
	bytes = make([][]byte, 13)
	for i := 0; i < 13; i++ {
		bytes[i] = []byte("a sequence at: " + strconv.Itoa(i))
	}

	merkleTree = MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)

	assert.NotNil(t, merkleTree.Root)
	assert.Equal(t, len(merkleTree.Leafs), 14)
	assert.Equal(t, merkleTree.Root.getHashString(), "a40e39f789c0b9359b0d11f32a0b87c7b5d2caf9")

}

// testing a tree with four elements
func TestCreateTreeWithFourElements(t *testing.T) {
	bytes := [][]byte{[]byte("first"), []byte("second"), []byte("first"), []byte("second")}
	merkleTree := MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)
	assert.NotNil(t, merkleTree.Root)
	assert.Equal(t, len(merkleTree.Leafs), 4)
	assert.Equal(t, merkleTree.Root.getHashString(), "97324ce4b16ed0f699ec059e618b981c497f2fd7")
}

// testing two approahces: recursive and non-recursive
func TestCreateTreeDifferentApproaches(t *testing.T) {
	nodes := []*Node{
		&Node{Content: []byte("first")},
		&Node{Content: []byte("second")},
		&Node{Content: []byte("third")},
		&Node{Content: []byte("fourth")},
		&Node{Content: []byte("fifth")},
	}
	merkleTree := MerkleTree{Leafs: nodes}
	merkleTreeRecursive := MerkleTree{Leafs: nodes}
	merkleTree.CreateTree()
	merkleTreeRecursive.CreateTreeRecursive()

	assert.NotNil(t, merkleTree.Root)
	assert.NotNil(t, merkleTreeRecursive.Root)
	assert.Equal(t, merkleTree.Root.getHashString(), merkleTreeRecursive.Root.getHashString())

}

// testing a tree with 1M elements
func TestCreateBigTree(t *testing.T) {
	bytes := make([][]byte, 1000000)
	for i := 0; i < 1000000; i++ {
		bytes[i] = []byte("a sequence at: " + strconv.Itoa(i))
	}

	merkleTree := MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)

	assert.NotNil(t, merkleTree.Root)
	assert.Equal(t, len(merkleTree.Leafs), 1000000)
	assert.Equal(t, merkleTree.Root.getHashString(), "87df3eed6314c26c3a95f64358a19c8bbdfcb65f")
}

func TestMerklePath(t *testing.T) {
	bytes := make([][]byte, 1000000)
	for i := 0; i < 1000000; i++ {
		bytes[i] = []byte("a sequence at: " + strconv.Itoa(i))
	}

	merkleTree := MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)
	leaf := merkleTree.Leafs[13]
	path := merkleTree.GetMerklePath(leaf.Hash)

	assert.Equal(t, len(path), 20)
	assert.Equal(t, merkleTree.Root.getHashString(), path.MerkleProof(leaf.Hash).ToString())

}

func TestMerklePathNonExistingNode(t *testing.T) {
	bytes := make([][]byte, 10)
	for i := 0; i < 10; i++ {
		bytes[i] = []byte("a sequence at: " + strconv.Itoa(i))
	}

	merkleTree := MerkleTree{}
	merkleTree.CreateTreeBytes(bytes)
	testNode := Node{Content: []byte("leaf.Hash")}
	path := merkleTree.GetMerklePath(testNode.GetHash())

	assert.Equal(t, len(path), 0)

}
