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
    line_split:=strings.Split(line,"")
    fmt.Printf("%v",line)
    
    if err != nil {
        // You may check here if err == io.EOF
        fmt.Println("ERROR")
    }
    trip:= make([]string,len(line_split)-3)
    for i,_ := range line_split {
       fmt.Printf("%v->%q\n",i,line_split[i])
       if i==len(line_split)-3{
           break
       }
    trip[i]=string(line_split[i])+string(line_split[i+1])+string(line_split[i+2])
    fmt.Printf("%q %q %q\n",line_split[i],line_split[i+1],line_split[i+2])
    fmt.Println(trip[i])
    fmt.Println("______________")
        
    }
    
    
    
}
