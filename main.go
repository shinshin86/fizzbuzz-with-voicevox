package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func fizzbuzz() error {
	file, err := os.Create("fizzbuzz.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Fprint(writer, "ふぃず")
		}

		if i%5 == 0 {
			fmt.Fprint(writer, "ばず")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Fprint(writer, i)
		}

		fmt.Fprint(writer, "、")
	}

	return nil
}

func makeQuery() error {
	b, err := ioutil.ReadFile("fizzbuzz.txt")
	if err != nil {
		return err
	}

	query := url.QueryEscape(string(b))
	data := "speaker=1&text=" + query

	url := "http://localhost:50021/audio_query?" + data

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	err = ioutil.WriteFile("query.json", body, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Success: Write query.json")

	return nil
}

func makeAudio() error {
	query, err := ioutil.ReadFile("query.json")
	if err != nil {
		return err
	}

	url := "http://localhost:50021/synthesis?speaker=1"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("audio.wav", body, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Success: Write audio.wav")

	return nil
}

func main() {
	err := fizzbuzz()
	if err != nil {
		fmt.Println(err)
	}

	err = makeQuery()
	if err != nil {
		fmt.Println(err)
	}

	err = makeAudio()
	if err != nil {
		fmt.Println(err)
	}
}
