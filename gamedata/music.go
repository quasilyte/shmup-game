package gamedata

type MusicChannelKind int

const (
	ChannelUnknown MusicChannelKind = iota
	ChannelPlayerAttack
	ChannelEnemyAttack
	ChannelEnemyAltAttack
	ChannelEnemySpecialAttack
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

		9: {ChannelEnemySpecialAttack, 70},

		4: {ChannelEnemyAltAttack, 70},

		10: {ChannelEnemyAttack, 70},
		11: {ChannelEnemyAttack, 70},
	},
}
