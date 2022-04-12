package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var newlineMode = false

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func randomNumber(min, max int, seed int64) int {
	rand.Seed(seed)
	return rand.Intn(max-min) + min
}

func getRandomWord(seed int64) string {
	file, err := os.Open("english-quran.txt")
	handleError(err)
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, strings.Split(line, " ")...)
	}
	wordCount := len(words)
	var recieved string = (words[randomNumber(0, wordCount, seed)])
	recieved = strings.Replace(recieved, ",", "", -1)
	recieved = strings.Replace(recieved, ".", "", -1)
	recieved = strings.ToLower(recieved)
	return recieved
}

func main() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "-ln" {
			newlineMode = true
		}
		if os.Args[i] == "-1" {
			fmt.Println(getRandomWord(int64(time.Now().Nanosecond())))
			return
		}
		if os.Args[i] == "-w" {
			var word string = os.Args[i+1]
			fmt.Println(getRandomWord(int64(word[0])))
			return
		}
	}

	if newlineMode {
		for {
			println(getRandomWord(int64(time.Now().Nanosecond())))
			time.Sleep(time.Second * 1)
		}
	} else {
		for {
			print(getRandomWord(int64(time.Now().Nanosecond())) + " ")
			time.Sleep(time.Second * 1)
		}
	}
}
