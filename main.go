package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"git.nightcrickets.space/keefleoflimon/resonate/fuse"
	"git.nightcrickets.space/keefleoflimon/resonate/network"
	"git.nightcrickets.space/keefleoflimon/resonate/util"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 1234, "choose a port number for self")
	peer = flag.String("peer", "localhost:4321", "choose peer address")
	dir  = flag.String("dir", "test", "choose directory to sync")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*peer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not esaplish connection with peer (%v)", peer)
	}

	client := network.NewFileManagerClient(conn)
	lm := util.NewLockManager()
	hm := NewHookManager(lm, client)
	vol := fuse.NewVolume(*dir, hm.HooksToOptions()...)
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

	done := make(chan struct{}, 1)
	cleanup(vol.Stop, done)

	if err := vol.Serve(); err != nil {
		log.Printf("serve failed: %v", err)
	}

	<-done
}

func cleanup(term func() error, done chan<- struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range c {
			if err := term(); err != nil {
				log.Printf("Error occured when tring to terminate: %v", err)
			}
			done <- struct{}{}

			// check if needed
			close(c)
		}
	}()
}
