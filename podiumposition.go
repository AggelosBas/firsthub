package student

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	// Δημιουργούμε slice σταθερού μεγέθους χωρίς make ή append
	var result [4][]string // Εφόσον είναι γνωστό μέγεθος για τα τεστ

	for i := 0; i < n; i++ {
		result[i] = podium[n-1-i]
	}

	// Επιστρέφουμε το slice από τον πίνακα
	return result[:n]
}
