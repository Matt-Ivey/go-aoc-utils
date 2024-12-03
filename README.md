# aocutils

A Go utility package to simplify fetching input data for [Advent of Code](https://adventofcode.com) challenges.

## Overview

The `aocutils` package provides a convenient function `GetInput` that automates the process of retrieving your personal puzzle input for a given day and year from the Advent of Code website. It handles session authentication, input caching, and file management, allowing you to focus on solving the puzzles.

## Features

- **Automatic Input Fetching**: Retrieves the input for the current day and year based on your directory structure.
- **Caching**: Saves the input locally to avoid unnecessary network requests.
- **Session Management**: Uses your Advent of Code session token securely stored in an environment variable.

## Setup

### Environment Variable

Set the `AOC` environment variable with your Advent of Code session token. This token is required to authenticate your requests to the website.

**Important**: Keep your session token secure. Do not share it or commit it to version control.

#### How to Set the Environment Variable

- **Unix/Linux/macOS**:

  ```bash
  export AOC=your_session_token_here
  ```

- **Windows Command Prompt**:

  ```cmd
  set AOC=your_session_token_here
  ```

- **Windows PowerShell**:

  ```powershell
  $env:AOC="your_session_token_here"
  ```

### Directory Structure

The `GetInput` function relies on a specific directory naming convention to determine the current day and year:

- **Year Directory**: `year-YYYY`
- **Day Directory**: `day-DD`

**Example**:

```
/path/to/your/projects/year-2023/day-01/
```

Ensure your working directory follows this structure before running your program.

### Fetch Input Data

Use the `GetInput` function to retrieve the puzzle input:

```go
input, err := aocutils.GetInput()
if err != nil {
    fmt.Println("Error fetching input:", err)
    return
}
fmt.Println("Puzzle Input:", input)
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
