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
	"github.com/xiaoma20082008/httl-go"
	"github.com/xiaoma20082008/httl-go/spi"
	"github.com/xiaoma20082008/httl-go/spi/loaders/resources"
)

type StringLoader struct {
	resources map[string]resources.StringResource
	spi.Loader
}

func (l *StringLoader) Exists(name string, locale string, encoding string) bool {
	_, ok := l.resources[name]
	return ok
}

func (l *StringLoader) Load(name string, locale string, encoding string) (httl.Resource, error) {
	key := getTemplateKey(name, locale, encoding)
	if r, ok := l.resources[key]; ok {
		return &r, nil
	}
	return nil, nil
}

func (l *StringLoader) Add(name string, locale string, encoding string) {
	l.resources[getTemplateKey(name, locale, encoding)] = resources.StringResource{}
}

func getTemplateKey(name, locale, encoding string) string {
	return ""
}
