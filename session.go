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

package httl

import "context"

type Session struct {
	data map[string]any
	ctx  context.Context
}

func (s *Session) Get(key string) any {
	v, _ := s.data[key]
	return v
}

func (s *Session) Remove(key string) any {
	return nil
}

func (s *Session) Size() int {
	return len(s.data)
}

func (s *Session) Keys() []string {
	var keys = make([]string, 16)
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}
