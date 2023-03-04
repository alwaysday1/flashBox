// # Copyright 2023 LightTree <alwaysday1@qq.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/alwaysday1/flashBox.

package post

import (
	"flash_box_server/internal/pkg/core"
	"flash_box_server/internal/pkg/known"
	"flash_box_server/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// Get 获取指定的博客.
func (ctrl *PostController) Get(c *gin.Context) {
	log.C(c).Infow("Get post function called")

	post, err := ctrl.b.Posts().Get(c, c.GetString(known.XUsernameKey), c.Param("postID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
