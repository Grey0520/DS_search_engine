package invertedindex

import (
	"fmt"
	"regexp"
	"strings"
)

// InvertedIndexEntry contains the term followed by the
// number of times it has appeared across all documents
// and an array of documents it is persent in
type InvertedIndexEntry struct {
	Term            string
	Frequency       int
	DocumentListing []int
	Index           []int
}

// InvertedIndex contains a hash map to easily check if the
// term is present and an array of InvertedIndexEntry
type InvertedIndex struct {
	HashMap map[string]*InvertedIndexEntry
	Items   []*InvertedIndexEntry
}

// FindItem returns the position of a given
// Item in an Inverted Index
func (invertedIndex *InvertedIndex) FindItem(Term string) int {
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			return index

		}
	}
	panic("Not Found")
}

// AddItem works by first checking if a given term is already present
// in the inverse index or not by checking the hashmap. If it is
// present it updates the Items by increasing the frequency and
// adding the document it is found in. If it is not present it
// adds it to the hash map and adds it to the items list
func (invertedIndex *InvertedIndex) AddItem(Term string, Document int, pos []int) {
	if invertedIndex.HashMap[Term] != nil {
		FoundItemPosition := invertedIndex.FindItem(Term)
		invertedIndex.Items[FoundItemPosition].Index = pos
		invertedIndex.Items[FoundItemPosition].Frequency++
		invertedIndex.Items[FoundItemPosition].DocumentListing = append(invertedIndex.Items[FoundItemPosition].DocumentListing, Document)
	} else {
		// log.Println("Index item", Term, " does not exist :: creating new item")

		InvertedIndexEntry := &InvertedIndexEntry{
			Term:            Term,
			Frequency:       1,
			DocumentListing: []int{Document},
		}

		invertedIndex.HashMap[Term] = InvertedIndexEntry
		invertedIndex.Items = append(invertedIndex.Items, InvertedIndexEntry)
	}
}

// CreateInvertedIndex initializes an
// empty Inverted Index
func CreateInvertedIndex() *InvertedIndex {
	invertedIndex := &InvertedIndex{
		HashMap: make(map[string]*InvertedIndexEntry),
		Items:   []*InvertedIndexEntry{},
	}
	return invertedIndex
}

// RemoveDuplicates filters out all duplicate
// words from each document
func RemoveDuplicates(wordList []string) []string {
	keys := make(map[string]bool)
	uniqueWords := []string{}

	for _, entry := range wordList {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueWords = append(uniqueWords, entry)
		}
	}

	return uniqueWords
}

// Preprocessing converts each word to lowercase
// TODO: Clean up each word for symbols
func Preprocessing(wordList []string) []string {
	ProcessedWordList := []string{}

	// Convert each string to lowercase from
	// wordList and add to ProcessedWordList
	for _, word := range wordList {
		ProcessedWordList = append(ProcessedWordList, strings.ToLower(word))
	}

	return ProcessedWordList
}

// Tokenize gets the individual words from each
// document and generates a wordlist
func Tokenize(Doc string) []string {
	wordList := []string{}

	// The following regexp finds individual
	// words in a sentence
	r := regexp.MustCompile("[^\\s]+")
	wordList = r.FindAllString(Doc, -1)

	wordList = Preprocessing(wordList)
	wordList = RemoveDuplicates(wordList)

	return wordList
}

// GenerateDocMap creates a hash map of
// each word in the document
func GenerateDocMap(token []string) map[string][]int {
	docMap := make(map[string][]int)

	for ind, word := range token {
		docMap[word] = append(docMap[word], ind)
		//if value == nil {
		//	docMap[word] = append(docMap[word], ind)
		//	if word == "cherry" {
		//		fmt.Println(ind)
		//	}
		//}
	}
	return docMap
}

// GenerateInvertedIndex for each document list
// gets each word as a token, processes it and
// generates a hash map for each document
// using them it then generates the
// inverted index of all words
func GenerateInvertedIndex(DocList []string) InvertedIndex {
	globalDocMap := make([]map[string][]int, 0)

	for _, Doc := range DocList {
		token := Tokenize(Doc)
		docMap := GenerateDocMap(token)
		globalDocMap = append(globalDocMap, docMap)
	}

	// Create an empty inverted index
	invertedIndex := CreateInvertedIndex()
	// Using the generated hash maps add
	// each word to the inverted index
	for DocMapIndex, DocMap := range globalDocMap {
		for DocEntry, bo := range DocMap {
			invertedIndex.AddItem(DocEntry, DocMapIndex, bo)
		}
	}
	return *invertedIndex
}

// Find for a given inverted index and search term
// checks if the term exists and then
// outputs the documents the
// term is in
func Find(index InvertedIndex, searchTerm string) {
	Term := strings.ToLower(searchTerm)

	if index.HashMap[Term] != nil {
		itemPosition := index.FindItem(Term)
		item := index.Items[itemPosition]

		fmt.Println("Found:", searchTerm, "in documents:", item.DocumentListing)
	} else {
		fmt.Println("Not Found:", searchTerm)
	}
}
func Sp(pattern string) []string {
	l := len(pattern)
	s := "()&|"
	var ans []string
	var k string = ""
	for i := 0; i < l; i++ {
		tmp := string(pattern[i])
		if strings.Contains(s, tmp) {
			if k != "" {
				ans = append(ans, k)
				k = ""
			}
			ans = append(ans, tmp)
		} else {
			k += tmp
		}

	}
	if k != "" {
		ans = append(ans, k)
	}
	return ans
}

//func Search(index InvertedIndex, pattern []string) []interface{} {
//	s := "()&|"
//	h := [][]int{}
//	i := []string{}
//	for _, v := range pattern {
//		if !strings.Contains(s, v) {
//			h = append(h, index.HashMap[v].DocumentListing)
//		} else {
//			i = append(i, v)
//		}
//	}
//	return h, i
//}

//func Tostr(arr []int) []string {
//	t := ""
//	for _,v :=range arr {
//		t += string(v)
//	}
//}
