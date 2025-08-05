// design_patterns.go - SOLUTION
// Learn common design patterns implementation in Go

package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton Pattern
type Logger struct{}

var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

// Implement singleton getInstance
func GetLogger() *Logger {
	loggerOnce.Do(func() {
		loggerInstance = &Logger{}
	})
	return loggerInstance
}

func (l *Logger) Log(message string) {
	fmt.Printf("[LOG] %s: %s\n", time.Now().Format("15:04:05"), message)
}

// Factory Pattern
type Animal interface {
	MakeSound() string
	GetType() string
}

type Dog struct {
	Name string
}

func (d Dog) MakeSound() string {
	return "Woof!"
}

func (d Dog) GetType() string {
	return "Dog"
}

type Cat struct {
	Name string
}

func (c Cat) MakeSound() string {
	return "Meow!"
}

func (c Cat) GetType() string {
	return "Cat"
}

// Implement animal factory
func CreateAnimal(animalType, name string) Animal {
	switch animalType {
	case "dog":
		return Dog{Name: name}
	case "cat":
		return Cat{Name: name}
	default:
		return nil
	}
}

// Observer Pattern
type Observer interface {
	Update(message string)
}

type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Notify(message string)
}

type NewsAgency struct {
	observers []Observer
}

// Implement Subject interface for NewsAgency
func (na *NewsAgency) Subscribe(observer Observer) {
	na.observers = append(na.observers, observer)
}

func (na *NewsAgency) Unsubscribe(observer Observer) {
	for i, obs := range na.observers {
		if obs == observer {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			break
		}
	}
}

func (na *NewsAgency) Notify(message string) {
	for _, observer := range na.observers {
		observer.Update(message)
	}
}

func (na *NewsAgency) PublishNews(news string) {
	fmt.Printf("NewsAgency: Publishing news - %s\n", news)
	na.Notify(news)
}

type NewsChannel struct {
	Name string
}

// Implement Observer interface
func (nc *NewsChannel) Update(message string) {
	fmt.Printf("[%s] Received news: %s\n", nc.Name, message)
}

// Strategy Pattern
type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCard struct {
	Number string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card ending in %s", 
		amount, cc.Number[len(cc.Number)-4:])
}

type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, pp.Email)
}

type BitCoin struct {
	Address string
}

func (bc BitCoin) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Bitcoin to %s", amount, bc.Address[:8]+"...")
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
}

// Implement strategy pattern methods
func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	sc.paymentStrategy = strategy
}

func (sc *ShoppingCart) Checkout(amount float64) {
	if sc.paymentStrategy != nil {
		result := sc.paymentStrategy.Pay(amount)
		fmt.Println(result)
	} else {
		fmt.Println("No payment method selected")
	}
}

// Decorator Pattern
type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (sc SimpleCoffee) Cost() float64 {
	return 2.0
}

func (sc SimpleCoffee) Description() string {
	return "Simple Coffee"
}

type CoffeeDecorator struct {
	coffee Coffee
}

type MilkDecorator struct {
	CoffeeDecorator
}

// Implement decorator methods
func (md MilkDecorator) Cost() float64 {
	return md.coffee.Cost() + 0.5
}

func (md MilkDecorator) Description() string {
	return md.coffee.Description() + ", Milk"
}

type SugarDecorator struct {
	CoffeeDecorator
}

func (sd SugarDecorator) Cost() float64 {
	return sd.coffee.Cost() + 0.2
}

func (sd SugarDecorator) Description() string {
	return sd.coffee.Description() + ", Sugar"
}

// Command Pattern
type Command interface {
	Execute()
	Undo()
}

type Light struct {
	Location string
	IsOn     bool
}

func (l *Light) TurnOn() {
	l.IsOn = true
	fmt.Printf("%s light is ON\n", l.Location)
}

func (l *Light) TurnOff() {
	l.IsOn = false
	fmt.Printf("%s light is OFF\n", l.Location)
}

type LightOnCommand struct {
	light *Light
}

// Implement Command interface
func (loc *LightOnCommand) Execute() {
	loc.light.TurnOn()
}

func (loc *LightOnCommand) Undo() {
	loc.light.TurnOff()
}

type LightOffCommand struct {
	light *Light
}

func (lfc *LightOffCommand) Execute() {
	lfc.light.TurnOff()
}

func (lfc *LightOffCommand) Undo() {
	lfc.light.TurnOn()
}

type RemoteControl struct {
	commands []Command
	lastCommand Command
}

// Implement remote control methods
func (rc *RemoteControl) SetCommand(slot int, command Command) {
	if slot >= len(rc.commands) {
		// Extend slice if needed
		for len(rc.commands) <= slot {
			rc.commands = append(rc.commands, nil)
		}
	}
	rc.commands[slot] = command
}

func (rc *RemoteControl) PressButton(slot int) {
	if slot < len(rc.commands) && rc.commands[slot] != nil {
		rc.commands[slot].Execute()
		rc.lastCommand = rc.commands[slot]
	}
}

func (rc *RemoteControl) PressUndo() {
	if rc.lastCommand != nil {
		rc.lastCommand.Undo()
		rc.lastCommand = nil
	}
}

func main() {
	fmt.Println("=== Design Patterns in Go ===")
	
	fmt.Println("\n=== Singleton Pattern ===")
	
	// Test singleton pattern
	logger1 := GetLogger()
	logger2 := GetLogger()
	
	logger1.Log("First message")
	logger2.Log("Second message")
	
	fmt.Printf("Same instance: %t\n", logger1 == logger2)
	
	fmt.Println("\n=== Factory Pattern ===")
	
	// Test factory pattern
	animals := []struct {
		animalType string
		name       string
	}{
		{"dog", "Buddy"},
		{"cat", "Whiskers"},
		{"dog", "Rex"},
		{"bird", "Tweety"}, // This should handle unknown type
	}
	
	fmt.Println("Creating animals:")
	for _, spec := range animals {
		animal := CreateAnimal(spec.animalType, spec.name)
		if animal != nil {
			fmt.Printf("Created %s named %s: %s\n", 
				animal.GetType(), spec.name, animal.MakeSound())
		} else {
			fmt.Printf("Unknown animal type: %s\n", spec.animalType)
		}
	}
	
	fmt.Println("\n=== Observer Pattern ===")
	
	// Test observer pattern
	newsAgency := &NewsAgency{}
	
	// Create news channels
	cnn := &NewsChannel{Name: "CNN"}
	bbc := &NewsChannel{Name: "BBC"}
	fox := &NewsChannel{Name: "Fox"}
	
	// Subscribe channels
	newsAgency.Subscribe(cnn)
	newsAgency.Subscribe(bbc)
	newsAgency.Subscribe(fox)
	
	// Publish news
	newsAgency.PublishNews("Breaking: Go 1.22 Released!")
	
	// Unsubscribe one channel
	fmt.Println("\nUnsubscribing BBC...")
	newsAgency.Unsubscribe(bbc)
	
	newsAgency.PublishNews("Update: New Go features announced")
	
	fmt.Println("\n=== Strategy Pattern ===")
	
	// Test strategy pattern
	cart := &ShoppingCart{}
	
	// Test different payment methods
	paymentMethods := []PaymentStrategy{
		CreditCard{Number: "1234-5678-9012-3456"},
		PayPal{Email: "user@example.com"},
		BitCoin{Address: "1A2B3C4D5E6F7G8H9I0J"},
	}
	
	amount := 99.99
	for i, method := range paymentMethods {
		fmt.Printf("\nPayment method %d:\n", i+1)
		cart.SetPaymentStrategy(method)
		cart.Checkout(amount)
	}
	
	fmt.Println("\n=== Decorator Pattern ===")
	
	// Test decorator pattern
	coffee := SimpleCoffee{}
	fmt.Printf("Base: %s - $%.2f\n", coffee.Description(), coffee.Cost())
	
	// Add decorators
	coffeeWithMilk := MilkDecorator{CoffeeDecorator{coffee: coffee}}
	fmt.Printf("With milk: %s - $%.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())
	
	coffeeWithMilkAndSugar := SugarDecorator{CoffeeDecorator{coffee: coffeeWithMilk}}
	fmt.Printf("With milk and sugar: %s - $%.2f\n", 
		coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())
	
	fmt.Println("\n=== Command Pattern ===")
	
	// Test command pattern
	livingRoomLight := &Light{Location: "Living Room", IsOn: false}
	kitchenLight := &Light{Location: "Kitchen", IsOn: false}
	
	// Create commands
	livingRoomLightOn := &LightOnCommand{light: livingRoomLight}
	livingRoomLightOff := &LightOffCommand{light: livingRoomLight}
	kitchenLightOn := &LightOnCommand{light: kitchenLight}
	
	// Create remote control
	remote := &RemoteControl{}
	
	// Set commands
	remote.SetCommand(0, livingRoomLightOn)
	remote.SetCommand(1, livingRoomLightOff)
	remote.SetCommand(2, kitchenLightOn)
	
	// Test remote control
	fmt.Println("Testing remote control:")
	remote.PressButton(0) // Turn on living room light
	remote.PressButton(2) // Turn on kitchen light
	remote.PressButton(1) // Turn off living room light
	
	fmt.Println("\nTesting undo:")
	remote.PressUndo() // Undo last command (turn living room light back on)
	remote.PressUndo() // No effect (no previous command)
}