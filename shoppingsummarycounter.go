package student

// ShoppingSummaryCounter takes a string containing grocery items separated by spaces
// and returns a map summarizing the count of each item.
func ShoppingSummaryCounter(str string) map[string]int {
    // Create a map to store item counts.
    summary := make(map[string]int)

    // Manually split the input string into words using only built-in functions.
    word := ""
    for i := 0; i < len(str); i++ {
        if str[i] != ' ' {
            word += string(str[i])
        } else {
            summary[word]++
            word = ""
        }
    }
    // Add the last word if there is one
    summary[word]++

    // Return the summary map.
    return summary
}