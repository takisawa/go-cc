package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  var val = stdin.Text()

  fmt.Printf("\t.text\n\t.global _mymain\n_mymain:\n\tmov $%s, %%eax\n\tret\n", val)
}
