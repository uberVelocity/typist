package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func generateText(words []string, number int) []string {
	var history []int
	var selected []string
	rand.Seed(time.Now().UnixNano())
	for len(selected) < number {
		pos := rand.Intn(3000)  // 3000 total words to choose from
		history = append(history, pos)
		for randomPos := range history {
			if pos == randomPos {
				continue
			}
		}
		selected = append(selected, words[pos])
	}
	return selected
}

func readWords() []string {
	content, err := ioutil.ReadFile("words")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func showStats(correctness []bool, text []string, elapsed time.Duration, fastestTime float64, fastestTimeWord string) {
	fmt.Println("===== WORDS =====")
	correct := 0
	for index, result := range correctness {
		if result == true {
			color.Green("%s", text[index])
			correct += 1
		} else {
			color.Red("%s", text[index])
		}
	}
	fmt.Println("===== STATS =====")
	fmt.Println("correct / total:", correct, "/", len(correctness))
	color.Red("mistakes: %d", len(correctness) - correct)
	var accuracy float32
	accuracy = 100 * float32(correct) / float32(len(correctness))
	if len(correctness) - correct < len(text) {
		color.Blue("Fastest word: %s (%fs)\n", fastestTimeWord, fastestTime)
	}
	color.Green("Accuracy: %.2f%%\n", accuracy)
	wpm := len(correctness) * 60 / int(elapsed.Seconds())
	color.Yellow("wpm: %d\n", wpm)
}

func gameLoop(text []string) ([]bool, time.Duration, float64, string) {
	correctness := make([]bool, len(text))
	start := time.Now()

	var fastestTime = 420.0
	var fastestWordIndex int
	for index, word := range text {
		fmt.Println(word)
		individualWord := time.Now()
		var input string
		if _, err := fmt.Scanln(&input); err != nil {
			if input == "\n" {
				continue
			}
		}
		timePerWord := time.Since(individualWord)
		if input != word {
			correctness[index] = false
		} else {
			if timePerWord.Seconds() < fastestTime {
				fastestTime = timePerWord.Seconds()
				fastestWordIndex = index
			}
			correctness[index] = true
		}
	}
	elapsed := time.Since(start)
	return correctness, elapsed, fastestTime, text[fastestWordIndex]
}

func main() {
	fmt.Println("Select number of words: (ex: 10, 25, 50, 100)")
	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal(err)
	}

	words := readWords()
	text := generateText(words, input)
	correctness, elapsed, fastestTime, fastestTimeWord := gameLoop(text)
	showStats(correctness, text, elapsed, fastestTime, fastestTimeWord)
}
