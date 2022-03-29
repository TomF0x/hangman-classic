package hangman_classic

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func Message(state, randomword, choose string, try *int, jose []string) {
	if state == "fail" {
		*try--
		Clear()
		fmt.Print(jose[9-*try])
		fmt.Printf("La lettre %v n'est pas comprise dans le mot, il ne te reste plus que : %v essais\n", Nomalize(choose), *try)
	} else if state == "usedletter" {
		fmt.Printf("Lettre déjà utiliser\n")
	} else if state == "good" {
		fmt.Printf("La lettre %v est bien comprise dans le mot\n", Nomalize(choose))
	} else if state == "wordinvalid" {
		fmt.Printf("Le format n'est pas valide, veuillez rentrer une lettre ou un mot de bonne taille\n")
	} else if state == "wordgood" {
		Win()
		fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", *try, randomword)
		os.Exit(0)
	} else if state == "error" {
		fmt.Println("La lettre est invalide, veuillez recommencer")
	} else if state == "wordwrong" {
		*try -= 2
		Clear()
		fmt.Print(jose[9-*try])
		fmt.Printf("Le mot proposé n'est pas le bon, il te reste %v essais\n", *try)
	}
}

func Printasciiart(wordtoprint string, listascii []string) {
	for i := 0; i < 9; i++ {
		for _, letter := range wordtoprint {
			newletter := Nomalize(string(letter))
			//newletter = strings.ToUpper(newletter) Ascii Maj
			fmt.Print(Split(listascii[newletter[0]-32], "\n")[i])
		}
		fmt.Print("\n")
	}
}

func Loose() {
	AsciiDeath := Loadressource("loose.txt", 1)
	for _, col := range AsciiDeath {
		fmt.Println(col)
	}
}

func Win() {
	for i := 0; i < 150; i++ {
		Clear()
		Firework := Loadressource("Firework/"+strconv.Itoa(i)+".txt", 1)
		for _, col := range Firework {
			fmt.Println(col)
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
