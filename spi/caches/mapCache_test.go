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

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewMapCache(t *testing.T) {
	cache := NewMapCache()
	v1 := 123
	cache.Put("Tom", v1)
	v2 := cache.Get("Tom")
	assert.Equal(t, 123, v2)
	v3 := cache.Get("Tim")
	assert.Equal(t, v3, nil)
	var v4 any = ""
	v5 := cache.TryGet("Tim", &v4)
	assert.Equal(t, v5, false)
}
