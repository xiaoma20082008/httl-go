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

package resolvers

import "github.com/xiaoma20082008/httl/spi"

var global map[string]any

func init() {
	global = make(map[string]any)
}

type GlobalResolver struct {
	spi.Resolver
}

func (r *GlobalResolver) Resolve(key string) any {
	if v, ok := global[key]; ok {
		return v
	}
	return nil
}

func NewGlobalResolver() *GlobalResolver {
	return &GlobalResolver{}
}