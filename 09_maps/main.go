package main

import "fmt"

func main() {
	fmt.Println("maps in go")

	var languages = make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["py"])

	delete(languages, "rb")
	fmt.Println(languages)

	// go lang provides for loop to deal with maps
	for key, value := range languages {
		fmt.Println("for key", key, "value is", value)
	}
}
