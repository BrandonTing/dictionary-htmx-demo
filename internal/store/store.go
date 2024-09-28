package store

type Dictionary struct {
	List []Vocabulary
}

type Vocabulary struct {
	Word    string
	Example string
}

var Basic_Dictionary Dictionary

func GetDictionary() Dictionary {
	return Basic_Dictionary
}

func AddWord(word string, example string) {
	Basic_Dictionary.List = append(Basic_Dictionary.List, Vocabulary{
		Word:    word,
		Example: example,
	})
}

func DeleteWord(word string) {
	index := 0
	for i := 0; i < len(Basic_Dictionary.List); i++ {
		if Basic_Dictionary.List[i].Word == word {
			index = i
			break
		}
	}
	Basic_Dictionary.List = removeAtIndex(Basic_Dictionary.List, index)
}

func removeAtIndex(source []Vocabulary, index int) []Vocabulary {
	// STEP 1：取得最後一個元素的 index
	lastIndex := len(source) - 1

	// STEP 2：把要移除的元素換到最後一個位置
	source[index], source[lastIndex] = source[lastIndex], source[index]
	// STEP 3：除了最後一個位置的元素其他回傳出去
	return source[:lastIndex]
}
