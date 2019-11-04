// Package tree contains a solution to the Tree Building exercism.io exercise.
package tree

import "errors"

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

	// Return if any records are outside the range and sore the records.
	for i, r := range records {
		if r.ID < 0 || r.ID >= len(records) {
			return nil, errors.New("invalid record")
		}
		if r.ID != i {
			records[i], records[r.ID] = records[r.ID], records[i]
		}
	}

	// Place each node in a slice. Append the children to the respective parent.
	nodes := make([]Node, len(records))
	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("non-continuous records")
		}
		if !((r.ID > r.Parent) || (r.ID == 0 && r.Parent == 0)) {
			return nil, errors.New("invalid parent")
		}

		nodes[i].ID = i
		if i != 0 {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	// The first node in the slice is the root node. Return it to the caller.
	return &nodes[0], nil
}
