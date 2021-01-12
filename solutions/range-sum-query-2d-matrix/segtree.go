/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package range_sum_query_2d_matrix

import "math"

type SegTreeNode struct {
	ur, uc, lr, lc, sum int
	children            [4]*SegTreeNode
}

func buildTree(matrix [][]int, ur, uc, lr, lc int) *SegTreeNode {
	if ur > lr || uc > lc {
		return nil
	}
	root := &SegTreeNode{ur: ur, uc: uc, lr: lr, lc: lc, sum: matrix[ur][uc]}
	if ur == lr && uc == lc {
		return root
	}

	mr, mc := ur+(lr-ur)/2, uc+(lc-uc)/2
	root.children[0] = buildTree(matrix, ur, uc, mr, mc)
	root.children[1] = buildTree(matrix, ur, mc+1, mr, lc)
	root.children[2] = buildTree(matrix, mr+1, uc, lr, mc)
	root.children[3] = buildTree(matrix, mr+1, mc+1, lr, lc)
	sum := 0
	for i := 0; i < 4; i++ {
		if root.children[i] != nil {
			sum += root.children[i].sum
		}
	}
	root.sum = sum

	return root
}

func (s *SegTreeNode) Update(r, c, val int) {
	if s.ur == s.lr && s.uc == s.lc {
		s.sum = val
		return
	}

	mr, mc := s.ur+(s.lr-s.ur)/2, s.uc+(s.lc-s.uc)/2
	if r <= mr && c <= mc {
		s.children[0].Update(r, c, val)
	} else if r <= mr && c > mc {
		s.children[1].Update(r, c, val)
	} else if r > mr && c <= mc {
		s.children[2].Update(r, c, val)
	} else {
		s.children[3].Update(r, c, val)
	}
	sum := 0
	for i := 0; i < 4; i++ {
		if s.children[i] != nil {
			sum += s.children[i].sum
		}
	}
	s.sum = sum
}

func (s *SegTreeNode) Query(ur, uc, lr, lc int) int {
	if s.ur == ur && s.uc == uc && s.lr == lr && s.lc == lc {
		return s.sum
	}

	mr, mc := s.ur+(s.lr-s.ur)/2, s.uc+(s.lc-s.uc)/2
	sum := 0
	if ur <= mr && uc <= mc {
		sum += s.children[0].Query(ur, uc, min(lr, mr), min(lc, mc))
	}
	if ur <= mr && lc > mc {
		sum += s.children[1].Query(ur, max(uc, mc+1), min(lr, mr), lc)
	}
	if lr > mr && uc <= mc {
		sum += s.children[2].Query(max(ur, mr+1), uc, lr, min(lc, mc))
	}
	if lr > mr && lc > mc {
		sum += s.children[3].Query(max(ur, mr+1), max(uc, mc+1), lr, lc)
	}

	return sum
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
