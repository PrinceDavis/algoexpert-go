package algos

func NodeDepths(root *BinaryTree) int {
	depths := []int{}

	depths = collectDepths(root, depths, 0)

	sums := 0

	for _, depth := range depths {
		sums += depth
	}

	return sums
}

func collectDepths(root *BinaryTree, depths []int, currentDepth int) []int {
	depths = append(depths, currentDepth)
	currentDepth += 1
	if root.Left != nil {
		depths = collectDepths(root.Left, depths, currentDepth)
	}
	if root.Right != nil {
		depths = collectDepths(root.Right, depths, currentDepth)
	}

	return depths
}
