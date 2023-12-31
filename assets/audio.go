package assets

import (
	"fmt"
	"io"
	"strings"

	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/xm"
	"github.com/quasilyte/xm/xmfile"
)

const (
	SoundGroupEffect uint = iota
	SoundGroupMusic
)

func VolumeMultiplier(level int) float64 {
	switch level {
	case 1:
		return 0.01
	case 2:
		return 0.15
	case 3:
		return 0.45
	case 4:
		return 0.8
	case 5:
		return 1.0
	default:
		return 0
	}
}

func registerAudioResources(ctx *ge.Context, config Config) {
	audioResources := map[resource.AudioID]resource.AudioInfo{
		AudioMusic1:    {Path: "music/drozerix-leisurely_voice.xm", Group: SoundGroupMusic},
		AudioMusic2:    {Path: "music/drozerix-porn_industry.xm", Group: SoundGroupMusic},
		AudioMusic3:    {Path: "music/drozerix-playful_girl.xm", Group: SoundGroupMusic},
		AudioMusicMenu: {Path: "music/drozerix-crush.xm", Group: SoundGroupMusic},

		AudioFragImpact:     {Path: "sfx/frag_impact.wav"},
		AudioMegaBombImpact: {Path: "sfx/megabomb_impact.wav"},
	}

	xmParser := xmfile.NewParser(xmfile.ParserConfig{})
	ctx.Loader.CustomAudioLoader = func(r io.Reader, info resource.AudioInfo) io.ReadSeeker {
		if !strings.HasSuffix(info.Path, ".xm") {
			return nil
		}
		m, err := xmParser.Parse(r)
		if err != nil {
			panic(fmt.Sprintf("parse %q module: %v", info.Path, err))
		}
		s := xm.NewStream()
		s.SetEventHandler(config.PlayerListener)
		s.SetLooping(true)
		config := xm.LoadModuleConfig{
			LinearInterpolation: true,
		}
		if err := s.LoadModule(m, config); err != nil {
			panic(fmt.Sprintf("load %q module: %v", info.Path, err))
		}
		return s
	}

	for id, res := range audioResources {
		ctx.Loader.AudioRegistry.Set(id, res)
		ctx.Loader.LoadAudio(id)
	}

	ctx.Loader.CustomAudioLoader = nil
}

func NumSamples(a resource.AudioID) int {
	switch a {
	default:
		return 1
	}
}

const (
	AudioNone resource.AudioID = iota

	AudioFragImpact
	AudioMegaBombImpact

	AudioMusic1
	AudioMusic2
	AudioMusic3
	AudioMusicMenu
)
