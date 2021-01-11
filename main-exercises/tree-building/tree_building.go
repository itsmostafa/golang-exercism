package tree

import (
	"fmt"
	"sort"
)

// Record object.
type Record struct {
	ID, Parent int
}

// Node is a branch of a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build a tree logic with the given records.
func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node)
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("Invalid record: Parent %d ID %d", r.Parent, r.ID)
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID > 0 {
			p := nodes[r.Parent]
			p.Children = append(p.Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
