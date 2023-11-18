# MakeGo : Project Layout Builder
Introduction

MakeGo is a command-line tool designed to simplify the process of setting up a new Go project. It automates the creation of a standard project structure, allowing developers to focus on building their application right away.
Features

    Easy-to-use CLI interface.
    Generates a basic Go project structure, including common directories like cmd, pkg, internal, and more.
    Customizable project layout to suit different needs.

Installation
Prerequisites

    Go (Version 1.xx or higher) Download Go

Steps

    Clone the Repository: First, clone the repository to your local machine using Git:

    git clone https://github.com/yourusername/MakeGo.git



Build the Executable: Navigate to the cloned directory and build the executable:


    cd MakeGo
    go build -o makego.exe

Add to PATH (Windows): Add the directory containing makego.exe to your system's PATH to use the command from anywhere:
Right-click on 'This PC' or 'My Computer' → 'Properties' → 'Advanced system settings' → 'Environment Variables'.
Under 'System variables', find and select the 'Path' variable, then click 'Edit'.
Click 'New' and add the full path to the directory containing makego.exe.
Click 'OK' to save and close all windows.
Open a new command prompt to use the makego command.

Usage

To create a new Go project with a standard layout, run the following command:

    makego [project-name]

Replace [project-name] with your desired project name. This will create a new directory with the specified name and the standard Go project layout.
Contributing

Contributions to MakeGo are welcome! Please refer to the contributing guidelines for more information on how to submit pull requests, report issues, or suggest enhancements.
License

This project is licensed under the MIT License - see the LICENSE file for details.