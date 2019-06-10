package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"git.edraj.io/mo/resonate/gitcluster"
	"github.com/fsnotify/fsnotify"
)

func main() {

	done := make(chan bool)
	var (
		fDir = flag.String("f1", ".", "specifies the folder1 to sync")
		sDir = flag.String("f2", ".", "specifies the folder2 to sync")
	)
	flag.Parse()

	var err error
	*fDir, err = filepath.Abs(*fDir)
	*sDir, err = filepath.Abs(*sDir)

	cluster, err := gitcluster.NewCluster(*fDir, *sDir)
	if err != nil {
		log.Fatal(fmt.Errorf("could not init git cluster: %v", err))
	}

	go WatchNode(*fDir, cluster)
	go WatchNode(*sDir, cluster)
	<-done
}

func WatchNode(url string, cluster *gitcluster.Cluster) {

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	defer watcher.Close()

	if err := watcher.Add(url); err != nil {
		log.Fatal(fmt.Errorf("could not watch %v: %v", url, err))
	}

	err = filepath.Walk(url, func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		if info.IsDir() {
			watcher.Add(path)
		}

		return nil
	})

	for {
		select {
		case event := <-watcher.Events:
			go cluster.PullFrom(url)
			log.Print(fmt.Sprintf("%v: %v", event.Op, filepath.Base(event.Name)))
		case <-watcher.Errors:
			fmt.Println("ERROR", url, err)
		}
	}
}
