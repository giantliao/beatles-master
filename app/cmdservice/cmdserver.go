package cmdservice

import (
	"github.com/giantliao/beatles-master/db"
	"github.com/giantliao/beatles-master/webserver"
	"google.golang.org/grpc"
	"sync"
	"time"

	"net"

	"errors"
	"google.golang.org/grpc/reflection"
	"log"

	"github.com/giantliao/beatles-master/app/cmdpb"
	"github.com/giantliao/beatles-master/app/cmdservice/api"
	"github.com/giantliao/beatles-master/config"
)

type cmdServer struct {
	localaddr  string
	grpcServer *grpc.Server
}

type CmdServerInter interface {
	StartCmdService()
	StopCmdService()
}

var (
	cmdServerInst     CmdServerInter
	cmdServerInstLock sync.Mutex
	stopflag          bool
	stopflagLock      sync.Mutex
)

func GetCmdServerInst() CmdServerInter {
	if cmdServerInst == nil {
		cmdServerInstLock.Lock()
		defer cmdServerInstLock.Unlock()
		if cmdServerInst == nil {
			cfg := config.GetCBtlm()
			cmdServerInst = &cmdServer{localaddr: cfg.CmdListenPort}
		}
	}

	return cmdServerInst
}

func (cs *cmdServer) checklocaladdress() error {
	if cs.localaddr == "" {
		return errors.New("No Server Listen address")
	}

	return nil
}

func (cs *cmdServer) StartCmdService() {
	if err := cs.checklocaladdress(); err != nil {
		log.Fatal("Start Cmd Service Failed", err)
		return
	}

	lis, err := net.Listen("tcp", cs.localaddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cs.grpcServer = grpc.NewServer()

	cmdpb.RegisterDefaultcmdsrvServer(cs.grpcServer, &api.CmdDefaultServer{stop})
	cmdpb.RegisterStringopsrvServer(cs.grpcServer, &api.CmdStringOPSrv{})

	reflection.Register(cs.grpcServer)
	log.Println("Commamd line server will start at", cs.localaddr)
	if err := cs.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (cs *cmdServer) StopCmdService() {
	config.GetCBtlm().Save()

	time.Sleep(time.Second * 5)

	cs.grpcServer.Stop()
	log.Println("Command line server stoped")
}

func stop() {

	if !stopflag {
		stopflagLock.Lock()
		if !stopflag {
			stopflag = true
		} else {
			stopflagLock.Unlock()
			return
		}
		stopflagLock.Unlock()
	} else {
		return
	}

	db.GetMinersDb().Close()
	webserver.StopWebDaemon()

	GetCmdServerInst().StopCmdService()

}
