package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"./filetree"

	"log"

	"github.com/fsnotify/fsnotify"
)

const (
	BUFFERSIZE = 1024
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	var (
		aconns = make(map[net.Conn]int)

		done = make(chan bool)
		send = make(chan string)
	)

	go func() {
		i := 0
		for {
			conn, err := ln.Accept()
			if err != nil {
				continue
			}
			aconns[conn] = i
			i++
		}
	}()

	go listenSend(send, &aconns)

	filetree.SetRoot(".")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if info.IsDir() {
			watcher.Add(path)
		}
		return filetree.InsertNode(path)
	})

	if err != nil {
		fmt.Println("ERROR", err)
	}

	go func() {
		var (
			first string
			i     = 0
		)

		for {
			select {
			case event := <-watcher.Events:
				switch event.Op {
				case fsnotify.Write:
					s, err := ioutil.ReadFile(event.Name)
					if err != nil {
						log.Print(err)
						break
					}

					hash := sha256.Sum256(s)
					oldHash, err := filetree.GetNodeHash(event.Name)
					if err != nil {
						log.Print(err)
						break
					}

					if oldHash != hash {
						filetree.UpdateNodeHash(event.Name, hash)
						send <- event.Name
					}

				case fsnotify.Create:

					fileinfo, err := os.Stat(event.Name)
					if err == nil && fileinfo.IsDir() {
						if i == 1 {
							i = -1
							filetree.MoveNode(first, event.Name)
							fmt.Printf("moving: %v , to: %v\n", first, event.Name)
							break
						}
						watcher.Add(event.Name)
					}

					if err == nil && !fileinfo.IsDir() && i == 1 {
						i = 0
						filetree.MoveNode(first, event.Name)
						fmt.Printf("moving: %v , to: %v\n", first, event.Name)
						break
					}

					filetree.InsertNode(event.Name)
				case fsnotify.Remove:
					filetree.DeleteNode(event.Name)
					watcher.Remove(event.Name)
				case fsnotify.Rename:
					i++
					first = event.Name
				}

			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func listenSend(send <-chan string, aconns *map[net.Conn]int) {
	for {
		name := <-send
		fileInfo, err := os.Stat(name)
		if err != nil {
			log.Print(err)
			continue
		}

		fileSize := strconv.FormatInt(fileInfo.Size(), 10)
		fileName := name

		s, err := ioutil.ReadFile(name)
		if err != nil {
			log.Print(err)
			continue
		}

		go func() {
			for conn, _ := range *aconns {

				go func(conn net.Conn) {

					conn.Write([]byte("size: "))
					conn.Write([]byte(fileSize))
					conn.Write([]byte("\nname: "))
					conn.Write([]byte(fileName))
					conn.Write([]byte("\n"))

					buf := bytes.NewBuffer(s)
					sendBuffer := make([]byte, BUFFERSIZE)

					for {
						_, err := buf.Read(sendBuffer)
						if err == io.EOF {
							break
						}
						conn.Write(sendBuffer)
					}
				}(conn)
			}
		}()
	}
}
