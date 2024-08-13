package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Constants for file contents
const (
	makefileTemplate = `CC = gcc
CFLAGS = -Wall -Wextra -Werror -Iincludes

# List source files explicitly or use a more controlled pattern
SRCS = $(wildcard src/*.c) $(wildcard src/**/*.c)
OBJS = $(patsubst src/%%.c, build/%%.o, $(SRCS))
NAME = build/%s
RM = rm -f

# Default target
all: $(NAME)

# Link the final executable
$(NAME): $(OBJS)
	$(CC) $(CFLAGS) -o $(NAME) $(OBJS)

# Compile source files into object files, creating necessary directories
build/%%.o: src/%%.c | build
	mkdir -p $(dir $@)
	$(CC) $(CFLAGS) -c $< -o $@

# Ensure the build directory exists
build:
	mkdir -p build
	mkdir -p build/utils

# Clean up object files
clean:
	$(RM) $(OBJS)

# Clean up all artifacts
fclean: clean
	$(RM) $(NAME)

# Rebuild everything
re: fclean all

.PHONY: all clean fclean re`

	infoShTemplate = `#!/bin/bash
echo "Project %s Information"
echo "==============================="
echo "Structure:"
tree
echo "==============================="
echo "Header files:"
ls -lptm includes/
`

	gitignoreContent = `# Compiled object files
*.o

# Executables
*.exe
*.out

# Libraries
*.a
*.so

# Build directories
build/
bin/

# VS Code
.vscode/

# macOS
.DS_Store
`
)

// createDirectory creates a directory with the given path
func createDirectory(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", path, err)
		os.Exit(1)
	}
}

// createFile creates a file with the given path and writes content to it
func createFile(path, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", path, err)
		os.Exit(1)
	}
	defer file.Close()

	if content != "" {
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", path, err)
			os.Exit(1)
		}
	}
}

// createNewEntity creates a new entity with a source file and a header file
func createNewEntity(projectName, entityName string) {
	entityNameUpper := strings.ToUpper(entityName)
	entityDir := filepath.Join(projectName, "src", entityName)
	headerFileName := filepath.Join(projectName, "includes", entityName+".h")
	headerContent := fmt.Sprintf(`#ifndef %s_H_
#define %s_H_

// %s entity functions

#endif
`, entityNameUpper, entityNameUpper, entityNameUpper)

	createDirectory(entityDir)
	createFile(filepath.Join(entityDir, entityName+".c"), "")
	createFile(headerFileName, headerContent)

	fmt.Printf("Entity '%s' created successfully.\n", entityName)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <project_name> | new-entity <entity_name>")
		os.Exit(1)
	}

	if os.Args[1] == "new-entity" {
		if len(os.Args) != 3 {
			fmt.Println("Usage: new-entity <entity_name>")
			os.Exit(1)
		}
		entityName := os.Args[2]
		projectName := "."

		// Check if the Makefile exists in the current directory to infer project name
		if _, err := os.Stat("Makefile"); err == nil {
			projectName = filepath.Base(filepath.Dir("Makefile"))
		}

		createNewEntity(projectName, entityName)
		return
	}

	projectName := os.Args[1]

	// Create directories
	dirs := []string{
		filepath.Join(projectName, "includes"),
		filepath.Join(projectName, "src"),
		filepath.Join(projectName, "src", "utils"), // utils folder inside src
		filepath.Join(projectName, "tools"),
	}

	for _, dir := range dirs {
		createDirectory(dir)
	}

	// Create Makefile
	makefileContent := fmt.Sprintf(makefileTemplate, projectName)
	createFile(filepath.Join(projectName, "Makefile"), makefileContent)

	// Create tools/info.sh
	infoShContent := fmt.Sprintf(infoShTemplate, projectName)
	createFile(filepath.Join(projectName, "tools/info.sh"), infoShContent)

	// Make tools/info.sh executable
	err := os.Chmod(filepath.Join(projectName, "tools/info.sh"), 0755)
	if err != nil {
		fmt.Printf("Error making tools/info.sh executable: %v\n", err)
		os.Exit(1)
	}

	// Create .gitignore
	createFile(filepath.Join(projectName, ".gitignore"), gitignoreContent)

	// Create utils source and header files in their own folders
	createFile(filepath.Join(projectName, "src", "main.c"), `#include <stdio.h>
#include "utils.h"
	
int main(void) {
	print_hello_world();
	return 0;
}`)
	createFile(filepath.Join(projectName, "src", "utils", "utils.c"), `#include <stdio.h>

void print_hello_world() {
    printf("Hello, world!\n");
}`)
	createFile(filepath.Join(projectName, "includes", "utils.h"), `#ifndef UTILS_H_
#define UTILS_H_

// Here your functions
void print_hello_world();

#endif
`)

	fmt.Printf("Project structure for '%s' created successfully.\n", projectName)
}
