// design_patterns.go
// Learn common design patterns implementation in Go

package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Singleton Pattern
type Logger struct {
	// Define singleton logger
}

var (
	loggerInstance *Logger
	loggerOnce     sync.Once
)

// TODO: Implement singleton getInstance
func GetLogger() *Logger {
	// TODO: Implement thread-safe singleton
}

func (l *Logger) Log(message string) {
	fmt.Printf("[LOG] %s: %s\n", time.Now().Format("15:04:05"), message)
}

// TODO: Factory Pattern
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

// TODO: Implement animal factory
func CreateAnimal(animalType, name string) Animal {
	// TODO: Create animals based on type
}

// TODO: Observer Pattern
type Observer interface {
	Update(message string)
}

type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Notify(message string)
}

type NewsAgency struct {
	// TODO: Define observers list
}

// TODO: Implement Subject interface for NewsAgency
func (na *NewsAgency) Subscribe(observer Observer) {
	// TODO: Add observer to list
}

func (na *NewsAgency) Unsubscribe(observer Observer) {
	// TODO: Remove observer from list
}

func (na *NewsAgency) Notify(message string) {
	// TODO: Notify all observers
}

func (na *NewsAgency) PublishNews(news string) {
	fmt.Printf("NewsAgency: Publishing news - %s\n", news)
	/* notify observers */
}

type NewsChannel struct {
	Name string
}

// TODO: Implement Observer interface
func (nc *NewsChannel) Update(message string) {
	fmt.Printf("[%s] Received news: %s\n", nc.Name, message)
}

// TODO: Strategy Pattern
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
	// TODO: Add payment strategy field
}

// TODO: Implement strategy pattern methods
func (sc *ShoppingCart) SetPaymentStrategy(strategy PaymentStrategy) {
	// TODO: Set payment strategy
}

func (sc *ShoppingCart) Checkout(amount float64) {
	// TODO: Use current payment strategy
}

// TODO: Decorator Pattern
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
	// TODO: Embed Coffee interface
}

type MilkDecorator struct {
	// TODO: Embed CoffeeDecorator
}

// TODO: Implement decorator methods
func (md MilkDecorator) Cost() float64 {
	// TODO: Add milk cost to base coffee
}

func (md MilkDecorator) Description() string {
	// TODO: Add milk to description
}

type SugarDecorator struct {
	// TODO: Embed CoffeeDecorator
}

func (sd SugarDecorator) Cost() float64 {
	// TODO: Add sugar cost
}

func (sd SugarDecorator) Description() string {
	// TODO: Add sugar to description
}

// TODO: Command Pattern
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
	// TODO: Add light reference
}

// TODO: Implement Command interface
func (loc *LightOnCommand) Execute() {
	// TODO: Turn light on
}

func (loc *LightOnCommand) Undo() {
	// TODO: Turn light off
}

type LightOffCommand struct {
	// TODO: Add light reference
}

func (lfc *LightOffCommand) Execute() {
	// TODO: Turn light off
}

func (lfc *LightOffCommand) Undo() {
	// TODO: Turn light on
}

type RemoteControl struct {
	// TODO: Add command slots and history
}

// TODO: Implement remote control methods
func (rc *RemoteControl) SetCommand(slot int, command Command) {
	// TODO: Set command for slot
}

func (rc *RemoteControl) PressButton(slot int) {
	// TODO: Execute command and add to history
}

func (rc *RemoteControl) PressUndo() {
	// TODO: Undo last command
}

func main() {
	fmt.Println("=== Design Patterns in Go ===")
	
	fmt.Println("\n=== Singleton Pattern ===")
	
	// TODO: Test singleton pattern
	logger1 := /* get logger instance */
	logger2 := /* get logger instance */
	
	logger1.Log("First message")
	logger2.Log("Second message")
	
	fmt.Printf("Same instance: %t\n", /* compare logger instances */)
	
	fmt.Println("\n=== Factory Pattern ===")
	
	// TODO: Test factory pattern
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
		animal := /* create animal using factory */
		if animal != nil {
			fmt.Printf("Created %s named %s: %s\n", 
				animal.GetType(), spec.name, animal.MakeSound())
		} else {
			fmt.Printf("Unknown animal type: %s\n", spec.animalType)
		}
	}
	
	fmt.Println("\n=== Observer Pattern ===")
	
	// TODO: Test observer pattern
	newsAgency := /* create news agency */
	
	// TODO: Create news channels
	cnn := /* create CNN channel */
	bbc := /* create BBC channel */
	fox := /* create Fox channel */
	
	// TODO: Subscribe channels
	/* subscribe CNN */
	/* subscribe BBC */
	/* subscribe Fox */
	
	// TODO: Publish news
	newsAgency.PublishNews("Breaking: Go 1.22 Released!")
	
	// TODO: Unsubscribe one channel
	fmt.Println("\nUnsubscribing BBC...")
	/* unsubscribe BBC */
	
	newsAgency.PublishNews("Update: New Go features announced")
	
	fmt.Println("\n=== Strategy Pattern ===")
	
	// TODO: Test strategy pattern
	cart := /* create shopping cart */
	
	// TODO: Test different payment methods
	paymentMethods := []PaymentStrategy{
		/* create credit card payment */,
		/* create PayPal payment */,
		/* create Bitcoin payment */,
	}
	
	amount := 99.99
	for i, method := range paymentMethods {
		fmt.Printf("\nPayment method %d:\n", i+1)
		/* set payment strategy */
		/* checkout with amount */
	}
	
	fmt.Println("\n=== Decorator Pattern ===")
	
	// TODO: Test decorator pattern
	coffee := /* create simple coffee */
	fmt.Printf("Base: %s - $%.2f\n", coffee.Description(), coffee.Cost())
	
	// TODO: Add decorators
	coffeeWithMilk := /* add milk decorator */
	fmt.Printf("With milk: %s - $%.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())
	
	coffeeWithMilkAndSugar := /* add sugar decorator */
	fmt.Printf("With milk and sugar: %s - $%.2f\n", 
		coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())
	
	fmt.Println("\n=== Command Pattern ===")
	
	// TODO: Test command pattern
	livingRoomLight := /* create living room light */
	kitchenLight := /* create kitchen light */
	
	// TODO: Create commands
	livingRoomLightOn := /* create light on command */
	livingRoomLightOff := /* create light off command */
	kitchenLightOn := /* create kitchen light on command */
	
	// TODO: Create remote control
	remote := /* create remote control */
	
	// TODO: Set commands
	/* set command for slot 0 */
	/* set command for slot 1 */
	/* set command for slot 2 */
	
	// TODO: Test remote control
	fmt.Println("Testing remote control:")
	/* press button 0 */
	/* press button 2 */
	/* press button 1 */
	
	fmt.Println("\nTesting undo:")
	/* press undo */
	/* press undo */
}