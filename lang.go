package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"



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
    //reader:=bufio.NewReader(f)
    scanner:=bufio.NewScanner(f)
    total_counter:=0.0
    var temp []string
    for scanner.Scan() {
        a:=scanner.Text()
        _=a
        temp=strings.Split(scanner.Text(),"\t")
        data[temp[0]],_=strconv.ParseFloat(temp[1],64)
        total_counter = data[temp[0]] + total_counter
    }
    for trigram := range data{
        data[trigram]=data[trigram]/total_counter
    }
    return data
}

func main(){
    trip:=get_trip()
    _=trip
    PT_data:="/home/rami/coolstuff/go/lang_detector/pt_trigram_count_pruned_100000.tsv"
    PT:=get_data(PT_data)
    _=PT
}
