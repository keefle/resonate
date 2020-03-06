package main

import (
	"context"
	"log"
	"path/filepath"

	"git.nightcrickets.space/keefleoflimon/resonate/network"
	rfs "git.nightcrickets.space/keefleoflimon/resonatefuse"
	"github.com/pkg/errors"
)

func genCreateHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("create hook running")

	_, err := client.Create(context.Background(), &network.CreateRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func (hm *HookManager) genCreateHook(req *rfs.GeneralRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeCreated := filepath.Join(req.Path, req.Name)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)

	return genCreateHook(req, hm.client)
}

func genWriteHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("write hook running")

	_, err := client.Write(context.Background(), &network.WriteRequest{Path: req.Path, Data: req.Data, Offset: req.Offset})
	log.Println(err)
	return err
}

func (hm *HookManager) genWriteHook(req *rfs.GeneralRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)
	return genWriteHook(req, hm.client)
}

func genRemoveHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("remove hook running")

	_, err := client.Remove(context.Background(), &network.RemoveRequest{Path: req.Path, Name: req.Name})
	log.Println(err)
	return err
}

func (hm *HookManager) genRemoveHook(req *rfs.GeneralRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)
	return genRemoveHook(req, hm.client)
}

func genRenameHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("rename hook running")

	_, err := client.Rename(context.Background(), &network.RenameRequest{Path: req.Path, Oldname: req.OldName, Newname: req.NewName, Newdirpath: req.NewDir})
	log.Println(err)
	return err
}

func (hm *HookManager) genRenameHook(req *rfs.GeneralRequest) error {
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

	return genRenameHook(req, hm.client)
}

func genMkdirHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("mkdir hook running")

	_, err := client.Mkdir(context.Background(), &network.MkdirRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func (hm *HookManager) genMkdirHook(req *rfs.GeneralRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)

	toBeCreated := filepath.Join(req.Path, req.Name)
	if !hm.lm.Lock(toBeCreated) {
		return errors.Errorf("could not lock file (%v)", toBeCreated)
	}
	defer hm.lm.Unlock(toBeCreated)
	return genMkdirHook(req, hm.client)
}

func genLinkHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("link hook running")

	_, err := client.Link(context.Background(), &network.LinkRequest{Path: req.Path, Newname: req.NewName, Old: req.Old})
	log.Println(err)
	return err
}

func (hm *HookManager) genLinkHook(req *rfs.GeneralRequest) error {
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

	return genLinkHook(req, hm.client)
}

func genSymlinkHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("symlink hook running")

	_, err := client.Symlink(context.Background(), &network.SymlinkRequest{Path: req.Path, Newname: req.NewName, Target: req.Target})
	log.Println(err)
	return err
}

func (hm *HookManager) genSymlinkHook(req *rfs.GeneralRequest) error {
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
	return genSymlinkHook(req, hm.client)
}

func genSetattrHook(req *rfs.GeneralRequest, client network.FileManagerClient) error {
	log.Println("setattr hook running")

	_, err := client.Setattr(context.Background(), &network.SetattrRequest{Path: req.Path, Mode: uint32(req.Mode), Atime: req.Atime.Unix(), Mtime: req.Mtime.Unix()})
	log.Println(err)
	return err
}

func (hm *HookManager) genSetattrHook(req *rfs.GeneralRequest) error {
	if !hm.lm.Lock(req.Path) {
		return errors.Errorf("could not lock file (%v)", req.Path)
	}
	defer hm.lm.Unlock(req.Path)
	return genSetattrHook(req, hm.client)
}
