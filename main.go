package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func Print(msg string) {
	fmt.Println(msg)
}
func Say(msg string) {
	fmt.Println(msg)
}
func Random(per int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(per)
}
func ArrStr(arr []string) string {
	return strings.Join(arr, "")
}
func ArrStrSpace(arr []string) string {
	return strings.Join(arr, " ")
}
func ConncetDb(name string) {
	db, err := sql.Open("sqlite3", "./"+name)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
