package gamedata

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/shmup-game/assets"
)

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
	Name     string
	AudioID  resource.AudioID
	Channels [][]MusicChannelVariant
}

var MusicList = []*MusicInfo{
	Music2,
	Music1,
	Music3,
}

var Music1 = &MusicInfo{
	Name:    "Leisurely Voice",
	AudioID: assets.AudioMusic1,
	Channels: [][]MusicChannelVariant{
		0: {{ChannelPlayerAttack, 70, 2}, {ChannelPlayerAttack, 70, 3}},
		1: {{ChannelPlayerAttack, 70, 2}, {ChannelPlayerAttack, 70, 3}},
		2: {{ChannelPlayerAttack, 70, 2}, {ChannelPlayerAttack, 70, 3}},

		9: {{ChannelEnemySpecialAttack, 70, 1}},

		4: {{ChannelEnemyAltAttack, 70, 4}},

		10: {{ChannelEnemyAttack, 70, 1}},
		11: {{ChannelEnemyAttack, 70, 0}},
	},
}

var Music2 = &MusicInfo{
	Name:    "Industrial",
	AudioID: assets.AudioMusic2,
	Channels: [][]MusicChannelVariant{
		0: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemyAltAttack, 70, 1}, {ChannelEnemySpecialAttack, 70, 1}},
		1: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemyAltAttack, 70, 1}, {ChannelEnemySpecialAttack, 70, 1}},
		2: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemySpecialAttack, 70, 1}, {ChannelEnemyAttack, 70, 0}},
		3: {{ChannelEnemyAltAttack, 70, 6}, {ChannelEnemyAltAttack, 70, 1}, {ChannelEnemySpecialAttack, 70, 1}, {ChannelEnemyAttack, 70, 0}},

		6: {{ChannelEnemySpecialAttack, 70, 5}},

		8: {{ChannelPlayerAttack, 50, 2}},
		9: {{ChannelPlayerAttack, 50, 2}, {ChannelPlayerAttack, 50, 3}},
	},
}

var Music3 = &MusicInfo{
	Name:    "Playful",
	AudioID: assets.AudioMusic3,
	Channels: [][]MusicChannelVariant{
		0: {{ChannelPlayerAttack, 68, 5}, {ChannelPlayerAttack, 50, 1}, {ChannelPlayerAttack, 50, 0}},

		5: {{ChannelEnemyAltAttack, 70, 5}, {ChannelEnemyAltAttack, 70, 1}},
		6: {{ChannelEnemyAltAttack, 70, 5}, {ChannelEnemyAltAttack, 70, 1}},

		7: {{ChannelEnemySpecialAttack, 70, 3}, {ChannelEnemySpecialAttack, 70, 4}},
		8: {{ChannelEnemyAttack, 70, 5}},
		9: {{ChannelEnemyAttack, 70, 5}},
	},
}
