package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var numberDesired uint64 = 1
	firstDesired := true
	secondDesired := true
	lastDesired := true
	secondEntered := false
	numberEntered := false
	var err error
	possibleArgs := []string{"--help", "-n", "--number", "-f", "-first", "-s", "--second", "-l", "--last"}
	for i, s := range os.Args {
		switch s {
		case "--help":
			help()
			os.Exit(0)
		case "-n", "--number":
			numberDesired, err = strconv.ParseUint(os.Args[i+1], 10, 64)
			if err != nil {
				fmt.Println("Expected positive integer, got '" + os.Args[i+1] + "'. If this is a positive integer, please try a smaller number. Negative integers are not supported.")
				os.Exit(1)
			}
			numberEntered = true
		case "-f", "--first":
			firstDesired, err = strconv.ParseBool(os.Args[i+1])
			if err != nil {
				fmt.Println("Expected boolean, got '" + os.Args[i+1] + "'")
				os.Exit(1)
			}
		case "-s", "--second":
			secondDesired, err = strconv.ParseBool(os.Args[i+1])
			if err != nil {
				fmt.Println("Expected boolean, got '" + os.Args[i+1] + "'")
				os.Exit(1)
			}
			secondEntered = true
		case "-l", "--last":
			lastDesired, err = strconv.ParseBool(os.Args[i+1])
			if err != nil {
				fmt.Println("Expected boolean, got '" + os.Args[i+1] + "'")
				os.Exit(1)
			}
		default:
			if i != 0 {
				throwError := true
				for _, a := range possibleArgs {
					if os.Args[i-1] == a {
						throwError = false
						break
					}
				}
				if throwError {
					fmt.Println("Unexpected argument '" + s + "'")
					os.Exit(1)
				}
			}
		}
	}
	if !numberEntered {
		fmt.Println("Press 1 to generate another Frisian, or press 2 to exit")
	}
	var i uint64
	for i = 0; i < numberDesired; i++ {
		rand.Seed(time.Now().UnixNano())
		firstPieces := []string{"B", "Bl", "Br", "D", "Dr", "F", "Fl", "Fr", "G", "Gl", "Gr", "H", "J", "K", "Kl", "Kr", "L", "M", "N", "P", "Pl", "Pr", "R", "S", "Sl", "St", "Str", "Sp", "Spr", "Sj", "Sk", "Skr", "Sw", "T", "Tr", "W", "", "", "", "", "", "", "", "", ""}
		secondPieces := []string{"a", "e", "i", "o", "u", "y", "â", "ê", "ô", "û", "aa", "ee", "ie", "oo", "aai", "ei", "ooi", "ai", "ú", "á", "é", "ó", "ú", "ieu", "ie", "eu"}
		thirdPieces := []string{"bbe", "bke", "dde", "dse", "dsje", "ffe", "fje", "fke", "ge", "kke", "kje", "lde", "lle", "lje", "lke", "ltsje", "mme", "mke", "nne", "nke", "ppe", "pke", "pkje", "rre", "rke", "sse", "ske", "tte", "tke", "tsje", "vo", "wo", "wke", "t", "r"}
		placeNames := []string{"Ljouwert", "Snit", "Harn", "Frjentsjer", "Ljouwerteradiel", "Frjentsjerteradiel", "Dokkum", "Terp", "Dyk", "Sân"}
		surnameEndPieces := []string{"ma", "ma", "stra", "stra", "inga", "enga"}
		firstPiece := firstPieces[rand.Intn(len(firstPieces))]
		secondPiece := secondPieces[rand.Intn(len(secondPieces))]
		thirdPiece := thirdPieces[rand.Intn(len(thirdPieces))]
		if firstPiece == "" {
			secondPiece = strings.Title(secondPiece)
		}
		firstName := ""
		if firstDesired {
			firstName = firstPiece + secondPiece + thirdPiece
		}
		secondName := ""
		if secondDesired && secondEntered {
			firstPiece = firstPieces[rand.Intn(len(firstPieces))]
			secondPiece = secondPieces[rand.Intn(len(secondPieces))]
			thirdPiece = thirdPieces[rand.Intn(len(thirdPieces))]
			if firstPiece == "" {
				secondPiece = strings.Title(secondPiece)
			}
			secondName = " " + firstPiece + secondPiece + thirdPiece + "s"
		} else if !secondDesired && secondEntered {
		} else if rand.Int()%2 == 0 && !secondEntered {
			firstPiece = firstPieces[rand.Intn(len(firstPieces))]
			secondPiece = secondPieces[rand.Intn(len(secondPieces))]
			thirdPiece = thirdPieces[rand.Intn(len(thirdPieces))]
			if firstPiece == "" {
				secondPiece = strings.Title(secondPiece)
			}
			secondName = " " + firstPiece + secondPiece + thirdPiece + "s"
		}
		surname := ""
		if lastDesired {
			surnameStartPiece := ""
			surnameEndPiece := surnameEndPieces[rand.Intn(len(surnameEndPieces))]
			if surnameEndPiece == "ma" {
				firstPiece = firstPieces[rand.Intn(len(firstPieces))]
				secondPiece = secondPieces[rand.Intn(len(secondPieces))]
				thirdPiece = thirdPieces[rand.Intn(len(thirdPieces))]
				if firstPiece == "" {
					secondPiece = strings.Title(secondPiece)
				}
				surnameStartPiece = firstPiece + secondPiece + thirdPiece
			} else if surnameEndPiece == "stra" {
				surnameStartPiece = placeNames[rand.Intn(len(placeNames))]
			} else {
				firstPiece = firstPieces[rand.Intn(len(firstPieces))]
				secondPiece = secondPieces[rand.Intn(len(secondPieces))]
				thirdPiece = thirdPieces[rand.Intn(len(thirdPieces))]
				if firstPiece == "" {
					secondPiece = strings.Title(secondPiece)
				}
				thirdPiece = strings.TrimSuffix(thirdPiece, "e")
				thirdPiece = strings.TrimSuffix(thirdPiece, "o")
				surnameStartPiece = firstPiece + secondPiece + thirdPiece
			}
			surname = " " + surnameStartPiece + surnameEndPiece
		}
		fmt.Println(firstName + secondName + surname)
		if !numberEntered {
			tty, err := tty.Open()
			if err != nil {
				log.Fatal(err)
			}
			for {
				r, err := tty.ReadRune()
				if err != nil {
					log.Fatal(err)
				}
				if string(r) == "1" {
					break
				} else if string(r) == "2" {
					os.Exit(0)
				}
			}
			i--
		}
		time.Sleep(1)
	}
}

func help() {
	fmt.Print("Usage: genfrisian [OPTIONS]\n" +
		"\n" +
		"-f, --first <bool>	Whether the program should generate a first name (default: true)\n" +
		"-s, --second <bool>	Whether the program should generate a second name (default: random)\n" +
		"-l, --last <bool>	Whether the program should generate a last/surname (default: true)\n" +
		"-n, --number <number>	How many names should be generated (default: Generates new name when 1 is pressed)")
}
