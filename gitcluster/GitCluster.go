package GitCluster

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	object "gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Cluster struct {
	urls map[string]int

	i int
}

// var (
// 	fDir = flag.String("f1", ".", "specifies the folder1 to sync")
// 	sDir = flag.String("f2", ".", "specifies the folder2 to sync")
// )
// flag.Parse()

func NewCluster(fDir, sDir string) (*Cluster, error) {

	r, err := git.PlainInit(fDir, false)
	checkIfError(err)

	w, err := r.Worktree()
	checkIfError(err)

	err = ioutil.WriteFile(filepath.Join(fDir, "info"), []byte(time.Now().String()), 0644)
	checkIfError(err)

	_, err = w.Add("info")
	checkIfError(err)

	username, err := user.Current()
	checkIfError(err)

	host, err := os.Hostname()
	checkIfError(err)
	_, err = w.Commit("initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  username.Name,
			Email: fmt.Sprintf("%v@%v", username.Username, host),
			When:  time.Now(),
		},
	})
	checkIfError(err)

	_, err = git.PlainClone(sDir, false, &git.CloneOptions{
		URL: fDir,
	})
	checkIfError(err)

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{sDir},
	})
	checkIfError(err)

	err = r.CreateBranch(&config.Branch{
		Name:   "master",
		Remote: "origin",

		Merge: plumbing.NewBranchReferenceName("master"),
	})
	checkIfError(err)

	return &Cluster{urls: map[string]int{fDir: 1, sDir: 2}}, nil
}

func (c *Cluster) AddNode(dir string) error {

	if _, ok := c.urls[dir]; ok {
		return nil
	}

	initNode(dir, &c.urls)

	for url := range c.urls {
		r, err := git.PlainOpen(url)

		checkIfError(err)
		_, err = r.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: []string{dir},
		})

		checkIfError(err)
	}

	c.urls[dir] = c.i
	c.i++
	return nil
}

func initNode(dir string, remotes *map[string]int) {

	if len(*remotes) == 0 {
		r, err := git.PlainInit(dir, false)
		checkIfError(err)

		err = r.CreateBranch(&config.Branch{
			Name:   "master",
			Remote: "origin",

			Merge: plumbing.NewBranchReferenceName("master"),
		})
		checkIfError(err)
		return
	}

	n := "hello"

	for val := range *remotes {
		n = val
	}

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: n,
	})

	checkIfError(err)
}

func (c *Cluster) PullFrom(s string) {
	for url, _ := range c.urls {

		if url == s {
			continue
		}
		r, err := git.PlainOpen(url)
		checkIfError(err)

		w, err := r.Worktree()
		checkIfError(err)

		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		checkIfError(err)
	}
}

func checkIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
