package pentomino

var pentominoPieces = []piece{
	{name: "I", points: []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}},
	{name: "X", points: []point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}},
	{name: "T", points: []point{{0, 0}, {1, 0}, {2, 0}, {1, 1}, {1, 2}}},
	{name: "U", points: []point{{0, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}}},
	{name: "V", points: []point{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}},
	{name: "W", points: []point{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}}},
	{name: "Z", points: []point{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}}},
	{name: "L", points: []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}}},
	{name: "Y", points: []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 1}}},
	{name: "N", points: []point{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {1, 3}}},
	{name: "P", points: []point{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}}},
	{name: "F", points: []point{{1, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 2}}},
}

var pieceNameToIndex = make(map[string]int)
var IndexToPieceName = make(map[int]string)

type point struct {
	x int
	y int
}

type piece struct {
	name   string
	points []point
}

type choices struct {
	N   string
	Pos []int
}

func rotate(p piece) piece {
	newPoints := make([]point, len(p.points))
	for i, pt := range p.points {
		newPoints[i] = point{-pt.y, pt.x}
	}

	return piece{name: p.name, points: newPoints}
}

func flip(p piece) piece {
	newPoints := make([]point, len(p.points))
	for i, pt := range p.points {
		newPoints[i] = point{-pt.x, pt.y}
	}

	return piece{name: p.name, points: newPoints}
}

func pieceOrientationInfo(n string) (bool, int) {
	if n == "I" || n == "X" {
		return false, 2
	}

	if n == "T" || n == "U" || n == "V" || n == "W" || n == "Z" {
		return false, 4
	}

	return true, 4
}

func genOrientations(p piece) []piece {
	var orientations []piece
	flipable, rots := pieceOrientationInfo(p.name)

	curr := p
	for i := 0; i < rots; i++ {
		curr = rotate(curr)
		orientations = append(orientations, curr)

		if flipable {
			flipped := flip(curr)
			orientations = append(orientations, flipped)
		}
	}

	return orientations
}

func isValidPlacement(i, j, w, h int, o piece) bool {
	for k := range o.points {
		nwx := o.points[k].x + i
		nwy := o.points[k].y + j

		if nwx < 0 || nwx >= w || nwy < 0 || nwy >= h {
			return false
		}
	}

	return true
}

func findPieces(pieces []string) []piece {
	found := make([]piece, 0, len(pieces))
	set := make(map[string]struct{})
	
	for _, name := range pieces {
		set[name] = struct{}{}
	}

	k := 0
	for _, p := range pentominoPieces {
		if _, ok := set[p.name]; ok {
			found = append(found, p)
			pieceNameToIndex[p.name] = k
			IndexToPieceName[k] = p.name
			k++
		}
	}

	return found
}

func GenChoices(w, h int, pieces []string) []choices {
	c := make([]choices, 0)
	p := findPieces(pieces)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			for k := range p {
				orientations := genOrientations(p[k])
				for l := range orientations {
					if isValidPlacement(i, j, w, h, orientations[l]) {
						pos := make([]int, len(orientations[l].points))
						for m := range orientations[l].points {
							pos[m] = (orientations[l].points[m].y + j)*w + (orientations[l].points[m].x + i)
						}

						c = append(c, choices{N: orientations[l].name, Pos: pos})
					}
				}
			}
		}
	}

	return c
}
