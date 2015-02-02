package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main(){
    fmt.Println("Insert text")
    fmt.Print("=>")
    
    //reader é estrutura que le do standart input
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    
    if err != nil {
        // You may check here if err == io.EOF
        fmt.Println("ERROR")
    }
    line_split:=strings.Split(line,"")
    fmt.Println(line_split[0]+line_split[1])
    
}