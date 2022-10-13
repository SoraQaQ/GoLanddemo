package main 
import(
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File,m map[string]map[string]int){
	text := bufio.NewScanner(f)
	for text.Scan(){
		if m[text.Text()] == nil {
			m[text.Text()] = make(map[string]int)
		}
		m[text.Text()][f.Name()]++
	}
	f.Close()
}

func main() {
	linesMap := make(map[string]map[string]int)
	if len(os.Args) == 1 {
		countLines(os.Stdin, linesMap)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println("open file failed")
			}
			countLines(f, linesMap)
		}
	}
	//print result
	for line, fileName := range linesMap {
		if len(fileName) > 1 {
			fmt.Printf("context: %s\ntimes: %d\nfiles:", line, len(fileName))
			for k, _ := range fileName {
				fmt.Print(k + " ")
			}
			fmt.Printf("\n\n")
		} else {
			for k, v := range fileName {
				if v > 1 {
					fmt.Printf("%s\ntimes:%d\nfiles:%s\n\n", line, v, k)
				}
			}
		}
	}
}
