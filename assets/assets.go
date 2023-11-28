package assets

import (
	"embed"
	"io"

	"github.com/quasilyte/ge"
	"github.com/quasilyte/xm"
)

//go:embed all:_data
var gameAssets embed.FS

type Config struct {
	PlayerListener func(e xm.StreamEvent)
}

func MakeOpenAssetFunc(ctx *ge.Context) func(path string) io.ReadCloser {
	return func(path string) io.ReadCloser {
		f, err := gameAssets.Open("_data/" + path)
		if err != nil {
			ctx.OnCriticalError(err)
		}
		return f
	}
}

func RegisterResources(ctx *ge.Context, config Config) {
	registerImageResources(ctx)
	registerRawResources(ctx)
	registerAudioResources(ctx, config)
	registerShaderResources(ctx)
}
