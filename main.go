package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"google.golang.org/genai"

	"github.com/truongd3/golang-gpt-analyzer/helpers"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <input_file_path> <output_file_path>")
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing API Key")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println("Parsing file:", inputFile)

	promptPrefix := "Given a Python code snippet. Extract a list of libraries that are used in the code:\n```python\n"
	promptSuffix := "\n```\nList the libraries in a comma-separated format. If parent and child libraries are used, only list the parent library. Only list the libraries, do not provide any additional information."
	prompt := promptPrefix + string(fileBytes) + promptSuffix

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())

	helpers.ConvertLibrariesToOutputFile(result.Text(), outputFile)
	fmt.Println("DONE ✅")
}
