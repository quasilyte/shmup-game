package gamedata

type MusicChannelKind int

const (
	ChannelUnknown MusicChannelKind = iota
	ChannelPlayerAttack
	ChannelEnemyAttack
)

type MusicChannelInfo struct {
	Kind     MusicChannelKind
	HighNote int
}

type MusicInfo struct {
	Channels []MusicChannelInfo
}

var Music1 = &MusicInfo{
	Channels: []MusicChannelInfo{
		0: {ChannelPlayerAttack, 70},
		1: {ChannelPlayerAttack, 70},
		2: {ChannelPlayerAttack, 70},
	},
}
