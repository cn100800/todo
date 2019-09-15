package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func addUrl(id string, url string) error {
	os.Chdir(workDir)
	f, err := os.OpenFile(urlFile, os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	if err != nil {
		return nil
	}
	//fmt.Println(f.Stat())
	//a, _ := f.Stat()

       	//fmt.Println(f.Read([a.Size()]byte))
	return nil
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	f.WriteString(id + "|" + url + "\n")
	return nil
}

func removeUrl(id int, url string) error {
	return nil
}

func infoUrl(id int) error {
	return nil
}
