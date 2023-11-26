package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"bambang", "cahyo", "level2"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
