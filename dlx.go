package pentomino

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
