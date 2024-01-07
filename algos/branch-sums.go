package algos

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	return calculateSums(root, sums, 0)

}

func calculateSums(root *BinaryTree, sums []int, currentTotal int) []int {
	currentTotal += root.Value

	if root.Left == nil && root.Right == nil {
		sums = append(sums, currentTotal)
	}

	if root.Left != nil {
		sums = calculateSums(root.Left, sums, currentTotal)
	}

	if root.Right != nil {
		sums = calculateSums(root.Right, sums, currentTotal)
	}
	return sums
}
