package network

import (
	context "context"
	"log"
	"os"
	"path/filepath"

	rfs "git.iyi.cz/mo/resonatefuse"
)

type FileManager struct {
	fs *rfs.FS
}

func NewFileManager(fs *rfs.FS) *FileManager {
	return &FileManager{fs: fs}
}

func (fm *FileManager) Create(ctx context.Context, req *CreateRequest) (*Void, error) {
	node, _ := fm.fs.Root()
	child := node.(*rfs.File).Child(req.GetPath())
	if child == nil {
		log.Fatal("what da bell man")
	}
	_, err := child.FFNode.Create(req.GetName(), os.FileMode(req.GetMode()))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("filemanager creating", filepath.Join(fm.fs.Location(), req.GetPath(), req.GetName()))
	log.Println("filemanager:", rfs.Touch(filepath.Join(fm.fs.Location(), req.GetPath(), req.GetName()), os.FileMode(req.GetMode())))
	return &Void{}, nil
}

func (fm *FileManager) Write(ctx context.Context, req *WriteRequest) (*Void, error) {
	node, _ := fm.fs.Root()
	child := node.(*rfs.File).Child(req.GetPath())
	_, err := child.FFNode.Write(req.GetData(), req.GetOffset())

	if err != nil {
		log.Fatal(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Remove(ctx context.Context, req *RemoveRequest) (*Void, error) {
	node, _ := fm.fs.Root()
	child := node.(*rfs.File).Child(req.GetPath())
	err := child.FFNode.Remove(req.GetName())

	if err != nil {
		log.Fatal(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Mkdir(ctx context.Context, req *MkdirRequest) (*Void, error) {
	node, _ := fm.fs.Root()
	child := node.(*rfs.File).Child(req.GetPath())
	_, err := child.FFNode.Mkdir(req.GetName(), os.FileMode(req.GetMode()))

	if err != nil {
		log.Fatal(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Rename(ctx context.Context, req *RenameRequest) (*Void, error) {
	node, _ := fm.fs.Root()
	child := node.(*rfs.File).Child(req.GetPath())
	newDir := node.(*rfs.File).Child(req.GetNewdirpath())
	err := child.FFNode.Rename(req.GetOldname(), req.GetNewname(), newDir.FFNode)

	if err != nil {
		log.Fatal(err)
	}

	return &Void{}, nil
}
