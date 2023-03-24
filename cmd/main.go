package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type Config struct {
	Data map[string]string `json:"data"`
}

func readConfig(filePath string) (*Config, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func readStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

func main() {
	openaiToken, ok := os.LookupEnv("OPENAI_TOKEN")
	if !ok {
		log.Fatal("Failed to read OPENAI_TOKEN environment variable.")
	}

	typeFlag := flag.String("type", "", "Type to be retrieved from conf.json")
	flag.Parse()

	if *typeFlag == "" {
		log.Fatal("Error: '--type' flag must be provided.")
	}

	config, err := readConfig("conf.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	prompt, ok := config.Data[*typeFlag]
	if !ok {
		log.Fatalf("Error: Type '%s' not found in conf.json", *typeFlag)
	}

	kubeOutput, err := readStdin()
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}

	client := openai.NewClient(openaiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt + "\n" + kubeOutput,
				},
			},
		},
)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
