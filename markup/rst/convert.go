// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rst converts content to HTML using the RST external helper.
package rst

import (
	"github.com/cli/safeexec"
	"github.com/tychoish/shimgo"

	"github.com/gohugoio/hugo/identity"

	"github.com/gohugoio/hugo/markup/converter"
)

// Provider is the package entry point.
var Provider converter.ProviderProvider = provider{}

type provider struct {
}

func (p provider) New(cfg converter.ProviderConfig) (converter.Provider, error) {
	return converter.NewProvider("rst", func(ctx converter.DocumentContext) (converter.Converter, error) {
		return &rstConverter{
			ctx: ctx,
			cfg: cfg,
		}, nil
	}), nil
}

type rstConverter struct {
	ctx converter.DocumentContext
	cfg converter.ProviderConfig
}

func (c *rstConverter) Convert(ctx converter.RenderContext) (converter.Result, error) {
	return converter.Bytes(c.getRstContent(ctx.Src, c.ctx)), nil
}

func (c *rstConverter) Supports(feature identity.Identity) bool {
	return false
}

// getRstContent calls the Python script rst2html as an external helper
// to convert reStructuredText content to HTML.
func (c *rstConverter) getRstContent(src []byte, ctx converter.DocumentContext) []byte {
	out, err := shimgo.ConvertFromRst(src)
	if err != nil {
		c.cfg.Logger.Errorln("Problem rendering", ctx.DocumentName+":", err.Error())
		return src
	}

	return out
}

func getRstExecPath() string {
	path, err := safeexec.LookPath("rst2html")
	if err != nil {
		path, err = safeexec.LookPath("rst2html.py")
		if err != nil {
			return ""
		}
	}
	return path
}

// Supports returns whether rst is installed on this computer.
func Supports() bool {
	return getRstExecPath() != ""
}
