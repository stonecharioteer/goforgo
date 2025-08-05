// json_advanced.go
// Learn advanced JSON techniques: custom marshaling, streaming, validation

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

// Custom JSON marshaling for time
type CustomTime struct {
	time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	// Format time as "YYYY-MM-DD HH:MM:SS"
	formatted := ct.Time.Format("2006-01-02 15:04:05")
	return json.Marshal(formatted)
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	// Parse JSON string
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}
	
	// Parse time using custom format
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return err
	}
	
	// Set ct.Time
	ct.Time = parsedTime
	return nil
}

// Struct with custom JSON behavior
type Product struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	InStock     bool        `json:"in_stock"`
	CreatedAt   CustomTime  `json:"created_at"`
	Tags        []string    `json:"tags,omitempty"`
	private     string      // Not exported - won't be marshaled
}

// Custom marshaling for the entire struct
func (p Product) MarshalJSON() ([]byte, error) {
	// Create a map to control JSON output
	data := map[string]interface{}{
		"id":         p.ID,
		"name":       p.Name,
		"price":      p.Price,
		"in_stock":   p.InStock,
		"created_at": p.CreatedAt,
	}
	
	// Only include tags if not empty
	if len(p.Tags) > 0 {
		data["tags"] = p.Tags
	}
	
	// Add computed field
	data["display_name"] = fmt.Sprintf("%s (#%d)", p.Name, p.ID)
	
	return json.Marshal(data)
}

// Dynamic JSON structure for flexible data
type DynamicData struct {
	Type   string                 `json:"type"`
	Data   map[string]interface{} `json:"data"`
	Meta   json.RawMessage        `json:"meta,omitempty"`
}

func main() {
	fmt.Println("=== Advanced JSON Processing ===")
	
	// Demonstrate custom marshaling
	demonstrateCustomMarshaling()
	
	// Demonstrate streaming JSON
	demonstrateStreamingJSON()
	
	// Demonstrate dynamic JSON handling
	demonstrateDynamicJSON()
	
	// Demonstrate JSON validation
	demonstrateJSONValidation()
}

func demonstrateCustomMarshaling() {
	fmt.Println("\n=== Custom JSON Marshaling ===")
	
	// Create product with custom time
	product := Product{
		ID:        123,
		Name:      "Gaming Laptop",
		Price:     1299.99,
		InStock:   true,
		CreatedAt: CustomTime{time.Now()},
		Tags:      []string{"electronics", "computers", "gaming"},
		private:   "This won't appear in JSON",
	}
	
	// Marshal to JSON
	jsonData, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Marshaled JSON:\n%s\n", jsonData)
	
	// Unmarshal back
	var unmarshaled Product
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled product: %+v\n", unmarshaled)
	fmt.Printf("Created at: %s\n", unmarshaled.CreatedAt.Format(time.RFC3339))
}

func demonstrateStreamingJSON() {
	fmt.Println("\n=== Streaming JSON ===")
	
	// JSON stream with multiple objects
	jsonStream := `
	{"id": 1, "name": "Product 1", "price": 10.50}
	{"id": 2, "name": "Product 2", "price": 25.99}
	{"id": 3, "name": "Product 3", "price": 5.75}
	`
	
	reader := strings.NewReader(jsonStream)
	decoder := json.NewDecoder(reader)
	
	fmt.Println("Processing JSON stream:")
	
	var products []map[string]interface{}
	
	for {
		var product map[string]interface{}
		err := decoder.Decode(&product)
		
		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
			continue
		}
		
		products = append(products, product)
		fmt.Printf("Processed product: %v\n", product)
	}
	
	fmt.Printf("Total products processed: %d\n", len(products))
}

func demonstrateDynamicJSON() {
	fmt.Println("\n=== Dynamic JSON Handling ===")
	
	// Different types of dynamic data
	userData := DynamicData{
		Type: "user",
		Data: map[string]interface{}{
			"name":  "Alice Johnson",
			"email": "alice@example.com",
			"age":   30,
			"active": true,
		},
		Meta: json.RawMessage(`{"source": "api", "version": "1.0"}`),
	}
	
	orderData := DynamicData{
		Type: "order",
		Data: map[string]interface{}{
			"order_id": "ORD-12345",
			"items":    []string{"laptop", "mouse", "keyboard"},
			"total":    299.99,
		},
	}
	
	// Process dynamic data
	processData := []DynamicData{userData, orderData}
	
	for _, data := range processData {
		fmt.Printf("\nProcessing %s data:\n", data.Type)
		
		// Marshal to see the full structure
		jsonBytes, _ := json.MarshalIndent(data, "", "  ")
		fmt.Printf("JSON: %s\n", jsonBytes)
		
		// Process based on type
		switch data.Type {
		case "user":
			if name, ok := data.Data["name"].(string); ok {
				fmt.Printf("User name: %s\n", name)
			}
			if email, ok := data.Data["email"].(string); ok {
				fmt.Printf("User email: %s\n", email)
			}
			
		case "order":
			if orderID, ok := data.Data["order_id"].(string); ok {
				fmt.Printf("Order ID: %s\n", orderID)
			}
			if total, ok := data.Data["total"].(float64); ok {
				fmt.Printf("Order total: $%.2f\n", total)
			}
		}
		
		// Process meta if present
		if len(data.Meta) > 0 {
			var meta map[string]interface{}
			if err := json.Unmarshal(data.Meta, &meta); err == nil {
				fmt.Printf("Meta: %v\n", meta)
			}
		}
	}
}

func demonstrateJSONValidation() {
	fmt.Println("\n=== JSON Validation ===")
	
	// Valid JSON samples
	validJSONSamples := []string{
		`{"name": "John", "age": 25}`,
		`{"products": [{"id": 1, "name": "Item"}]}`,
		`{"nested": {"data": {"value": 42}}}`,
	}
	
	// Invalid JSON samples
	invalidJSONSamples := []string{
		`{"name": "John", "age": }`,           // Missing value
		`{"name": "John" "age": 25}`,          // Missing comma
		`{"name": "John", "age": 25`,          // Missing closing brace
	}
	
	fmt.Println("Validating JSON samples:")
	
	fmt.Println("\nValid JSON:")
	for i, sample := range validJSONSamples {
		if isValidJSON(sample) {
			fmt.Printf("  %d. ✓ Valid: %s\n", i+1, sample)
		} else {
			fmt.Printf("  %d. ✗ Invalid: %s\n", i+1, sample)
		}
	}
	
	fmt.Println("\nInvalid JSON:")
	for i, sample := range invalidJSONSamples {
		if isValidJSON(sample) {
			fmt.Printf("  %d. ✓ Valid: %s\n", i+1, sample)
		} else {
			fmt.Printf("  %d. ✗ Invalid: %s\n", i+1, sample)
		}
	}
}

func isValidJSON(jsonStr string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(jsonStr), &js) == nil
}

// JSON transformation example
func transformJSON(input []byte) ([]byte, error) {
	var data map[string]interface{}
	
	err := json.Unmarshal(input, &data)
	if err != nil {
		return nil, err
	}
	
	// Transform: add timestamp and normalize keys
	transformed := make(map[string]interface{})
	transformed["timestamp"] = time.Now().Unix()
	transformed["version"] = "2.0"
	
	// Copy and normalize original data
	for key, value := range data {
		// Convert keys to lowercase
		normalizedKey := strings.ToLower(key)
		transformed[normalizedKey] = value
	}
	
	return json.Marshal(transformed)
}