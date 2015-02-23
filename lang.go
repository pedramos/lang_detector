package main

import "fmt"
import "bufio"
import "os"
import "strings"



//FUNCAO QUE SEPARA O INPUT EM TRIPLETES
func get_trip () []string{
    fmt.Println("Insert text")
    fmt.Print("=>")
    
    //reader é estrutura que le do standart input
    reader := bufio.NewReader(os.Stdin)
    line, err := reader.ReadString('\n')
    line_split:=strings.Split(line,"")
    //fmt.Printf("%v",line)
    
    if err != nil {
        // You may check here if err == io.EOF
        fmt.Println("ERROR")
    }
    trip:= make([]string,len(line_split)-3)
    for i,_ := range line_split {
       //fmt.Printf("%v->%q\n",i,line_split[i])
       if i==len(line_split)-3{
           break
        }
        trip[i]=string(line_split[i])+string(line_split[i+1])+string(line_split[i+2])
        //fmt.Printf("%q %q %q\n",line_split[i],line_split[i+1],line_split[i+2])
        fmt.Println(trip[i])
        //fmt.Println("______________")
    }
    return trip    
}

func get_data(file string) map[string]float64{
    var data map[string]float64
    data=make(map[string]float64)
    f,err:=os.Open(file)
    if err!=nil{
        panic(err)
    }
    r:=bufio.NewReader(f)
    buff:=make([]byte, 1024)
    for {
        n,err:=f.ReadLine(buff)
        
    }
    return data
}

func main(){
    trip:=get_trip()
    
    
}