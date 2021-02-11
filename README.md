# BEDA

[![Build Status](https://travis-ci.org/hyperjumptech/beda.svg?branch=master)](https://travis-ci.org/hyperjumptech/beda)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Get BEDA

```
go get github.com/hyperjumptech/beda
```

## Introduction 

**BEDA** is a golang library to detect differences or similarities between two words or string.
Some time you want to detect whether a string is "the same" or "somehow similar to" another string. 
Suppose your system wants to detect whenever the user is putting bad-word as their user name, or 
to forbid them from using unwanted words in their postings. You need to implement some, *not so easy* , 
algorithm to do this task.

**BEDA**  contains implementation of algorithm for detecting word differences. They are 

1. Levenshtein Distance :  A string metric for measuring the difference between two sequences. [Wikipedia](https://en.wikipedia.org/wiki/Levenshtein_distance)
2. Trigram or n-gram : A contiguous sequence of n items from a given sample of text or speech. [Wikipedia](https://en.wikipedia.org/wiki/N-gram)
3. Jaro & Jaro Winkler Distance : A string metric measuring an edit distance between two sequences. [Wikipedia](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)

**BEDA** is an Indonesia word for "different". 

## Usage

```go
import "github.com/hyperjumptech/beda"

sd := beda.NewStringDiff("The First String", "The Second String")
lDist := sd.LevenshteinDistance()
tDiff := sd.TrigramCompare()
jDiff := sd.JaroDistance()
jwDiff := sd.JaroWinklerDistance(0.1)

fmt.Printf("Levenshtein Distance is %d \n", lDist)
fmt.Printf("Trigram Compare is is %f \n", lDist)
fmt.Printf("Jaro Distance is is %d \n", jDiff)
fmt.Printf("Jaro Wingkler Distance is %d \n", jwDiff)
```

## Algorithms and APIs

String comparison is not so easy. 
There are a couple of algorithm to do this comparison, and each of them yield different result. 
Thus may suited for one purposses compared to the other. 

To understand how and when or which algorithm should benefit your string comparisson quest,
Please read this [String similarity algorithms compared](https://medium.com/@appaloosastore/string-similarity-algorithms-compared-3f7b4d12f0ff).
Read them through, they will help you, a lot.

```go
type StringDiff struct {
    S1 string
	S2 string
}
```

### Levenshtein Distance

LevenshteinDistance is the minimum number of single-character edits
required to change one word into the other, so the result is a positive
integer. The algorithm is sensitive to string length. Which make it more difficult to draw pattern.

Reading :

- [https://github.com/mhutter/string-similarity](https://github.com/mhutter/string-similarity)
- [https://en.wikipedia.org/wiki/Levenshtein_distance](https://en.wikipedia.org/wiki/Levenshtein_distance)

API :

```go
func LevenshteinDistance(s1, s2 string) int
func (sd *StringDiff) LevenshteinDistance() int
```

`s1` is the first string to compare<br>
`s2` is the second string to compare<br>
The closer return value to 0 means the more similar the two words.

Example :

```go
sd := beda.NewStringDiff("abcd", "bc")
lDist := sd.LevenshteinDistance()
fmt.Printf("Distance is %d \n", lDist) // prints : Distance is 2
```

or

```go
fmt.Printf("Distance is %d \n", beda.LevenshteinDistance("abcd", "bc"))
```


### Damerau-Levenshtein Distance

(From [Wikipedia](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance))
Damerau-Levenshtein Distance is a string metric for measuring the edit distance between two 
sequences. Informally, the Damerau–Levenshtein distance between two words is the minimum 
number of operations (consisting of insertions, deletions or substitutions of a single 
character, or transposition of two adjacent characters) required to change one word into the other.

The Damerau–Levenshtein distance differs from the classical Levenshtein distance by 
including transpositions among its allowable operations in addition to the three classical 
single-character edit operations (insertions, deletions and substitutions).

Reading :

- [https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)

API :

```go
func DamerauLevenshteinDistance(s1, s2 string) int
func (sd *StringDiff) DamerauLevenshteinDistance(deleteCost, insertCost, replaceCost, swapCost int) int
```

`func DamerauLevenshteinDistance` take 2 arguments,<br>
`s1` is the first string to compare<br>
`s2` is the second string to compare<br>
The closer return value to 0 means the more similar the two words.
This function uses the default value of 1 for all `deleteCost`, `insertCost`, `replaceCost` and `swapCost`

`func (sd *StringDiff) DamerauLevenshteinDistance` takes 4 arguments,<br>
`deleteCost` is multiplier factor for delete operation<br>
`insertCost` is multiplier factor for insert operation<br>
`replaceCost` is multiplier factor for replace operation<br>
`swapCost` is multiplier factor for swap operation<br>
A multiplier value enable us to weight on how impactful each of the operation 
contributing to the change distance.


Example :

```go
sd := beda.NewStringDiff("abcd", "bc")
lDist := sd.DamerauLevenshteinDistance(1,1,1,1)
fmt.Printf("Distance is %d \n", lDist) // prints : Distance is 2
```

or

```go
fmt.Printf("Distance is %d \n", beda.DamerauLevenshteinDistance("abcd", "bc"))
```


### TriGram Compare

TrigramCompare  is a case of n-gram, a contiguous sequence of n (three, in this case) items from a given sample.
In our case, an application name is a sample and a character is an item.

Reading:

- [https://github.com/milk1000cc/trigram/blob/master/lib/trigram.rb](https://github.com/milk1000cc/trigram/blob/master/lib/trigram.rb)
- [http://search.cpan.org/dist/String-Trigram/Trigram.pm](http://search.cpan.org/dist/String-Trigram/Trigram.pm)
- [https://en.wikipedia.org/wiki/N-gram](https://en.wikipedia.org/wiki/N-gram)

API :

```go
func TrigramCompare(s1, s2 string) float32
func (sd *StringDiff) TrigramCompare() float32
```

`s1` is the first string to compare<br>
`s2` is the second string to compare<br>
The closer the result to 1 (one) means that the word is closer 100% similarities in 3 grams sequence.

Example :

```go
sd := beda.NewStringDiff("martha", "marhta")
diff := sd.TrigramCompare()
fmt.Printf("Differences is %f \n", diff) 
```

or

```go
fmt.Printf("Distance is %f \n", beda.TrigramCompare("martha", "marhta"))
```

### Jaro Distance

JaroDistance distance between two words is the minimum number
of single-character transpositions required to change one word
into the other.

API :

```go
func JaroDistance(s1, s2 string) float32
func (sd *StringDiff) JaroDistance() float32
```

`s1` is the first string to compare<br>
`s2` is the second string to compare<br>
The closer the result to 1 (one) means that the word is closer 100% similarities

Example :

```go
sd := beda.NewStringDiff("martha", "marhta")
diff := sd.JaroDistance()
fmt.Printf("Differences is %f \n", diff) 
```

or

```go
fmt.Printf("Distance is %f \n", beda.JaroDistance("martha", "marhta"))
```

### Jaro Wingkler Distance

JaroWinklerDistance uses a prefix scale which gives more
favourable ratings to strings that match from the beginning
for a set prefix length

Reading :

- [https://github.com/flori/amatch](https://github.com/flori/amatch)
- [https://fr.wikipedia.org/wiki/Distance_de_Jaro-Winkler](https://fr.wikipedia.org/wiki/Distance_de_Jaro-Winkler)
- [https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)

API : 

```go
func JaroWinklerDistance(s1, s2 string) float32
func (sd *StringDiff) JaroWinklerDistance(p float32) float32
```

or

```go
fmt.Printf("Distance is %f \n", beda.JaroWinklerDistance("martha", "marhta"))
```

`s1` is the first string to compare<br>
`s2` is the second string to compare<br>
`p` argument is constant scaling factor for how much the score is adjusted upwards for having common prefixes.
The standard value for this constant in Winkler’s work is `p = 0.1`

The closer the result to 1 (one) means that the word is closer 100% similarities

Example :

```go
sd := beda.NewStringDiff("martha", "marhta")
diff := sd.JaroWinklerDistance(0.1)
fmt.Printf("Differences is %f \n", diff) 
```

# Tasks and Help Wanted.

Yes. We need contributor to make **BEDA** even better and useful to Open Source Community.

If you really want to help us, simply `Fork` the project and apply for Pull Request.
Please read our [Contribution Manual](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCTS.md)