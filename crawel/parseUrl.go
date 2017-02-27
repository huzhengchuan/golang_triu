package main

import (
    "fmt"
    "regexp"
    "os"
//    "log"
    "bufio"
    "io"
)


func parseUrl(file string) {

    inputFile, inputError := os.Open(file)
    if inputError != nil {
        fmt.Println("An error occurred on opening the inputfile")
        return
    }
   
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    lineCounter := 0
    for {
    
        inputString, readerError := inputReader.ReadString('\n')
        if readerError == io.EOF {
            break        
        }
        lineCounter++
        //fmt.Printf("%d:%s", lineCounter, inputString)

        var hrefRegExp = regexp.MustCompile(".*<a class=\"research-area-pubs maia-button\" href=\"(.*?)\"(.*?)>.*?</a>.*?")
        match := hrefRegExp.FindStringSubmatch(inputString)
        if match != nil {
            for i, v := range match {
                if i == 1 {
                    fmt.Println("[", i, "]-", v)
                    }
            }
        }
    }
}

func main() {
    file := "./Research_at_Google.html"
    parseUrl(file)

}
