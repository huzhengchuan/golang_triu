package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
	"math/rand"
	"time"
	"strconv"
    	"regexp"
	"bufio"
        "io"
	"strings"
)



func parseUrl(file string) map[string]string {
    
    var firstLevelMap map[string]string

    firstLevelMap = make(map[string]string)

    inputFile, inputError := os.Open(file)
    if inputError != nil {
        fmt.Println("An error occurred on opening the inputfile")
        return firstLevelMap
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

        //var hrefRegExp = regexp.MustCompile(".*<a class=\"research-area-pubs maia-button\" href=\"(.*?)\"(.*?)>.*?</a>.*?")
        var hrefRegExp = regexp.MustCompile(".*\"(/pubs/.*?)\".*")
        match := hrefRegExp.FindStringSubmatch(inputString)
        if match != nil {
            for i, v := range match {
                if i == 1 {
                    //fmt.Println("[", i, "]-", v)
                    url := "https://research.google.com" + v
                    urlSplit := strings.Split(url, "/")
                    firstLevelMap[strings.Split(urlSplit[len(urlSplit) - 1], ".")[0]] = url
                }
            }
        }
    }
    return firstLevelMap
}


func parseUrlTwoLevel(file string) []string {
    

    pdfSlice := make([]string, 0)

    inputFile, inputError := os.Open(file)
    if inputError != nil {
        fmt.Println("An error occurred on opening the inputfile")
        return pdfSlice
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

        var hrefRegExp = regexp.MustCompile(".*\"(.*?pdf)\".*?")
        match := hrefRegExp.FindStringSubmatch(inputString)
        if match != nil {
            for i, v := range match {
                if i == 1 {
                    if strings.HasPrefix(v, "http") {
                        pdfSlice = append(pdfSlice, v)
                    }else if strings.HasPrefix(v, "https") {
                        pdfSlice = append(pdfSlice, v)
                    }else if strings.HasPrefix(v, "www") {
                        pdfSlice = append(pdfSlice, v)
                    }else {
                        url := "https://research.google.com" + v
                        pdfSlice = append(pdfSlice, url)
                    }
                }
            }
        }
    }
    return pdfSlice
}

func getUrlBody(url string) string{
    
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }
    
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        fmt.Println(resp.StatusCode)
    }

    buf := make([]byte, 65535)
    
    randTool := rand.New(rand.NewSource(time.Now().UnixNano()))

    tmpUrlFile := "/tmp/" + strconv.Itoa(randTool.Intn(10000))
    f, err1 := os.OpenFile(tmpUrlFile, os.O_RDWR|os.O_CREATE|os.O_APPEND,os.ModePerm)
    if err1 != nil {
        panic(err1)
        return ""
    }
    
    defer f.Close()

    for {
        n, _ := resp.Body.Read(buf)
        if 0 == n {
            break
        }

        f.WriteString(string(buf[:n]))
    }
    return tmpUrlFile
}

func crawl(url string) {
    tmpFile := getUrlBody(url)
    fmt.Println(tmpFile)
    firstLevelMap := parseUrl(tmpFile)
    
    twoLevelTempFile := make([]string, 0)
    for _, v := range firstLevelMap {
    //    fmt.Println(k, v)
	twoLevelTempFile = append(twoLevelTempFile, getUrlBody(v))
    }  
	
    for _, tempFile := range twoLevelTempFile {
	fmt.Println(tempFile)
    	pdfSlice := parseUrlTwoLevel(tempFile)	
    	for k, v := range pdfSlice {
        	fmt.Println(k, v)
    	}
	os.Remove(tempFile)
    }
}

func main() {
    crawl("https://research.google.com/pubs/papers.html")
}


