/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cache

import "time"

// Key ...
type Key interface {
	Key() string
}

// StringKey ...
type StringKey struct {
	key string
}

// NewStringKey ...
func NewStringKey(key string) StringKey {
	return StringKey{
		key: key,
	}
}

// Key ...
func (s StringKey) Key() string {
	return s.key
}

// Cache ...
type Cache interface {
	// 写缓存
	Set(key Key, value interface{}, duration time.Duration) error
	// 检查缓存
	Exists(key Key) bool
	// 读缓存
	Get(key Key, value interface{}) error
	// 清理缓存
	Delete(key Key) error
}