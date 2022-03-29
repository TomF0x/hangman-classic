package main

import (
	"hangman_classic"
	"math/rand"
	"os"
	"time"
)

import (
	"fmt"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 || len(args) > 5 {
		os.Exit(0)
	}

	randomword, randomwordhide, try, wordlist, usedletter, asciart, listasci, choose, state, savename := "", "", 10, []string{}, []string{}, false, []string{}, "", "", ""

	if len(args) == 3 && args[1] == "--letterFile" {
		asciart = true
		listasci = hangman_classic.Loadressource(args[2], 9)
		rand.Seed(time.Now().UnixNano())
		wordlist = hangman_classic.Loadressource(args[0], 1)
		randomword = wordlist[rand.Intn(len(wordlist))]
		randomwordhide = hangman_classic.GenerateWord(randomword)
	} else if len(args) == 4 && args[0] == "--startWith" && args[2] == "--letterFile" {
		randomword, randomwordhide, try, usedletter = hangman_classic.LoadSave(args)
		savename = args[1]
		asciart = true
		listasci = hangman_classic.Loadressource(args[3], 9)
	} else if len(args) == 2 && args[0] == "--startWith" {
		randomword, randomwordhide, try, usedletter = hangman_classic.LoadSave(args)
		savename = args[1]
	} else {
		rand.Seed(time.Now().UnixNano())
		wordlist = hangman_classic.Loadressource(args[0], 1)
		randomword = wordlist[rand.Intn(len(wordlist))]
		randomwordhide = hangman_classic.GenerateWord(randomword)
	}
	jose := hangman_classic.Loadressource("hangman.txt", 7)
	hangman_classic.Clear()
	fmt.Printf("La partie commence, tu possèdes actuellement %v essais !\n", try)
	if asciart {
		hangman_classic.Printasciiart(randomwordhide, listasci)
	} else {
		fmt.Println(randomwordhide)
	}
	for {
		fmt.Print("Choose: ")
		fmt.Scanln(&choose)
		if choose == "STOP" {
			hangman_classic.Clear()
			fmt.Print("Comment voulez vous appellez votre sauvegarde : ")
			fmt.Scanln(&choose)
			hangman_classic.CreateSave(choose, randomword, randomwordhide, try, usedletter)
			os.Exit(0)
		}
		randomwordhide, state = hangman_classic.Finder(choose, randomword, randomwordhide, &usedletter)
		hangman_classic.Clear()
		if try != 10 {
			fmt.Print(jose[9-try])
		}
		hangman_classic.Message(state, randomword, choose, &try, jose)
		if try <= 0 {
			hangman_classic.Clear()
			hangman_classic.Loose()
			fmt.Printf("Tu as perdu, le mot était : %v", randomword)
			os.Exit(0)
		}
		if asciart {
			hangman_classic.Printasciiart(randomwordhide, listasci)
		} else {
			fmt.Println(randomwordhide)
		}
		if randomwordhide == randomword {
			hangman_classic.Win()
			fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", try, randomword)
			if savename != "" {
				hangman_classic.DeleteSave(savename)
			}
			os.Exit(0)
		}
	}
}
