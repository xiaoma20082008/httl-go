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

package loaders

import (
	"github.com/xiaoma20082008/httl"
)

type BaseLoader struct {
	engine   httl.Engine
	reload   bool
	locale   string
	encoding string
	prefix   []string
	suffix   []string
}

func (l *BaseLoader) SetEngine(e httl.Engine) {
	l.engine = e
}

func (l *BaseLoader) SetReload(reload bool) {
	l.reload = reload
}

func (l *BaseLoader) SetLocale(locale string) {
	l.locale = locale
}

func (l *BaseLoader) SetEncoding(encoding string) {
	l.encoding = encoding
}

func (l *BaseLoader) SetPrefix(prefix []string) {
	l.prefix = prefix
}

func (l *BaseLoader) SetSuffix(suffix []string) {
	l.suffix = suffix
}

func (l *BaseLoader) ToPath(name string) string {
	return ""
}
