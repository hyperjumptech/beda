package beda

// NewStringDiff will create a new instance of StringDiff
func NewStringDiff(s1, s2 string) *StringDiff {
	return &StringDiff{
		S1: s1,
		S2: s2,
	}
}

// StringDiff is a utility struct to compare similarity between two string.
//
// read https://medium.com/@appaloosastore/string-similarity-algorithms-compared-3f7b4d12f0ff
type StringDiff struct {
	S1 string
	S2 string
}

// LevenshteinDistance is the minimum number of single-character edits
// required to change one word into the other, so the result is a positive
// integer, sensitive to string length .
// Which make it more difficult to draw pattern.
//
// Read https://github.com/mhutter/string-similarity and
// https://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(s1, s2 string) int {
	sd := NewStringDiff(s1, s2)
	return sd.LevenshteinDistance()
}

// LevenshteinDistance is the minimum number of single-character edits
// required to change one word into the other, so the result is a positive
// integer, sensitive to string length .
// Which make it more difficult to draw pattern.
//
// Read https://github.com/mhutter/string-similarity and
// https://en.wikipedia.org/wiki/Levenshtein_distance
func (sd *StringDiff) LevenshteinDistance() int {
	s := []byte(sd.S1)
	t := []byte(sd.S2)
	m := len(s)
	n := len(t)
	// for all i and j, d[i,j] will hold the Levenshtein distance between
	// the first i characters of s and the first j characters of t
	// note that d has (m+1)*(n+1) values
	d := make([][]byte, m+1)
	for i := range d {
		d[i] = make([]byte, n+1)
	}
	// source prefixes can be transformed into empty string by
	// dropping all characters
	for i := 1; i <= m; i++ {
		d[i][0] = byte(i)
	}
	// target prefixes can be reached from empty source prefix
	// by inserting every character
	for j := 1; j <= n; j++ {
		d[0][j] = byte(j)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			var substitutionCost byte
			if s[i] == t[j] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			d[i+1][j+1] = byte(minimum(int(d[i][j+1]+1), // deletion
				int(d[i+1][j]+1),               // insertion
				int(d[i][j]+substitutionCost))) // substitution
		}
	}
	return int(d[m][n])
}

type trigram struct {
	Data []byte
}
type trigramuniqueset struct {
	Set []*trigram
}

func (tus *trigramuniqueset) Add(t *trigram) {
	for _, c := range tus.Set {
		if c.Equals(t) {
			return
		}
	}
	tus.Set = append(tus.Set, t)
}

func (t *trigram) Equals(that *trigram) bool {
	if len(t.Data) != len(that.Data) {
		return false
	}
	for i, b := range t.Data {
		if that.Data[i] != b {
			return false
		}
	}
	return true
}

func maketrigrams(d []byte) []*trigram {
	ret := make([]*trigram, 0)
	if len(d) == 0 {
		return ret
	}
	dd := make([]byte, 0)
	dd = append(dd, []byte(" ")...)
	dd = append(dd, d...)
	dd = append(dd, []byte(" ")...)

	for i := 0; i < len(dd)-2; i++ {
		tg := &trigram{Data: dd[i : i+3]}
		ret = append(ret, tg)
	}
	return ret
}

// TrigramCompare  is a case of n-gram, a contiguous sequence
// of n (three, in this case) items from a given sample.
// In our case, an application name is a sample and a
// character is an item.
func TrigramCompare(s1, s2 string) float32 {
	sd := NewStringDiff(s1, s2)
	return sd.TrigramCompare()
}

// TrigramCompare  is a case of n-gram, a contiguous sequence
// of n (three, in this case) items from a given sample.
// In our case, an application name is a sample and a
// character is an item.
//
// Read https://github.com/milk1000cc/trigram/blob/master/lib/trigram.rb
// Read http://search.cpan.org/dist/String-Trigram/Trigram.pm
// Read https://en.wikipedia.org/wiki/N-gram
func (sd *StringDiff) TrigramCompare() float32 {
	s := []byte(sd.S1)
	t := []byte(sd.S2)
	sSet := maketrigrams(s)
	tSet := maketrigrams(t)
	matching := 0.0
	unique := 0.0
	for _, s := range sSet {
		for _, t := range tSet {
			if s.Equals(t) {
				matching++
				//fmt.Printf("Match '%s'\n", string(s.Data))
			}
		}
	}
	tus := &trigramuniqueset{Set: make([]*trigram, 0)}
	for _, s := range sSet {
		tus.Add(s)
	}
	for _, t := range tSet {
		tus.Add(t)
	}
	unique = float64(len(tus.Set))
	//fmt.Printf("Matching is %f, Unique is %f\n", matching, unique )
	return float32(matching / unique)
}

func minimum(args ...int) int {
	var min int
	for i, v := range args {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

func nonmatching(a, b []byte) int {
	ret := 0
	var s, l []byte
	if len(a) > len(b) {
		l = a
		s = b
	} else {
		l = b
		s = a
	}
	ret += len(l) - len(s)
	for i, ca := range s {
		if l[i] != ca {
			ret++
		}
	}
	return ret
}

func matching(a, b []byte) int {
	var s, l []byte
	if len(a) > len(b) {
		l = a
		s = b
	} else {
		l = b
		s = a
	}
	ret := 0
	for _, ca := range s {
		for _, cb := range l {
			if ca == cb {
				ret++
				break
			}
		}
	}
	return ret
}

// JaroDistance distance between two words is the minimum number
// of single-character transpositions required to change one word
// into the other.
func JaroDistance(s1, s2 string) float32 {
	sd := NewStringDiff(s1, s2)
	return sd.JaroDistance()
}

// JaroDistance distance between two words is the minimum number
// of single-character transpositions required to change one word
// into the other.
func (sd *StringDiff) JaroDistance() float32 {
	s := []byte(sd.S1)
	t := []byte(sd.S2)
	m := float32(matching(s, t))
	tt := float32(nonmatching(s, t)) / 2
	s1 := float32(len(s))
	s2 := float32(len(t))

	dj := (1.0 / 3.0) * ((m / s1) + (m / s2) + ((m - tt) / m))

	return dj
}

// JaroWinklerDistance uses a prefix scale which gives more
// favourable ratings to strings that match from the beginning
// for a set prefix length
//
// p argument is constant scaling factor for how much the score
// is adjusted upwards for having common prefixes.
// The standard value for this constant in Winkler’s work is p=0.1
func JaroWinklerDistance(s1, s2 string, p float32) float32 {
	sd := NewStringDiff(s1, s2)
	return sd.JaroWinklerDistance(p)
}

// JaroWinklerDistance uses a prefix scale which gives more
// favourable ratings to strings that match from the beginning
// for a set prefix length
//
// p argument is constant scaling factor for how much the score
// is adjusted upwards for having common prefixes.
// The standard value for this constant in Winkler’s work is p=0.1
//
// Read https://github.com/flori/amatch
// Read https://fr.wikipedia.org/wiki/Distance_de_Jaro-Winkler
// Read https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func (sd *StringDiff) JaroWinklerDistance(p float32) float32 {
	a := []byte(sd.S1)
	b := []byte(sd.S2)
	dj := sd.JaroDistance()
	sim := 0
	var s, l []byte
	if len(a) > len(b) {
		l = a
		s = b
	} else {
		l = b
		s = a
	}
	for i, c := range s {
		if c == l[i] {
			sim++
			if sim > 4 {
				break
			}
		} else {
			break
		}
	}

	dw := dj + ((p * float32(sim)) * (1.0 - dj))

	return dw
}
