// GoForGo Solution: Hadoop YARN
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type YARNClient struct {
	baseURL string
	client  *http.Client
}

type ApplicationInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

func NewYARNClient(resourceManagerURL string) *YARNClient {
	return &YARNClient{
		baseURL: resourceManagerURL,
		client:  &http.Client{},
	}
}

func (y *YARNClient) listApplications(ctx context.Context) ([]ApplicationInfo, error) {
	url := fmt.Sprintf("%s/ws/v1/cluster/apps", y.baseURL)
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	resp, err := y.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var result struct {
		Apps struct {
			App []ApplicationInfo `json:"app"`
		} `json:"apps"`
	}
	
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	return result.Apps.App, nil
}

func main() {
	client := NewYARNClient("http://localhost:8088")
	
	apps, err := client.listApplications(context.Background())
	if err != nil {
		log.Printf("Error listing applications: %v", err)
	} else {
		fmt.Printf("Found %d applications\n", len(apps))
		for _, app := range apps {
			fmt.Printf("App: %s (%s) - %s\n", app.Name, app.ID, app.State)
		}
	}
	
	fmt.Println("Hadoop YARN operations completed!")
}