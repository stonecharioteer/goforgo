// nosql_embedded.go
// Learn embedded NoSQL database operations with BoltDB

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	
	"go.etcd.io/bbolt"
)

// Document structure for NoSQL operations
type Document struct {
	ID        string                 `json:"id"`
	Title     string                 `json:"title"`
	Content   string                 `json:"content"`
	Tags      []string               `json:"tags"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// Search result structure
type SearchResult struct {
	Document  Document `json:"document"`
	Score     int      `json:"score"`
	MatchType string   `json:"match_type"`
}

func main() {
	fmt.Println("=== Embedded NoSQL Database (BoltDB) ===")
	
	// Open BoltDB database
	db, err := bbolt.Open("documents.db", 0600, nil)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()
	
	// Initialize database buckets
	if err := initializeBuckets(db); err != nil {
		log.Fatal("Failed to initialize buckets:", err)
	}
	
	// Demonstrate document operations
	fmt.Println("1. Creating documents...")
	if err := createDocuments(db); err != nil {
		log.Printf("Error creating documents: %v", err)
	}
	
	fmt.Println("\n2. Reading documents...")
	if err := readDocuments(db); err != nil {
		log.Printf("Error reading documents: %v", err)
	}
	
	fmt.Println("\n3. Updating document...")
	if err := updateDocument(db, "doc1"); err != nil {
		log.Printf("Error updating document: %v", err)
	}
	
	fmt.Println("\n4. Searching documents...")
	if err := searchDocuments(db, "golang"); err != nil {
		log.Printf("Error searching documents: %v", err)
	}
	
	fmt.Println("\n5. Listing by tags...")
	if err := listDocumentsByTag(db, "tutorial"); err != nil {
		log.Printf("Error listing by tag: %v", err)
	}
	
	fmt.Println("\n6. Database statistics...")
	showDatabaseStats(db)
	
	fmt.Println("\n7. Deleting document...")
	if err := deleteDocument(db, "doc2"); err != nil {
		log.Printf("Error deleting document: %v", err)
	}
	
	fmt.Println("\n8. Final document list...")
	if err := readDocuments(db); err != nil {
		log.Printf("Error reading documents: %v", err)
	}
	
	fmt.Println("\nNoSQL operations completed successfully!")
}

// Initialize database buckets
func initializeBuckets(db *bbolt.DB) error {
	return db.Update(func(tx *bbolt.Tx) error {
		// Create buckets for documents and indices
		_, err := tx.CreateBucketIfNotExists([]byte("documents"))
		if err != nil {
			return err
		}
		
		_, err = tx.CreateBucketIfNotExists([]byte("tags_index"))
		if err != nil {
			return err
		}
		
		_, err = tx.CreateBucketIfNotExists([]byte("metadata"))
		if err != nil {
			return err
		}
		
		fmt.Println("Database buckets initialized")
		return nil
	})
}

// Create sample documents
func createDocuments(db *bbolt.DB) error {
	documents := []Document{
		{
			ID:      "doc1",
			Title:   "Go Programming Basics",
			Content: "Learn Go programming language fundamentals including variables, functions, and control flow.",
			Tags:    []string{"go", "programming", "tutorial", "basics"},
			Metadata: map[string]interface{}{
				"author":     "John Doe",
				"difficulty": "beginner",
				"rating":     4.5,
			},
		},
		{
			ID:      "doc2",
			Title:   "Advanced Go Concurrency",
			Content: "Deep dive into goroutines, channels, and advanced concurrency patterns in Go.",
			Tags:    []string{"go", "concurrency", "advanced", "goroutines"},
			Metadata: map[string]interface{}{
				"author":     "Jane Smith",
				"difficulty": "advanced",
				"rating":     4.8,
			},
		},
		{
			ID:      "doc3",
			Title:   "Database Integration with Go",
			Content: "Complete guide to working with SQL and NoSQL databases in Go applications.",
			Tags:    []string{"go", "database", "sql", "nosql", "tutorial"},
			Metadata: map[string]interface{}{
				"author":     "Bob Johnson",
				"difficulty": "intermediate",
				"rating":     4.3,
			},
		},
	}
	
	return db.Update(func(tx *bbolt.Tx) error {
		// Get buckets
		docBucket := tx.Bucket([]byte("documents"))
		tagsBucket := tx.Bucket([]byte("tags_index"))
		
		for _, doc := range documents {
			// Set timestamps
			now := time.Now()
			doc.CreatedAt = now
			doc.UpdatedAt = now
			
			// Serialize document to JSON
			docJSON, err := json.Marshal(doc)
			if err != nil {
				return fmt.Errorf("failed to marshal document %s: %w", doc.ID, err)
			}
			
			// Store document
			if err := docBucket.Put([]byte(doc.ID), docJSON); err != nil {
				return fmt.Errorf("failed to store document %s: %w", doc.ID, err)
			}
			
			// Update tag indices
			for _, tag := range doc.Tags {
				// Get existing documents for this tag
				var docIDs []string
				if tagData := tagsBucket.Get([]byte(tag)); tagData != nil {
					json.Unmarshal(tagData, &docIDs)
				}
				
				// Add current document ID if not already present
				found := false
				for _, id := range docIDs {
					if id == doc.ID {
						found = true
						break
					}
				}
				if !found {
					docIDs = append(docIDs, doc.ID)
				}
				
				// Store updated tag index
				updatedTagData, err := json.Marshal(docIDs)
				if err != nil {
					return fmt.Errorf("failed to marshal tag data: %w", err)
				}
				
				if err := tagsBucket.Put([]byte(tag), updatedTagData); err != nil {
					return fmt.Errorf("failed to update tag index: %w", err)
				}
			}
			
			fmt.Printf("Created document: %s - %s\n", doc.ID, doc.Title)
		}
		
		return nil
	})
}

// Read all documents
func readDocuments(db *bbolt.DB) error {
	return db.View(func(tx *bbolt.Tx) error {
		// Get documents bucket
		bucket := tx.Bucket([]byte("documents"))
		
		fmt.Println("All documents:")
		fmt.Println("ID    | Title                      | Tags                    | Author")
		fmt.Println("------|----------------------------|-------------------------|----------")
		
		// Iterate through all documents
		return bucket.ForEach(func(k, v []byte) error {
			var doc Document
			if err := json.Unmarshal(v, &doc); err != nil {
				log.Printf("Failed to unmarshal document %s: %v", k, err)
				return nil
			}
			
			// Format tags for display
			tagsStr := strings.Join(doc.Tags, ", ")
			if len(tagsStr) > 20 {
				tagsStr = tagsStr[:20] + "..."
			}
			
			// Get author from metadata
			author := "Unknown"
			if authorVal, ok := doc.Metadata["author"]; ok {
				if authorStr, ok := authorVal.(string); ok {
					author = authorStr
				}
			}
			
			fmt.Printf("%-5s | %-26s | %-23s | %s\n",
				doc.ID, doc.Title, tagsStr, author)
			return nil
		})
	})
}

// Update existing document
func updateDocument(db *bbolt.DB, docID string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		// Get documents bucket
		bucket := tx.Bucket([]byte("documents"))
		
		// Retrieve existing document
		docData := bucket.Get([]byte(docID))
		if docData == nil {
			return fmt.Errorf("document %s not found", docID)
		}
		
		// Parse existing document
		var doc Document
		if err := json.Unmarshal(docData, &doc); err != nil {
			return fmt.Errorf("failed to unmarshal document: %w", err)
		}
		
		// Update document fields
		doc.Title = doc.Title + " (Updated)"
		doc.Content = doc.Content + "\n\nThis document has been updated with additional information."
		doc.UpdatedAt = time.Now()
		doc.Tags = append(doc.Tags, "updated")
		
		// Update metadata
		if doc.Metadata == nil {
			doc.Metadata = make(map[string]interface{})
		}
		doc.Metadata["last_updated_by"] = "system"
		doc.Metadata["version"] = 2
		
		// Serialize and store updated document
		updatedJSON, err := json.Marshal(doc)
		if err != nil {
			return fmt.Errorf("failed to marshal updated document: %w", err)
		}
		
		if err := bucket.Put([]byte(docID), updatedJSON); err != nil {
			return fmt.Errorf("failed to store updated document: %w", err)
		}
		
		fmt.Printf("Successfully updated document: %s\n", docID)
		return nil
	})
}

// Search documents by content
func searchDocuments(db *bbolt.DB, searchTerm string) error {
	var results []SearchResult
	
	err := db.View(func(tx *bbolt.Tx) error {
		// Get documents bucket
		bucket := tx.Bucket([]byte("documents"))
		
		// Iterate and search documents
		return bucket.ForEach(func(k, v []byte) error {
			var doc Document
			if err := json.Unmarshal(v, &doc); err != nil {
				return nil // Skip invalid documents
			}
			
			// Simple text search (case-insensitive)
			score := 0
			matchType := ""
			searchLower := strings.ToLower(searchTerm)
			
			// Search in title
			if strings.Contains(strings.ToLower(doc.Title), searchLower) {
				score += 3
				matchType = "title"
			}
			
			// Search in content
			if strings.Contains(strings.ToLower(doc.Content), searchLower) {
				score += 2
				if matchType == "" {
					matchType = "content"
				} else {
					matchType = "title+content"
				}
			}
			
			// Search in tags
			for _, tag := range doc.Tags {
				if strings.Contains(strings.ToLower(tag), searchLower) {
					score += 1
					if matchType == "" {
						matchType = "tags"
					}
					break
				}
			}
			
			// Add to results if match found
			if score > 0 {
				results = append(results, SearchResult{
					Document:  doc,
					Score:     score,
					MatchType: matchType,
				})
			}
			
			return nil
		})
	})
	
	if err != nil {
		return err
	}
	
	// Display search results
	if len(results) == 0 {
		fmt.Printf("No documents found matching '%s'\n", searchTerm)
		return nil
	}
	
	fmt.Printf("Search results for '%s':\n", searchTerm)
	for i, result := range results {
		fmt.Printf("%d. %s (Score: %d, Match: %s)\n",
			i+1, result.Document.Title, result.Score, result.MatchType)
	}
	
	return nil
}

// List documents by tag
func listDocumentsByTag(db *bbolt.DB, tag string) error {
	return db.View(func(tx *bbolt.Tx) error {
		// Get buckets
		tagsBucket := tx.Bucket([]byte("tags_index"))
		docsBucket := tx.Bucket([]byte("documents"))
		
		// Get document IDs for the tag
		tagData := tagsBucket.Get([]byte(tag))
		if tagData == nil {
			fmt.Printf("No documents found for tag '%s'\n", tag)
			return nil
		}
		
		// Parse document IDs
		var docIDs []string
		if err := json.Unmarshal(tagData, &docIDs); err != nil {
			return fmt.Errorf("failed to unmarshal tag data: %w", err)
		}
		
		// Display documents with this tag
		fmt.Printf("Documents tagged with '%s':\n", tag)
		for _, docID := range docIDs {
			// Get document by ID
			docData := docsBucket.Get([]byte(docID))
			if docData == nil {
				continue
			}
			
			// Parse and display document
			var doc Document
			if err := json.Unmarshal(docData, &doc); err != nil {
				continue
			}
			
			fmt.Printf("- %s: %s\n", doc.ID, doc.Title)
		}
		
		return nil
	})
}

// Show database statistics
func showDatabaseStats(db *bbolt.DB) {
	db.View(func(tx *bbolt.Tx) error {
		// Get database stats
		stats := db.Stats()
		
		fmt.Printf("Database Statistics:\n")
		fmt.Printf("Page Size: %d bytes\n", stats.PageSize)
		fmt.Printf("Free Pages: %d\n", stats.FreePageN)
		fmt.Printf("Pending Pages: %d\n", stats.PendingPageN)
		
		// Count documents
		docsBucket := tx.Bucket([]byte("documents"))
		docCount := 0
		docsBucket.ForEach(func(k, v []byte) error {
			docCount++
			return nil
		})
		
		// Count tags
		tagsBucket := tx.Bucket([]byte("tags_index"))
		tagCount := 0
		tagsBucket.ForEach(func(k, v []byte) error {
			tagCount++
			return nil
		})
		
		fmt.Printf("Documents: %d\n", docCount)
		fmt.Printf("Tags: %d\n", tagCount)
		
		return nil
	})
}

// Delete document
func deleteDocument(db *bbolt.DB, docID string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		// Get buckets
		docsBucket := tx.Bucket([]byte("documents"))
		tagsBucket := tx.Bucket([]byte("tags_index"))
		
		// Get document to delete
		docData := docsBucket.Get([]byte(docID))
		if docData == nil {
			return fmt.Errorf("document %s not found", docID)
		}
		
		// Parse document to get tags
		var doc Document
		if err := json.Unmarshal(docData, &doc); err != nil {
			return fmt.Errorf("failed to unmarshal document: %w", err)
		}
		
		// Remove from tag indices
		for _, tag := range doc.Tags {
			// Get existing document IDs for tag
			var docIDs []string
			if tagData := tagsBucket.Get([]byte(tag)); tagData != nil {
				json.Unmarshal(tagData, &docIDs)
			}
			
			// Remove document ID from list
			var updatedIDs []string
			for _, id := range docIDs {
				if id != docID {
					updatedIDs = append(updatedIDs, id)
				}
			}
			
			// Update or delete tag index
			if len(updatedIDs) == 0 {
				tagsBucket.Delete([]byte(tag))
			} else {
				updatedTagData, _ := json.Marshal(updatedIDs)
				tagsBucket.Put([]byte(tag), updatedTagData)
			}
		}
		
		// Delete document
		if err := docsBucket.Delete([]byte(docID)); err != nil {
			return fmt.Errorf("failed to delete document: %w", err)
		}
		
		fmt.Printf("Successfully deleted document: %s\n", docID)
		return nil
	})
}