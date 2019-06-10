package gitcluster

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Cluster struct {
	mu   sync.Mutex
	urls map[string]int

	i int
}

func NewCluster(urls ...string) (*Cluster, error) {

	cluster := &Cluster{urls: make(map[string]int)}
	for _, url := range urls {
		if err := cluster.AddNode(url); err != nil {
			return nil, fmt.Errorf("could not add node %v to cluster: %v", url, err)
		}
	}

	for url := range cluster.urls {
		fmt.Println("added: ", url)
	}
	return cluster, nil
}

func (c *Cluster) AddNode(dir string) error {

	if _, ok := c.urls[dir]; ok {
		return nil
	}

	err := initNode(dir, &c.urls)
	if err != nil {
		return fmt.Errorf("could not add node to cluster: %v", err)
	}

	for url := range c.urls {
		r, err := git.PlainOpen(url)

		if err != nil {
			return fmt.Errorf("could not open repo %v: %v", url, err)
		}

		_, err = r.CreateRemote(&config.RemoteConfig{
			Name: "origin",
			URLs: []string{dir},
		})

		if err != nil {
			return fmt.Errorf("could not create remote for repo %v: %v", url, err)
		}
	}

	c.urls[dir] = c.i
	c.i++
	return nil
}

func initNode(dir string, remotes *map[string]int) error {

	if len(*remotes) == 0 {
		r, err := git.PlainInit(dir, false)
		if err != nil {
			return fmt.Errorf("could not init repo %v: %v", dir, err)
		}

		err = r.CreateBranch(&config.Branch{
			Name:   "master",
			Remote: "origin",

			Merge: plumbing.NewBranchReferenceName("master"),
		})
		if err != nil {
			return fmt.Errorf("could not reate branch on repo %v: %v", dir, err)
		}

		err = ioutil.WriteFile(filepath.Join(dir, "info"), []byte(time.Now().String()), 0644)
		if err != nil {
			return fmt.Errorf("could write info file on repo %v: %v", dir, err)
		}

		w, err := r.Worktree()
		if err != nil {
			return fmt.Errorf("could not open worktree on repo %v: %v", dir, err)
		}

		user, _ := user.Current()
		host, _ := os.Hostname()

		_, err = w.Add("info")
		if err != nil {
			return fmt.Errorf("could not add info file to stash on repo %v: %v", dir, err)
		}
		_, err = w.Commit("initial commit", &git.CommitOptions{All: true, Author: &object.Signature{
			Name:  user.Name,
			Email: fmt.Sprintf("%v@%v", user.Username, host),
			When:  time.Now(),
		}})
		if err != nil {
			return fmt.Errorf("could not commit initial commit on repo %v: %v", dir, err)
		}

		return nil
	}

	n := "hello"

	for val := range *remotes {
		n = val
	}

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: n,
	})

	if err != nil {
		return fmt.Errorf("could not clone repo %v to %v: %v", n, dir, err)

	}

	return nil
}

func (c *Cluster) PullFrom(s string) error {

	log.Print("attempting sync")

	c.mu.Lock()
	defer c.mu.Unlock()
	defer log.Print("finished sync")

	r, err := git.PlainOpen(s)
	if err != nil {

		return fmt.Errorf("could not open repo %v: %v", s, err)
	}

	w, err := r.Worktree()
	if err != nil {

		return fmt.Errorf("could not get working tree from repo %v: %v", s, err)
	}

	st, err := w.Status()

	if err == io.EOF {

		return fmt.Errorf("could not retrieve status of working tree from repo %v: %v", s, err)
	}

	if err != nil {

		return fmt.Errorf("could not retrieve status of working tree from repo %v: %v", s, err)
	}

	if st.IsClean() {

		return nil
	}

	err = w.AddGlob(".")
	if err == io.EOF {

		return fmt.Errorf("SOFT could not add files for commit in repo %v: %v", s, err)
	}

	if err != nil {

		return fmt.Errorf("could not add files for commit in repo %v: %v", s, err)
	}

	user, _ := user.Current()
	host, _ := os.Hostname()

	_, err = w.Commit("initial commit", &git.CommitOptions{All: true, Author: &object.Signature{
		Name:  user.Name,
		Email: fmt.Sprintf("%v@%v", user.Username, host),
		When:  time.Now(),
	}})

	if err != nil {

		return fmt.Errorf("could not commit changes in repo %v: %v", s, err)
	}

	for url, _ := range c.urls {

		if url == s {
			continue
		}
		tr, err := git.PlainOpen(url)
		if err != nil {

			return fmt.Errorf("could not open repo %v: %v", url, err)
		}

		tw, err := tr.Worktree()
		if err != nil {

			return fmt.Errorf("could not open worktree on repo %v: %v", url, err)
		}

		err = tw.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			return fmt.Errorf("could not perfom pull on repo %v from %v: %v", url, s, err)
		}
	}

	return nil
}
