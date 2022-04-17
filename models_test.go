package main

import (
	"os"
	"path/filepath"
	"testing"
)

const TEST_FOLDER string = "tmp_test"

var (
	checkTest = prefixedCheck("Test")
)

func TestUncategorizedFile(t *testing.T) {
	// Create temporal folder
	os.MkdirAll(TEST_FOLDER, 0755)
	// Create file to move
	fileName := "demo.notinlist"
	f, err := os.Create(filepath.Join(TEST_FOLDER, fileName))
	checkTest(err)
	f.Close()

	expected, _ := filepath.Abs(filepath.Join(TEST_FOLDER, "uncategorized", "notinlist", fileName))

	// Initialize folder
	folder := Folder{path: TEST_FOLDER, config: getConfig()}
	folder.findFiles()

	// Compare results
	operations := folder.findMoveOperations()

	os.RemoveAll(folder.path)

	if operations[0].target != expected {
		t.Fatalf("\nOperation target:\t%s\nExpected result:\t%s", operations[0].target, expected)
	}
}

func TestFindOperationsNoDuplicates(t *testing.T) {
	// Create temporal folder
	os.MkdirAll(TEST_FOLDER, 0755)
	// Create file to move
	fileName := "demo.json"
	f, err := os.Create(filepath.Join(TEST_FOLDER, fileName))
	checkTest(err)
	f.Close()

	expected, _ := filepath.Abs(filepath.Join(TEST_FOLDER, "data", "json", fileName))

	// Initialize folder
	folder := Folder{path: TEST_FOLDER, config: getConfig()}
	folder.findFiles()

	// Compare results
	operations := folder.findMoveOperations()

	os.RemoveAll(folder.path)

	if operations[0].target != expected {
		t.Fatalf("\nOperation target:\t%s\nExpected result:\t%s", operations[0].target, expected)
	}
}

func TestFindOperationsWithDuplicates(t *testing.T) {
	fileName := "demo.json"
	targetPath, _ := filepath.Abs(filepath.Join(TEST_FOLDER, "data", "json"))

	// Create temporal folder with initial file
	os.MkdirAll(targetPath, 0755)
	f, err := os.Create(filepath.Join(targetPath, fileName))
	checkTest(err)
	f.Close()
	// Create file to move
	f, err = os.Create(filepath.Join(TEST_FOLDER, fileName))
	checkTest(err)
	f.Close()

	// Initialize folder
	folder := Folder{path: TEST_FOLDER, config: getConfig()}
	folder.findFiles()
	expected := filepath.Join(targetPath, "demo-1.json")

	// Compare results
	operations := folder.findMoveOperations()

	os.RemoveAll(folder.path)

	if operations[0].target != expected {
		t.Fatalf("\nOperation target:\t%s\nExpected result:\t%s", operations[0].target, expected)
	}
}

func TestFindOperationsWithIgnoredFile(t *testing.T) {
	// Create temporal folder
	os.MkdirAll(TEST_FOLDER, 0755)
	// Create file to move
	fileName := ".DS_Store"
	f, err := os.Create(filepath.Join(TEST_FOLDER, fileName))
	checkTest(err)
	f.Close()

	// Initialize folder
	folder := Folder{path: TEST_FOLDER, config: getConfig()}
	folder.findFiles()

	// Compare results
	operations := folder.findMoveOperations()

	os.RemoveAll(folder.path)

	if len(operations) > 0 {
		t.Fatal(".DS_Store was moved")
	}
}
