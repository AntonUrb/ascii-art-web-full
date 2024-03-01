package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/asciiart", asciiArtHandler)
	fmt.Println("Starting server at http://localhost:8080")
	fmt.Println("Press CTRL + C to shut down server.")

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, "")
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	t, _ := template.ParseFiles("templates/index.html")
	textA := r.FormValue("text")
	font := r.FormValue("fonts")
	text := ""

	// Check if pressed enter for new line
	if strings.Contains(textA, "\r\n") {
		text = strings.ReplaceAll(textA, "\r\n", "\n")
	} else {
		text = textA
	}

	// Check for proper input
	for _, v := range text {
		if (v < 32 || v > 126) && v != 10 {
			http.Error(w, "ERROR-400\nBad request!", http.StatusBadRequest)
			return
		}
	}
	file, err := os.Open("./banners/" + font + ".txt")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)

	// ID the letters
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	asciiChars := make(map[rune][]string)

	decimal := rune(31)
	for _, line := range lines {
		if line == "" {
			decimal++
		} else {
			asciiChars[decimal] = append(asciiChars[decimal], line)
		}
	}

	print := printArt(text, asciiChars)

	// Check if the download button was clicked
	if r.FormValue("download") == "true" {
		w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(print)))
		w.Write([]byte(print))
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(print)))
	t.Execute(w, print)

}

func printArt(str string, asc map[rune][]string) string {
	slice := ""
	temp := strings.Split(str, "\n")

	for _, v := range temp {
		for i := 0; i < 8; i++ {
			for _, letter := range v {
				slice += asc[letter][i]
			}
			slice += "\n"
		}
	}
	return slice
}
