package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

// Compress a single file to .gz format
func compressFile(filePath string, wg *sync.WaitGroup) {
    defer wg.Done()

    // Open source file
    inputFile, err := os.Open(filePath)
    if err != nil {
        fmt.Printf("Failed to open file %s: %v\n", filePath, err)
        return
    }
    defer inputFile.Close()

    // Create output file with .gz extension
    outputFilePath := filePath + ".gz"
    outputFile, err := os.Create(outputFilePath)
    if err != nil {
        fmt.Printf("Failed to create file %s: %v\n", outputFilePath, err)
        return
    }
    defer outputFile.Close()

    // Create a gzip.Writer
    gzipWriter := gzip.NewWriter(outputFile)
    defer gzipWriter.Close()

    // Set the original file name inside the gzip metadata
    gzipWriter.Name = filepath.Base(filePath)

    // Copy content from input file to gzip writer
    _, err = io.Copy(gzipWriter, inputFile)
    if err != nil {
        fmt.Printf("Failed to compress file %s: %v\n", filePath, err)
        return
    }

    fmt.Printf("Compressed %s → %s\n", filePath, outputFilePath)
}

func main() {
    // List of files to compress (replace with your own)
    files := []string{
        "example1.txt",
        "example2.txt",
        "example3.txt",
    }

    var wg sync.WaitGroup

    // Start a goroutine for each file
    for _, file := range files {
        wg.Add(1)
        go compressFile(file, &wg)
    }

    // Wait for all compressions to complete
    wg.Wait()

    fmt.Println("All files have been compressed.")
}
