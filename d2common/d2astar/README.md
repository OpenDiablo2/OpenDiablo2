D2 A*
========

**A\* pathfinding implementation for OpenDiablo2**

***Forked from [go-astar](https://github.com/beefsack/go-astar)***

Changes
-------
* Used [sync.Pool](https://golang.org/pkg/sync/#Pool) to reuse objects created during path-finding.  This improves performance by roughly 30% by reducing allocations.
* Added a max cost to prevent searching the entire region for a path.
* If there is no path the target within the max cost, the path found that gets closest to target will be returned.  This allows the player to click in inaccessible areas causing the character to run along the edge.

TODO
------
* Evaluate bi-directional A*, specifically if it would more quickly identify if the user clicked an in inaccessible area (such as an island).


The [A\* pathfinding algorithm](http://en.wikipedia.org/wiki/A*_search_algorithm) is a pathfinding algorithm noted for its performance and accuracy and is commonly used in game development.  It can be used to find short paths for any weighted graph.

A fantastic overview of A\* can be found at [Amit Patel's Stanford website](http://theory.stanford.edu/~amitp/GameProgramming/AStarComparison.html).

Examples
--------

The following crude examples were taken directly from the automated tests.  Please see `path_test.go` for more examples.

### Key

*   `.` - Plain (movement cost 1)
*   `~` - River (movement cost 2)
*   `M` - Mountain (movement cost 3)
*   `X` - Blocker, unable to move through
*   `F` - From / start position
*   `T` - To / goal position
*   `●` - Calculated path

### Straight line

```
.....~......      .....~......
.....MM.....      .....MM.....
.F........T.  ->  .●●●●●●●●●●.
....MMM.....      ....MMM.....
............      ............
```

### Around a mountain

```
.....~......      .....~......
.....MM.....      .....MM.....
.F..MMMM..T.  ->  .●●●MMMM●●●.
....MMM.....      ...●MMM●●...
............      ...●●●●●....
```

### Blocked path

```
............      
.........XXX
.F.......XTX  ->  No path
.........XXX
............
```

### Maze

```
FX.X........      ●X.X●●●●●●..
.X...XXXX.X.      ●X●●●XXXX●X.
.X.X.X....X.  ->  ●X●X.X●●●●X.
...X.X.XXXXX      ●●●X.X●XXXXX
.XX..X.....T      .XX..X●●●●●●
```

### Mountain climber

```
..F..M......      ..●●●●●●●●●.
.....MM.....      .....MM...●.
....MMMM..T.  ->  ....MMMM..●.
....MMM.....      ....MMM.....
............      ............
```

### River swimmer

```
.....~......      .....~......
.....~......      ....●●●.....
.F...X...T..  ->  .●●●●X●●●●..
.....M......      .....M......
.....M......      .....M......
```

Usage
-----

### Import the package

```go
import "github.com/beefsack/go-astar"
```

### Implement Pather interface

An example implementation is done for the tests in `path_test.go` for the Tile type.

The `PathNeighbors` method should return a slice of the direct neighbors.

The `PathNeighborCost` method should calculate an exact movement cost for direct neighbors.

The `PathEstimatedCost` is a heuristic method for estimating the distance between arbitrary tiles.  The examples in the test files use [Manhattan distance](http://en.wikipedia.org/wiki/Taxicab_geometry) to estimate orthogonal distance between tiles.

```go
type Tile struct{}

func (t *Tile) PathNeighbors() []astar.Pather {
	return []astar.Pather{
		t.Up(),
		t.Right(),
		t.Down(),
		t.Left(),
	}
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return to.MovementCost
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return t.ManhattanDistance(to)
}
```

### Call Path function

```go
// t1 and t2 are *Tile objects from inside the world.
path, distance, found := astar.Path(t1, t2)
if !found {
	log.Println("Could not find path")
}
// path is a slice of Pather objects which you can cast back to *Tile.
```

Authors
-------

Michael Alexander <beefsack@gmail.com>
Robin Ranjit Chauhan <robin@pathwayi.com>
