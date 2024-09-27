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
	Basic_Dictionary.AddWord("test", "test is hard")
	return Basic_Dictionary
}

func (d *Dictionary) AddWord(word string, example string) {
	d.List = append(d.List, Vocabulary{
		Word:    word,
		Example: example,
	})
}

func (d *Dictionary) DeleteWord(word string) {
	index := 0
	for i := 0; i < len(d.List); i++ {
		if d.List[i].Word == word {
			index = i
			break
		}
	}
	removeAtIndex(d.List, index)
}

func removeAtIndex(source []Vocabulary, index int) []Vocabulary {
	// STEP 1：取得最後一個元素的 index
	lastIndex := len(source) - 1

	// STEP 2：把要移除的元素換到最後一個位置
	source[index], source[lastIndex] = source[lastIndex], source[index]

	// STEP 3：除了最後一個位置的元素其他回傳出去
	return source[:lastIndex]
}
