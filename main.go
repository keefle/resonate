package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"git.nightcrickets.space/keefleoflimon/resonate/network"
	"git.nightcrickets.space/keefleoflimon/resonate/util"
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

	lm := util.NewLockManager()
	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not esaplish connection with peer (%v)", peer)
	}

	client := network.NewFileManagerClient(conn)
	hm := NewHookManager(lm, client)
	vol := rfs.NewVolume(*dir, hm.HooksToOptions()...)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	network.RegisterFileManagerServer(grpcServer, network.NewFileManager(vol.Fuse(), lm))
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	vol.Serve()
}
