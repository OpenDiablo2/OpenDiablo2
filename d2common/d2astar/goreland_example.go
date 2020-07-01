package d2astar

// goreland_example.go implements implements Pather for
// the sake of testing.  This functionality forms the back end for
// goreland_test.go, and serves as an example for how to use A* for a graph.

// The Magical World of Goreland, is where Ted Stevens and Al Gore are from.
//
// It is composed of Big Trucks, and a Series of Tubes!
//
// Ok, it is basically just a Graph.
// Nodes are called "Trucks" and they have X, Y coordinates
// Edges are called "Tubes", they connect Trucks, and they have a cost
//
// The key differences between this example and the Tile world:
// 1) There is no grid.  Trucks have arbitrary coordinates.
// 2) Edges are not implied by the grid positions.  Instead edges are explicitly
//    modelled as Tubes.
//
// The key similarities between this example and the Tile world:
// 1) They both use Manhattan distance as their heuristic
// 2) Both implement Pather

//Goreland represents a world of trucks and tubes.
type Goreland struct {
	//	trucks map[int]*Truck		// not needed really
}

//Tube is an edge. They connect Trucks, and have a cost.
type Tube struct {
	from *Truck
	to   *Truck
	Cost float64
}

// A Truck is a Truck in a grid which implements Grapher.
type Truck struct {

	// X and Y are the coordinates of the truck.
	X, Y int

	// array of tubes going to other trucks
	outTo []Tube

	label string
}

// PathNeighbors returns the neighbors of the Truck
func (t *Truck) PathNeighbors() []Pather {

	neighbors := []Pather{}

	for _, tubeElement := range t.outTo {
		neighbors = append(neighbors, Pather(tubeElement.to))
	}
	return neighbors
}

// PathNeighborCost returns the cost of the tube leading to Truck.
func (t *Truck) PathNeighborCost(to Pather) float64 {

	for _, tubeElement := range (t).outTo {
		if Pather((tubeElement.to)) == to {
			return tubeElement.Cost
		}
	}
	return 10000000
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Truck) PathEstimatedCost(to Pather) float64 {

	toT := to.(*Truck)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)

	return r
}

// RenderPath renders a path on top of a Goreland world.
func (w Goreland) RenderPath(path []Pather) string {

	s := ""
	for _, p := range path {
		pT := p.(*Truck)
		s = pT.label + " " + s
	}
	return s
}
