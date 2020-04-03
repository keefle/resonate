package main

import (
	"context"
	"log"

	"git.nightcrickets.space/keefleoflimon/resonate/fuse"
	"git.nightcrickets.space/keefleoflimon/resonate/network"
	"git.nightcrickets.space/keefleoflimon/resonate/util"
)

type HookManager struct {
	lm     *util.LockManager
	client network.FileManagerClient
}

func NewHookManager(lm *util.LockManager, client network.FileManagerClient) *HookManager {
	return &HookManager{lm: lm, client: client}
}

func (hm *HookManager) HooksToOptions() []fuse.Option {
	return []fuse.Option{
		fuse.GeneralOption(fuse.CreateType, hm.genCreateHook),
		fuse.GeneralOption(fuse.WriteType, hm.genWriteHook),
		fuse.GeneralOption(fuse.RemoveType, hm.genRemoveHook),
		fuse.GeneralOption(fuse.MkdirType, hm.genMkdirHook),
		fuse.GeneralOption(fuse.RenameType, hm.genRenameHook),
		fuse.GeneralOption(fuse.LinkType, hm.genLinkHook),
		fuse.GeneralOption(fuse.SymlinkType, hm.genSymlinkHook),
		fuse.GeneralOption(fuse.SetattrType, hm.genSetattrHook),
	}
}

func createHook(req *fuse.CreateRequest, client network.FileManagerClient) error {
	log.Println("create hook running")

	_, err := client.Create(context.Background(), &network.CreateRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func writeHook(req *fuse.WriteRequest, client network.FileManagerClient) error {
	log.Println("write hook running")

	_, err := client.Write(context.Background(), &network.WriteRequest{Path: req.Path, Data: req.Data, Offset: req.Offset})
	log.Println(err)
	return err
}

func removeHook(req *fuse.RemoveRequest, client network.FileManagerClient) error {
	log.Println("remove hook running")

	_, err := client.Remove(context.Background(), &network.RemoveRequest{Path: req.Path, Name: req.Name})
	log.Println(err)
	return err
}

func renameHook(req *fuse.RenameRequest, client network.FileManagerClient) error {
	log.Println("rename hook running")

	_, err := client.Rename(context.Background(), &network.RenameRequest{Path: req.Path, Oldname: req.OldName, Newname: req.NewName, Newdirpath: req.NewDir})
	log.Println(err)
	return err
}

func mkdirHook(req *fuse.MkdirRequest, client network.FileManagerClient) error {
	log.Println("mkdir hook running")

	_, err := client.Mkdir(context.Background(), &network.MkdirRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func linkHook(req *fuse.LinkRequest, client network.FileManagerClient) error {
	log.Println("link hook running")

	_, err := client.Link(context.Background(), &network.LinkRequest{Path: req.Path, Newname: req.NewName, Old: req.Old})
	log.Println(err)
	return err
}

func symlinkHook(req *fuse.SymlinkRequest, client network.FileManagerClient) error {
	log.Println("symlink hook running")

	_, err := client.Symlink(context.Background(), &network.SymlinkRequest{Path: req.Path, Newname: req.NewName, Target: req.Target})
	log.Println(err)
	return err
}

func setattrHook(req *fuse.SetattrRequest, client network.FileManagerClient) error {
	log.Println("setattr hook running")

	_, err := client.Setattr(context.Background(), &network.SetattrRequest{Path: req.Path, Mode: uint32(req.Mode), Atime: req.Atime.Unix(), Mtime: req.Mtime.Unix()})
	log.Println(err)
	return err
}
