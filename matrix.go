package pentomino

func GenMatrix(w int, h int, pieces []string) [][]bool {
	choices := GenChoices(w, h, pieces)
	matrix := make([][]bool, len(choices))
	
	for i := range matrix {
		matrix[i] = make([]bool, len(pieces) + w*h)
	}

	for j := range choices {
		matrix[j][w*h + pieceNameToIndex[choices[j].N]] = true
		for k := range choices[j].Pos {
			matrix[j][choices[j].Pos[k]] = true
		}
	}

	return matrix
}