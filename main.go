package main

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var xoxsToken string = ""

func main() {
	var backupDir string = "./backup-emojis"
	emojis, err := slack.New(xoxsToken).GetEmoji()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = os.Stat(backupDir)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(backupDir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}

	}
	for key, each := range emojis {
		if strings.Contains(each, "https") {
			fmt.Printf("%s : %s\n", key, each)
			extension := filepath.Ext(each)
			if err = downloadFile(each, backupDir+"/"+key+extension); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func downloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
