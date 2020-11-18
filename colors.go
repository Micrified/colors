package colors

import (
	"math/rand"
	"errors"
)

var codes = []string {
	"#f5f5dc", 		// beige
	"#000000", 		// black
	"#0000ff", 		// blue
	"#a52a2a", 		// brown
	"#00008b", 		// darkblue
	"#008b8b", 		// darkcyan
	"#a9a9a9", 		// darkgrey
	"#006400", 		// darkgreen
	"#bdb76b", 		// darkkhaki
	"#8b008b", 		// darkmagenta
	"#556b2f", 		// darkolivegreen
	"#ff8c00", 		// darkorange
	"#9932cc", 		// darkorchid
	"#8b0000", 		// darkred
	"#e9967a", 		// darksalmon
	"#9400d3", 		// darkviolet
	"#ff00ff", 		// fuchsia
	"#ffd700", 		// gold
	"#008000", 		// green
	"#4b0082", 		// indigo
	"#f0e68c", 		// khaki
	"#00ff00", 		// lime
	"#ff00ff", 		// magenta
	"#800000", 		// maroon
	"#000080", 		// navy
	"#808000", 		// olive
	"#ffa500", 		// orange
	"#ffc0cb", 		// pink
	"#800080", 		// purple
	"#800080", 		// violet
	"#ff0000", 		// red
	"#ffff00", 		// yellow
}


// Returns a sequence of random RGB hex codes of length
// 'count'. If the requested number of colors is too large
// an error is returned
func RandomColors (count int) ([]string, error) {

	// Check count
	if count > len(codes) {
		return []string{}, errors.New("Insufficient colors to accomodate request!")
	}

	// Create a random permutation
	seq := make([]string, len(codes))
	copy(seq, codes)
	rand.Shuffle(len(seq), func(i, j int) { seq[i], seq[j] = seq[j], seq[i] })

	// Return the first 'count'
	return seq[:count], nil
}

