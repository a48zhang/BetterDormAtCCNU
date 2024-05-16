package pdfgen

import (
	"context"
	"encoding/json"
	"errors"
	"main/pkg/conf"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type PdfForm struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	School  string `json:"school"`
	Sid     string `json:"sid"`
	Contact string `json:"contact"`
	From    string `json:"from"`
	To      string `json:"to"`
	Text0   string `json:"text0"`
	Text1   string `json:"text1"`
	Text2   string `json:"text2"`
	Text3   string `json:"text3"`
	Text4   string `json:"text4"`
}

func GenPDFForForm(ctx context.Context, p PdfForm) (string, error) {
	if !CleanerWorking {
		go WaitListCleaner()
	}
	url := "http://localhost:5000/gen/"
	p.ID = strconv.Itoa(time.Now().Nanosecond())
	Waitlist[p.ID] = 1
	res, _ := json.Marshal(p)
	_, err := http.Post(url, "application/json", strings.NewReader(string(res)))
	if err != nil {
		return "", err
	}

	start := time.Now()
	for Waitlist[p.ID] != 2 {
		if time.Since(start) > 10*time.Second {
			Waitlist[p.ID] = 4
			return "", errors.New("PDF generator timeout")
		}
	}

	return p.ID, nil
}

func GetFilePath(id string) string {
	return "./" + conf.GetConf("bd_pdfgen_path") + id + ".pdf"
}

var Waitlist = make(map[string]int)

var CleanerWorking = false

func WaitListCleaner() {
	CleanerWorking = true
	for {
		time.Sleep(10 * time.Minute)
		for k, v := range Waitlist {
			if v >= 4 {
				err := os.Remove("./pdfgen/" + k + ".pdf")
				if err != nil {
					continue
				}
				delete(Waitlist, k)
			}
		}
	}

}
