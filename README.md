# MakeGo : Project Layout Builder
Introduction

MakeGo is a command-line tool designed to simplify the process of setting up a new Go project. It automates the creation of a standard project structure, allowing developers to focus on building their application right away.
Features

    Easy-to-use CLI interface.
    Generates a basic Go project structure, including common directories like cmd, pkg, internal, and more.
    Customizable project layout to suit different needs.

Installation
### Prerequisites
- Go (Version 1.21.4 or higher) [Download Go](https://golang.org/dl/)

### Steps
1. **Install the CLI**: Run the following command:
   ```bash
   go install github.com/Fulenn/MakeGo@latest



2. **Build the Executable**: Navigate to the cloned directory and build the executable:

    ```bash
    cd MakeGo
    go build -o makego.exe
   
### Usage
**To create a new Go project with a standard layout, run the following command:**
    
    ```bash
    makego [project-name]

Replace [project-name] with your desired project name. This will create a new directory with the specified name and the standard Go project layout.
Contributing

Contributions to MakeGo are welcome! Please refer to the contributing guidelines for more information on how to submit pull requests, report issues, or suggest enhancements.
License

This project is licensed under the MIT License - see the LICENSE file for details.