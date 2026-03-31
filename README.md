# Go AI Agent Framework

A lightweight and extensible framework for building AI agents in Go, focusing on modularity and performance.

## Features
- **Modular Agent Design:** Easily define and combine different agent components (e.g., perception, decision-making, action).
- **Concurrency Support:** Leverages Go's goroutines and channels for efficient concurrent execution of agent tasks.
- **Event-Driven Architecture:** Agents can react to events and communicate asynchronously.
- **Pluggable Modules:** Supports easy integration of various AI models and algorithms.

## Getting Started

### Prerequisites
- Go 1.18+

### Installation

```bash
git clone https://github.com/Dras1950/go-ai-agent-framework.git
cd go-ai-agent-framework
go mod tidy
```

### Usage

```go
package main

import (
	"fmt"
	"time"
)

type Agent struct {
	Name string
}

func (a *Agent) Perceive(input string) {
	fmt.Printf("Agent %s perceiving: %s\n", a.Name, input)
}

func (a *Agent) Decide() string {
	return "Perform action"
}

func (a *Agent) Act(action string) {
	fmt.Printf("Agent %s acting: %s\n", a.Name, action)
}

func main() {
	agent := Agent{Name: "Alpha"}

	agent.Perceive("New data arrived")
	action := agent.Decide()
	agent.Act(action)

	fmt.Println("Agent Alpha completed its cycle.")
}
```

## Contributing

We welcome contributions! Please see `CONTRIBUTING.md` for details.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
