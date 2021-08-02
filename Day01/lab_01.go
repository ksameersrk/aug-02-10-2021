package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main()  {
	rand.Seed(time.Now().UTC().UnixNano())
	target := rand.Intn(100)
	attempts, message, guess, gameOver := 0, "", -1, false

	//fmt.Printf("Target: %d\n", target) //or %v
	fmt.Println("Enter a number between 1 and 100")
	for !gameOver {
		fmt.Scanf("%d", &guess)
		attempts++
		if guess > target {
			message = "Aim Lower"
		} else if guess < target {
			message = "Aim Higher"
		} else {
			message = "You've got it in " +  strconv.Itoa(attempts) + " attempts"
			gameOver = true
		}
		fmt.Println(message)
	}

}
