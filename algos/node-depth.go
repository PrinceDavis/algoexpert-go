package algos

func NodeDepths(root *BinaryTree) int {
	depths := []int{}

	collectDepths(root, depths, 0)

	sums := 0

	for _, depth := range depths {
		sums += depth
	}

	return sums
}

func collectDepths(root *BinaryTree, depths []int, currentDepth int) {
	depths = append(depths, currentDepth)
	currentDepth += 1
	if root.Left != nil {
		calculateSums(root.Left, depths, currentDepth)
	}
	if root.Right != nil {
		calculateSums(root.Right, depths, currentDepth)
	}
}
