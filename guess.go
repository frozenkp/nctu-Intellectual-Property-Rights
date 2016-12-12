package main

import(
  "fmt"
  "os"
  "io"
  "os/signal"
  "bufio"
  "strings"
  "github.com/fatih/color"
)

func main(){
  //color
  red:=color.New(color.FgHiRed)
  green:=color.New(color.FgHiGreen)
  //get data
  f,_:=os.OpenFile("data",os.O_RDONLY,0777)
  fin:=bufio.NewReader(f)
  var s string
  var a int
  data:=make(map[string]int)
  for {
    i,_,err:=fin.ReadRune();
    a=int(i)-'0'
    sin,_:=fin.ReadSlice('\n')
    s=strings.TrimLeft(strings.TrimRight(string(sin),"\r\n")," ")
    if err==io.EOF{
      break;
    }
    data[s]=a
  }
  f.Close()
  f,_=os.OpenFile("data",os.O_WRONLY,0777)
  defer f.Close()

  //ctrl+c
  c:=make(chan os.Signal,1)
  signal.Notify(c,os.Interrupt)
  go func(){
    for sig:=range c {
      red.Printf("received ctrl+c(%v)\n",sig)
      store(data,f)
      os.Exit(0)
    }
  }()

  //process
  stdin:=bufio.NewReader(os.Stdin)
  var mode string
  for true {
    fmt.Scanf("%s",&mode)
    if mode=="find"{
      sin,_:=stdin.ReadSlice('\n')
      s=strings.TrimLeft(strings.TrimRight(string(sin),"\r\n")," ")
      if data[s]==0 {
        fmt.Println("no data")
      }else{
        if data[s]==1 {
          green.Println("答案：是")
        }else{
          red.Println("答案：否")
        }
      }
    }else if mode=="revise" {
      fmt.Scanf("%d",&a)
      sin,_:=stdin.ReadSlice('\n')
      s=strings.TrimLeft(strings.TrimRight(string(sin),"\r\n")," ")
      data[s]=a
      color.Set(color.FgHiMagenta)
      fmt.Println("revise finish")
      color.Unset()
    }else if mode=="end" {
      store(data,f)
      break
    }
  }
}

func store(data map[string]int,f *os.File){
  color.Set(color.FgHiMagenta)
  fmt.Printf("Storing data...")
  for name,ans:=range data{
    fmt.Fprintf(f,"%d %s\n",ans,name)
  }
  fmt.Printf("finished\n")
  color.Unset()
}

