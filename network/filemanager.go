package network

import (
	context "context"
	"log"
	"os"
	"path/filepath"
	"time"

	"git.nightcrickets.space/keefleoflimon/resonate/fuse"
	"git.nightcrickets.space/keefleoflimon/resonate/util"
	"github.com/pkg/errors"
)

var _ FileManagerServer = (*FileManager)(nil)

type FileManager struct {
	fs *fuse.FS
	lm *util.LockManager
}

func NewFileManager(fs *fuse.FS, lm *util.LockManager) *FileManager {
	return &FileManager{fs: fs, lm: lm}
}

func (fm *FileManager) Create(ctx context.Context, req *CreateRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	toBeCreated := filepath.Join(req.GetPath(), req.GetName())
	if !fm.lm.Lock(toBeCreated) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer fm.lm.Unlock(toBeCreated)

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	if child == nil {
		log.Print("could not find child")
	}

	_, err := child.FFNode.Create(req.GetName(), os.FileMode(req.GetMode()))
	if err != nil {
		log.Print(err)
	}
	log.Println("filemanager creating", filepath.Join(fm.fs.Location(), req.GetPath(), req.GetName()))
	return &Void{}, nil
}

func (fm *FileManager) Write(ctx context.Context, req *WriteRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	_, err := child.FFNode.Write(req.GetData(), req.GetOffset())

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Remove(ctx context.Context, req *RemoveRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	err := child.FFNode.Remove(req.GetName())

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Mkdir(ctx context.Context, req *MkdirRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	toBeCreated := filepath.Join(req.GetPath(), req.GetName())
	if !fm.lm.Lock(toBeCreated) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer fm.lm.Unlock(toBeCreated)

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	_, err := child.FFNode.Mkdir(req.GetName(), os.FileMode(req.GetMode()))

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Rename(ctx context.Context, req *RenameRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	toBeRenamed := filepath.Join(req.GetPath(), req.GetOldname())
	if !fm.lm.Lock(toBeRenamed) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeRenamed)
	}
	defer fm.lm.Unlock(toBeRenamed)

	toBeCreated := filepath.Join(req.GetPath(), req.GetNewname())
	if !fm.lm.Lock(toBeCreated) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer fm.lm.Unlock(toBeCreated)

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	newDir := node.(*fuse.File).Child(req.GetNewdirpath())

	err := child.FFNode.Rename(req.GetOldname(), req.GetNewname(), newDir.FFNode)

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Link(ctx context.Context, req *LinkRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	toBeLinked := filepath.Join(req.GetPath(), req.GetOld())
	if !fm.lm.Lock(toBeLinked) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeLinked)
	}
	defer fm.lm.Unlock(toBeLinked)

	toBeCreated := filepath.Join(req.GetPath(), req.GetNewname())
	if !fm.lm.Lock(toBeCreated) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer fm.lm.Unlock(toBeCreated)

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	old := node.(*fuse.File).Child(req.GetOld())
	_, err := child.FFNode.Link(req.GetNewname(), old.FFNode)

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Symlink(ctx context.Context, req *SymlinkRequest) (*Void, error) {
	if !fm.lm.Lock(req.GetPath()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", req.GetPath())
	}
	defer fm.lm.Unlock(req.GetPath())

	toBeLinked := filepath.Join(req.GetPath(), req.GetTarget())
	if !fm.lm.Lock(toBeLinked) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeLinked)
	}
	defer fm.lm.Unlock(toBeLinked)

	toBeCreated := filepath.Join(req.GetPath(), req.GetNewname())
	if !fm.lm.Lock(toBeCreated) {
		return &Void{}, errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer fm.lm.Unlock(toBeCreated)

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())
	_, err := child.FFNode.Symlink(req.GetTarget(), req.GetNewname())

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}

func (fm *FileManager) Setattr(ctx context.Context, req *SetattrRequest) (*Void, error) {

	node, _ := fm.fs.Root()
	child := node.(*fuse.File).Child(req.GetPath())

	if !fm.lm.Lock(child.FFNode.Path()) {
		return &Void{}, errors.Errorf("could not lock file (%v)", child.FFNode.Path())
	}
	defer fm.lm.Unlock(child.FFNode.Path())
	err := child.FFNode.Setattr(os.FileMode(req.GetMode()), time.Unix(req.GetAtime(), 0), time.Unix(req.GetMtime(), 0))

	if err != nil {
		log.Print(err)
	}

	return &Void{}, nil
}
