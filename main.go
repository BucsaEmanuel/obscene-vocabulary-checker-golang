package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordIsTaboo(word string, tabooList []string) bool {
	for _, tabooWord := range tabooList {
		if strings.ToLower(word) == tabooWord {
			return true
		}
	}
	return false
}

func main() {
	// Get filename from user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		var err = file.Close()
		if err != nil {

		}
	}(file)

	tabooWords := makeTabooWordsList(scanner, file)

	scanner = bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		sentence := strings.TrimSpace(scanner.Text())
		switch sentence {
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			sentence = processSentence(sentence, tabooWords)
			fmt.Println(sentence)
		}
	}
}

func processSentence(sentence string, tabooWords []string) string {
	var wordList []string
	hasEndingPunctuation := false
	if strings.HasSuffix(sentence, ".") {
		hasEndingPunctuation = true
	}

	for _, word := range strings.Split(sentence, " ") {
		if strings.HasSuffix(word, ".") {
			word = word[:len(word)-1]
		}
		word = censorIfTaboo(word, tabooWords)
		wordList = append(wordList, word)
	}

	if hasEndingPunctuation {
		return strings.Join(wordList, " ") + "."
	}
	return strings.Join(wordList, " ")
}

func censorWord(word string) string {
	var runeList []rune
	for i := range word {
		runeList = append(runeList, rune(word[i]))
	}

	return strings.Repeat("*", len(runeList))
}

func censorIfTaboo(word string, tabooWords []string) string {
	if wordIsTaboo(word, tabooWords) {
		return censorWord(word)
	} else {
		return word
	}
}

func makeTabooWordsList(scanner *bufio.Scanner, file *os.File) []string {
	var tabooWords []string
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tabooWords = append(tabooWords, strings.ToLower(strings.TrimSpace(scanner.Text())))
	}
	return tabooWords
}
