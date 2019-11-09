// Package tree contains a solution to the Tree Building exercism.io exercise.
package tree

import (
	"errors"
	"sort"
)

// Record represents a single record.
type Record struct {
	ID     int
	Parent int
}

// Node represents a single node in the tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build parses a slice of Record structs and outputs the Node struct.
func Build(records []Record) (*Node, error) {
	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	nodes := make(map[int]*Node, len(records))
	for i, r := range records {
		if (r.ID != i) || !((r.ID > r.Parent) || (r.ID == 0 && r.Parent == 0)) {
			return nil, errors.New("invalid record")
		}
		nodes[i] = &Node{
			ID: i,
		}
		if i != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[i])
		}
	}
	return nodes[0], nil
}
