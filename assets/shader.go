package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
)

func registerShaderResources(ctx *ge.Context) {
	shaderResources := map[resource.ShaderID]resource.ShaderInfo{
		ShaderCRT: {Path: "shader/crt.go"},
	}

	for id, res := range shaderResources {
		ctx.Loader.ShaderRegistry.Set(id, res)
		ctx.Loader.LoadShader(id)
	}
}

const (
	ShaderNone resource.ShaderID = iota

	ShaderCRT
)
