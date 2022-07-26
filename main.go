package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Press 1 to generate another Frisian, or press 2 to exit")
	for {
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
		firstName := firstPiece + secondPiece + thirdPiece
		secondName := ""
		if rand.Int()%2 == 0 {
			firstPiece = firstPieces[rand.Intn(len(firstPieces))]
			secondPiece = secondPieces[rand.Intn(len(secondPieces))]
			thirdPiece = thirdPieces[rand.Intn(len(thirdPieces))]
			if firstPiece == "" {
				secondPiece = strings.Title(secondPiece)
			}
			secondName = " " + firstPiece + secondPiece + thirdPiece + "s"
		}
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
		surname := " " + surnameStartPiece + surnameEndPiece
		fmt.Println(firstName + secondName + surname)
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
	}
}
