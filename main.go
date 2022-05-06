package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const notes = "notes.txt"

func main() {
	first := firstStart()
	if first == -1 {
		entryFirstNotes()
	}

	create := flag.Bool("c", false, "creat")
	read := flag.Bool("r", false, "read")
	update := flag.Bool("u", false, "update")
	deleted := flag.Bool("d", false, "delete")
	flag.Parse()

	switch {
	case *create:
		createNotes()
	case *update:
		updateNotes()
	case *read:
		readNotes()
	case *deleted:
		deleteNotes()
	}
}

func firstStart() int {
	file, err := os.Open(notes)
	if err != nil {
		//log.Println(err)
		return -1
	}
	defer file.Close()
	return 1
}

func entryFirstNotes() {
	text := "-> Hello. This is your first post creating automatically\n"
	file, err := os.Create("notes.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err = file.WriteString(text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

func createNotes() {
	file, err := os.OpenFile(notes, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("enter note text:")
	var b bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	b.WriteString(fmt.Sprintf("-> %s \n", s))

	if _, err = file.WriteString(string(b.Bytes())); err != nil {
		log.Fatal(err)
	}
	fmt.Println("create new notes successfully")
}

func openReedFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()

	readFile, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	return readFile
}

func readNotes() {
	fmt.Println("YOUR NOTES:")
	readFile := openReedFile(notes)
	fmt.Println(string(readFile))
}

func updateNotes() {
	readFile := openReedFile(notes)
	line := strings.Split(string(readFile), "\n")

	var index int
	fmt.Println("enter the note number to be update:")
	_, err := fmt.Scan(&index)
	if err != nil || index == 0 {
		log.Println("invalid value entered")
		return
	}
	fmt.Printf("your note number %d, text: %s \n", index, line[index-1])
	fmt.Println("enter update text:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	line[index-1] = s
	joinNotes := strings.Join(line, "\n")
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("-> %s", joinNotes))

	if err := ioutil.WriteFile(notes, b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Println("update notes successfully")
}

func deleteNotes() {
	readFile := openReedFile(notes)
	line := strings.Split(string(readFile), "\n")

	var index int
	fmt.Println("enter the note number to be deleted:")
	_, err := fmt.Scan(&index)
	if err != nil || index == 0 {
		log.Println("invalid value entered")
		return
	}

	newNotes := append(line[:index-1], line[index:]...)
	joinNotes := strings.Join(newNotes, "\n")
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%s", joinNotes))

	if err := ioutil.WriteFile(notes, b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete notes successfully")
}
