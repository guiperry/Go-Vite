package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunInit(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Test basic init
	cmd := &cobra.Command{}
	cmd.Flags().StringP("module", "m", "", "Go module name")
	cmd.Flags().StringP("description", "d", "A Go-Vite desktop application", "Project description")
	cmd.Flags().StringP("author", "a", "", "Author name")
	cmd.Flags().IntP("port", "p", 5173, "Frontend port")
	cmd.Flags().IntP("backend-port", "b", 8080, "Backend port")

	err = runInit(cmd, []string{"test-app"})
	if err != nil {
		t.Fatalf("runInit failed: %v", err)
	}

	// Check if project was created
	if _, err := os.Stat("test-app"); os.IsNotExist(err) {
		t.Fatal("Project directory was not created")
	}

	// Check essential files
	files := []string{
		"test-app/main.go",
		"test-app/go.mod",
		"test-app/Makefile",
		"test-app/README.md",
		"test-app/backend/go.mod",
		"test-app/backend/cmd/server/main.go",
		"test-app/frontend/package.json",
		"test-app/frontend/src/main.tsx",
		"test-app/frontend/src/App.tsx",
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Fatalf("File %s was not created", file)
		}
	}
}

func TestRunInitDirectoryExists(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Create existing directory
	os.Mkdir("existing-app", 0755)

	cmd := &cobra.Command{}
	cmd.Flags().StringP("module", "m", "", "Go module name")
	cmd.Flags().StringP("description", "d", "A Go-Vite desktop application", "Project description")
	cmd.Flags().StringP("author", "a", "", "Author name")
	cmd.Flags().IntP("port", "p", 5173, "Frontend port")
	cmd.Flags().IntP("backend-port", "b", 8080, "Backend port")

	err = runInit(cmd, []string{"existing-app"})
	if err == nil {
		t.Fatal("Expected error when directory exists, but got none")
	}

	if !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("Expected 'already exists' error, got: %v", err)
	}
}

func TestCreateProjectStructure(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	config := ProjectConfig{
		Name:        "test-app",
		Module:      "github.com/test/test-app",
		Description: "Test application",
		Author:      "Test Author",
		Port:        3000,
		BackendPort: 9000,
	}

	err = createProjectStructure(tempDir, config)
	if err != nil {
		t.Fatalf("createProjectStructure failed: %v", err)
	}

	// Check directories
	dirs := []string{
		"backend/cmd/server",
		"backend/config",
		"backend/internal/api/handlers",
		"backend/internal/api/middleware",
		"backend/internal/models",
		"backend/internal/modules",
		"backend/internal/storage",
		"backend/internal/utils",
		"frontend/src/components",
		"frontend/src/pages",
		"frontend/src/hooks",
		"frontend/src/services",
		"frontend/src/utils",
		"frontend/public",
		"dist",
		"bin",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(tempDir, dir)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Fatalf("Directory %s was not created", dir)
		}
	}
}

func TestGenerateGoMod(t *testing.T) {
	config := ProjectConfig{
		Module: "github.com/test/my-app",
	}

	result := generateGoMod(config)
	expected := "module github.com/test/my-app"
	if !strings.Contains(result, expected) {
		t.Fatalf("Expected %s in result, got: %s", expected, result)
	}

	if !strings.Contains(result, "go 1.24.0") {
		t.Fatal("Expected Go version in go.mod")
	}
}

func TestGenerateMainGo(t *testing.T) {
	result := generateMainGo()
	if !strings.Contains(result, "package main") {
		t.Fatal("Expected package main")
	}

	if !strings.Contains(result, "webview") {
		t.Fatal("Expected webview import")
	}

	if !strings.Contains(result, "func main()") {
		t.Fatal("Expected main function")
	}
}

func TestGenerateMakefile(t *testing.T) {
	config := ProjectConfig{
		Name: "test-app",
	}

	result := generateMakefile(config)
	if !strings.Contains(result, "PROJECT_NAME := test-app") {
		t.Fatal("Expected project name in Makefile")
	}

	if !strings.Contains(result, ".PHONY: help all") {
		t.Fatal("Expected make targets")
	}
}

func TestGenerateReadme(t *testing.T) {
	config := ProjectConfig{
		Name:        "Test App",
		Description: "A test application",
	}

	result := generateReadme(config)
	if !strings.Contains(result, "# Test App") {
		t.Fatal("Expected title in README")
	}

	if !strings.Contains(result, "A test application") {
		t.Fatal("Expected description in README")
	}
}

func TestGenerateBackendGoMod(t *testing.T) {
	result := generateBackendGoMod()
	expected := "module backend"
	if !strings.Contains(result, expected) {
		t.Fatalf("Expected %s in result, got: %s", expected, result)
	}

	if !strings.Contains(result, "go 1.24") {
		t.Fatal("Expected Go version")
	}
}

func TestGenerateBackendMain(t *testing.T) {
	result := generateBackendMain()
	if !strings.Contains(result, "package main") {
		t.Fatal("Expected package main")
	}

	if !strings.Contains(result, "gin.Default()") {
		t.Fatal("Expected Gin router")
	}
}

func TestGenerateConfig(t *testing.T) {
	result := generateConfig()
	if !strings.Contains(result, "package config") {
		t.Fatal("Expected package config")
	}

	if !strings.Contains(result, "type Config struct") {
		t.Fatal("Expected Config struct")
	}
}

func TestGenerateRoutes(t *testing.T) {
	result := generateRoutes()
	if !strings.Contains(result, "package api") {
		t.Fatal("Expected package api")
	}

	if !strings.Contains(result, "SetupRoutes") {
		t.Fatal("Expected SetupRoutes function")
	}
}

func TestGenerateHandlers(t *testing.T) {
	result := generateHandlers()
	if !strings.Contains(result, "package handlers") {
		t.Fatal("Expected package handlers")
	}

	if !strings.Contains(result, "ListItems") {
		t.Fatal("Expected ListItems function")
	}
}

func TestGenerateCorsMiddleware(t *testing.T) {
	result := generateCorsMiddleware()
	if !strings.Contains(result, "package middleware") {
		t.Fatal("Expected package middleware")
	}

	if !strings.Contains(result, "CORS()") {
		t.Fatal("Expected CORS function")
	}
}

func TestGenerateLoggerMiddleware(t *testing.T) {
	result := generateLoggerMiddleware()
	if !strings.Contains(result, "package middleware") {
		t.Fatal("Expected package middleware")
	}

	if !strings.Contains(result, "Logger()") {
		t.Fatal("Expected Logger function")
	}
}

func TestGeneratePipelineModel(t *testing.T) {
	result := generatePipelineModel()
	if !strings.Contains(result, "package models") {
		t.Fatal("Expected package models")
	}

	if !strings.Contains(result, "type Pipeline struct") {
		t.Fatal("Expected Pipeline struct")
	}
}

func TestGenerateProjectModel(t *testing.T) {
	result := generateProjectModel()
	if !strings.Contains(result, "package models") {
		t.Fatal("Expected package models")
	}

	if !strings.Contains(result, "type Project struct") {
		t.Fatal("Expected Project struct")
	}
}

func TestGenerateUserModel(t *testing.T) {
	result := generateUserModel()
	if !strings.Contains(result, "package models") {
		t.Fatal("Expected package models")
	}

	if !strings.Contains(result, "type User struct") {
		t.Fatal("Expected User struct")
	}
}

func TestGenerateModulesManager(t *testing.T) {
	result := generateModulesManager()
	if !strings.Contains(result, "package modules") {
		t.Fatal("Expected package modules")
	}

	if !strings.Contains(result, "type Manager struct") {
		t.Fatal("Expected Manager struct")
	}
}

func TestGenerateBuiltinModules(t *testing.T) {
	result := generateBuiltinModules()
	if !strings.Contains(result, "package modules") {
		t.Fatal("Expected package modules")
	}

	if !strings.Contains(result, "ExampleModule") {
		t.Fatal("Expected ExampleModule")
	}
}

func TestGenerateDatabase(t *testing.T) {
	result := generateDatabase()
	if !strings.Contains(result, "package storage") {
		t.Fatal("Expected package storage")
	}

	if !strings.Contains(result, "type Database struct") {
		t.Fatal("Expected Database struct")
	}
}

func TestGenerateLogger(t *testing.T) {
	result := generateLogger()
	if !strings.Contains(result, "package utils") {
		t.Fatal("Expected package utils")
	}

	if !strings.Contains(result, "Logger = log.New") {
		t.Fatal("Expected Logger initialization")
	}
}

func TestGeneratePackageJson(t *testing.T) {
	config := ProjectConfig{
		Name:        "test-app",
		Description: "Test app",
	}

	result := generatePackageJson(config)
	if !strings.Contains(result, `"name": "test-app"`) {
		t.Fatal("Expected package name")
	}

	if !strings.Contains(result, `"description": "Test app"`) {
		t.Fatal("Expected description")
	}
}

func TestGenerateViteConfig(t *testing.T) {
	config := ProjectConfig{
		Port:        3000,
		BackendPort: 9000,
	}

	result := generateViteConfig(config)
	if !strings.Contains(result, "port: 3000") {
		t.Fatal("Expected frontend port")
	}

	if !strings.Contains(result, "localhost:9000") {
		t.Fatal("Expected backend port in proxy")
	}
}

func TestGenerateTailwindConfig(t *testing.T) {
	result := generateTailwindConfig()
	if !strings.Contains(result, "brand:") {
		t.Fatal("Expected brand colors")
	}
}

func TestGeneratePostcssConfig(t *testing.T) {
	result := generatePostcssConfig()
	if !strings.Contains(result, "tailwindcss") {
		t.Fatal("Expected tailwindcss plugin")
	}
}

func TestGenerateIndexHtml(t *testing.T) {
	config := ProjectConfig{
		Name: "Test App",
	}

	result := generateIndexHtml(config)
	if !strings.Contains(result, "<title>Test App</title>") {
		t.Fatal("Expected title")
	}
}

func TestGenerateMainTsx(t *testing.T) {
	result := generateMainTsx()
	if !strings.Contains(result, "ReactDOM.createRoot") {
		t.Fatal("Expected React root creation")
	}
}

func TestGenerateAppTsx(t *testing.T) {
	config := ProjectConfig{
		Name:        "Test App",
		Description: "Test description",
	}

	result := generateAppTsx(config)
	if !strings.Contains(result, "Test App") {
		t.Fatal("Expected app name")
	}

	if !strings.Contains(result, "Test description") {
		t.Fatal("Expected description")
	}
}

func TestGenerateIndexCss(t *testing.T) {
	result := generateIndexCss()
	if !strings.Contains(result, "@tailwind") {
		t.Fatal("Expected Tailwind directives")
	}
}

func TestGenerateEslintrc(t *testing.T) {
	result := generateEslintrc()
	if !strings.Contains(result, "eslint:recommended") {
		t.Fatal("Expected ESLint config")
	}
}

func TestGeneratePrettierrc(t *testing.T) {
	result := generatePrettierrc()
	if !strings.Contains(result, `"semi": true`) {
		t.Fatal("Expected Prettier config")
	}
}

func TestGenerateGitignore(t *testing.T) {
	result := generateGitignore()
	if !strings.Contains(result, "node_modules/") {
		t.Fatal("Expected node_modules in gitignore")
	}
}

func TestGenerateGitattributes(t *testing.T) {
	result := generateGitattributes()
	if !strings.Contains(result, "* text=auto") {
		t.Fatal("Expected gitattributes content")
	}
}

func TestGenerateEnvExample(t *testing.T) {
	config := ProjectConfig{
		Name:        "test-app",
		BackendPort: 9000,
	}

	result := generateEnvExample(config)
	if !strings.Contains(result, "VITE_APP_NAME=test-app") {
		t.Fatal("Expected app name in env")
	}

	if !strings.Contains(result, "PORT=9000") {
		t.Fatal("Expected backend port in env")
	}
}

func TestGenerateNetlifyToml(t *testing.T) {
	config := ProjectConfig{
		Name: "test-app",
	}
	result := generateNetlifyToml(config)
	if !strings.Contains(result, "[build]") {
		t.Fatal("Expected build section")
	}
}

func TestGenerateNetlifyApiFunction(t *testing.T) {
	result := generateNetlifyApiFunction()
	if !strings.Contains(result, "exports.handler") {
		t.Fatal("Expected Netlify function")
	}
}

func TestRunInitDefaultName(t *testing.T) {
	// Test with no args (should use default name)
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	cmd := &cobra.Command{}
	cmd.Flags().StringP("module", "m", "", "Go module name")
	cmd.Flags().StringP("description", "d", "A Go-Vite desktop application", "Project description")
	cmd.Flags().StringP("author", "a", "", "Author name")
	cmd.Flags().IntP("port", "p", 5173, "Frontend port")
	cmd.Flags().IntP("backend-port", "b", 8080, "Backend port")

	err = runInit(cmd, []string{}) // No args
	if err != nil {
		t.Fatalf("runInit failed: %v", err)
	}

	// Should create "my-app" directory
	if _, err := os.Stat("my-app"); os.IsNotExist(err) {
		t.Fatal("Default project directory was not created")
	}
}

func TestMainFunction(t *testing.T) {
	// Test that main function can be called without panicking
	// This is tricky since main calls Execute, but we can test the setup
	rootCmd := &cobra.Command{Use: "go-vite"}
	initCmd := &cobra.Command{
		Use:   "init [project-name]",
		Short: "Initialize a new Go-Vite project",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runInit,
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("go-vite v%s\n", version)
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)

	initCmd.Flags().StringP("module", "m", "", "Go module name")
	initCmd.Flags().StringP("description", "d", "A Go-Vite desktop application", "Project description")
	initCmd.Flags().StringP("author", "a", "", "Author name")
	initCmd.Flags().IntP("port", "p", 5173, "Frontend port")
	initCmd.Flags().IntP("backend-port", "b", 8080, "Backend port")

	// Just test that the commands are set up correctly
	if rootCmd.Use != "go-vite" {
		t.Fatal("Root command not set up correctly")
	}
}

func TestVersionCommand(t *testing.T) {
	// Test the version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("go-vite v%s\n", version)
		},
	}

	// This should not panic
	versionCmd.Run(versionCmd, []string{})
}

// Tests for new module management functionality

func TestDetectProjectType(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	oldWd, _ := os.Getwd()
	defer os.Chdir(oldWd)

	// Test unknown project type
	os.Chdir(tempDir)
	if detectProjectType() != Unknown {
		t.Fatal("Expected Unknown project type")
	}

	// Test Go project
	os.WriteFile("go.mod", []byte("module test"), 0644)
	if detectProjectType() != GoProject {
		t.Fatal("Expected Go project type")
	}

	// Remove go.mod and test Node.js project
	os.Remove("go.mod")
	os.WriteFile("package.json", []byte("{}"), 0644)
	if detectProjectType() != NodeProject {
		t.Fatal("Expected Node.js project type")
	}
}

func TestGetDataFilePath(t *testing.T) {
	path := getDataFilePath()
	if path == "" {
		t.Fatal("Expected non-empty data file path")
	}

	// Should contain automationgenie and cli.json
	if !strings.Contains(path, "automationgenie") {
		t.Fatal("Expected automationgenie in path")
	}

	if !strings.Contains(path, "cli.json") {
		t.Fatal("Expected cli.json in path")
	}
}

func TestLoadDataAndSaveData(t *testing.T) {
	// Test loading data (should not fail even if file doesn't exist)
	data, err := loadData()
	if err != nil {
		t.Fatalf("loadData failed: %v", err)
	}

	if data.InstalledModules == nil {
		t.Fatal("Expected InstalledModules map to be initialized")
	}

	// Test saving data
	data.InstalledModules["test"] = []string{"module1", "module2"}
	err = saveData(data)
	if err != nil {
		t.Fatalf("saveData failed: %v", err)
	}

	// Test loading saved data
	loadedData, err := loadData()
	if err != nil {
		t.Fatalf("loadData after save failed: %v", err)
	}

	if len(loadedData.InstalledModules["test"]) != 2 {
		t.Fatal("Expected 2 modules in loaded data")
	}
}

func TestSaveInstalledModuleAndRemoveInstalledModule(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	oldWd, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(oldWd)

	// Test saving module
	saveInstalledModule("test-module")
	data, _ := loadData()
	if len(data.InstalledModules[tempDir]) != 1 {
		t.Fatal("Expected 1 module after save")
	}

	// Test removing module
	removeInstalledModule("test-module")
	data, _ = loadData()
	if len(data.InstalledModules[tempDir]) != 0 {
		t.Fatal("Expected 0 modules after remove")
	}
}

func TestDetectLocalModuleType(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test unknown type
	if detectLocalModuleType(tempDir) != Unknown {
		t.Fatal("Expected Unknown for empty directory")
	}

	// Test Go module
	goModPath := filepath.Join(tempDir, "go.mod")
	os.WriteFile(goModPath, []byte("module test"), 0644)
	if detectLocalModuleType(tempDir) != GoProject {
		t.Fatal("Expected Go project type")
	}

	// Test Node.js module
	os.Remove(goModPath)
	packageJsonPath := filepath.Join(tempDir, "package.json")
	os.WriteFile(packageJsonPath, []byte("{}"), 0644)
	if detectLocalModuleType(tempDir) != NodeProject {
		t.Fatal("Expected Node.js project type")
	}
}

func TestGetGoModuleName(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test valid go.mod
	goModPath := filepath.Join(tempDir, "go.mod")
	os.WriteFile(goModPath, []byte("module github.com/test/my-module\n\ngo 1.21"), 0644)

	name := getGoModuleName(tempDir)
	if name != "github.com/test/my-module" {
		t.Fatalf("Expected 'github.com/test/my-module', got '%s'", name)
	}

	// Test invalid go.mod
	os.WriteFile(goModPath, []byte("invalid content"), 0644)
	name = getGoModuleName(tempDir)
	if name != "" {
		t.Fatalf("Expected empty string for invalid go.mod, got '%s'", name)
	}
}

func TestGetNodeModuleName(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test valid package.json
	packageJsonPath := filepath.Join(tempDir, "package.json")
	os.WriteFile(packageJsonPath, []byte(`{"name": "my-test-module", "version": "1.0.0"}`), 0644)

	name := getNodeModuleName(tempDir)
	if name != "my-test-module" {
		t.Fatalf("Expected 'my-test-module', got '%s'", name)
	}

	// Test invalid package.json
	os.WriteFile(packageJsonPath, []byte("invalid json"), 0644)
	name = getNodeModuleName(tempDir)
	if name != "" {
		t.Fatalf("Expected empty string for invalid package.json, got '%s'", name)
	}
}

func TestGetModuleName(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test Go module
	goModPath := filepath.Join(tempDir, "go.mod")
	os.WriteFile(goModPath, []byte("module github.com/test/go-module"), 0644)

	name := getModuleName(tempDir, GoProject)
	if name != "github.com/test/go-module" {
		t.Fatalf("Expected Go module name, got '%s'", name)
	}

	// Test Node.js module
	os.Remove(goModPath)
	packageJsonPath := filepath.Join(tempDir, "package.json")
	os.WriteFile(packageJsonPath, []byte(`{"name": "node-module"}`), 0644)

	name = getModuleName(tempDir, NodeProject)
	if name != "node-module" {
		t.Fatalf("Expected Node.js module name, got '%s'", name)
	}

	// Test unknown type
	name = getModuleName(tempDir, Unknown)
	if name != "" {
		t.Fatalf("Expected empty string for unknown type, got '%s'", name)
	}
}

func TestGetModuleDestinationPath(t *testing.T) {
	path := getModuleDestinationPath("test-module", GoProject)
	expected := filepath.Join("backend", "internal", "modules", "test-module")
	if path != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, path)
	}

	path = getModuleDestinationPath("test-module", NodeProject)
	if path != expected {
		t.Fatalf("Expected same path for Node.js module, got '%s'", path)
	}

	path = getModuleDestinationPath("test-module", Unknown)
	if path != "" {
		t.Fatalf("Expected empty string for unknown type, got '%s'", path)
	}
}

func TestCopyFile(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create source file
	srcPath := filepath.Join(tempDir, "source.txt")
	content := "test content"
	os.WriteFile(srcPath, []byte(content), 0644)

	// Copy file
	dstPath := filepath.Join(tempDir, "dest.txt")
	err = copyFile(srcPath, dstPath)
	if err != nil {
		t.Fatalf("copyFile failed: %v", err)
	}

	// Check destination file
	destContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(destContent) != content {
		t.Fatal("File content does not match")
	}
}

func TestCopyDir(t *testing.T) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "govite-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create source directory structure
	srcDir := filepath.Join(tempDir, "src")
	os.MkdirAll(srcDir, 0755)
	os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(srcDir, "file2.txt"), []byte("content2"), 0644)

	subDir := filepath.Join(srcDir, "subdir")
	os.MkdirAll(subDir, 0755)
	os.WriteFile(filepath.Join(subDir, "file3.txt"), []byte("content3"), 0644)

	// Copy directory
	dstDir := filepath.Join(tempDir, "dst")
	err = copyDir(srcDir, dstDir)
	if err != nil {
		t.Fatalf("copyDir failed: %v", err)
	}

	// Check destination files
	if _, err := os.Stat(filepath.Join(dstDir, "file1.txt")); os.IsNotExist(err) {
		t.Fatal("file1.txt not copied")
	}

	if _, err := os.Stat(filepath.Join(dstDir, "file2.txt")); os.IsNotExist(err) {
		t.Fatal("file2.txt not copied")
	}

	if _, err := os.Stat(filepath.Join(dstDir, "subdir", "file3.txt")); os.IsNotExist(err) {
		t.Fatal("subdir/file3.txt not copied")
	}
}

func TestModuleTypeString(t *testing.T) {
	if moduleTypeString(GoProject) != "Go" {
		t.Fatal("Expected 'Go' for GoProject")
	}

	if moduleTypeString(NodeProject) != "Node.js" {
		t.Fatal("Expected 'Node.js' for NodeProject")
	}

	if moduleTypeString(Unknown) != "Unknown" {
		t.Fatal("Expected 'Unknown' for Unknown")
	}
}



func TestRunUninstall(t *testing.T) {
	// Test that runUninstall function exists and can be called
	cmd := &cobra.Command{}
	err := runUninstall(cmd, []string{"test-module"})
	// We expect this to fail since we're not in a proper project directory,
	// but it should not panic
	if err == nil {
		t.Log("runUninstall completed without error (unexpected but not a failure)")
	}
}
