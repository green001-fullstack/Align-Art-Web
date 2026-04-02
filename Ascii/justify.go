package Ascii

import (
    "fmt"
    "io"
    "os"
    "strings"
    "bufio"
)

func Justify(text string, align string, banner string) (string, error){
    
        
    word := strings.Split(text, `\n`)
    file, err := os.Open("Ascii/banners/" + banner + ".txt")
    if err != nil {
        return "", fmt.Errorf("error opening file")
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var bannerString []string
    for{
        line, err:= reader.ReadString('\n')
        line = strings.TrimRight(line, "\r\n")
        bannerString = append(bannerString, line)
        if err!=nil{
            if err == io.EOF{
                break
            }
            return "", fmt.Errorf("Error reading file")
        }
    }

    terminalWidth := 160

    var result strings.Builder
    for i, _ := range word{
        if word[i] == ""{
            continue
        }

        // looking for the total width of each word
        var totalWidth int
        splitWord := strings.Fields(word[i])
            for k, _ := range splitWord{
                width := 0
                for _, char := range splitWord[k]{
                start := int(char-32) * 9 + 1
                width += len(bannerString[start])
            }
            totalWidth += width 
            }

        for j:=0; j < 8; j++{
            var line string
            oneLineString := ""
            for _, char := range word[i]{
            start := int(char-32) * 9 + 1
            oneLineString += bannerString[start + j]
        }

        // fmt.Println(len(oneLineString))

        // var wordWidth string
        var space string
        var spaceWidth int
        if align == "right"{
            spaceWidth = terminalWidth - len(oneLineString)
            space = strings.Repeat(" ", spaceWidth)
            line = space + oneLineString
            result.WriteString(line + "\n") 
        } else if align == "center"{
            spaceWidth = (terminalWidth - len(oneLineString))/2
            space = strings.Repeat(" ", spaceWidth)
            line = space + oneLineString + space
            result.WriteString(line + "\n")
        } else if align == "justify"{
            spaceWidth = terminalWidth - totalWidth
            noOfGaps := len(splitWord) - 1
            if noOfGaps <= 0{
                result.WriteString(oneLineString + "\n")
                continue
            }
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
                result.WriteString(line + "\n")
            }else{
            result.WriteString(oneLineString + "\n")
            }
        }
    }
    final := result.String() 
    return final, nil
}

