package api

import (
	"context"
	"encoding/json"
	"github.com/giantliao/beatles-master/app/cmdcommon"
	"github.com/giantliao/beatles-master/app/cmdpb"
	"github.com/giantliao/beatles-master/config"
	"time"
)

type CmdDefaultServer struct {
	Stop func()
}

func (cds *CmdDefaultServer) DefaultCmdDo(ctx context.Context,
	request *cmdpb.DefaultRequest) (*cmdpb.DefaultResp, error) {

	msg := ""

	switch request.Reqid {
	case cmdcommon.CMD_STOP:
		msg = cds.stop()
	case cmdcommon.CMD_CONFIG_SHOW:
		msg = cds.configShow()
	case cmdcommon.CMD_BOOTSTRAP_SHOW:
		msg = cds.bootstrapShow()
	}

	if msg == "" {
		msg = "No Results"
	}

	resp := &cmdpb.DefaultResp{}
	resp.Message = msg

	return resp, nil

}

func (cds *CmdDefaultServer) stop() string {

	go func() {
		time.Sleep(time.Second * 2)
		cds.Stop()
	}()

	return "beatles master stopped"
}

func encapResp(msg string) *cmdpb.DefaultResp {
	resp := &cmdpb.DefaultResp{}
	resp.Message = msg

	return resp
}

func (cds *CmdDefaultServer) configShow() string {
	cfg := config.GetCBtlm()

	bapc, err := json.MarshalIndent(*cfg, "", "\t")
	if err != nil {
		return "Internal error"
	}

	return string(bapc)
}

func (cds *CmdDefaultServer) bootstrapShow() string {
	cfg := config.GetCBtlm()

	msg := cfg.BootstrapString()

	return msg
}
