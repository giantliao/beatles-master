package wallet

import (
	"encoding/json"
	"errors"
	"github.com/giantliao/beatles-master/config"
	"github.com/kprc/libeth/account"
	"github.com/kprc/libeth/wallet"
	"github.com/kprc/nbsnetwork/tools"
)

var (
	beatlesMasterWallet wallet.WalletIntf
)

func GetWallet() (wallet.WalletIntf, error) {
	if beatlesMasterWallet == nil {
		return nil, errors.New("no wallet, please load it")
	}
	return beatlesMasterWallet, nil
}

func newWallet(auth, savepath, remoteeth string) wallet.WalletIntf {
	w := wallet.CreateWallet(savepath, remoteeth)

	if w == nil {
		return nil
	}

	w.Save(auth)

	return w
}

func LoadWallet(auth string) error {
	cfg := config.GetCBtlm()

	if !tools.FileExists(cfg.GetWalletSavePath()) {
		beatlesMasterWallet = newWallet(auth, cfg.GetWalletSavePath(), cfg.EthAccessPoint)
		if beatlesMasterWallet == nil {
			return errors.New("create wallet failed ")
		}
	} else {
		var err error
		beatlesMasterWallet, err = wallet.RecoverWallet(cfg.GetWalletSavePath(), cfg.EthAccessPoint, auth)
		if err != nil {
			return errors.New("load wallet failed : " + err.Error())
		}
	}
	return nil
}

func ShowWallet() (string, error) {
	cfg := config.GetCBtlm()
	if !tools.FileExists(cfg.GetWalletSavePath()) {
		return "", errors.New("no wallet")
	}

	var (
		data []byte
		err  error
	)

	if data, err = tools.OpenAndReadAll(cfg.GetWalletSavePath()); err != nil {
		return "", err
	}

	wsj := &wallet.WalletSaveJson{}

	if err = json.Unmarshal(data, wsj); err != nil {
		return "", err
	}

	var (
		ethAcct *account.AccountJson
		btlAcct *account.CryptBTLJson
	)
	if wsj.EthAcct != "" {
		if ethAcct, err = account.EthUnmarshal([]byte(wsj.EthAcct)); err != nil {
			return "", err
		}
	}
	if wsj.BtlAcct != "" {
		if btlAcct, err = account.BeatlesUnmarshal([]byte(wsj.BtlAcct)); err != nil {
			return "", err
		}
	}

	var (
		jeth []byte
		jbtl []byte
	)

	msg := "Wallet Save Path: " + cfg.GetWalletSavePath() + "\r\n"

	if ethAcct != nil {
		jeth, _ = json.MarshalIndent(ethAcct, " ", "\t")
		msg += "Eth Account:\r\n"
		msg += string(jeth) + "\r\n"
	}
	if btlAcct != nil {
		jbtl, _ = json.MarshalIndent(btlAcct, " ", "\t")
		msg += "Beatles Account:\r\n"
		msg += string(jbtl) + "\r\n"
	}

	return msg, nil
}
