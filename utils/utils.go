package utils

// SlicesToMap Takes 2 slices (of equal length) that contain strings
// returns a map derived from taking the first slice's elements as keys and the second slice's as values.
// Such that map[a1[i]] = a2[i] where 0 < i < n.
//
// PARAMS:
//
//	a1 - Slice from which to get keys of map
//	a2 - Slice from which to get values of map
func SlicesToMap(a1 []string, a2 []string) map[string]string {

	res := make(map[string]string)
	for i := 0; i < len(a1); i++ {
		res[a1[i]] = a2[i]
	}

	return res
}

// IsSubset Check if a slice of strings is a subset of another slice of strings.
// If is a subset, returns true and an empty string
// If not a subset, returns false along with the first string found not to be in the superset
//
// PARAMS:
//
//	a1 - smaller slice (the possible subset)
//	a2 - bigger slice (the possible superset)
//
// RETURNS:
//
//	bool - is a1 a subset?
//	string - first string in a1 that is found to not be in a2 (if a1 found to be a subset, this is nil)
func IsSubset(a1 []string, a2 []string) (bool, string) {

	// Make set-type structure from a2
	a2Set := make(map[string]bool)
	for _, s2 := range a2 {
		a2Set[s2] = true
	}

	// Check if all strings in a1 are present in a2
	for _, s1 := range a1 {
		// If a key is not present in the map
		// Golang initializes that key's value to the zero value of it's type
		// e.g. for bool-type values, this will be false
		if a2Set[s1] != true {
			return false, s1
		}
	}

	return true, ""
}
