package main

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func GetAPIKey() (string, error) {
	// Start Playwright
	pw, err := playwright.Run()
	if err != nil {
		return "Playwright failed to start.", err
	}
	defer pw.Stop()

	// Launch Chromium browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // must be false for SSO to work because it triggers a login!
	})
	if err != nil {
		return "Failed to launch browser.", err
	}
	defer browser.Close()

	// Create a fresh context
	context, err := browser.NewContext()
	if err != nil {
		return "Failed to create context.", err
	}
	defer context.Close()

	// Open new page
	page, err := context.NewPage()
	if err != nil {
		return "Failed to open new page.", err
	}

	// Start SSO flow
	_, err = page.Goto("https://support.stedwards.edu/TDWebApi/api/auth/loginsso")
	if err != nil {
		return "Failed to navigate to login page.", err
	}

	// Wait until redirected to the final API key page
	// Replace with your actual expected final URL
	err = page.WaitForURL("https://support.stedwards.edu/TDWebApi/api/auth/loginsso", playwright.PageWaitForURLOptions{
		Timeout: playwright.Float(60000),
	})
	if err != nil {
		return "Failed to wait for final API key page.", err
	}

	// Extract API key from body
	apiKey, err := page.TextContent("body")
	if err != nil {
		return "Failed to extract API key.", err
	}
	if apiKey == "" {
		return "API key not found in body text.", fmt.Errorf("API key not found in body text")
	}

	return apiKey, nil
}
