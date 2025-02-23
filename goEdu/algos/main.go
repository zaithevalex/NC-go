package main

import (
	"fmt"
	"slices"
	"strconv"
)

//type Tree struct {
//	content int64
//	parent  *Tree
//	nodes   []*Tree
//}
//
//func AmountDigits(num string) int {
//	digits := 0
//	saveNum, _ := strconv.ParseInt(num, 10, 64)
//	for {
//		if saveNum == 0 {
//			break
//		} else {
//			digits++
//		}
//
//		saveNum /= 10
//	}
//
//	return digits
//}
//
//func InitTree(tree **Tree, t int64) {
//	if t == 1 {
//		return
//	}
//
//	for i := int64(2); i <= 9; i++ {
//		if t%i == 0 {
//			newNode := &Tree{content: i, parent: *tree}
//			(*tree).nodes = append((*tree).nodes, newNode)
//		}
//	}
//
//	for _, node := range (*tree).nodes {
//		InitTree(&node, t/node.content)
//	}
//}
//
//func ByPass(tree *Tree, graph *string, numbers *[]string) {
//	if tree == nil {
//		return
//	}
//
//	if tree.parent != nil {
//		*graph += fmt.Sprintf("node%p[shape=record, label=\"%d\"];\n", tree.parent, tree.parent.content)
//		*graph += fmt.Sprintf("node%p[shape=record, label=\"%d\"];\n", tree, tree.content)
//		*graph += fmt.Sprintf("node%p->node%p[color=black];\n", tree.parent, tree)
//	}
//
//	for _, node := range tree.nodes {
//		ByPass(node, graph, numbers)
//	}
//
//	if len(tree.nodes) == 0 {
//		treeNull := ""
//		bypassPtr := tree
//		for {
//			treeNull += fmt.Sprintf("%d", bypassPtr.content)
//			if bypassPtr.parent.parent == nil {
//				break
//			}
//			bypassPtr = bypassPtr.parent
//		}
//
//		runes := []rune(treeNull)
//		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//			runes[i], runes[j] = runes[j], runes[i]
//		}
//		*numbers = append(*numbers, string(runes))
//	}
//}
//
//func smallestNumber(num string, t int64) string {
//	tree := &Tree{
//		content: 1,
//		parent:  nil,
//	}
//	InitTree(&tree, t)
//
//	var numbers []string
//	graph := "digraph G {\n\trankdir=L\n"
//	ByPass(tree, &graph, &numbers)
//	graph += "}\n"
//
//	inputAmountDigits := AmountDigits(num)
//
//	fmt.Println("GRAPH:")
//	fmt.Println(graph)
//	fmt.Println("NUMBERS:", numbers)
//	fmt.Println("inputAmountDigits:", inputAmountDigits)
//
//	return ""
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

// l1 =  2 4 3
// l2 =  5 6 4
//func sum(l1 *ListNode, l2 *ListNode, sumNode **ListNode, part int) {
//	if l1 == nil && l2 == nil {
//		if part != 0 {
//			*sumNode = &ListNode{part, nil}
//		}
//
//		return
//	}
//
//	var columnSum int
//	if l1 != nil {
//		if l2 != nil {
//			columnSum = l1.Val + l2.Val + part
//		} else {
//			columnSum = l1.Val + part
//		}
//	} else {
//		columnSum = l2.Val + part
//	}
//
//	*sumNode = &ListNode{columnSum % 10, nil}
//	if delta := columnSum / 10; delta != 0 {
//		part = delta
//	} else {
//		part = 0
//	}
//
//	if l1 != nil {
//		if l2 != nil {
//			sum(l1.Next, l2.Next, &(*sumNode).Next, part)
//		} else {
//			sum(l1.Next, nil, &(*sumNode).Next, part)
//		}
//	} else {
//		sum(nil, l2.Next, &(*sumNode).Next, part)
//	}
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	sumNode := &ListNode{}
//	sum(l1, l2, &sumNode, 0)
//
//	return sumNode
//}

//func threeSum(nums []int) [][]int {
//	set := map[string]struct{}{}
//
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			for k := j + 1; k < len(nums); k++ {
//				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
//					set[fmt.Sprintf("%d %d %d", nums[i], nums[j], nums[k])] = struct{}{}
//				}
//			}
//		}
//	}
//
//	var setToSlice [][]int
//	for k := range set {
//		strNumbers := strings.Split(k, " ")
//		num1, _ := strconv.Atoi(strNumbers[0])
//		num2, _ := strconv.Atoi(strNumbers[1])
//		num3, _ := strconv.Atoi(strNumbers[2])
//
//		mn := min(num1, num2, num3)
//		mx := max(num1, num2, num3)
//		av := num1 + num2 + num3 - mn - mx
//		orderedSlice := []int{mn, av, mx}
//
//		flag := false
//		for _, setTS := range setToSlice {
//			if slices.Equal(orderedSlice, setTS) {
//				flag = true
//			}
//		}
//
//		if !flag {
//			setToSlice = append(setToSlice, []int{mn, av, mx})
//			flag = false
//		}
//	}
//
//	return setToSlice
//}
//
//func myPow(x float64, n int) float64 {
//	var result float64 = 1
//
//	minus := 1
//	if n < 0 {
//		minus = -1
//		n = -n
//	}
//
//	if x == 1.0 || x == -1.0 {
//		if n%2 == 0 {
//			return 1
//		}
//		return x
//	}
//
//	if n == 0 {
//		return 1
//	}
//
//	for {
//		if n == 0 {
//			break
//		}
//
//		result *= x
//		n--
//	}
//	if minus < 0 {
//		return 1 / result
//	}
//	return result
//}

//func IsVowel(r rune) bool {
//	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
//		return true
//	}
//	return false
//}

//func vowelStrings(words []string, queries [][]int) []int {
//	var vowelIndexes []int
//	var result []int
//
//	for index, word := range words {
//		if IsVowel(rune(word[0])) && IsVowel(rune(word[len(word)-1])) {
//			vowelIndexes = append(vowelIndexes, index)
//		}
//	}
//
//	for index, query := range queries {
//		fmt.Println(slices.Repeat(query, vowelIndexes[0]))
//	}
//
//	return result
//}

//func divide(dividend int, divisor int) int {
//	minus := 1
//
//	if dividend == -2147483648 {
//		if divisor == -1 {
//			return 2147483647
//		} else if divisor == 1 {
//			return dividend
//		} else if divisor == -2147483648 {
//			return 1
//		}
//	}
//
//	if dividend < 0 {
//		dividend = -dividend
//		minus = -minus
//	}
//
//	if divisor < 0 {
//		divisor = -divisor
//		minus = -minus
//	}
//
//	if divisor == 1 {
//		return dividend * minus
//	}
//
//	result := 0
//	for {
//		if dividend < divisor {
//			break
//		}
//		dividend -= divisor
//		result++
//	}
//
//	return result * minus
//}

//func reverse(x int) int {
//	result := 0
//	dec := 1
//	for {
//		if x == 0 {
//			break
//		}
//
//		result = result*10 + x%10
//		dec *= 10
//		x /= 10
//	}
//
//	if result > 2147483647 || result < -2147483648 {
//		return 0
//	}
//
//	return result
//}

//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	var vals []int
//
//	for {
//		if head == nil {
//			break
//		}
//
//		vals = append(vals, head.Val)
//		head = head.Next
//	}
//
//	slices.Sort(vals)
//
//	var save **ListNode
//	newHead := &ListNode{vals[0], nil}
//	save = &newHead
//	for i := 1; i < len(vals); i++ {
//		(*save).Next = &ListNode{vals[i], nil}
//		save = &(*save).Next
//	}
//
//	return newHead
//}

//func mySqrt(x int) int {
//	if x == 1 {
//		return 1
//	}
//
//	l := 0
//	r := x
//	for l+1 < r {
//		m := (l + r) / 2
//		if m*m >= x {
//			if m*m == x {
//				return m
//			}
//
//			r = m
//		} else {
//			l = m
//		}
//	}
//
//	return l
//}

//func sortColors(nums []int) {
//	for i := 0; i < len(nums); i++ {
//		fmt.Println("BEFORE:", nums)
//		for j := i; j > 0; j-- {
//			if nums[j] < nums[j-1] {
//				temp := nums[j]
//				nums[j] = nums[j-1]
//				nums[j-1] = temp
//				break
//			}
//		}
//		fmt.Println("AFTER:", nums)
//	}
//}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	var list *ListNode
//	var save **ListNode
//
//	if list1 == nil {
//		if list2 != nil {
//			return list2
//		}
//		return nil
//	} else {
//		if list2 == nil {
//			return list1
//		}
//	}
//
//	if list1.Val <= list2.Val {
//		save = &list
//		*save = &ListNode{list1.Val, nil}
//		list1 = list1.Next
//		save = &(*save).Next
//	} else if list2.Val <= list1.Val {
//		save = &list
//		*save = &ListNode{list2.Val, nil}
//		list2 = list2.Next
//		save = &(*save).Next
//	}
//
//	for {
//		if list1 == nil || list2 == nil {
//			break
//		}
//
//		if list1.Val <= list2.Val {
//			*save = &ListNode{list1.Val, nil}
//			list1 = list1.Next
//			save = &(*save).Next
//		} else if list2.Val <= list1.Val {
//			*save = &ListNode{list2.Val, nil}
//			list2 = list2.Next
//			save = &(*save).Next
//		}
//	}
//
//	if list1 == nil {
//		for {
//			if list2 == nil {
//				break
//			}
//
//			*save = &ListNode{list2.Val, nil}
//			list2 = list2.Next
//			save = &(*save).Next
//		}
//	} else if list2 == nil {
//		for {
//			if list1 == nil {
//				break
//			}
//
//			*save = &ListNode{list1.Val, nil}
//			list1 = list1.Next
//			save = &(*save).Next
//		}
//	}
//
//	return list
//}

//func searchRange(nums []int, target int) []int {
//	if len(nums) == 0 {
//		return []int{-1, -1}
//	}
//
//	var result []int
//
//	l := 0
//	r := len(nums) - 1
//	for l < r {
//		m := (l + r) / 2
//		if target > nums[m] {
//			l = m + 1
//		} else {
//			r = m
//		}
//	}
//
//	if nums[0] > target || nums[len(nums)-1] < target || nums[l] != target {
//		return []int{-1, -1}
//	}
//
//	result = append(result, l)
//	l = 0
//	r = len(nums) - 1
//	for l < r {
//		m := (l + r + 1) / 2
//		if target >= nums[m] {
//			l = m
//		} else {
//			r = m - 1
//		}
//	}
//
//	result = append(result, l)
//	return result
//}

//func search(nums []int, target int) bool {
//	if len(nums) == 0 {
//		return false
//	}
//
//	if len(nums) == 1 {
//		if nums[0] == target {
//			return true
//		} else {
//			return false
//		}
//	}
//
//	l := 0
//	r := len(nums) - 1
//	m := 0
//
//	for nums[m] <= nums[m+1] {
//		m = (l + r + 1) / 2
//
//		if nums[m]-nums[l] > 0 {
//			l = m
//		} else {
//			r = m - 1
//		}
//
//		if l > len(nums)-1 {
//			break
//		}
//	}
//
//	if target <= nums[m] && target >= nums[0] {
//		l = 0
//		r = m
//	} else {
//		l = m + 1
//		r = len(nums) - 1
//	}
//
//	for l < r {
//		m = (l + r) / 2
//
//		if target > nums[m] {
//			l = m + 1
//		} else {
//			r = m
//		}
//	}
//
//	if l > len(nums)-1 || nums[l] != target {
//		return false
//	}
//
//	return true
//}

//func searchInsert(nums []int, target int) int {
//	l := 0
//	r := len(nums) - 1
//
//	for l < r {
//		m := (l + r) / 2
//
//		if target > nums[m] {
//			l = m + 1
//		} else {
//			r = m
//		}
//	}
//
//	if nums[l] != target {
//		return -1
//	}
//
//	return l
//}

//func searchMatrix(matrix [][]int, target int) bool {
//	l := 0
//	r := len(matrix)*len(matrix[0]) - 1
//
//	for l < r {
//		m := (l + r) / 2
//
//		if target > matrix[m/len(matrix[0])][m%len(matrix[0])] {
//			l = m + 1
//		} else {
//			r = m
//		}
//	}
//
//	if matrix[l/len(matrix[0])][l%len(matrix[0])] != target {
//		return false
//	}
//	return true
//}

//func search(nums []int, target int) bool {
//	l := 0
//	r := len(nums) - 1
//	m := 0
//
//	for l < r {
//		m = (l + r) / 2
//
//		if target > nums[m] {
//			r = m - 1
//		} else {
//			l = m
//		}
//
//		fmt.Println(l, r, m)
//	}
//
//	return true
//}

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//
//	m := make(map[rune]int)
//	runes := []rune(s)
//
//	maxLen := 0
//	startIndex := 0
//	endIndex := 0
//	for i := 0; i < len(runes); i++ {
//		if _, ex := m[runes[i]]; !ex {
//			m[runes[i]] = i
//			endIndex = i
//		} else {
//			endIndex = i
//
//			if startIndex <= m[runes[i]] {
//				if delta := endIndex - startIndex; delta > maxLen {
//					maxLen = delta
//				}
//
//				startIndex = m[runes[i]] + 1
//				m[runes[i]] = i
//			} else {
//				if delta := endIndex - startIndex; delta > maxLen {
//					maxLen = delta + 1
//				}
//				m[runes[i]] = i
//			}
//		}
//	}
//
//	if delta := endIndex - startIndex; delta >= maxLen {
//		maxLen = delta + 1
//	}
//
//	return maxLen
//}

//func isPalindrome(s string) bool {
//	if len(s)%2 == 0 {
//		left := []rune(s)[0 : len(s)/2]
//		right := []rune(s)[len(s)/2 : len(s)]
//
//		slices.Reverse(right)
//		if string(left) == string(right) {
//			return true
//		} else {
//			return false
//		}
//	} else {
//		left := []rune(s)[0 : len(s)/2]
//		right := []rune(s)[len(s)/2+1 : len(s)]
//
//		slices.Reverse(right)
//		if string(left) == string(right) {
//			return true
//		} else {
//			return false
//		}
//	}
//}

//func longestPalindrome(s string) string {
//	if len(s) == 0 {
//		return ""
//	}
//
//	runes := []rune(s)
//	maxPal := string(runes[0])
//	for i := 0; i < len(runes); i++ {
//		for j := i + 1; j < len(runes); j++ {
//			if res := string([]rune(runes)[i : j+1]); isPalindrome(res) && len(res) > len(maxPal) {
//				maxPal = res
//			}
//		}
//	}
//
//	return maxPal
//}

//	func removeOccurrences(s string, part string) string {
//		stack := ""
//
//		for i := 0; i < len(s); i++ {
//			stack += string(s[i])
//
//			if len(stack)-len(part) >= 0 && stack[len(stack)-len(part):] == part {
//				stack = stack[:len(stack)-len(part)]
//			}
//		}
//		return stack
//	}
//func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
//}

//func findMin(nums []int) int {
//	l := 0
//	r := len(nums) - 1
//	m := 0
//
//	if nums[l] <= nums[r] {
//		return nums[l]
//	}
//
//	for nums[m] < nums[m+1] {
//		m = (l + r) / 2
//
//		if nums[m] < nums[r] {
//			r = m
//		} else {
//			l = m + 1
//		}
//	}
//
//	return nums[m+1]
//}

//func findPeakElement(nums []int) int {
//	if len(nums) == 1 {
//		return 0
//	}
//
//	if len(nums) == 2 {
//		if nums[0] > nums[1] {
//			return 0
//		}
//		return 1
//	}
//
//	if nums[0] < nums[len(nums)-1] && nums[len(nums)-2] < nums[len(nums)-1] {
//		return len(nums) - 1
//	}
//
//	if nums[len(nums)-1] < nums[0] && nums[1] < nums[0] {
//		return 0
//	}
//
//	l := 0
//	r := len(nums) - 1
//	m := 0
//
//	for {
//		m = (l + r) / 2
//
//		if m-1 >= 0 && m+1 < len(nums)-1 && nums[m-1] < nums[m] && nums[m] < nums[m+1] {
//			return m
//		}
//
//		if nums[m-1] < nums[m] && nums[m] < nums[m+1] {
//			l = m + 1
//		} else {
//			r = m
//		}
//	}
//
//	return m
//}

//func fourSum(nums []int, target int) [][]int {
//	slices.Sort(nums)
//
//	if len(nums) < 4 {
//		return [][]int{}
//	}
//
//	if nums[0] == nums[len(nums)-1] {
//		if nums[0] == 0 && nums[0] == target {
//			return [][]int{{0, 0, 0, 0}}
//		}
//	}
//
//	m := make(map[string][]int)
//
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			p1 := j + 1
//			p2 := len(nums) - 1
//			for p1 < p2 {
//				if nums[i]+nums[j]+nums[p1]+nums[p2] < target {
//					p1++
//				} else if nums[i]+nums[j]+nums[p1]+nums[p2] > target {
//					p2--
//				} else {
//					m[fmt.Sprintf("%d%d%d%d", nums[i], nums[j], nums[p1], nums[p2])] = []int{nums[i], nums[j], nums[p1], nums[p2]}
//					p1++
//					p2 = len(nums) - 1
//				}
//			}
//		}
//	}
//
//	var res [][]int
//	for _, v := range m {
//		res = append(res, v)
//	}
//
//	return res
//}

// BAG, 3 test
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if n == 1 {
//		return head.Next
//	}
//
//	start := head
//	for i := 0; i < n; i++ {
//		if head == nil {
//			return start
//		}
//		head = head.Next
//	}
//
//	save := head
//	if save.Next == nil {
//		return start
//	}
//
//	head = head.Next.Next
//	save.Next = head
//
//	return start
//}

//func strStr(haystack string, needle string) int {
//	re := regexp.MustCompilePOSIX(needle)
//	indexes := re.FindStringIndex(haystack)
//
//	if len(indexes) == 0 {
//		return -1
//	}
//	return indexes[0]
//}

//func nextPermutation(nums []int) {
//	pick := len(nums) - 1
//	for i := len(nums) - 2; i >= 0; i-- {
//		if nums[i] >= nums[pick] {
//			pick--
//		} else {
//			break
//		}
//	}
//
//	if pick == 0 {
//		slices.Sort(nums)
//		return
//	}
//
//	for i := len(nums) - 1; i >= pick; i-- {
//		if nums[i] > nums[pick-1] {
//			tmp := nums[i]
//			nums[i] = nums[pick-1]
//			nums[pick-1] = tmp
//			break
//		}
//	}
//
//	slices.Sort(nums[pick:])
//}

func backtracking(nums []int, str string, result *[]string) {
	if len(str) == 4 {
		*result = append(*result, str)
		return
	}

	for i := 0; i < len(nums); i++ {
		str += strconv.Itoa(nums[i])
		backtracking(nums, str, result)
		str = str[:len(str)-1]
	}
}

func toString(nums []int) string {
	conv := ""
	for _, num := range nums {
		conv += strconv.Itoa(num)
	}
	return conv
}

func backtracking2(nums []int, str *[]int, result *[]string) {
	if len(*str) == 4 {
		*result = append(*result, toString(*str))
		return
	}

	for i := 0; i < len(nums); i++ {
		if !slices.Contains(*str, nums[i]) {
			*str = append(*str, nums[i])
		} else {
			continue
		}
		
		backtracking2(nums, str, result)
		*str = (*str)[:len(*str)-1]
	}
}

func main() {
	//num := "10"
	//t := int64(320)

	//minNum := smallestNumber(num, t)
	//fmt.Println(minNum)

	//n := 4 // Общее количество ячеек
	//k := 3 // Количество выбираемых элементов

	//used := make([]bool, n)
	//current := make([]int, 0)
	//
	//generatePermutations(n, k, current, used)

	//l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	//l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	//sumNode := addTwoNumbers(l1, l2)
	//for {
	//	if sumNode == nil {
	//		break
	//	}
	//
	//	fmt.Println(sumNode.Val)
	//	sumNode = sumNode.Next
	//}

	//fmt.Println(threeSum([]int{0, 1, 0}))
	//fmt.Println(myPow(2.0, -2))

	//var a int = -2147483648
	//fmt.Println(a)

	//words := []string{"a", "e", "i"}
	//queries := [][]int{{0, 2}, {0, 1}, {2, 2}}
	//fmt.Println("Vowel Strings:", vowelStrings(words, queries))

	//fmt.Println(divide(7, -3))
	//fmt.Println(reverse(1534236469))
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//rotate(matrix)
	//fmt.Println(matrix)

	//listNode := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	//h := sortList(listNode)
	//
	//for {
	//	if h == nil {
	//		break
	//	}
	//
	//	fmt.Println(h.Val)
	//	h = h.Next
	//}

	//nums := []int{2, 1, 4, 3, 5, 1}
	//sortColors(nums)
	//fmt.Println(nums)

	//nums1 := []int{}
	//nums2 := []int{}
	//var nums []int
	//
	//pointer1 := 0
	//pointer2 := 0
	//for {
	//	if pointer1 == len(nums1) || pointer2 == len(nums2) {
	//		break
	//	}
	//
	//	if nums1[pointer1] <= nums2[pointer2] {
	//		nums = append(nums, nums1[pointer1])
	//		pointer1++
	//	} else if nums2[pointer2] <= nums1[pointer1] {
	//		nums = append(nums, nums2[pointer2])
	//		pointer2++
	//	}
	//}
	//
	//if pointer1 == len(nums1) {
	//	for i := pointer2; i < len(nums2); i++ {
	//		nums = append(nums, nums2[i])
	//	}
	//} else if pointer2 == len(nums2) {
	//	for i := pointer1; i < len(nums1); i++ {
	//		nums = append(nums, nums1[i])
	//	}
	//}
	//
	//fmt.Println(nums)

	//list1 := &ListNode{2, nil}
	//list2 := &ListNode{1, nil}
	//list := mergeTwoLists(list1, list2)
	//
	//for list != nil {
	//	fmt.Println(list.Val)
	//	list = list.Next
	//}

	//fmt.Println(mySqrt(0))

	//fmt.Println(search([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}, 2))

	// 1 2 3 4 5
	// 5 1 2 3 4
	//fmt.Println(search([]int{1, 3}, 4))

	//fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))

	//fmt.Println(lengthOfLongestSubstring(""))

	//fmt.Println(longestPalindrome("jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel"))
	//
	//s := ""
	//s += "ss"
	//fmt.Println(s[0])

	//list1 := &ListNode{10, &ListNode{1, &ListNode{13, &ListNode{6, &ListNode{9, &ListNode{5, nil}}}}}}
	////list2 := &ListNode{1000000, &ListNode{1000001, &ListNode{1000002, nil}}}
	//
	//a := 3
	//b := 4
	//
	//lborder := &list1
	//counter := 0
	//for counter != a {
	//	lborder = &(*lborder).Next
	//	counter++
	//}
	//
	//rborder := &(*lborder)
	//for counter != b {
	//	rborder = &(*rborder).Next
	//	counter++
	//}

	//l2 := &list2
	//for *l2 != nil {
	//	(*lborder).Val = (*l2).Val
	//	(*lborder).Next = (*l2).Next
	//	l2 = &(*l2).Next
	//	lborder = &(*lborder).Next
	//}

	//for *rborder != nil {
	//	//(*lborder).Val = (*rborder).Val
	//	//(*lborder).Next = (*rborder).Next
	//	fmt.Println((*rborder).Val, (*rborder).Next)
	//	rborder = &(*rborder).Next
	//	//lborder = &(*lborder).Next
	//}

	//for list1 != nil {
	//	fmt.Println(list1.Val)
	//	list1 = list1.Next
	//}

	//fmt.Println(findMin([]int{}))
	//fmt.Println(findPeakElement([]int{1, 2, 3, 1}))

	//fmt.Println(fourSum([]int{0, 0, 0, 0}, 1)

	//fmt.Println(strStr("leetcode", "leeto"))

	//arr := []int{4, 2, 3, 1}
	//fmt.Println("BEFORE:", arr)
	//nextPermutation(arr)
	//fmt.Println("AFTER:", arr)

	var result []string
	arr := []int{1, 2, 3, 4}
	//backtracking(arr, "", &result)
	backtracking2(arr, &[]int{}, &result)
	fmt.Println(result)
}
