package treebuilding

import (
    "errors"
    "math"
    "slices"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
        return nil, nil
    }

    mapParentToChildren := map[int][]int{}
    mapChildToParent := map[int]int{}

    root := math.MaxFloat64
    max := math.SmallestNonzeroFloat64

    for _, record := range records {
		if record.Parent > record.ID {
            return nil, errors.New("parent value cannot be greater than child value")
        }

        max = math.Max(max, float64(record.ID))
        
        if record.Parent == record.ID {
            if root == math.MaxFloat64 {
                root = float64(record.ID)
            } else {
            	return nil, errors.New("duplicate root found")
            }
        } else {
            if parentChildren, parentExists := mapParentToChildren[record.Parent]; parentExists {
                if slices.Contains(parentChildren, record.ID) {
                    return nil, errors.New("duplicate node found")
                }

                if _, mappingToParentExists := mapChildToParent[record.ID]; mappingToParentExists {
                    return nil, errors.New("duplicate node found or cycle detected")
                }

                parentChildren = append(parentChildren, record.ID)
                mapParentToChildren[record.Parent] = parentChildren
            } else {
            	mapParentToChildren[record.Parent] = []int{record.ID}
            }
        }

        mapChildToParent[record.ID] = record.Parent
    }

    if root == math.MaxFloat64 {
        return nil, errors.New("root node does not exist or has a parent")
    }

    if int(max) + 1 != len(mapChildToParent) {
        return nil, errors.New("non-continuous tree detected")
    }
    
    runningTree := []*Node{}
	
    rootNode := &Node{
        ID: int(root),
        Children: nil,
    }

	runningTree = append(runningTree, rootNode)

    for len(runningTree) > 0 {
        curr := runningTree[0]
        runningTree = slices.Delete(runningTree, 0, 1)

        if childrenNodes, childrenExist := mapParentToChildren[curr.ID]; childrenExist {
            children := []*Node{}
            slices.Sort(childrenNodes)
            
            for _, childNode := range childrenNodes {
                child := &Node{
                    ID: childNode,
                    Children: nil,
                }

                children = append(children, child)
                runningTree = append(runningTree, child)
            }
            
            curr.Children = children
        }
    }

    return rootNode, nil
}
