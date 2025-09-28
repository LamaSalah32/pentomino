package pentomino

import (
	"fmt"
	"math/rand"
)

type Node struct{
	L, R *Node
	U, D *Node
	C *header
}

type header struct{
	Node
	L, R *header

	S int
	N int
}

func chooseColumn(h *header) *header {
    c := h.R
	s := c.S

	for j := c.R; j != h; j = j.R {
		col := j
		
		if col.S < s {
			c = col
			s = col.S
		}
	}

	return c
}

func cover(h *header) {
	h.R.L = h.L
	h.L.R = h.R

	for i := h.D; i != &h.Node; i = i.D {
		for j := i.R; j != i; j = j.R {
			j.D.U = j.U
			j.U.D = j.D
			j.C.S--
		}
	}
}

func uncover (h *header){
	for i := h.U; i != &h.Node; i = i.U {
		for j := i.L; j != i; j = j.L {
			j.C.S++ 
			j.D.U = j
			j.U.D = j
		}
	}

	h.R.L = h
	h.L.R = h
}

func SolveDLX(h *header, k int, solution []*Node) [][]*Node {
    if h.R == h {
        solCopy := make([]*Node, len(solution))
        copy(solCopy, solution)
        return [][]*Node{solCopy}
    }

    var res [][]*Node
    c := chooseColumn(h)
    cover(c)

    for r := c.D; r != &c.Node; r = r.D {
        solution = append(solution, r)

        for j := r.R; j != r; j = j.R {
            cover(j.C)
        }

        res = append(res, SolveDLX(h, k+1, solution)...)
        solution = solution[:len(solution)-1]

        for j := r.L; j != r; j = j.L {
            uncover(j.C)
        }
    }

    uncover(c)
    return res
}

func BuildDLX(matrix [][]bool) *header {
	w := len(matrix[0])
	
	root := &header{N: -1}
	root.L = root
	root.R = root
	root.U = &root.Node
	root.D = &root.Node
	root.C = root

	headers := make([]*header, w)
	prev := root

	for i := 0; i < w; i++ {
		h := &header{N: i}
		h.C = h
		h.S = 0

		h.U = &h.Node
		h.D = &h.Node

		h.L = prev
		h.R = root
		prev.R = h
		root.L = h

		headers[i] = h
		prev = h
	}


	for _, row := range matrix {
		var first *Node
		var last *Node

		for j, val := range row {
			if val{
				h := headers[j]
				n := &Node{C: h}

				n.D = &h.Node
				n.U = h.Node.U
				h.Node.U.D = n
				h.Node.U = n
				h.S++

				if first == nil {
					first = n
					last = n
					n.L = n
					n.R = n
				} else {
					n.L = last
					n.R = first
					last.R = n
					first.L = n
					last = n
				}
			}
		}
	}

	return root
}

func print(width, height int, solutions [][]*Node) {
	board := make([][]string, height)
	for i := range board {
		board[i] = make([]string, width)
	}

	i := rand.Intn(len(solutions))
	sol := solutions[i]
	for _, node := range sol {
		var ch string
		for j := node; ; j = j.R {
			if j.C.N >= width*height {
				ch = IndexToPieceName[j.C.N - width*height]
				break
			}

			if j.R == node {
				break
			}
		}
		
		for j := node; ; j = j.R {
			if j.C.N < width*height {
				pos := j.C.N
				x := pos % width
				y := pos / width
				board[y][x] = ch
			}

			if j.R == node {
				break
			}
		}
	}

	for i := range board {
		fmt.Println(board[i])  
	}
}

func Solve(width, height int, pieces []string) {
	matrix := GenMatrix(width, height, pieces)
	root := BuildDLX(matrix)
	solutions := SolveDLX(root, 0, nil)

	if len(solutions) > 0{
		print(width, height, solutions)
	} else {
		println("No solutions found")
	}
}