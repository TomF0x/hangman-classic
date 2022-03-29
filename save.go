package hangman_classic

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var Xorstring = "salut_je_suis_la_phrase_qui_xor_cette_save_je_suis_tres_long_cest_normal_salut_les_mentor_je_suis_un_xor"

type Save struct {
	Word     string
	Wordhide string
	Try      int
	Usedlist []string
}

type LoadSavest struct {
	Word     string   `json:"Word"`
	Wordhide string   `json:"Wordhide"`
	Try      int      `json:"Try"`
	Usedlist []string `json:"Usedlist"`
}

func LoadSave(args []string) (string, string, int, []string) {
	datatemp, _ := ioutil.ReadFile(args[1])
	var unxor []byte
	for i := 0; i < len(datatemp); i++ {
		unxor = append(unxor, datatemp[i]^Xorstring[i])
	}
	savedata := LoadSavest{}
	_ = json.Unmarshal(unxor, &savedata)
	return savedata.Word, savedata.Wordhide, savedata.Try, savedata.Usedlist
}

func CreateSave(filename, word, wordhide string, try int, used []string) {
	data := Save{
		Word:     word,
		Wordhide: wordhide,
		Try:      try,
		Usedlist: used}
	file, _ := json.Marshal(data)
	var xor []byte
	for i := 0; i < len(file); i++ {
		xor = append(xor, file[i]^Xorstring[i])
	}
	_ = ioutil.WriteFile(filename+".txt", xor, 0644)
}

func DeleteSave(filename string) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "del", filename)
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux":
		cmd := exec.Command("rm", filename)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
