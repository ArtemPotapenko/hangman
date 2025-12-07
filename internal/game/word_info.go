package game

type WordInfo struct {
	wordLength        int
	hiddenWord        []rune
	guessLittersCount int
	guessLitters      []bool
	numberLetterMap   map[rune][]int
}

func newWordInfo(word string) *WordInfo {
	runes := []rune(word)
	n := len(runes)
	numberLetterMap := make(map[rune][]int)
	for i, letter := range runes {
		if _, ok := numberLetterMap[letter]; !ok {
			numberLetterMap[letter] = make([]int, 0, 2)
		}
		numberLetterMap[letter] = append(numberLetterMap[letter], i)
	}
	return &WordInfo{
		wordLength:        n,
		hiddenWord:        runes,
		guessLitters:      make([]bool, n),
		guessLittersCount: 0,
		numberLetterMap:   numberLetterMap}
}

func (gameWord *WordInfo) guessLetter(char rune) bool {
	if numbers, ok := gameWord.numberLetterMap[char]; ok {
		gameWord.guessLittersCount += len(numbers)
		for _, number := range numbers {
			gameWord.guessLitters[number] = true
		}
		return true
	} else {
		return false
	}
}
