package ccsc

import (
	"bytes"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type merkleProof struct {
	value []byte
	path  []byte
	nodes []byte
}

func MerkleProof(list types.DerivableList, index uint) (*merkleProof, error) {
	merkleTrie := new(trie.Trie)
	_ = types.DeriveSha(list, merkleTrie)

	b := new(bytes.Buffer)
	err := rlp.Encode(b, index)
	if err != nil {
		return nil, err
	}
	path := b.Bytes()

	nodes, err := LeafProofByKey(merkleTrie, path)
	if err != nil {
		return nil, err
	}

	return &merkleProof{
		value: merkleTrie.Get(path),
		path:  path,
		nodes: nodes,
	}, nil
}

func LeafProofByKey(merkleTrie *trie.Trie, key []byte) ([]byte, error) {
	iter := merkleTrie.NodeIterator(nil)
	nodes := new(bytes.Buffer)
	for iter.Next(true) {
		if iter.Leaf() && bytes.Equal(iter.LeafKey(), key) {
			err := rlp.Encode(nodes, iter.LeafProof())
			if err != nil {
				return nil, err
			}
			break
		}
	}
	return nodes.Bytes(), nil
}
