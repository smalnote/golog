package leetcode

// Given n pairs of parentheses, wirte a function
// to generate all combinations of well-formed parentheses.
// For example, given n = 3, a solution set is:
// [ "((()))", "(()())", "(())()", "()(())" , "()()()" ]
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	const le, rh = '(', ')'
	ss := make([]string, n)

	for i := 0; i < 2*n; i++ {

	}

	return ss
}
