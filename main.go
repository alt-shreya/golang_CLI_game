package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func PrintSnoopy() {
	fmt.Println(`
    ,-~~-.___.
   / |  '     \
  (  )         0    brb writing your name in the Hall of Game!
   \_/-, ,----'
      ====           //
     /  \-'~;    /~~~(O)
    /  __/~|   /       |
  =(  _____| (_________|
`)
}

func main() {
	rand.NewSource(time.Now().Unix())
	answer := rand.Intn(100) + 1
	fmt.Println("Guess which number I'm thinking of! (HINT: it's between 1 and 100)")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter your guess ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, e := strconv.Atoi(input)

		if e != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}

		if guess == answer {
			fmt.Println("You read my mind! ")
			PrintSnoopy()
			break
		} else if guess < answer {
			fmt.Println("Not quite! (HINT: guess higher) ")
		} else {
			fmt.Println("Not quite! (HINT: guess lower) ")
		}
	}
}
