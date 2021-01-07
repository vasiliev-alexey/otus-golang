package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

var regexpEndSymbols = regexp.MustCompile(`^[[:punct:]]*|[[:punct:]]*?$`)

func Top10(sourceString string) []string {
	sourceString = strings.ToLower(sourceString)

	var result []string

	mapRate := make(map[string]int64)

	for _, match := range strings.Fields(sourceString) {
		if len(match) > 0 {
			if len(match) > 0 {
				match = regexpEndSymbols.ReplaceAllString(match, "")
			}

			if len(match) > 0 {
				mapRate[match]++
			}
		}
	}

	rez := rankByWordCount(mapRate)

	for ind, pair := range rez {
		if ind < 10 {
			result = append(result, pair.Key)
		}
	}
	return result
}

func rankByWordCount(wordFrequencies map[string]int64) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

type Pair struct {
	Key   string
	Value int64
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	}
	return p[i].Value > p[j].Value
}
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
