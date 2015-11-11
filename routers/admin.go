// Copyright 2015 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routers

import (
	//"fmt"
	//"strings"

	"github.com/peachdocs/peach/models"
	"github.com/peachdocs/peach/modules/middleware"
	"github.com/peachdocs/peach/modules/setting"
)

func TocManager(ctx *middleware.Context) {
	models.ReloadTocs()
	toc := models.Tocs[ctx.Locale.Language()]
	if toc == nil {
		toc = models.Tocs[setting.Docs.Langs[0]]
	}
	ctx.Data["Toc"] = toc

	//ctx.Data["Title"] = node.Title
	//ctx.Data["Content"] = fmt.Sprintf(`<script type="text/javascript" src="/%s/%s?=%d"></script>`, langVer, node.DocumentPath+".js", node.LastBuildTime)
	ctx.Data["KuuYee"] = "光子"

	ctx.HTML(200, "admin")
}
