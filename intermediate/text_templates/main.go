package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))
	// // if err != nil {
	// // 	panic(err)
	// // }
	//
	// // Define data for the welcome message template
	// data := map[string]any{
	// 	"name": "pdp0w",
	// }
	//
	// if err := tmpl.Execute(os.Stdout, data); err != nil {
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// define name template for different types of
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad to joined.",
		"notification": "{{.name}}, you have a new notification: {{.notify}}",
		"error":        "Oops! An error occured: {{.errMessage}}",
	}

	// Parse and store templates
	parseTemplate := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parseTemplate[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// show menu
		fmt.Println("\n Menu:")
		fmt.Println("1. John")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")

		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parseTemplate["welcome"]
			data = map[string]any{
				"name": name,
			}
		case "2":
			fmt.Println("Enter you notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parseTemplate["notification"]
			data = map[string]any{"name": name, "notify": notification}
		case "3":
			fmt.Println("Enter you error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parseTemplate["error"]
			data = map[string]any{"name": name, "errMessage": errorMessage}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid options")
			continue
		}

		// render and print template to the console
		if err := tmpl.Execute(os.Stdout, data); err != nil {
			fmt.Println("Error executing template :", err)
		}

	}

}
