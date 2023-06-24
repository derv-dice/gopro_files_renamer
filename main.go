package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	goproFileRegexp = regexp.MustCompile(`^G[X|H](?P<chapter_num>\d{2})(?P<file_number>\d{4})\.`)
	nums            = make([][]string, 10000, 10000)
)

const (
	rxpFileNum    = "file_number"
	rxpChapterNum = "chapter_num"
)

func main() {
	fmt.Println("Running...")

	err := renameFiles()
	if err != nil {
		fmt.Printf("...Error: %s\n", err.Error())
	}

	fmt.Println("...Success")
}

func renameFiles() (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(pwd)
	if err != nil {
		return err
	}

	for i := range entries {
		if entries[i].IsDir() {
			fmt.Printf("\tSkip directory %s\n", entries[i].Name())
			continue
		}

		rxParams := getRxParams(goproFileRegexp, entries[i].Name())

		if rxParams[rxpChapterNum] == "" || rxParams[rxpFileNum] == "" {
			fmt.Printf("\tSkip file %s\n", entries[i].Name())
			continue
		}

		fileNum, _ := strconv.Atoi(rxParams[rxpFileNum])
		chapterNum, _ := strconv.Atoi(rxParams[rxpChapterNum])

		if len(nums[fileNum]) == 0 {
			nums[fileNum] = make([]string, 10000, 10000)
		}

		nums[fileNum][chapterNum] = entries[i].Name()
	}

	for i := range nums {
		for k := range nums[i] {
			if nums[i][k] != "" {
				oldName := nums[i][k]
				newName := fmt.Sprintf("%d_%d%s", i, k, filepath.Ext(nums[i][k]))
				if err = os.Rename(oldName, newName); err != nil {
					fmt.Printf("\tCan't rename file '%s' --> '%s'\n", oldName, newName)
					continue
				}
			}
		}
	}

	return
}

// getRxParams - Get all regexp params from string with provided regular expression
func getRxParams(rx *regexp.Regexp, str string) (pm map[string]string) {
	if !rx.MatchString(str) {
		return nil
	}
	p := rx.FindStringSubmatch(str)
	n := rx.SubexpNames()
	pm = map[string]string{}
	for i := range n {
		if i == 0 {
			continue
		}

		if n[i] != "" && p[i] != "" {
			pm[n[i]] = p[i]
		}
	}

	return
}
