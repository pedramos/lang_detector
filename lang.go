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
    
    //####################################
    //codigo para separar letras da string
    line_split:=strings.Split(line,"")
    fmt.Println(line_split[0]+line_split[1])
    
    
    //####################################
    //codigo para alucar slice com o tamanho certo
    if len(line_split)%3==0{
        trip:=make(string,len(line_split)/3)
    } else {
        trip:=make(string,(len(line_split)/3)+1)
    }
    
    
    //#####################################
    //codigo para agrupar strings em grupos de 3
    i:=0
    j:=0
    for i=0;i<len(trip);i++ {
        trip[i]=line_split[0+j]+line_split[1+j]+line_split[2+j]
        j=j+3
    }
    j=j-3
     
    if len(line_split)%3 != 0 {
        last_size:=len(line_split)%3
        i++
        switch last_size:
        case 1: trip[i] = line_split[j+1]+ " " + " "
        case 2: trip[i] = line_split[j+1]+ line_split[j+2] + " "
        case default : PrintLn("DEU MERDA")
    }
        
}
