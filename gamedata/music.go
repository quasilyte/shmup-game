package gamedata

type MusicChannelKind int

const (
	ChannelUnknown MusicChannelKind = iota
	ChannelPlayerAttack
	ChannelEnemyAttack
	ChannelEnemyAltAttack
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

		11: {ChannelEnemyAltAttack, 70},

		12: {ChannelEnemyAttack, 70},
		13: {ChannelEnemyAttack, 70},
		14: {ChannelEnemyAttack, 70},
	},
}
