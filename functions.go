package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func MakeMap(s string) map[string]rune {
	simvol := make(map[string]rune)
	temp := ""
	count := 0
	j := rune(32)
	for _, data := range s {
		temp += string(data)
		if data == '\n' {
			count++
		}
		if count == 9 {
			count = 0
			simvol[temp[1:len(temp)-1]] = j
			temp = ""
			j++
		}
	}
	return simvol
}

func SplitArray(file string) []string {
	arr := strings.Split(file, "\n")
	str := ""
	count := 0
	result := []string{}
	for _, simvol := range arr {
		if simvol == "" {
			result = append(result, " ")
		} else {
			str += simvol + "\n"
			count++
			if count == 8 {
				str = str[:len(str)-1]
				result = append(result, str)
				str = ""
				count = 0
			}
		}
	}
	for i := 0; i < len(result); i++ {
		splitResult := strings.Split(result[i], "\n")
		for j := 0; j < len(splitResult); j++ {
			if j != len(splitResult)-1 && len(splitResult[j]) != len(splitResult[j+1]) {
				fmt.Println("ERROR")
				return nil
			}
		}
	}

	return result
}

func Convert(text string, ascii map[string]rune) string {
	arr := strings.Split(string(text), "\n")
	start := 0
	result := ""
	for j := 0; j < len(arr[0]); j++ {
		column := ""
		for i := 0; i < len(arr); i++ {
			column += string(arr[i][start:j]) + "\n"
		}
		column = column[:len(column)-1]
		_, ok := ascii[column]
		if ok {
			start = j
			result += string(ascii[column])
		}
	}

	return result + "\n"
}

func Hash(content []byte, hash string) bool {
	newhash := md5.New()
	newhash.Write([]byte(content))
	return hex.EncodeToString(newhash.Sum(nil)) == hash
}

func ReadFile(arg string) (string, map[string]rune) {
	data, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		return "ERROR: read standard.txt file", nil
	}
	hashOfStandard := "ac85e83127e49ec42487f272d9b9db8b"
	hash := Hash(data, hashOfStandard)
	if !hash {
		return "ERROR: standard.txt file changed", nil
	}

	ascii := MakeMap(string(data))

	reverse := "--reverse="
	if len(arg) < len(reverse) || arg[:len(reverse)] != reverse {
		return "Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>", nil
	}

	fileName := arg[len(reverse):]
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "ERROR: read file", nil
	}
	fileData := string(text)
	return fileData, ascii
}

func Run(arg string) {
	text, ascii := ReadFile(arg)
	array := SplitArray(text)
	result := ""
	for _, w := range array {
		if w == "" {
			result += "\n"
			continue
		}
		result += Convert(w, ascii)
	}
	fmt.Print(result)
}
