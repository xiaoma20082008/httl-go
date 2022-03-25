/*
 *  Copyright 2022 HTTL-Go Team
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package caches

import "github.com/xiaoma20082008/httl-go/spi"

type mapCache struct {
	cache map[string]any
	spi.Cache
}

func (c *mapCache) Has(key string) bool {
	return c.TryGet(key, nil)
}

func (c *mapCache) Put(key string, value any) {
	c.cache[key] = value
}

func (c *mapCache) Get(key string) any {
	if value, ok := c.cache[key]; ok {
		return value
	}
	return nil
}

func (c *mapCache) TryGet(key string, out *any) bool {
	if value, ok := c.cache[key]; ok {
		if out != nil {
			*out = value
		}
		return true
	}
	return false
}

func NewMapCache() spi.Cache {
	return &mapCache{
		cache: map[string]any{},
	}
}
