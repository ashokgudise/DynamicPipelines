package replicator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func AddServiceToCompose(image, serviceName, composeFilePath string) (string, error) {
	// Check if the Docker Compose file already exists
	fileExists := checkFileExists(composeFilePath)

	if !fileExists {
		// Create a new Docker Compose file if it doesn't exist
		err := createComposeFile(composeFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to create Docker Compose file: %w", err)
		}
	}

	// Read the existing Docker Compose file
	yamlContent, err := ioutil.ReadFile(composeFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read Docker Compose file: %w", err)
	}

	// Generate the YAML content for the new service
	serviceYAML := generateServiceYAML(image, serviceName)

	// Append the new service definition to the existing content
	newContent := appendServiceToYAML(yamlContent, serviceYAML)

	// Write the updated content back to the Docker Compose file
	err = ioutil.WriteFile(composeFilePath, []byte(newContent), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write updated Docker Compose file: %w", err)
	}

	return newContent, nil
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createComposeFile(filePath string) error {
	defaultContent := []byte("version: '3'\n\nservices:\n")
	err := ioutil.WriteFile(filePath, defaultContent, 0644)
	if err != nil {
		return err
	}
	return nil
}

func generateServiceYAML(image, serviceName string) string {
	// Define the service YAML template
	serviceTemplate := `
  %s:
    image: %s
    # Add any other configurations you need for the service
	`

	// Fill in the template with the image and service name
	serviceYAML := fmt.Sprintf(serviceTemplate, serviceName, image)

	return serviceYAML
}

func appendServiceToYAML(existingContent []byte, serviceYAML string) string {
	// Convert existing content to string
	contentStr := string(existingContent)

	// Find the position of the last occurrence of "services:"
	index := strings.LastIndex(contentStr, "services:")
	if index == -1 {
		fmt.Println("Failed to find 'services' section in Docker Compose file")
		return contentStr
	}

	// Find the position of the next newline character after the last occurrence of "services:"
	newlineIndex := strings.Index(contentStr[index:], "\n")
	if newlineIndex == -1 {
		fmt.Println("Failed to find the position to insert the new service definition")
		return contentStr
	}

	// Calculate the position to insert the new service definition
	insertIndex := index + newlineIndex + 1

	// Insert the new service definition at the calculated position
	newContent := contentStr[:insertIndex] + serviceYAML + contentStr[insertIndex:]

	return newContent
}
