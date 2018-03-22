package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func run() ([]string, error) {
    searchDir := "/home/huanke/3rdProject/eth-static-analysis"

    fileList := make([]string, 0)
    e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, path)
        return err
    })
    
    if e != nil {
        panic(e)
    }

    for _, file := range fileList {
        fmt.Println(file)
    }

    return fileList, nil
}

func main() {
    run()
}