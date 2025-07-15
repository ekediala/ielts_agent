package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

type Agent struct {
	client *anthropic.Client
	tools  []ToolDefinition
}

func (a *Agent) RunInference(ctx context.Context, convo []anthropic.MessageParam) (*anthropic.Message, error) {
	rules, err := os.ReadFile("rules.xml")
	if err != nil {
		return nil, fmt.Errorf("reading rules.xml: %w", err)
	}

	anthropicTools := []anthropic.ToolUnionParam{}
	for _, tool := range a.tools {
		anthropicTools = append(anthropicTools, anthropic.ToolUnionParam{
			OfTool: &anthropic.ToolParam{
				Name:        tool.Name,
				Description: anthropic.String(tool.Description),
				InputSchema: tool.InputSchema,
			},
		})
	}

	return a.client.Messages.New(ctx,
		anthropic.MessageNewParams{
			System: []anthropic.TextBlockParam{
				{
					Text: string(rules),
				},
			},
			Messages:  convo,
			Model:     anthropic.ModelClaude3_5Haiku20241022,
			MaxTokens: 1024,
			Tools:     anthropicTools,
		},
	)
}

func (a *Agent) Run(ctx context.Context) error {
	convo := []anthropic.MessageParam{}
	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	r := bufio.NewScanner(os.Stdin)
	readUserInput := true

	for {
		if readUserInput {
			fmt.Print("\u001b[94mYou\u001b[0m: ")
			s, ok := a.getUserMessage(r)
			if !ok {
				return nil
			}

			userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(s))
			convo = append(convo, userMessage)
		}

		assistantMessage, err := a.RunInference(ctx, convo)
		if err != nil {
			return fmt.Errorf("running inference: %w", err)
		}

		convo = append(convo, assistantMessage.ToParam())
		toolResults := []anthropic.ContentBlockParamUnion{}

		for _, content := range assistantMessage.Content {
			switch content.Type {
			case "text":
				fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", content.Text)
			case "tool_use":
				result := a.executeTool(content.ID, content.Name, content.Input)
				toolResults = append(toolResults, result)
			}
		}

		if len(toolResults) == 0 {
			readUserInput = true
			continue
		}

		readUserInput = false
		convo = append(convo, anthropic.NewUserMessage(toolResults...))
	}
}

func (a *Agent) executeTool(id, name string, input json.RawMessage) anthropic.ContentBlockParamUnion {
	var toolDef ToolDefinition
	found := false
	for _, tool := range a.tools {
		if tool.Name == name {
			toolDef = tool
			found = true
			break
		}
	}

	if !found {
		return anthropic.NewToolResultBlock(id, "tool not found", true)
	}

	fmt.Printf("\u001b[92mtool\u001b[0m: %s(%s)\n", name, input)
	response, err := toolDef.Function(input)
	if err != nil {
		return anthropic.NewToolResultBlock(id, err.Error(), true)
	}

	return anthropic.NewToolResultBlock(id, response, false)
}

func (a *Agent) getUserMessage(r *bufio.Scanner) (string, bool) {
	if r.Scan() {
		return r.Text(), true
	}

	return "", false
}

func NewAgent(client *anthropic.Client, tools []ToolDefinition) *Agent {
	return &Agent{client: client, tools: tools}
}
