package aocutils

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

const aocURL = "https://adventofcode.com"

// GetInput fetches the Advent of Code input for the current day and year.
func GetInput() (string, error) {
    currentDir, err := os.Getwd()
    if err != nil {
        return "", err
    }

    // Extract day and year from directory names
    dayDir := filepath.Base(currentDir)
    parentDir := filepath.Base(filepath.Dir(currentDir))

    // Assuming directories are named like "day-XX" and "year-YYYY"
    dayParts := strings.Split(dayDir, "-")
    if len(dayParts) != 2 {
        return "", fmt.Errorf("unexpected day directory format: %s", dayDir)
    }
    day := dayParts[1]

    yearParts := strings.Split(parentDir, "-")
    if len(yearParts) != 2 {
        return "", fmt.Errorf("unexpected year directory format: %s", parentDir)
    }
    year := yearParts[1]

    fileName := fmt.Sprintf("day-%s-input.txt", day)
    if data, err := os.ReadFile(fileName); err == nil {
        return string(data), nil
    } else if !os.IsNotExist(err) {
        return "", err
    }

    // Fetch input from the Advent of Code website
    url := fmt.Sprintf("%s/%s/day/%s/input", aocURL, year, day)
    fmt.Println("Fetching input from URL:", url)

    session := os.Getenv("AOC")
    if session == "" {
        return "", fmt.Errorf("AOC session token not found in environment variable 'AOC'")
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }
    req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch input, status code: %d", resp.StatusCode)
    }

    content, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    err = os.WriteFile(fileName, content, 0666)
    if err != nil {
        return "", err
    }

    return string(content), nil
}

