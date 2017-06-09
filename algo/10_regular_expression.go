package algo

func isMatch(s string, p string) bool {
	ns := len(s)
	np := len(p)
	// make dp with two dimension array
	// +1 for not handling index out of range
	dp := make([][]bool, ns+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, np+1)
	}
	dp[0][0] = true

	// handle p is all by star case
	for i := 0; i < np; i++ {
		if p[i] == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}
	// dp[i][j] = true if
	// 1. dp[i-1][j-1]&& s[i] == p[j] || p[j]=='.'
	// p[j]== *
	// 1.  count as 0 time   (p[j-1]!= s[i]) dp[i][j] = dp[i][j-2] ;
	// 2.  count as 1 time   (p[j-1]==s[i] || p[j-1]=='.') dp[i][j] = dp[i][j-1]
	// 3.  count as n time   (s[i]==s[i-1] && p[j-1]==s[i]) dp[i][j] = dp[i-1][j]
	for i := 0; i < ns; i++ {
		for j := 0; j < np; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				if p[j-1] != '.' && p[j-1] != s[i] {
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1])
				}
			}
		}
	}
	return dp[ns][np]
}
