package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "bufio"
)

func main(){
    if len(os.Args) != 4{
        fmt.Println("Provide the required number of argument")
        return
    }

    inputText := os.Args[2]
    word := strings.Split(inputText, `\n`)
    options := os.Args[1]

    if !strings.HasPrefix(options, "--align="){
        fmt.Println("Example: go run . --align=right something standard")
    }

    options = strings.TrimPrefix(options, "--align=")

    banner := os.Args[3]

    file, err := os.Open(banner +".txt")
    if err != nil {
        fmt.Println("Error opening file", err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var bannerString []string
    for{
        line, err:= reader.ReadString('\n')
        line = strings.TrimSuffix(line, "\n")
        bannerString = append(bannerString, line)
        if err!=nil{
            if err == io.EOF{
                break
            }
            fmt.Println("Error reading file", err)
        }
    }

    terminalWidth := 184

    var result []string
    for i, _ := range word{
        if word[i] == ""{
            continue
        }

        // looking for the total width of each word
        var totalWidth int
        splitWord := strings.Fields(word[i])
            // fmt.Printf("%q\n",splitWord)
            for k, _ := range splitWord{
                width := 0
                for _, char := range splitWord[k]{
                start := int(char-32) * 9 + 1
                width += len(bannerString[start])
            }
            totalWidth += width 
            }

        for j:=0; j < 8; j++{
            oneLineString := ""
            for _, char := range word[i]{
            start := int(char-32) * 9 + 1
            oneLineString += bannerString[start + j]
        }

        // fmt.Println(len(oneLineString))

        // var wordWidth string
        var space string
        var spaceWidth int
        var line string
        if options == "right"{
            spaceWidth = terminalWidth - len(oneLineString)
            space = strings.Repeat(" ", spaceWidth)
            line = space + oneLineString
            result = append(result, line)
        } else if options == "center"{
            spaceWidth = (terminalWidth - len(oneLineString))/2
            space = strings.Repeat(" ", spaceWidth)
            line = space + oneLineString + space
            result = append(result, line)
        } else if options == "justify"{
            spaceWidth = terminalWidth - totalWidth
            noOfGaps := len(splitWord) - 1
            distributedSpace := spaceWidth/noOfGaps
            space = strings.Repeat(" ", distributedSpace)

            for index,_ := range splitWord{
                var oneJustifyLine string
                for _, char := range splitWord[index]{
                    start:= int(char-32) * 9 + 1
                    oneJustifyLine += bannerString[start + j]
                }
                if index != len(splitWord) - 1{
                    line += oneJustifyLine + space
                } else{
                    line += oneJustifyLine
                }
            }
                result = append(result, line)
            }else{
            result = append(result, oneLineString)
            }
        }
    }


    
    for _, line := range result {
    fmt.Println(line)
}
}