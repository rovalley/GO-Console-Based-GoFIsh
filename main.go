/*
NAME
main

DESCRIPTION
This module simulate a game of go fish with a AI player.

Created on March 5, 2024

@author: ryan_ovalley
*/

package main

import (
	"fmt"
)
import "math/rand"
import "time"
import "strings"

// function to create a deck
func createDeck(ranks [13]string, suits [4]string) []string {
	// create deck slice
	var deck []string
	// loop through suits and ranks to create and add cards to deck
	for s := 0; s < len(suits); s++ {
		for r := 0; r < len(ranks); r++ {
			card := ranks[r] + suits[s]
			deck = append(deck, card)
		}
	}
	return deck
}

// function to check if you have four of a kind
func hasFour(hand []string, rank string) bool {
	// create count and set to 0
	count := 0
	// loop through hand
	for c := 0; c < len(hand); c++ {
		card := hand[c]
		// if card is the rank
		if strings.HasPrefix(card, rank) {
			// increment count by 1
			count += 1
		}
	}
	// return count equal 4
	return count == 4
}

// function that removes a card
func remove(s []string, i int) []string {
	// swap value with the last index
	s[i] = s[len(s)-1]
	// return slice without last index
	return s[:len(s)-1]
}

// function to check if the game is over
func gameOver(deck []string, playerHand []string, aiHand []string) bool {
	// if the deck is empty return true
	if len(deck) == 0 {
		return true
	}
	// if the player hand is empty return true
	if len(playerHand) == 0 {
		return true
	}
	// if the ai hand is empty return true
	if len(aiHand) == 0 {
		return true
	}
	return false
}

// function to check if you have a rank
func containsRank(ranks [13]string, rankInput string) bool {
	// loop through ranks
	for i := 0; i < len(ranks); i++ {
		// create rank
		rank := ranks[i]
		// if rank equals rank input
		if rank == rankInput {
			return true
		}
	}
	return false
}

// function to check if you have a card
func hasCard(hand []string, rankInput string) bool {
	// loop through hand
	for i := 0; i < len(hand); i++ {
		// create card
		card := hand[i]
		// if card is the rank input
		if strings.HasPrefix(card, rankInput) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Go Fish Game")
	// create variables
	var playerScore int = 0
	var aiScore int = 0
	ranks := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := [4]string{"H", "D", "S", "C"}
	var deck = createDeck(ranks, suits)
	var playerHand []string
	var aiHand []string
	// set random seeds to shuffle deck
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	// add 7 random cards to player hand
	for i := 0; i < 7; i++ {
		card := deck[0]
		playerHand = append(playerHand, card)
		deck = deck[1:]
	}

	// add 7 random cards to AI hand
	for j := 0; j < 7; j++ {
		card := deck[0]
		aiHand = append(aiHand, card)
		deck = deck[1:]
	}

	// set turn to 0
	var turn = 0

	// loop through ranks
	for r := 0; r < len(ranks); r++ {
		// create rank
		rank := ranks[r]
		// if player hand has four of a kind
		if hasFour(playerHand, rank) {
			fmt.Println("Player has four of a kind: " + rank)
			// increment player score
			playerScore += 1
			// loop through player hand
			for c := 0; c < len(playerHand); c++ {
				// create card from player hand
				card := playerHand[c]
				// if card is the rank
				if strings.HasPrefix(card, rank) {
					// remove from player hand
					playerHand = remove(playerHand, c)
					c -= 1
				}
			}
		}
		// if AI hand has four of a kind
		if hasFour(aiHand, rank) {
			fmt.Println("AI has four of a kind: " + rank)
			// increment AI score
			aiScore += 1
			// loop through AI hand
			for c := 0; c < len(aiHand); c++ {
				// create card from AI hand
				card := aiHand[c]
				// if card is the rank
				if strings.HasPrefix(card, rank) {
					// remove from AI hand
					aiHand = remove(aiHand, c)
					c -= 1
				}
			}
		}
	}
	for !gameOver(deck, playerHand, aiHand) {
		// print in game displays
		fmt.Println()
		fmt.Println(playerHand)
		fmt.Println(aiHand)
		fmt.Printf("Remaining cards: %v", len(deck))
		fmt.Println()
		fmt.Printf("Player score: %v", playerScore)
		fmt.Println()
		fmt.Printf("AI score: %v", aiScore)
		fmt.Println()
		// if the turn is player's turn
		if turn == 0 {
			// create rank input and set to a empty string
			rankInput := ""
			// loop as long as rank is a empty string
			for rankInput == "" {
				fmt.Print("Enter a rank: ")
				// scan rank input
				fmt.Scan(&rankInput)
				// uppercase rank input
				rankInput = strings.ToUpper(rankInput)
				// if ranks does not contain the rank input
				if !containsRank(ranks, rankInput) {
					fmt.Println("Incorrect rank try again")
					// set rank input to an empty string
					rankInput = ""
				}
			}
			// if AI has the rank input
			if hasCard(aiHand, rankInput) {
				// loop through AI hand
				for c := 0; c < len(aiHand); c++ {
					// create card from AI hand
					card := aiHand[c]
					// if card is the rank input
					if strings.HasPrefix(card, rankInput) {
						// add card to player hand
						playerHand = append(playerHand, card)
						// remove card from AI hand
						aiHand = remove(aiHand, c)
						c -= 1
					}
				}
			} else {
				fmt.Println("Go Fish!")
				// get new card
				card := deck[0]
				// add card to player hand
				playerHand = append(playerHand, card)
				// set deck without that card
				deck = deck[1:]
			}
			// set turn to 1
			turn = 1
			// if it is the AI's turn
		} else {
			// create a random index from ranks
			randomIndex := rand.Intn(len(ranks))
			// create rank
			rank := ranks[randomIndex]
			fmt.Println("AI asked for: " + rank)
			// if player hand has the rank
			if hasCard(playerHand, rank) {
				// loop though the player hand
				for c := 0; c < len(playerHand); c++ {
					// create card from player hand
					card := playerHand[c]
					// if card is the rank
					if strings.HasPrefix(card, rank) {
						// add card to the AI hand
						aiHand = append(aiHand, card)
						// remove card from the player hand
						playerHand = remove(playerHand, c)
						c -= 1
					}
				}
			} else {
				fmt.Println("Go Fish!")
				// get new card
				card := deck[0]
				// add card to AI hand
				aiHand = append(aiHand, card)
				// set deck without that card
				deck = deck[1:]
			}
			// set turn to 0
			turn = 0
		}

		// loop through ranks
		for r := 0; r < len(ranks); r++ {
			// create rank
			rank := ranks[r]
			// if player hand has four of a kind
			if hasFour(playerHand, rank) {
				fmt.Println("Player has four of a kind: " + rank)
				// increment player score
				playerScore += 1
				// loop through player hand
				for c := 0; c < len(playerHand); c++ {
					// create card from player hand
					card := playerHand[c]
					// if card is the rank
					if strings.HasPrefix(card, rank) {
						// remove from player hand
						playerHand = remove(playerHand, c)
						c -= 1
					}
				}
			}
			// if AI hand has four of a kind
			if hasFour(aiHand, rank) {
				fmt.Println("AI has four of a kind: " + rank)
				// increment AI score
				aiScore += 1
				// loop through AI hand
				for c := 0; c < len(aiHand); c++ {
					// create card from AI hand
					card := aiHand[c]
					// if card is the rank
					if strings.HasPrefix(card, rank) {
						// remove from AI hand
						aiHand = remove(aiHand, c)
						c -= 1
					}
				}
			}
		}
	}
	// if player score is higher than AI score
	if playerScore > aiScore {
		// print player has won
		fmt.Println("Player has won the game!")
		// if AI score is higher than player score
	} else if aiScore > playerScore {
		// print AI has won
		fmt.Println("AI has won the game!")
	} else {
		// print the game is a tie
		fmt.Println("The game has ended in a tie.")
	}
}
