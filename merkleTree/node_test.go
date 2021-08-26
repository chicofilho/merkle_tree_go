package merkleTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// testing IsLeaf method
func TestIsLeaf(t *testing.T) {

	nodeEmpty := Node{}
	nodeNonLeaf := Node{Content: []byte("test"), Right: &Node{}, Left: &Node{}}
	nodeLeaf := Node{Content: []byte("test")}
	assert.Equal(t, nodeEmpty.IsLeaf(), true)
	assert.Equal(t, nodeNonLeaf.IsLeaf(), false)
	assert.Equal(t, nodeLeaf.IsLeaf(), true)
}

// testing isRoot method
func TestIsRoot(t *testing.T) {

	nodeEmpty := Node{}
	nodeNonRoot := Node{Content: []byte("test"), Parent: &Node{}}
	nodeRoot := Node{Content: []byte("test"), Right: &Node{}, Left: &Node{}}
	assert.Equal(t, nodeEmpty.isRoot(), true)
	assert.Equal(t, nodeNonRoot.isRoot(), false)
	assert.Equal(t, nodeRoot.isRoot(), true)
}

// testing isRoot method
func TestIsLeftChild(t *testing.T) {
	right := Node{Content: []byte("right")}
	left := Node{Content: []byte("left")}

	parent := Node{Right: &right, Left: &left}
	parent.addLeafs(&left, &right)

	assert.Equal(t, left.isLeftChild(), true)
	assert.Equal(t, right.isLeftChild(), false)
}

// testing isRoot method
func TestGetSibling(t *testing.T) {

	right := Node{Content: []byte("right")}
	left := Node{Content: []byte("left")}

	parent := Node{Right: &right, Left: &left}
	parent.addLeafs(&left, &right)

	assert.Equal(t, right.getSibling(), &left)
	assert.Equal(t, left.getSibling(), &right)
}

// testing isRoot method
func TestGetSiblingOfRoot(t *testing.T) {
	node := Node{Content: []byte("right")}
	assert.Nil(t, node.getSibling())
}

// testing hash generation by nodes
func TestSingleNodeHash(t *testing.T) {

	tests := [][]string{
		{"first", "e0996a37c13d44c3b06074939d43fa3759bd32c1"},
		{"second", "352f7829a2384b001cc12b0c2613c756454a1f6a"},
		{"third", "34fb3300b9a77bebdc988ec3edd0d4a6a42a26f9"},
		{"fourth", "2db18e1d98e7ab7f49dea56027312c2d97b1a2e0"},
		{"fifth", "5ad43148c90a8f2d296fc74f9f327538e964d000"},
		{"sixth", "04053a7b8a6957822a1a10641c094af04adc071e"}}

	for _, val := range tests {
		node := Node{Content: []byte(val[0])}
		assert.Equal(t, node.getHashString(), val[1])
	}
}

// testing hash generation by parent node
func TestWithChildrenHash(t *testing.T) {
	right := Node{Content: []byte("right")}
	left := Node{Content: []byte("left")}
	parent := Node{Right: &right, Left: &left}

	assert.Equal(t, parent.getHashString(), "05b94d588d9350bbc7d523871e6ee5a33a4e6f7a")
}

//Testing addleaf function
func TestAddNode(t *testing.T) {
	first, second := "first", "second"
	node := Node{}
	node.addLeafs(&Node{Content: []byte(first)}, &Node{Content: []byte(second)})
	assert.Equal(t, node.getHashString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")
	assert.Equal(t, node.Left.Parent.Hash.ToString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")
	assert.Equal(t, node.Right.Parent.Hash.ToString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")
}

//Testing add leafs by byte
func TestAddNodeByte(t *testing.T) {
	first, second := "first", "second"
	node := Node{}
	node.addLeafsBytes([]byte(first), []byte(second))
	assert.Equal(t, node.getHashString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")
	assert.Equal(t, node.Left.Parent.Hash.ToString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")
	assert.Equal(t, node.Right.Parent.Hash.ToString(), "4517358021ad7a8ed37d3aca96f3a70ac5d0868e")
}
