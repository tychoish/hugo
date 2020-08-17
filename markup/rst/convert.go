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
	"os/exec"

	"github.com/gohugoio/hugo/identity"
	"github.com/gohugoio/hugo/markup/converter"
	"github.com/tychoish/shimgo"
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

func (c *rstConverter) getRstContent(src []byte, ctx converter.DocumentContext) []byte {
	logger := c.cfg.Logger

	out, err := shimgo.ConvertFromRst(src)
	if out == nil {
		if err != nil {
			logger.ERROR.Printf("problem converting %s to rst; dependencies (docutils; flask) may not be installed (%v)",
				ctx.DocumentName, err)
			return nil
		}

		logger.ERROR.Printf("unknown error converting rst; %s not rendered", ctx.DocumentName)
		return nil
	}

	if err != nil {
		logger.ERROR.Printf("found rst error: %s:%v", ctx.DocumentName, err)
	}

	return out
}

func getRstExecPath() string {
	path, err := exec.LookPath("rst2html")
	if err != nil {
		path, err = exec.LookPath("rst2html.py")
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
