// # Copyright 2023 LightTree <alwaysday1@qq.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/alwaysday1/flashBox.

package errno

// ErrPostNotFound 表示未找到想法.
var ErrPostNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PostNotFound", Message: "Post was not found."}
