package cmdcommon

const (
	CMD_STOP           int32 = 1
	CMD_CONFIG_SHOW    int32 = 2
	CMD_BOOTSTRAP_SHOW int32 = 3
	CMD_RUN            int32 = 4
	CMD_MINER_SHOW     int32 = 5
	CMD_MINER_SAVE     int32 = 6

	CMD_BOOTSTRAP_PUSHALL int32 = 8
)

const (
	CMD_WALLET_CREATE  int32 = 2
	CMD_WALLET_LOAD    int32 = 3
	CMD_WALLET_SHOW    int32 = 4
	CMD_BOOTSTRAP_LIST int32 = 5
	CMD_BOOTSTRAP_ADD  int32 = 6
	CMD_BOOTSTRAP_DEL  int32 = 7
	CMD_BOOTSTRAP_PUSH int32 = 8
	CMD_MINER_REMOVE   int32 = 9
	CMD_DB_SHOW        int32 = 10
	CMD_DB_SAVE        int32 = 11

	////CMD_JOIN_GROUP         int32 = 9
	//CMD_QUIT_GROUP         int32 = 10
	//CMD_LIST_GROUPMBRS     int32 = 11
	//CMD_LISTEN_FRIEND      int32 = 12
	//CMD_LISTEN_GROUP       int32 = 13
	//CMD_SEND_P2PMSG        int32 = 14
	//CMD_SEND_GMSG          int32 = 15
	//CMD_QUIT_LISTEN_FRIEND int32 = 16
	//CMD_QUIT_LISTEN_GROUP  int32 = 17
)
