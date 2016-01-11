package fuzzy

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func indexOf(chars []rune, char rune) int {
	for i, c := range chars {
		if c == char {
			return i
		}
	}
	return -1
}

func ComputeScore(source string, pattern string) float32 {
	if(source == pattern) {
		return 1.0
	}

	query := []rune(pattern)
	candidate := []rune(source)

	queryLength := len(query)
	candidateLength := len(candidate)

	if queryLength == 0 || candidateLength == 0 {
		return 0.0
	}

	totalCharacterScore := float32(0.0)

	indexInQuery := 0

	for indexInQuery < queryLength {

		character := query[indexInQuery]
		indexInQuery = indexInQuery + 1

   		indexInCandidate := indexOf(candidate, character)

   		if indexInCandidate == -1 {
   			return 0.0
   		}

   		characterScore := float32(0.1)
   		if candidate[indexInCandidate] == character {
   			characterScore += 0.1
   		}

   		if indexInCandidate == 0 {
   			characterScore +=  0.8
   		} else {
	   		preChar := string(candidate[indexInCandidate - 1])
	   		wordSeperator := []string{"-", "_", " "}
		   	for _, b := range wordSeperator {
		        if b == preChar {
		            characterScore += 0.7
		            break
		        }
		    }	
		}

		candidate = candidate[indexInCandidate + 1 : ]

		totalCharacterScore += characterScore
	}

	queryScore := totalCharacterScore / float32(queryLength)
	candidateScore := totalCharacterScore / float32(candidateLength)
	return (candidateScore + queryScore) / 2.0
}