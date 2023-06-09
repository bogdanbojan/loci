package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	gowiki "github.com/trietmn/go-wiki"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Dictionary struct {
	Word       string
	Wordtype   string
	Definition string
}

func main() {
	ShowGUI()

	// fileName := "words.txt"
	// words := extractWords(fileName)
	// apiRequest(words)
}

//	func extractWords(fileName string) []string {
//		file, err := os.Open(fileName)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//
//		var words []string
//		for scanner.Scan() {
//			words = append(words, scanner.Text())
//		}
//
//		if err := scanner.Err(); err != nil {
//			log.Fatal(err)
//		}
//
//		return words
//	}
// func apiRequest(words []string) {
// 	reqURL := "https://api.dictionaryapi.dev/api/v2/entries/en/"
// 	for _, w := range words {
// 		resp, err := http.Get(reqURL + w)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
//
// 		handleLocalResponse(w)
// 		handleServerResponse(resp, w)
// 	}
// }

func handleLocalResponse(word string) ([]string, error) {
	var dict Dictionary
	var definitions []string

	file := "./sql-scripts/sqlite_websters_unabridged_dictionary.sql"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	// Setup for MYSQL with Docker-Compose.
	//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/webster_dictionary")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	res, err := db.Query(fmt.Sprintf(`SELECT entries.definition FROM entries WHERE entries.word = "%s"`, cases.Title(language.AmericanEnglish).String(word)))
	if err != nil {
		return nil, err
	}

	for res.Next() {
		err := res.Scan(&dict.Definition)

		if err != nil {
			return nil, err
		}
		definitions = append(definitions, dict.Definition)
		log.Println(dict.Definition)
	}

	return definitions, nil
}

//
//	func handleServerResponse(resp *http.Response, word string) {
//		defer resp.Body.Close()
//		body, err := io.ReadAll(resp.Body)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//		var r []Response
//		err = json.Unmarshal(body, &r)
//
//		if err != nil {
//			fmt.Println("Could not find word in the dictionary")
//			searchWiki(word)
//		}
//
//		s, _ := json.MarshalIndent(r, "", "\t")
//		fmt.Println(string(s))
//	}

func SearchWiki(word string) (string, error) {
	res, err := gowiki.Summary(word, 5, -1, false, true)
	if err != nil {
		return "", err
	}

	return res, nil
}
