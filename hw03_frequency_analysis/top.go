package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

func Top10(sourceString string) []string {
	sourceString = strings.ToLower(sourceString)
	regexpExp := regexp.MustCompile(`\s`) // panic если regexp невалидно
	regexpEndSymbols := regexp.MustCompile(`[[:punct:]]*?$`)
	regexpBeginSymbols := regexp.MustCompile(`^[[:punct:]]*`)
	//  dirtyMatches := regexpExp.Split(sourceString, -1)

	var result []string

	mapRate := make(map[string]int64)

	for _, match := range regexpExp.Split(sourceString, -1) {
		if len(match) > 0 {
			if len(match) > 0 {
				match = regexpEndSymbols.ReplaceAllString(match, "")
			}
			if len(match) > 1 {
				match = regexpBeginSymbols.ReplaceAllString(match, "")
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
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
