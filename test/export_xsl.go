package main

import (
	"encoding/csv"
	"os"
)

func export() error {
	//dir, err := os.Getwd()
	//if err != nil {
	//	return fmt.Errorf("os.Getwd err: %v", err)
	//}

	f, err := os.Create("./test/test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}

	w.WriteAll(data)
	return nil
}

func main() {
	export()
}
