package GameData

const (
	INIT_GAMEDATA_PROTO = iota
	C2SPlayerLoginProto // 1 Client - User Login
	S2CPlayerLoginProto // 2 Server - User Login
	C2SChooseRoomProto  // 3 Client Select Room
	S2CChooseRoomProto  // 4 Server Select Room
)

type PlayerSt struct {
	UID        int
	PlayerName string
	OpenID     string
}

type HeadProto struct {
	Protocol int // main protocol
	GameData int // child protocol
}

// Client - Server
type C2SPlayerLogin struct {
	HeadProto
	Code string // validate code
}

// Server - Client
type S2CPlayerLogin struct {
	HeadProto
	PlayerData *PlayerSt // Player struct
}
