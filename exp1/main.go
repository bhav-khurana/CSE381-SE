package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"github.com/ayush6624/go-chatgpt"
	"github.com/gofor-little/env"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Load data from file
	data, err := ioutil.ReadFile("exp1/data.txt")
	check(err)
	dataList := strings.Split(string(data), "\n")

	// Load the environment variables
	err = env.Load(".env")
	check(err)
	key := env.Get("OPENAI_API_KEY", "")

	// Set up the API client
	client, err := chatgpt.NewClient(key)
	check(err)
	messages := []chatgpt.ChatMessage{
		{
			Role: chatgpt.ChatGPTModelRoleUser,
			Content: "1. AIM: To extract design elements from a set of user stories. 2. THEORY: A user story is a tool in Agile software development used to capture a description of a software feature from a user's perspective. The user story describes the type of user, what they want and why.A user story helps to create a simplified description of a requirement. • Template: As a <type of user>, I want <some goal> so that <some reason>. • Example: As a faculty member, I want to mark attendance of students so that I can track their regularity. The design elements required to be modeled through UML diagrams are contained in these user stories and need to be extracted so as to further put in some UML diagram. The design elements, this experiment aims to extract are concepts/ classes, attributes, relationships (<subject> <predicate> <object>). 3. PROCEDURE: Generally, the noun terms are extracted to be the concepts/ classes, verbs are extracted to be relationships (predicates) and adjectives are extracted to be attributes. You are required to apply your intuition and knowledge also before putting an extracted element in to these categories. 4. RESULTS: Input: The set of user stories provided to you as per the group your roll no. falls into. Output: Excel file containing design elements (concepts/ classes, attributes, relationships (<subject> <predicate> <object>) extracted from the user stories. Sample Input: User story 1: As a faculty member, I want to mark attendance of students so that I can track their regularity. Sample Output: User Story 1: As a faculty member, I want to mark attendance of students so that I can track their regularity. Classes FacultyMember Student Attributes Attribute Associated Class Attendance Student Regularity Student Relationships Subject Predicate Object FacultyMember mark Attendance FacultyMember markAttendanceOf Student FacultyMember track Regularity FacultyMember trackRegularityOf Student.        This is the experiment .. i will give you the input (user story), provide me the desired output in json format (include the user story serial number also in ALL the subsequent outputs). ONLY JSON OUTPUT NO OTHER TEXT. INCLUDE THE SERIAL NUMBER as userStorySerialNumber, classes as classes, attributes as attributes, relationships as relationships in the json output. USE camelCase and a consistent output format in ALL THE subsequent responses. MAKE SURE NO FIELD IS EMPTY",
		},
	}
	ctx := context.Background()

	// Send the request
	res, err := client.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT35Turbo,
		Messages: messages,
	})
	check(err)
	fmt.Println(res.Choices[0].Message.Content)

	for _, dataItem := range dataList {
		messages = append(messages, chatgpt.ChatMessage{
			Role: chatgpt.ChatGPTModelRoleUser,
			Content: dataItem,
		})
		res, err := client.Send(ctx, &chatgpt.ChatCompletionRequest{
			Model: chatgpt.GPT35Turbo,
			Messages: messages,
		})
		check(err)
		response := res.Choices[0].Message.Content
		// Write the response to json file
		ioutil.WriteFile("exp1/output.json", []byte(response), 0644)

		// Convert the json file to excel file
		exec.Command("python3", "exp1/convert.py").Run()
	}
	
}