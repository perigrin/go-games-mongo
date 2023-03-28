package main

type node struct {
	Parent   *node
	Position *Position
	g        int
	h        int
	f        int
}

func newNode(parent *node, position *Position) *node {
	return &node{
		Parent:   parent,
		Position: position,
		g:        0,
		h:        0,
		f:        0,
	}
}

func (n *node) isEqual(other *node) bool {
	return n.Position.IsEqual(other.Position)
}

type AStar struct{}

func (as AStar) GetPath(level Level, start *Position, end *Position) []Position {
	gd := NewGameData()
	openList := make([]*node, 0)
	closedList := make([]*node, 0)

	startNode := newNode(nil, start)
	enp := newNode(nil, end)

	openList = append(openList, startNode)

	for {
		if len(openList) == 0 {
			break
		}

		currentNode := openList[0]
		currentIndex := 0

		for index, item := range openList {
			if item.f < currentNode.f {
				currentNode = item
				currentIndex = index
			}
		}

		openList = append(openList[:currentIndex], openList[currentIndex+1:]...)
		closedList = append(closedList, currentNode)

		if currentNode.isEqual(enp) {
			path := make([]Position, 0)
			current := currentNode
			for {
				if current == nil {
					break
				}
				path = append(path, *current.Position)
				current = current.Parent
			}

			reverseSlice(path)
			return path
		}

		edges := make([]*node, 0)
		if currentNode.Position.Y > 0 {
			idx := level.GetIndexFromXY(
				currentNode.Position.X,
				currentNode.Position.Y-1,
			)
			tile := level.Tiles[idx]
			if tile.TileType != WALL {
				newNodePosition := Position{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y - 1,
				}
				newNode := newNode(currentNode, &newNodePosition)
				edges = append(edges, newNode)
			}
		}
		if currentNode.Position.Y < gd.ScreenHeight {
			idx := level.GetIndexFromXY(
				currentNode.Position.X,
				currentNode.Position.Y+1,
			)
			tile := level.Tiles[idx]
			if tile.TileType != WALL {
				newNodePosition := Position{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y + 1,
				}
				newNode := newNode(currentNode, &newNodePosition)
				edges = append(edges, newNode)
			}
		}
		if currentNode.Position.X > 0 {
			idx := level.GetIndexFromXY(
				currentNode.Position.X-1,
				currentNode.Position.Y,
			)
			tile := level.Tiles[idx]
			if tile.TileType != WALL {
				newNodePosition := Position{
					X: currentNode.Position.X - 1,
					Y: currentNode.Position.Y,
				}
				newNode := newNode(currentNode, &newNodePosition)
				edges = append(edges, newNode)
			}
		}
		if currentNode.Position.X < gd.ScreenWidth {
			idx := level.GetIndexFromXY(
				currentNode.Position.X+1,
				currentNode.Position.Y,
			)
			tile := level.Tiles[idx]
			if tile.TileType != WALL {
				newNodePosition := Position{
					X: currentNode.Position.X + 1,
					Y: currentNode.Position.Y,
				}
				newNode := newNode(currentNode, &newNodePosition)
				edges = append(edges, newNode)
			}
		}
		for _, edge := range edges {
			if isInSlice(closedList, edge) {
				continue
			}
			edge.g = currentNode.g + 1
			edge.h = edge.Position.GetManhattanDistance(enp.Position)
			edge.f = edge.g + edge.h

			if isInSlice(openList, edge) {
				isFurther := false
				for _, n := range openList {
					if edge.g > n.g {
						isFurther = true
						break
					}
				}
				if isFurther {
					continue
				}
			}
			openList = append(openList, edge)
		}
	}
	return nil
}

func isInSlice(s []*node, target *node) bool {
	for _, n := range s {
		if n.isEqual(target) {
			return true
		}
	}
	return false
}

// for any data S such that S is a Slice of E's
func reverseSlice[S ~[]E, E any](data S) {
	// for i and j where is 0 and j is the last index of data
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		// swap data[i] and data[j]
		data[i], data[j] = data[j], data[i]
	}
}
