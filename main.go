package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func load_domains() []string {
	file, err := os.Open("/usr/share/partialguesser/domains.txt")
	if err != nil {
		fmt.Println("Unable to read domain file. Should be readable and located at /usr/share/partialguesser/domains.txt")
		os.Exit(1)
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}
	return domains
}

func validate_guess(partial string, comparison string) bool {
	if len(partial) != len(comparison) {
		return false
	}
	for i, c := range partial {
		if string(c) != string('*') {
			if string(comparison[i]) != string(c) {
				return false
			}
		}
	}
	return true
}

func guess_domain(partial string) []string {
	domains := load_domains()
	var possible []string
	for _, val := range domains {
		if validate_guess(partial, string(val)) {

			possible = append(possible, string(val))
		}
	}
	return possible
}

func main() {
	var partial string
	flag.StringVar(&partial, "partial", "nil", "The partial email to guess. (e.g test@g****.com)")
	flag.Parse()
	if flag.NArg() > 0 || partial == "nil" {
		flag.Usage()
	} else if !strings.Contains(partial, "@") {
		flag.Usage()
	} else {
		values := strings.Split(partial, "@")
		possible := guess_domain(values[1])
		for _, guess := range possible {
			fmt.Printf("%s@%s\n", values[0], guess)
		}
	}
}
