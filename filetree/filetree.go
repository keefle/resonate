package filetree

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type node struct {
	name        string
	parent      *node
	childrenMap map[string]*node

	hash [sha256.Size]byte
}

var (
	root *node
)

func SetRoot(name string) {
	root = &node{name: name, childrenMap: make(map[string]*node)}
}

func (n *node) AddChild(name string) {
	n.childrenMap[name] = &node{
		name:        name,
		parent:      n,
		childrenMap: make(map[string]*node),
	}
}

func (n *node) String() string {
	s := make([]string, 0)
	for n != nil {
		s = append(s, n.name)
		n = n.parent
	}

	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}

	return filepath.Join(s...)
}

func Walk() {
	walk(root, 0)
}

func walk(n *node, d int) {
	fmt.Printf("%*s%v\n", d, " ", n)
	for _, n := range n.childrenMap {
		walk(n, d+4)
	}
}

func CutPath(path string) []string {
	f := func(c rune) bool {
		return (c == os.PathSeparator)
	}
	return strings.FieldsFunc(path, f)
}

func getNode(path string) *node {
	n := get(root, path)
	if n == nil {
		return nil
	}

	return n
}

func get(r *node, path string) *node {
	paths := CutPath(path)

	if len(paths) == 0 {
		return nil
	}

	if paths[0] == root.name {
		paths = paths[1:]
	}

	n := r
	for _, v := range paths {
		if _, ok := n.childrenMap[v]; !ok {
			return nil
		}
		n = n.childrenMap[v]
	}

	return n
}

func InsertNode(path string) error {
	paths := CutPath(path)
	if len(paths) == 0 {
		return fmt.Errorf("unable to insert node: %v", path)
	}

	if paths[0] == root.name {
		paths = paths[1:]
	}

	cur := root
	for _, v := range paths {

		if cur.childrenMap[v] == nil || (len(cur.childrenMap[v].hash) == 0 && len(cur.childrenMap[v].childrenMap) == 0) {
			cur.childrenMap[v] = &node{name: v, parent: cur, childrenMap: make(map[string]*node)}
		}

		cur = cur.childrenMap[v]
	}

	return nil
}

func DeleteNode(path string) error {
	n := getNode(path)
	if n == nil {
		return fmt.Errorf("cannot delete %v, does not exist", path)
	}

	if n.parent == nil {
		n.childrenMap = make(map[string]*node)
		n.name = ""

		return nil
	}

	delete(n.parent.childrenMap, n.name)

	return nil
}

func MoveNode(path, target string) error {
	// cases: dir to another dir
	// cases: file to another dir
	p := getNode(path)
	t := getNode(target)
	if p != nil && t != nil {
		delete(p.parent.childrenMap, p.name)
		p.parent = t
		t.childrenMap[p.name] = p
		return nil
	}

	// cases: dir to another new dir inside a dir
	// cases: file to another new file inside a dir
	tp := CutPath(target)
	tn := getNode(filepath.Join(tp[:len(tp)-1]...))
	if p != nil && tn != nil {
		delete(p.parent.childrenMap, p.name)
		p.name = filepath.Base(filepath.Join(tp...))
		p.parent = tn
		tn.childrenMap[p.name] = p
		return nil
	}

	// cases: dir to new dir (rename)
	// cases: file to new file (rename)
	if p != nil && t == nil {
		delete(p.parent.childrenMap, p.name)
		p.name = filepath.Base(target)
		p.parent.childrenMap[p.name] = p
		return nil
	}

	return nil
}

func UpdateNodeName(path, name string) error {
	n := getNode(path)
	if n == nil {
		return fmt.Errorf("node with path %v was not found", path)
	}
	tmp := n.name
	n.name = name

	if n.parent != nil {
		delete(n.parent.childrenMap, tmp)
		n.parent.childrenMap[name] = n
	}

	return nil
}

func UpdateNodeHash(path string, hash [sha256.Size]byte) error {
	n := getNode(path)
	if n == nil {
		return fmt.Errorf("node with path %v was not found", path)
	}
	n.hash = hash

	return nil
}

func GetNodeHash(path string) ([sha256.Size]byte, error) {
	n := getNode(path)
	if n == nil {
		return [sha256.Size]byte{0}, fmt.Errorf("node with path %v was not found", path)
	}

	return n.hash, nil
}
