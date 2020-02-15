package main

import (
	"context"
	"log"
	"path/filepath"

	"git.nightcrickets.space/keefleoflimon/resonate/network"
	"git.nightcrickets.space/keefleoflimon/resonate/util"
	rfs "git.nightcrickets.space/keefleoflimon/resonatefuse"
	"github.com/pkg/errors"
)

type HookManager struct {
	lm     *util.LockManager
	client network.FileManagerClient
}

func NewHookManager(lm *util.LockManager, client network.FileManagerClient) *HookManager {
	// conn, err := grpc.Dial(*peer, grpc.WithInsecure())

	return &HookManager{lm: lm, client: client}
}

func (hm *HookManager) HooksToOptions() []rfs.Option {
	return []rfs.Option{
		rfs.CreateOption(hm.createHook),
		rfs.WriteOption(hm.writeHook),
		rfs.RemoveOption(hm.removeHook),
		rfs.MkdirOption(hm.mkdirHook),
		rfs.RenameOption(hm.renameHook),
		rfs.LinkOption(hm.linkHook),
		rfs.SymlinkOption(hm.symlinkHook),
		rfs.SetattrOption(hm.setattrHook),
	}
}

func createHook(req *rfs.CreateRequest, client network.FileManagerClient) error {
	log.Println("create hook running")

	_, err := client.Create(context.Background(), &network.CreateRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func (hm *HookManager) createHook(req *rfs.CreateRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeCreated := filepath.Join(req.Path, req.Name)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)

	return createHook(req, hm.client)
}

func writeHook(req *rfs.WriteRequest, client network.FileManagerClient) error {
	log.Println("write hook running")

	_, err := client.Write(context.Background(), &network.WriteRequest{Path: req.Path, Data: req.Data, Offset: req.Offset})
	log.Println(err)
	return err
}

func (hm *HookManager) writeHook(req *rfs.WriteRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)
	return writeHook(req, hm.client)
}

func removeHook(req *rfs.RemoveRequest, client network.FileManagerClient) error {
	log.Println("remove hook running")

	_, err := client.Remove(context.Background(), &network.RemoveRequest{Path: req.Path, Name: req.Name})
	log.Println(err)
	return err
}

func (hm *HookManager) removeHook(req *rfs.RemoveRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)
	return removeHook(req, hm.client)
}

func renameHook(req *rfs.RenameRequest, client network.FileManagerClient) error {
	log.Println("rename hook running")

	_, err := client.Rename(context.Background(), &network.RenameRequest{Path: req.Path, Oldname: req.OldName, Newname: req.NewName, Newdirpath: req.NewDir})
	log.Println(err)
	return err
}

func (hm *HookManager) renameHook(req *rfs.RenameRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeRenamed := filepath.Join(req.Path, req.OldName)
	if !hm.lm.Lock(toBeRenamed) {
		return errors.Errorf("could not lock file (%v)", toBeRenamed)
	}
	defer hm.lm.Unlock(toBeRenamed)

	toBeCreated := filepath.Join(req.Path, req.NewName)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)

	return renameHook(req, hm.client)
}

func mkdirHook(req *rfs.MkdirRequest, client network.FileManagerClient) error {
	log.Println("mkdir hook running")

	_, err := client.Mkdir(context.Background(), &network.MkdirRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func (hm *HookManager) mkdirHook(req *rfs.MkdirRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeCreated := filepath.Join(req.Path, req.Name)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)
	return mkdirHook(req, hm.client)
}

func linkHook(req *rfs.LinkRequest, client network.FileManagerClient) error {
	log.Println("link hook running")

	_, err := client.Link(context.Background(), &network.LinkRequest{Path: req.Path, Newname: req.NewName, Old: req.Old})
	log.Println(err)
	return err
}

func (hm *HookManager) linkHook(req *rfs.LinkRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeLinked := filepath.Join(req.Path, req.Old)
	if !hm.lm.Lock(toBeLinked) {
		return errors.Errorf("could not lock file (%v)", toBeLinked)
	}
	defer hm.lm.Unlock(toBeLinked)

	toBeCreated := filepath.Join(req.Path, req.NewName)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)

	return linkHook(req, hm.client)
}

func symlinkHook(req *rfs.SymlinkRequest, client network.FileManagerClient) error {
	log.Println("symlink hook running")

	_, err := client.Symlink(context.Background(), &network.SymlinkRequest{Path: req.Path, Newname: req.NewName, Target: req.Target})
	log.Println(err)
	return err
}

func (hm *HookManager) symlinkHook(req *rfs.SymlinkRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeLinked := filepath.Join(req.Path, req.Target)
	if !hm.lm.Lock(toBeLinked) {
		return errors.Errorf("could not lock file (%v)", toBeLinked)
	}
	defer hm.lm.Unlock(toBeLinked)

	toBeCreated := filepath.Join(req.Path, req.NewName)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)
	return symlinkHook(req, hm.client)
}

func setattrHook(req *rfs.SetattrRequest, client network.FileManagerClient) error {
	log.Println("setattr hook running")

	_, err := client.Setattr(context.Background(), &network.SetattrRequest{Path: req.Path, Atime: req.Atime.Unix(), Mtime: req.Mtime.Unix()})
	log.Println(err)
	return err
}

func (hm *HookManager) setattrHook(req *rfs.SetattrRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)
	return setattrHook(req, hm.client)
}
