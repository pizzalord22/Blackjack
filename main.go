package main

import (
	"time"
	"math/rand"
	"fmt"
	"sync"
)

var simulations = float64(100000)
var numDecks int
var scores = []int{0, 0, 0}
var threads = 1000
func main() {
	var wg sync.WaitGroup
	for i := 0; i < threads; i ++ {
		wg.Add(1)
		go playGame(&wg)
	}
	wg.Wait()
	fmt.Printf("win %.2f%s\nequal %.2f%s\nlose %.2f%s", float64(scores[0])/simulations*100, "%", float64(scores[1])/simulations*100, "%", float64(scores[2])/simulations*100, "%")
}

func playGame(wg *sync.WaitGroup) {
	for i := 0; float64(i) < simulations/float64(threads); i++ {
		stdDeck := newDeck()
		playHand(stdDeck)
	}
	wg.Done()
}

func newDeck() []int {
	stdDeck := []int{
		2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11,
		2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11,
		2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11,
		2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11,
	}
	stdDeckTemp := stdDeck

	for i := 0; i < numDecks; i++ {
		for _, v := range stdDeckTemp {
			stdDeck = append(stdDeck, v)
		}
	}

	stdDeck = Shuffle(stdDeck)
	return stdDeckTemp
}

func playHand(stdDeck []int) {
	var dealerCards []int
	var playerCards []int
	var exit bool
	playerCards = append(playerCards, stdDeck[0])
	stdDeck = stdDeck[1:]
	dealerCards = append(dealerCards, stdDeck[0])
	stdDeck = stdDeck[1:]
	playerCards = append(playerCards, stdDeck[0])
	stdDeck = stdDeck[1:]
	dealerCards = append(dealerCards, stdDeck[0])
	stdDeck = stdDeck[1:]

	for getCardScore(playerCards) < 12 {
		playerCards = append(playerCards, stdDeck[0])
		stdDeck = stdDeck[1:]
	}

	for getCardScore(dealerCards) < 18 {
		if getCardScore(dealerCards) == 17 {
			exit = true
			for k, v := range dealerCards {
				if v == 11 {
					exit = false
					dealerCards[k] = 1
				}
			}
		}
		if exit == true {
			break
		}
		dealerCards = append(dealerCards, stdDeck[0])
		stdDeck = stdDeck[1:]
	}

	pSum := getCardScore(playerCards)
	dSum := getCardScore(dealerCards)

	if dSum > 21 {
		scores[0] ++
		return
	}
	if dSum == pSum {
		scores[1] ++
		return
	}
	if dSum > pSum {
		scores[2] ++
		return
	}
	if dSum < pSum {
		scores[0] ++
		return
	}
}

func Shuffle(slice []int) []int {
	rand.Seed(int64(time.Now().UnixNano()))
	for n := len(slice); n > 0; n-- {
		randIndex := rand.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
	return slice
}

func getCardScore(cards []int) (x int) {
	for _, v := range cards {
		x += v
	}
	return
}
