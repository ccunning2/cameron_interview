package main

import (
"fmt"
"os"
sigar "github.com/cloudfoundry/gosigar" //Not sure if I'll lose or gain points by using this library
"time"
"bytes"
"os/exec"
"strconv"
"strings"
"log"
"bufio"
)

type Process struct {
    cpu float64
}


//Run 'ps aux' and sum up the cpu % from all processes displayed
func getCpuUse() float64 {
    cmd := exec.Command("ps", "aux")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    processes := make([]*Process, 0)
    for {
        line, err := out.ReadString('\n')
        if err!=nil {
            break;
        }
        tokens := strings.Split(line, " ")
        ft := make([]string, 0)
        for _, t := range(tokens) {
            if t!="" && t!="\t" {
                ft = append(ft, t)
            }
        }
       
        cpu, err := strconv.ParseFloat(ft[2], 64)
        if err!=nil {
            continue
        }
        processes = append(processes, &Process{cpu})
    }
    var total float64
    for _, p := range(processes) {
        total += p.cpu
    }
  
    return total
}






func main() {

	reader := bufio.NewReader(os.Stdin)
	enter := make(chan bool)

    //Need this to establish means of exiting upon hitting 'ENTER'
	go func() {
		reader.ReadString('\n')
		enter <- true
	}()



loop:

	for {

    //The sigar 'library' worked great for getting the memory information, but getting the CPU information was not as straightforward.
	mem := sigar.Mem{}


	mem.Get()
	
	
    usedMem := mem.ActualUsed
    freeMem := mem.ActualFree
    memPct := (float64(mem.ActualUsed)/float64(mem.Total)) * 100



    fmt.Fprintf(os.Stdout, "\rUsing %.0f%% of memory--- Used Memory: %d bytes    Available Memory: %d bytes ",memPct, usedMem, freeMem )
    fmt.Fprintf(os.Stdout, "Cpu Usage:%.2f%%", getCpuUse() )
    time.Sleep(500 * time.Millisecond)
    select {
    case <-enter:
    	break loop
    default:
    }
}
//}

}