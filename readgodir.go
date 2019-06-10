package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "bytes"
    "path/filepath"
)

func main() {
    // buffer to hold ~/go/src
    var pbuf bytes.Buffer
    // buffer to hold all directory names
    var buffer bytes.Buffer

    pbuf.WriteString(os.Getenv("GOPATH"))
    pbuf.WriteString("/src")
    // read names of all directories
    directories, err := ioutil.ReadDir(pbuf.String())
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range directories {
        // create path to search
        s := fmt.Sprintf("%s/%s", pbuf.String(), f.Name())
        // walk over every directory and subdirectory
        err := filepath.Walk(s, func(path string, info os.FileInfo, err error) error{
            if err != nil {
                fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
                return err
            }
            // skip .git directories
            if info.IsDir() && info.Name() != ".git" {
                buffer.WriteString(path + "\n")
            } else if info.Name() == ".git" {
                return filepath.SkipDir
            }
            return nil
        })
        if err != nil {
            fmt.Printf("error walking the path %q: %v\n", s, err)
            return
        }
    }

    // write to file
    /*f, err := os.OpenFile("directory_list.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }*/
    // write to stdout
    buffer.WriteTo(os.Stdout)
    /*if err != nil {
        log.Fatal(err)
    }*/

    return
}
