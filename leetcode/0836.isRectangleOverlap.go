package leetcode

import "log"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 防止数组越界
	if len(rec1) != 4 || len(rec2) != 4 {
		return false
	}

	// rename
	aX1 := rec1[0]
	aY1 := rec1[1]
	aX2 := rec1[2]
	aY2 := rec1[3]

	bX1 := rec2[0]
	bY1 := rec2[1]
	bX2 := rec2[2]
	bY2 := rec2[3]
	log.Printf("a x[%d, %d], y[%d, %d]", aX1, aX2, aY1, aY2)
	log.Printf("b x[%d, %d], y[%d, %d]", bX1, bX2, bY1, bY2)

	// x轴y轴都有交集的情况下，矩形才有交集
	if isLineOverlap(aX1, aX2, bX1, bX2) && isLineOverlap(aY1, aY2, bY1, bY2) {
		log.Printf("all line overlap")
		return true
	}

	return false
}

func isLineOverlap(a1, a2, b1, b2 int) bool {
	// a的任意一个端点在b这条线的中间
	log.Printf("a [%d, %d]", a1, a2)
	log.Printf("b [%d, %d]", b1, b2)
	if (a1 > b1 && a1 < b2) || (a2 > b1 && a2 < b2) {
		log.Printf("a in b, overlap")
		return true
	}
	if (b1 > a1 && b1 < a2) || (b2 > a1 && b2 < a2) {
		log.Printf("b in a, overlap")
		return true
	}
	// a和b完全重合
	if a1 == b1 && a2 == b2 {
		log.Printf("equal")
		return true
	}
	log.Printf("not overlap")

	return false
}
