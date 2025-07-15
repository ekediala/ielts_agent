package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/ledongthuc/pdf"
)

var ReadFileDefinition = ToolDefinition{
	Name:        "read_file",
	Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
	InputSchema: ReadFileInputSchema,
	Function:    ReadFile,
}

type ReadFileInput struct {
	Path string `json:"path" jsonschema_description:"The relative path of a file in the working directory."`
}

var ReadFileInputSchema = GenerateSchema[ReadFileInput]()

func ReadFile(input json.RawMessage) (string, error) {
	readFileInput := ReadFileInput{}
	err := json.Unmarshal(input, &readFileInput)
	if err != nil {
		return "", fmt.Errorf("unmarshalling: %w", err)
	}

	if path.Ext(readFileInput.Path) == ".pdf" {
		f, r, err := pdf.Open(readFileInput.Path)
		if err != nil {
			return "", fmt.Errorf("reading file: %w", err)
		}
		defer f.Close()

		var buf bytes.Buffer
		b, err := r.GetPlainText()
		if err != nil {
			return "", fmt.Errorf("getting pdf text content: %w", err)
		}

		_, err = buf.ReadFrom(b)
		if err != nil {
			return "", fmt.Errorf("reading pdf text content: %w", err)
		}

		fmt.Println(buf.String())

		return buf.String(), nil
	}

	content, err := os.ReadFile(readFileInput.Path)
	if err != nil {
		return "", fmt.Errorf("reading file: %w", err)
	}

	return string(content), nil
}
