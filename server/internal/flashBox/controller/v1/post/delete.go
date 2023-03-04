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

// Delete 删除指定的博客.
func (ctrl *PostController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete post function called")

	if err := ctrl.b.Posts().Delete(c, c.GetString(known.XUsernameKey), c.Param("postID")); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
