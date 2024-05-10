package maps

import "fmt"

func Maps() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Maps")
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Print(websites)
	fmt.Print(websites["Amazon Web Services"])

	websites["Linkedin"] = "https://linkedin.com"
	fmt.Print(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
