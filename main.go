package main

import (
  "fmt"
  "time")

type AWSRegion struct {
   Name string
   Timezone string
}

now := time.Now()

fmt.Println("=============================")
fmt.Println("Current UTC Time:", now.Format("2006-01-02 15:04:05 MST"))
