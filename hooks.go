package main

import (
	"context"
	"log"

	"git.nightcrickets.space/keefleoflimon/resonate/network"
	"git.nightcrickets.space/keefleoflimon/resonate/util"
	rfs "git.nightcrickets.space/keefleoflimon/resonatefuse"
)

type HookManager struct {
	lm     *util.LockManager
	client network.FileManagerClient
}

func NewHookManager(lm *util.LockManager, client network.FileManagerClient) *HookManager {
	return &HookManager{lm: lm, client: client}
}

func (hm *HookManager) HooksToOptions() []rfs.Option {
	return []rfs.Option{
		rfs.GeneralOption(rfs.CreateType, hm.genCreateHook),
		rfs.GeneralOption(rfs.WriteType, hm.genWriteHook),
		rfs.GeneralOption(rfs.RemoveType, hm.genRemoveHook),
		rfs.GeneralOption(rfs.MkdirType, hm.genMkdirHook),
		rfs.GeneralOption(rfs.RenameType, hm.genRenameHook),
		rfs.GeneralOption(rfs.LinkType, hm.genLinkHook),
		rfs.GeneralOption(rfs.SymlinkType, hm.genSymlinkHook),
		rfs.GeneralOption(rfs.SetattrType, hm.genSetattrHook),
	}
}

func createHook(req *rfs.CreateRequest, client network.FileManagerClient) error {
	log.Println("create hook running")

	_, err := client.Create(context.Background(), &network.CreateRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func writeHook(req *rfs.WriteRequest, client network.FileManagerClient) error {
	log.Println("write hook running")

	_, err := client.Write(context.Background(), &network.WriteRequest{Path: req.Path, Data: req.Data, Offset: req.Offset})
	log.Println(err)
	return err
}

func removeHook(req *rfs.RemoveRequest, client network.FileManagerClient) error {
	log.Println("remove hook running")

	_, err := client.Remove(context.Background(), &network.RemoveRequest{Path: req.Path, Name: req.Name})
	log.Println(err)
	return err
}

func renameHook(req *rfs.RenameRequest, client network.FileManagerClient) error {
	log.Println("rename hook running")

	_, err := client.Rename(context.Background(), &network.RenameRequest{Path: req.Path, Oldname: req.OldName, Newname: req.NewName, Newdirpath: req.NewDir})
	log.Println(err)
	return err
}

func mkdirHook(req *rfs.MkdirRequest, client network.FileManagerClient) error {
	log.Println("mkdir hook running")

	_, err := client.Mkdir(context.Background(), &network.MkdirRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
	return err
}

func linkHook(req *rfs.LinkRequest, client network.FileManagerClient) error {
	log.Println("link hook running")

	_, err := client.Link(context.Background(), &network.LinkRequest{Path: req.Path, Newname: req.NewName, Old: req.Old})
	log.Println(err)
	return err
}

func symlinkHook(req *rfs.SymlinkRequest, client network.FileManagerClient) error {
	log.Println("symlink hook running")

	_, err := client.Symlink(context.Background(), &network.SymlinkRequest{Path: req.Path, Newname: req.NewName, Target: req.Target})
	log.Println(err)
	return err
}

func setattrHook(req *rfs.SetattrRequest, client network.FileManagerClient) error {
	log.Println("setattr hook running")

	_, err := client.Setattr(context.Background(), &network.SetattrRequest{Path: req.Path, Mode: uint32(req.Mode), Atime: req.Atime.Unix(), Mtime: req.Mtime.Unix()})
	log.Println(err)
	return err
}
