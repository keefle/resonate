package filetree

import (
	"path/filepath"
	"testing"
)

func TestCutPath(t *testing.T) {
	tt := []struct {
		path string
		out  []string
	}{
		{
			path: "/home/user/directory/",
			out:  []string{"home", "user", "directory"},
		},
		{
			path: "/home/user/directory/",
			out:  []string{"home", "user", "directory"},
		},
		{
			path: "home/user/directory",
			out:  []string{"home", "user", "directory"},
		},
		{
			path: "123/456",
			out:  []string{"123", "456"},
		},
		{
			path: "/123/",
			out:  []string{"123"},
		},
		{
			path: "/123",
			out:  []string{"123"},
		},
		{
			path: "123",
			out:  []string{"123"},
		},
		{
			path: "",
			out:  []string{},
		},
		{
			path: " ",
			out:  []string{" "},
		},
		{
			path: "/",
			out:  []string{},
		},
		{
			path: "//",
			out:  []string{},
		},
		{
			path: "///",
			out:  []string{},
		},
		{
			path: "/ /",
			out:  []string{" "},
		},
		{
			path: "./",
			out:  []string{"."},
		},
	}

	for _, tc := range tt {
		out := CutPath(tc.path)
		if (out == nil) != (tc.out == nil) {
			t.Errorf("CutPath(\"%v\") = %v; want %v", tc.path, out, tc.out)
			continue
		}

		if len(out) != len(tc.out) {
			t.Errorf("CutPath(\"%v\") = %v; want %v", tc.path, out, tc.out)
			continue
		}
		for i, v := range out {
			if v != tc.out[i] {
				t.Errorf("CutPath(\"%v\") = %v; want %v", tc.path, out, tc.out)
			}
		}
	}
}

func TestGetNode(t *testing.T) {

	root = &node{
		name:        ".",
		childrenMap: make(map[string]*node),
		parent:      nil,
	}

	root.childrenMap["f1"] = &node{name: "f1", parent: root}
	root.childrenMap["f2"] = &node{name: "f2", parent: root}
	root.childrenMap["f3"] = &node{name: "f3", childrenMap: make(map[string]*node), parent: root}
	root.childrenMap["f3"].childrenMap["f4"] = &node{name: "f4", parent: root.childrenMap["f3"]}
	root.childrenMap["f3"].childrenMap["f5"] = &node{name: "f5", parent: root.childrenMap["f3"]}

	tt := []struct {
		path string
		node *node
	}{
		{path: "/", node: nil},
		{path: "//", node: nil},
		{path: "f1/", node: root.childrenMap["f1"]},
		{path: "/f1", node: root.childrenMap["f1"]},
		{path: "/f1/", node: root.childrenMap["f1"]},
		{path: "/f2", node: root.childrenMap["f2"]},
		{path: "/f3", node: root.childrenMap["f3"]},
		{path: "/f3/f4", node: root.childrenMap["f3"].childrenMap["f4"]},
		{path: "/f3/f5", node: root.childrenMap["f3"].childrenMap["f5"]},
		{path: "/f3/f6", node: nil},
		{path: " ", node: nil},
		{path: "", node: nil},
	}

	for _, tc := range tt {
		n := getNode(tc.path)
		if n != tc.node {
			t.Errorf("GetNode(\"%v\") = %v; want %v", tc.path, n, tc.node)
		}
	}
}

func TestInsertNode(t *testing.T) {

	root = &node{name: ".", childrenMap: make(map[string]*node)}

	tt := []struct {
		path  string
		added []string
	}{
		{path: "filetree/filetree.go", added: []string{"filetree", "filetree/filetree.go"}},
		{path: "moved", added: []string{"moved"}},
		{path: "readme.md", added: []string{"readme.md"}},
		{path: "home/user/happy/wow/213/djas", added: []string{"home", "home/user", "home/user/happy", "home/user/happy/wow", "home/user/happy/wow/213", "home/user/happy/wow/213/djas"}},
		{path: ".", added: []string{"."}},
	}

	for _, tc := range tt {
		err := InsertNode(tc.path)
		if err != nil {
			t.Errorf("%v", err)
		}
		for _, node := range tc.added {
			if getNode(node) == nil {
				t.Errorf("DeleteNode(%v) did not delete %v", tc.path, node)
			}
		}
		root = &node{name: ".", childrenMap: make(map[string]*node)}
	}
}

func TestDeleteNode(t *testing.T) {
	tree := []string{
		"filetree",
		"filetree/filetree.go",
		"filetree/filetree_test.go",
		"main.go",
		"moved",
		"dir",
		"readme.md",
		"tags",
		"todo.txt",
		"watchit",
		"watchit.go",
		".",
	}

	root = &node{name: ".", childrenMap: make(map[string]*node)}

	for _, v := range tree {
		err := InsertNode(v)
		if err != nil {
			t.Errorf("%v", err)
		}
	}

	tt := []struct {
		path    string
		deleted []string
	}{
		{path: "filetree", deleted: []string{"filetree", "filetree/filetree.go", "filetree/filetree_test.go"}},
		{path: "moved", deleted: []string{"moved"}},
		{path: ".", deleted: tree},
		{path: "readme.md", deleted: []string{"readme.md"}},
	}

	for _, tc := range tt {
		err := DeleteNode(tc.path)
		if err != nil {
			t.Errorf("%v", err)
		}
		for _, node := range tc.deleted {
			if getNode(node) != nil {
				t.Errorf("DeleteNode(%v) did not delete %v", tc.path, node)
			}
		}
		root = &node{name: ".", childrenMap: make(map[string]*node)}

		for _, v := range tree {
			err := InsertNode(v)
			if err != nil {
				t.Errorf("%v", err)
			}
		}
	}
}

func TestMoveNode(t *testing.T) {
	tree := []string{
		"filetree",
		"filetree/filetree.go",
		"filetree/filetree_test.go",
		"main.go",
		"moved",
		"dir",
		"readme.md",
		"test",
		"tags",
		"todo.txt",
		"watchit",
		"hip/hop",
		"watchit.go",
		".",
	}

	tt := []struct {
		path   string
		target string
		result string
	}{
		{path: "./filetree", target: "./dir", result: "./dir/filetree"},
		{path: "./test", target: "./dir", result: "./dir/test"},
		{path: "./moved", target: "./dir", result: "./dir/moved"},
		{path: "./dir", target: "./filetree/cool", result: "./filetree/cool"},
		{path: "./filetree/filetree.go", target: ".", result: "./filetree.go"},
		{path: "./readme.md", target: "./filetree", result: "./filetree/readme.md"},
		{path: "./tags", target: "./ctags", result: "./ctags"},
		{path: "./hip/hop", target: "./hop", result: "./hop"},
	}

	for _, tc := range tt {
		root = &node{name: ".", childrenMap: make(map[string]*node)}

		for _, v := range tree {
			err := InsertNode(v)
			if err != nil {
				t.Errorf("%v", err)
			}
		}

		err := MoveNode(tc.path, tc.target)
		if err != nil {
			t.Errorf("%v", err)
		}

		if getNode(tc.path) != nil && tc.path != tc.result {
			t.Errorf("MoveNode(%v, %v) did not remove %v", tc.path, tc.target, tc.path)
		}
		if getNode(tc.result) == nil || getNode(tc.result).name != filepath.Base(tc.result) {
			t.Errorf("MoveNode(%v, %v) did not create %v", tc.path, tc.target, tc.result)
		}
	}

}
