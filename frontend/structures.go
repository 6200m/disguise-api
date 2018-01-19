package frontend

var ConfigMagic = []byte("WcCf")

type ConfigFormat struct {
	Magic             [4]byte
	Version           int32
	FriendCode        int64
	AmountOfCreations int32
	HasRegistered     int32
	MailDomain        [40]byte
	Passwd            [20]byte
	Mlchkid           [24]byte
	AccountURL        [128]byte
	CheckURL          [128]byte
	ReceiveURL        [128]byte
	DeleteURL         [128]byte
	SendURL           [128]byte
	_                 [220]byte // Most likely reserved.
	TitleBooting      int32
	Checksum          [4]byte
}
