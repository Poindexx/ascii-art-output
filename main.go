package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func WriteStrings(filename string, aa []string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return err
	}
	defer file.Close()

	for _, str := range aa {
		_, err := file.WriteString(str)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}
	file.WriteString("\n")

	return nil
}

func massivkaSalu(s string) [][]string {
	mass := []string{}
	mass2 := [][]string{}
	word := ""

	for i := 1; i < len(s); i++ {
		if s[i] == '\n' && s[i-1] == '\n' {
			mass2 = append(mass2, mass)
			mass = []string{}
			word = ""
			continue
		}
		if s[i] == '\n' {
			mass = append(mass, word)
			word = ""
		} else {
			word = word + string(s[i])
		}
	}
	mass2 = append(mass2, mass)
	return mass2
}

func isItRightInMass(ind string, indexes string) bool {
	indexes1 := strings.Split(indexes, "")
	for _, ch := range indexes1 {
		if ind == ch {
			return true
		}
	}
	return false
}

func color_izdeu(s string) string {
	switch s {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	case "reset":
		return "\033[0m"
	case "black":
		return "\033[30m"
	case "orange":
		return "\033[38;5;214m"
	}
	return ""
}

func main() {
	soz := ""
	color := ""
	color_arip := ""
	banner := ""
	otput := ""
	take := []string(os.Args[1:])
	if take[len(take)-1] == "standard" || take[len(take)-1] == "shadow" || take[len(take)-1] == "thinkertoy" {
		if len(take) == 1 {
			soz = take[len(take)-1]
			take = take[:len(take)-1]
		} else {
			banner = take[len(take)-1]
			take = take[:len(take)-1]
			if len(take) != 0 {
				soz = take[len(take)-1]
				take = take[:len(take)-1]
			}
		}
	} else {
		soz = take[len(take)-1]
		take = take[:len(take)-1]
	}
	for i := 0; i < len(take); i++ {
		if len(take[i]) > 7 && take[i][:8] == "--color=" {
			color = take[i][8:]
			take = append(take[:i], take[i+1:]...)
			if len(take) != 0 {
				if len(take[i]) > 8 && take[i][:9] == "--output=" {
					continue
				} else {
					color_arip = take[i]
					take = append(take[:i], take[i+1:]...)
				}
			}
		}
	}
	for i := 0; i < len(take); i++ {
		if len(take[i]) > 8 && take[i][:9] == "--output=" {
			otput = take[i][9:]
			take = append(take[:i], take[i+1:]...)
		}
	}
	if len(take) == 0 {
		if len(color) == 0 {
			color = "st"
		}
		if len(banner) == 0 {
			banner = "standard"
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	aa, err := readFile(BANNER(banner))
	if err != nil {
		fmt.Println("No File")
		fmt.Println("sefs")
		return
	}
	aaa := massivkaSalu(aa)
	if len(aaa) != 95 {
		fmt.Println("Fail not correct")
		return
	}
	mass := []string{}
	word := ""
	soz = strings.Replace(soz, "\n", "\\n", -1)
	for i := 0; i < len(soz); i++ {
		if (soz[i] < 32 || soz[i] > 126) && soz[i] != '\n' {
			return
		}
		if soz[i] == '\n' {
			mass = append(mass, word)
			mass = append(mass, "\\n")
			word = ""
			continue
		}
		if soz[i] == 'n' && i > 0 && soz[i-1] == '\\' {
			mass = append(mass, "\\n")
			word = ""
			continue
		}
		if soz[i] == '\\' && soz[i+1] == 'n' {
			if word != "" {
				mass = append(mass, word)
				word = ""
				continue
			}
		} else {
			word = word + string(soz[i])
		}
	}
	if len(word) > 0 {
		mass = append(mass, word)
	}
	count := 0
	for i := 0; i < len(mass); i++ {
		if mass[i] == "\\n" {
			count++
		}
	}
	if count >= 1 && len(mass) != count {
		for i := 0; i < len(mass); i++ {
			if mass[i] == "\\n" && i != len(mass)-1 && mass[i+1] != "\\n" {
				mass = append(mass[:i], mass[i+1:]...)
			}
		}
	}
	if color != "st" && color_izdeu(color) == "" {
		fmt.Println("no color")
		return
	}
	massiv := []string{}
	for m := 0; m < len(mass); m++ {
		if mass[m] == "\\n" {
			massiv = append(massiv, "\n")
		} else {
			for s := 0; s < 8; s++ {
				for i := 0; i < len(mass[m]); i++ {
					if color == "st" {
						massiv = append(massiv, aaa[mass[m][i]-32][s])
					} else {
						if len(color_arip) > 0 {
							if isItRightInMass(string(mass[m][i]), color_arip) {
								massiv = append(massiv, color_izdeu(color), aaa[mass[m][i]-32][s], color_izdeu("reset"))
							} else {
								massiv = append(massiv, aaa[mass[m][i]-32][s])
							}
						} else {
							massiv = append(massiv, color_izdeu(color), aaa[mass[m][i]-32][s], color_izdeu("reset"))
						}
					}
				}
				massiv = append(massiv, "\n")
			}
		}
	}
	if len(otput) > 0 {
		errr := WriteStrings(otput, massiv)
		if errr != nil {
			fmt.Println("No File")
			return
		}
	} else {
		for _, e := range massiv {
			fmt.Print(e)
		}
	}

}

func BANNER(s string) string {
	switch s {
	case "standard":
		return "standard.txt"
	case "shadow":
		return "shadow.txt"
	case "thinkertoy":
		return "thinkertoy.txt"
	}
	return ""
}
