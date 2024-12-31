package main

func FindAllCombinations(input string, substrings []string) [][]string {
	var result [][]string

	var dfs func(input string, currentCombination []string)

	dfs = func(input string, currentCombination []string) {
		// Base case: if the input string is empty, we've found a valid combination
		if len(input) == 0 {
			result = append(result, append([]string{}, currentCombination...))
			return
		}

		// Try each substring to see if it can be the prefix of the input string
		for _, substring := range substrings {
			// Check if the input string starts with the current substring
			if len(input) >= len(substring) && input[:len(substring)] == substring {
				// Add the substring to the current combination
				currentCombination = append(currentCombination, substring)

				// Recur with the remaining part of the input string
				dfs(input[len(substring):], currentCombination)

				// Backtrack and remove the last added substring
				currentCombination = currentCombination[:len(currentCombination)-1]
			}
		}
	}

	// Start the recursion
	dfs(input, []string{})

	return result
}
