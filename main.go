package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"git.nightcrickets.space/keefleoflimon/resonate/network"
	rfs "git.nightcrickets.space/keefleoflimon/resonatefuse"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 1234, "choose a port number for self")
	peer = flag.String("peer", "127.0.0.1:4321", "choose peer address")
	dir  = flag.String("dir", "fake", "choose directory to sync")
)

func main() {
	flag.Parse()

	vol := rfs.NewVolume(*dir, createHook, writeHook, removeHook, mkdirHook, renameHook)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	network.RegisterFileManagerServer(grpcServer, network.NewFileManager(vol.Fuse()))
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	vol.Serve()
}

func createHook(req *rfs.CreateRequest) {
	log.Println("hook running")
	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	log.Println(err)

	client := network.NewFileManagerClient(conn)

	_, err = client.Create(context.Background(), &network.CreateRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
}

func writeHook(req *rfs.WriteRequest) {
	log.Println("hook running")
	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	log.Println(err)

	client := network.NewFileManagerClient(conn)
	_, err = client.Write(context.Background(), &network.WriteRequest{Path: req.Path, Data: req.Data, Offset: req.Offset})
	log.Println(err)
}

func removeHook(req *rfs.RemoveRequest) {
	log.Println("hook running")
	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	log.Println(err)

	client := network.NewFileManagerClient(conn)
	_, err = client.Remove(context.Background(), &network.RemoveRequest{Path: req.Path, Name: req.Name})
	log.Println(err)
}

func renameHook(req *rfs.RenameRequest) {
	log.Println("hook running")
	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	log.Println(err)

	client := network.NewFileManagerClient(conn)
	_, err = client.Rename(context.Background(), &network.RenameRequest{Path: req.Path, Oldname: req.OldName, Newname: req.NewName, Newdirpath: req.NewDir})
	log.Println(err)
}

func mkdirHook(req *rfs.MkdirRequest) {
	log.Println("hook running")
	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	log.Println(err)

	client := network.NewFileManagerClient(conn)
	_, err = client.Mkdir(context.Background(), &network.MkdirRequest{Path: req.Path, Name: req.Name, Mode: uint32(req.Mode)})
	log.Println(err)
}
