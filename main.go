package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/madxiii/dream_tz/jobs"
	"github.com/madxiii/dream_tz/models"
)

const filename = "values.json"

func main() {
	if err := run(); err != nil {
		log.Fatalf("run -> %v\n", err)
	}
}

func run() error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("os.Open file: %s, err: %w", filename, err)
	}
	defer file.Close()

	bb, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("io.ReadAll err: %w", err)
	}

	var vals []models.Values

	if err := json.Unmarshal(bb, &vals); err != nil {
		return fmt.Errorf("json.Unmarshal err: %w", err)
	}

	result := jobs.GetSum(vals, 3)

	fmt.Println("Result:", result)
	return nil
}
