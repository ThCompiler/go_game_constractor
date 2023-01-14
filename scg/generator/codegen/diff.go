package codegen

// Based on https://github.com/pmezard/go-difflib

type Tag byte

const (
	Delete = Tag(iota)
	Replace
	Insert
	Equal
)

type LinesState struct {
	Tag         Tag
	firstStart  int
	firstEnd    int
	secondStart int
	secondEnd   int
}

type Match struct {
	firstIndex  int
	secondIndex int
	Size        int
}

type Differ struct {
	first          []string
	second         []string
	secondIndexes  map[string][]int
	matchingBlocks []Match
	lineStates     []LinesState
}

func NewDiffer(first, second []string) *Differ {
	diff := Differ{}
	diff.SetSequences(first, second)

	return &diff
}

func (diff *Differ) SetSequences(first, second []string) {
	diff.SetFirstSequences(first)
	diff.SetSecondSequences(second)
}

func (diff *Differ) SetFirstSequences(first []string) {
	if &first == &diff.first {
		return
	}

	diff.first = first
	diff.matchingBlocks = nil
	diff.lineStates = nil
}

func (diff *Differ) SetSecondSequences(second []string) {
	if &second == &diff.second {
		return
	}

	diff.second = second
	diff.matchingBlocks = nil
	diff.lineStates = nil
	diff.bChain()
}

func (diff *Differ) bChain() {
	// Populate line -> index mapping
	diff.secondIndexes = map[string][]int{}
	for i, s := range diff.second {
		indices := diff.secondIndexes[s]
		indices = append(indices, i)
		diff.secondIndexes[s] = indices
	}
}

// Find longest matching block in first[firstStart:firstEnd] and second[secondStart:secondEnd].
func (diff *Differ) findLongestMatch(firstStart, firstEnd, secondStart, secondEnd int) Match {
	maxMatch := Match{firstStart, secondStart, 0}

	// Find longest junk-free match
	lens := map[int]int{}

	for i := firstStart; i != firstEnd; i++ {
		// look at all instances of first[i] in second; note that because
		// b2j has no junk keys, the loop is skipped if first[i] is junk
		newLens := map[int]int{}

		for _, j := range diff.secondIndexes[diff.first[i]] {
			// first[i] matches second[j]
			if j < secondStart {
				continue
			}

			if j >= secondEnd {
				break
			}

			k := lens[j-1] + 1
			newLens[j] = k

			if k > maxMatch.Size {
				maxMatch = Match{i - k + 1, j - k + 1, k}
			}
		}

		lens = newLens
	}

	return maxMatch
}

func (diff *Differ) matchBlocks(firstStart, firstEnd, secondStart, secondEnd int, matched []Match) []Match {
	match := diff.findLongestMatch(firstStart, firstEnd, secondStart, secondEnd)
	i, j, k := match.firstIndex, match.secondIndex, match.Size

	if match.Size > 0 {
		if firstStart < i && secondStart < j {
			matched = diff.matchBlocks(firstStart, i, secondStart, j, matched)
		}

		matched = append(matched, match)
		if i+k < firstEnd && j+k < secondEnd {
			matched = diff.matchBlocks(i+k, firstEnd, j+k, secondEnd, matched)
		}
	}

	return matched
}

func (diff *Differ) GetMatchingBlocks() []Match {
	if diff.matchingBlocks != nil {
		return diff.matchingBlocks
	}

	matched := diff.matchBlocks(0, len(diff.first), 0, len(diff.second), nil)

	// It's possible that we have adjacent equal blocks in the
	// matching_blocks list now.
	diff.matchingBlocks = make([]Match, 0)
	old := Match{0, 0, 0}

	for _, second := range matched {
		// Is this block adjacent to i1, j1, k1?
		current := second
		if old.firstIndex+old.Size == current.firstIndex && old.secondIndex+old.Size == current.secondIndex {
			// Yes, so collapse them -- this just increases the length of
			// the first block by the length of the second, and the first
			// block so lengthened remains the block to compare against.
			old.Size += current.Size
		} else {
			// Not adjacent.  Remember the first block (k1==0 means it's
			// the dummy we started with), and make the second block the
			// new block to compare against.
			if old.Size > 0 {
				diff.matchingBlocks = append(diff.matchingBlocks, old)
			}
			old = current
		}
	}

	if old.Size > 0 {
		diff.matchingBlocks = append(diff.matchingBlocks, old)
	}

	diff.matchingBlocks = append(diff.matchingBlocks, Match{len(diff.first), len(diff.second), 0})

	return diff.matchingBlocks
}

func (diff *Differ) GetLineStates() []LinesState {
	if diff.lineStates != nil {
		return diff.lineStates
	}

	i, j := 0, 0
	matches := diff.GetMatchingBlocks()
	diff.lineStates = make([]LinesState, 0, len(matches))

	for _, m := range matches {
		tag := Tag(0)

		switch {
		case i < m.firstIndex && j < m.secondIndex:
			tag = Replace
		case i < m.firstIndex:
			tag = Delete
		case j < m.secondIndex:
			tag = Insert
		}

		if tag > 0 {
			diff.lineStates = append(diff.lineStates, LinesState{tag, i, m.firstIndex, j, m.secondIndex})
		}

		i, j = m.firstIndex+m.Size, m.secondIndex+m.Size

		if m.Size > 0 {
			diff.lineStates = append(diff.lineStates, LinesState{Equal, m.firstIndex, i, m.secondIndex, j})
		}
	}

	return diff.lineStates
}
