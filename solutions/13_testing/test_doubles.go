// test_doubles.go - SOLUTION
// Learn about test doubles: mocks, stubs, fakes, and dependency injection

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Domain interfaces that we'll create test doubles for

// EmailService interface for sending emails
type EmailService interface {
	SendEmail(to, subject, body string) error
}

// Database interface for data operations  
type Database interface {
	SaveUser(id int, name string) error
	GetUser(id int) (string, error)
	DeleteUser(id int) error
}

// Logger interface for logging
type Logger interface {
	Log(level, message string)
	Error(message string)
}

// Production implementations

// Real email service implementation
type SMTPEmailService struct {
	host string
	port int
}

// Implement EmailService interface
func (s *SMTPEmailService) SendEmail(to, subject, body string) error {
	// This would normally connect to SMTP server
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	return nil
}

// Real database implementation
type PostgreSQLDatabase struct {
	connectionString string
}

// Implement Database interface methods
func (db *PostgreSQLDatabase) SaveUser(id int, name string) error {
	// This would normally execute SQL INSERT
	fmt.Printf("Saving user %d: %s to PostgreSQL\n", id, name)
	return nil
}

func (db *PostgreSQLDatabase) GetUser(id int) (string, error) {
	// This would normally execute SQL SELECT
	return fmt.Sprintf("User%d", id), nil
}

func (db *PostgreSQLDatabase) DeleteUser(id int) error {
	// This would normally execute SQL DELETE
	fmt.Printf("Deleting user %d from PostgreSQL\n", id)
	return nil
}

// Real logger implementation
type FileLogger struct {
	filename string
}

// Implement Logger interface
func (l *FileLogger) Log(level, message string) {
	// This would normally write to file
	fmt.Printf("[%s] %s\n", level, message)
}

func (l *FileLogger) Error(message string) {
	l.Log("ERROR", message)
}

// Business logic that depends on external services

// UserService that depends on Database, EmailService, and Logger
type UserService struct {
	db     Database
	email  EmailService
	logger Logger
}

// Constructor for UserService
func NewUserService(db Database, email EmailService, logger Logger) *UserService {
	return &UserService{
		db:     db,
		email:  email,
		logger: logger,
	}
}

// Business method that uses all dependencies
func (us *UserService) RegisterUser(id int, name, email string) error {
	// Validate input
	if name == "" {
		us.logger.Error("Name cannot be empty")
		return errors.New("name cannot be empty")
	}
	
	// Save user to database
	if err := us.db.SaveUser(id, name); err != nil {
		us.logger.Error("Failed to save user: " + err.Error())
		return err
	}
	
	// Send welcome email (don't fail registration if this fails)
	if err := us.email.SendEmail(email, "Welcome!", "Welcome to our service!"); err != nil {
		us.logger.Error("Failed to send welcome email: " + err.Error())
		// Continue - don't fail registration
	}
	
	// Log the operation
	us.logger.Log("INFO", "User registered successfully: "+name)
	return nil
}

func (us *UserService) GetUser(id int) (string, error) {
	us.logger.Log("INFO", fmt.Sprintf("Getting user %d", id))
	return us.db.GetUser(id)
}

// Test Doubles

// Mock EmailService - records method calls for verification
type MockEmailService struct {
	called    bool
	lastTo    string
	lastSubject string
	lastBody  string
	returnError error
}

// Implement EmailService interface for mock
func (m *MockEmailService) SendEmail(to, subject, body string) error {
	m.called = true
	m.lastTo = to
	m.lastSubject = subject
	m.lastBody = body
	return m.returnError
}

// Add verification methods
func (m *MockEmailService) WasCalled() bool {
	return m.called
}

func (m *MockEmailService) GetLastCall() (string, string, string) {
	return m.lastTo, m.lastSubject, m.lastBody
}

func (m *MockEmailService) SetReturnError(err error) {
	m.returnError = err
}

// Stub Database - returns predetermined responses
type StubDatabase struct {
	saveError error
	getResult string
	getError  error
	deleteError error
}

// Implement Database interface for stub
func (s *StubDatabase) SaveUser(id int, name string) error {
	return s.saveError
}

func (s *StubDatabase) GetUser(id int) (string, error) {
	return s.getResult, s.getError
}

func (s *StubDatabase) DeleteUser(id int) error {
	return s.deleteError
}

// Fake Logger - simple working implementation for testing
type FakeLogger struct {
	logs []string
}

// Implement Logger interface for fake
func (f *FakeLogger) Log(level, message string) {
	f.logs = append(f.logs, fmt.Sprintf("[%s] %s", level, message))
}

func (f *FakeLogger) Error(message string) {
	f.Log("ERROR", message)
}

// Add methods to inspect logged messages
func (f *FakeLogger) GetLogs() []string {
	return f.logs
}

func (f *FakeLogger) ContainsLog(message string) bool {
	for _, log := range f.logs {
		if strings.Contains(log, message) {
			return true
		}
	}
	return false
}

// Test helper functions

// Test successful user registration
func testSuccessfulRegistration() {
	fmt.Println("Testing successful user registration...")
	
	// Create test doubles
	mockEmail := &MockEmailService{}
	stubDB := &StubDatabase{} // No errors configured, so operations succeed
	fakeLogger := &FakeLogger{}
	
	// Create UserService with test doubles
	userService := NewUserService(stubDB, mockEmail, fakeLogger)
	
	// Call the method being tested
	err := userService.RegisterUser(123, "John Doe", "john@example.com")
	
	// Verify the results
	if err != nil {
		fmt.Printf("❌ Expected no error, got: %v\n", err)
		return
	}
	
	// Verify mock was called correctly
	if !mockEmail.WasCalled() {
		fmt.Println("❌ Expected email service to be called")
		return
	}
	
	// Verify specific method parameters
	to, subject, body := mockEmail.GetLastCall()
	expectedTo := "john@example.com"
	if to != expectedTo {
		fmt.Printf("❌ Expected email to %s, got %s\n", expectedTo, to)
		return
	}
	
	// Verify logging occurred
	if !fakeLogger.ContainsLog("User registered successfully") {
		fmt.Println("❌ Expected success log message")
		return
	}
	
	fmt.Println("✅ Successful registration test passed")
}

// Test registration with database error
func testRegistrationWithDatabaseError() {
	fmt.Println("Testing registration with database error...")
	
	// Create test doubles
	mockEmail := &MockEmailService{}
	stubDB := &StubDatabase{
		saveError: errors.New("database connection failed"),
	}
	fakeLogger := &FakeLogger{}
	
	// Create UserService with test doubles
	userService := NewUserService(stubDB, mockEmail, fakeLogger)
	
	// Call the method being tested
	err := userService.RegisterUser(123, "John Doe", "john@example.com")
	
	// Verify error handling
	if err == nil {
		fmt.Println("❌ Expected error, got nil")
		return
	}
	
	// Verify email was NOT sent (since database failed)
	if mockEmail.WasCalled() {
		fmt.Println("❌ Expected email service NOT to be called when database fails")
		return
	}
	
	// Verify error was logged
	if !fakeLogger.ContainsLog("Failed to save user") {
		fmt.Println("❌ Expected error log message")
		return
	}
	
	fmt.Println("✅ Database error test passed")
}

// Test registration with email service error
func testRegistrationWithEmailError() {
	fmt.Println("Testing registration with email service error...")
	
	// Create test doubles - email service that returns error
	mockEmail := &MockEmailService{}
	mockEmail.SetReturnError(errors.New("SMTP server unavailable"))
	stubDB := &StubDatabase{} // No error, so save succeeds
	fakeLogger := &FakeLogger{}
	
	// Create UserService and test
	userService := NewUserService(stubDB, mockEmail, fakeLogger)
	
	// Call the method and verify behavior
	err := userService.RegisterUser(123, "John Doe", "john@example.com")
	
	// Verify the registration still succeeds even if email fails
	if err != nil {
		fmt.Printf("Registration should succeed even if email fails, got: %v\n", err)
		return
	}
	
	// Verify warning was logged about email failure
	if !fakeLogger.ContainsLog("Failed to send welcome email") {
		fmt.Println("❌ Expected email failure log message")
		return
	}
	
	fmt.Println("✅ Email error test passed")
}

func main() {
	fmt.Println("=== Test Doubles: Mocks, Stubs, and Fakes ===")
	
	fmt.Println("\nDemonstrating dependency injection and test doubles...")
	
	// Run all test scenarios
	testSuccessfulRegistration()
	testRegistrationWithDatabaseError()
	testRegistrationWithEmailError()
	
	fmt.Println("\n=== Test Doubles Summary ===")
	fmt.Println("✅ Mock: Verifies behavior (method calls, parameters)")
	fmt.Println("✅ Stub: Provides predetermined responses")
	fmt.Println("✅ Fake: Simple working implementation for testing")
	fmt.Println("✅ Dependency Injection: Enables easy substitution of implementations")
}