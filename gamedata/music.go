package gamedata

type MusicChannelKind int

const (
	ChannelUnknown MusicChannelKind = iota
	ChannelPlayerAttack
	ChannelEnemyAttack
	ChannelEnemyAltAttack
	ChannelEnemySpecialAttack
)

type MusicChannelVariant struct {
	Kind     MusicChannelKind
	HighNote int
	Inst     int
}

type MusicInfo struct {
	Channels [][]MusicChannelVariant
}

// var Music1 = &MusicInfo{
// 	Channels: []MusicChannelInfo{
// 		0: {ChannelPlayerAttack, 70},
// 		1: {ChannelPlayerAttack, 70},
// 		2: {ChannelPlayerAttack, 70},

// 		9: {ChannelEnemySpecialAttack, 70},

// 		4: {ChannelEnemyAltAttack, 70},

// 		10: {ChannelEnemyAttack, 70},
// 		11: {ChannelEnemyAttack, 70},
// 	},
// }

var Music2 = &MusicInfo{
	Channels: [][]MusicChannelVariant{
		0: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemyAltAttack, 70, 1}, {ChannelEnemySpecialAttack, 70, 1}},
		1: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemyAltAttack, 70, 1}, {ChannelEnemySpecialAttack, 70, 1}},
		2: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemySpecialAttack, 70, 1}, {ChannelEnemyAttack, 70, 0}},
		3: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemyAltAttack, 70, 1}, {ChannelEnemySpecialAttack, 70, 1}, {ChannelEnemyAttack, 70, 0}},

		6: {{ChannelEnemySpecialAttack, 70, 5}},

		// 10: {ChannelEnemyAltAttack, 70},

		8: {{ChannelPlayerAttack, 50, 2}},
		9: {{ChannelPlayerAttack, 50, 2}, {ChannelPlayerAttack, 50, 3}},
	},
}
