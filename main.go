package main

import (
	"fmt"
	"time"
)

// Agent represents a simple AI agent.
type Agent struct {
	Name string
}

// Perceive simulates the agent perceiving its environment.
func (a *Agent) Perceive(input string) {
	fmt.Printf("Agent %s perceiving: %s\n", a.Name, input)
}

// Decide simulates the agent making a decision.
func (a *Agent) Decide() string {
	// In a real scenario, this would involve complex AI logic.
	// For this example, it's a simple decision.
	return "Perform action"
}

// Act simulates the agent performing an action.
func (a *Agent) Act(action string) {
	fmt.Printf("Agent %s acting: %s\n", a.Name, action)
}

func main() {
	fmt.Println("Starting AI Agent Framework example...")

	// Create a new agent
	agent := Agent{Name: "Alpha"}

	// Simulate agent's cycle
	agent.Perceive("New data stream detected")
	time.Sleep(500 * time.Millisecond) // Simulate processing time

	action := agent.Decide()
	fmt.Printf("Agent %s decided to: %s\n", agent.Name, action)
	time.Sleep(500 * time.Millisecond) // Simulate decision time

	agent.Act(action)
	time.Sleep(500 * time.Millisecond) // Simulate action time

	fmt.Println("Agent Alpha completed its cycle.")

	// Another agent example
	agent2 := Agent{Name: "Beta"}
	agent2.Perceive("User input received")
	fmt.Println("Agent Beta is processing...")

	// Simulate a more complex decision process
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("Agent %s acting: Responding to user input\n", agent2.Name)
	}()

	fmt.Println("Main function continuing...")
	time.Sleep(3 * time.Second) // Wait for agent2's action to complete
	fmt.Println("AI Agent Framework example finished.")
}
