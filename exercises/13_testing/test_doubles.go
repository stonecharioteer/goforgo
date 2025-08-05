// test_doubles.go
// Learn about test doubles: mocks, stubs, fakes, and dependency injection

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Domain interfaces that we'll create test doubles for

// TODO: EmailService interface for sending emails
type EmailService interface {
	// TODO: Define method to send email
}

// TODO: Database interface for data operations  
type Database interface {
	// TODO: Define methods for database operations
}

// TODO: Logger interface for logging
type Logger interface {
	// TODO: Define logging methods
}

// Production implementations

// TODO: Real email service implementation
type SMTPEmailService struct {
	// TODO: Define fields for SMTP configuration
}

// TODO: Implement EmailService interface
func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	// TODO: Implement real email sending logic
	// This would normally connect to SMTP server
}

// TODO: Real database implementation
type PostgreSQLDatabase struct {
	// TODO: Define fields for database connection
}

// TODO: Implement Database interface methods
func (db *PostgreSQLDatabase) SaveUser(id int, name string) error {
	// TODO: Implement real database save logic
}

func (db *PostgreSQLDatabase) GetUser(id int) (string, error) {
	// TODO: Implement real database get logic
}

func (db *PostgreSQLDatabase) DeleteUser(id int) error {
	// TODO: Implement real database delete logic
}

// TODO: Real logger implementation
type FileLogger struct {
	// TODO: Define fields for file logging
}

// TODO: Implement Logger interface
func (l *FileLogger) Log(level, message string) {
	// TODO: Implement real file logging
}

func (l *FileLogger) Error(message string) {
	// TODO: Implement error logging
}

// Business logic that depends on external services

// TODO: UserService that depends on Database, EmailService, and Logger
type UserService struct {
	// TODO: Define dependencies as interfaces
}

// TODO: Constructor for UserService
func NewUserService(db Database, email EmailService, logger Logger) *UserService {
	// TODO: Create UserService with dependencies
}

// TODO: Business method that uses all dependencies
func (us *UserService) RegisterUser(id int, name, email string) error {
	// TODO: Implement user registration logic:
	// 1. Validate input
	// 2. Save user to database
	// 3. Send welcome email
	// 4. Log the operation
	// Handle errors appropriately
}

func (us *UserService) GetUser(id int) (string, error) {
	// TODO: Implement get user logic with logging
}

// Test Doubles

// TODO: Mock EmailService - records method calls for verification
type MockEmailService struct {
	// TODO: Define fields to track method calls
}

// TODO: Implement EmailService interface for mock
func (m *MockEmailService) SendEmail(to, subject, body string) error {
	// TODO: Record the method call and return configured response
}

// TODO: Add verification methods
func (m *MockEmailService) WasCalled() bool {
	// TODO: Return whether SendEmail was called
}

func (m *MockEmailService) GetLastCall() (string, string, string) {
	// TODO: Return parameters of last call
}

// TODO: Stub Database - returns predetermined responses
type StubDatabase struct {
	// TODO: Define predetermined responses
}

// TODO: Implement Database interface for stub
func (s *StubDatabase) SaveUser(id int, name string) error {
	// TODO: Return predetermined response
}

func (s *StubDatabase) GetUser(id int) (string, error) {
	// TODO: Return predetermined response
}

func (s *StubDatabase) DeleteUser(id int) error {
	// TODO: Return predetermined response
}

// TODO: Fake Logger - simple working implementation for testing
type FakeLogger struct {
	// TODO: Define simple in-memory storage for logs
}

// TODO: Implement Logger interface for fake
func (f *FakeLogger) Log(level, message string) {
	// TODO: Store log in memory
}

func (f *FakeLogger) Error(message string) {
	// TODO: Store error log in memory
}

// TODO: Add methods to inspect logged messages
func (f *FakeLogger) GetLogs() []string {
	// TODO: Return all logged messages
}

func (f *FakeLogger) ContainsLog(message string) bool {
	// TODO: Check if a specific message was logged
}

// Test helper functions

// TODO: Test successful user registration
func testSuccessfulRegistration() {
	fmt.Println("Testing successful user registration...")
	
	// TODO: Create test doubles
	mockEmail := /* create mock email service */
	stubDB := /* create stub database that succeeds */
	fakeLogger := /* create fake logger */
	
	// TODO: Create UserService with test doubles
	userService := /* create user service with test doubles */
	
	// TODO: Call the method being tested
	err := /* register user with test data */
	
	// TODO: Verify the results
	if err != nil {
		fmt.Printf("❌ Expected no error, got: %v\n", err)
		return
	}
	
	// TODO: Verify mock was called correctly
	if !mockEmail.WasCalled() {
		fmt.Println("❌ Expected email service to be called")
		return
	}
	
	// TODO: Verify specific method parameters
	to, subject, body := mockEmail.GetLastCall()
	expectedTo := "john@example.com"
	if to != expectedTo {
		fmt.Printf("❌ Expected email to %s, got %s\n", expectedTo, to)
		return
	}
	
	// TODO: Verify logging occurred
	if !fakeLogger.ContainsLog("User registered successfully") {
		fmt.Println("❌ Expected success log message")
		return
	}
	
	fmt.Println("✅ Successful registration test passed")
}

// TODO: Test registration with database error
func testRegistrationWithDatabaseError() {
	fmt.Println("Testing registration with database error...")
	
	// TODO: Create test doubles
	mockEmail := /* create mock email service */
	stubDB := /* create stub database that returns error */
	fakeLogger := /* create fake logger */
	
	// TODO: Create UserService with test doubles
	userService := /* create user service with test doubles */
	
	// TODO: Call the method being tested
	err := /* register user with test data */
	
	// TODO: Verify error handling
	if err == nil {
		fmt.Println("❌ Expected error, got nil")
		return
	}
	
	// TODO: Verify email was NOT sent (since database failed)
	if mockEmail.WasCalled() {
		fmt.Println("❌ Expected email service NOT to be called when database fails")
		return
	}
	
	// TODO: Verify error was logged
	if !fakeLogger.ContainsLog("Failed to save user") {
		fmt.Println("❌ Expected error log message")
		return
	}
	
	fmt.Println("✅ Database error test passed")
}

// TODO: Test registration with email service error
func testRegistrationWithEmailError() {
	fmt.Println("Testing registration with email service error...")
	
	// TODO: Create test doubles - email service that returns error
	mockEmail := /* create mock email service that returns error */
	stubDB := /* create stub database that succeeds */
	fakeLogger := /* create fake logger */
	
	// TODO: Create UserService and test
	userService := /* create user service with test doubles */
	
	// TODO: Call the method and verify behavior
	err := /* register user with test data */
	
	// TODO: Verify the registration still succeeds even if email fails
	if err != nil {
		fmt.Printf("Registration should succeed even if email fails, got: %v\n", err)
		return
	}
	
	// TODO: Verify warning was logged about email failure
	if !fakeLogger.ContainsLog("Failed to send welcome email") {
		fmt.Println("❌ Expected email failure log message")
		return
	}
	
	fmt.Println("✅ Email error test passed")
}

func main() {
	fmt.Println("=== Test Doubles: Mocks, Stubs, and Fakes ===")
	
	fmt.Println("\nDemonstrating dependency injection and test doubles...")
	
	// TODO: Run all test scenarios
	testSuccessfulRegistration()
	testRegistrationWithDatabaseError()
	testRegistrationWithEmailError()
	
	fmt.Println("\n=== Test Doubles Summary ===")
	fmt.Println("✅ Mock: Verifies behavior (method calls, parameters)")
	fmt.Println("✅ Stub: Provides predetermined responses")
	fmt.Println("✅ Fake: Simple working implementation for testing")
	fmt.Println("✅ Dependency Injection: Enables easy substitution of implementations")
}