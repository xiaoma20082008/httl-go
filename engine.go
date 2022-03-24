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

type Engine interface {
	GetName() string
	GetVersion() string
	GetProperty(key string, dv string) string
	GetAny(key string) any
	GetBool(key string, dv bool) bool
	GetInt32(key string, dv string) int32
	GetInt64(key string, dv string) int64
	GetTemplate(name string, locale string, encoding string) (Template, error)
	GetResource(name string, locale string, encoding string) (Resource, error)
	NewContext(parent map[string]any, current map[string]any) map[string]any
}

func NewEngine() Engine {
	return nil
}
