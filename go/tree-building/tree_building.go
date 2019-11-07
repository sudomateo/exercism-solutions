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

	// Return if no data is passed.
	if len(records) == 0 {
		return nil, nil
	}

	// Sort the records.
	sort.SliceStable(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// Place each node in a slice. Append the children to the respective parent.
	nodes := make([]Node, len(records))
	for i, r := range records {
		if (r.ID != i) || !((r.ID > r.Parent) || (r.ID == 0 && r.Parent == 0)) {
			return nil, errors.New("invalid record")
		}
		if i != 0 {
			nodes[i].ID = i
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, &nodes[i])
		}
	}

	// The first node in the slice is the root node. Return it to the caller.
	return &nodes[0], nil
}
