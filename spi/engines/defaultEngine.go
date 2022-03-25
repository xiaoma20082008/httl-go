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

package engines

import (
	"github.com/xiaoma20082008/httl-go"
	"github.com/xiaoma20082008/httl-go/spi"
	"github.com/xiaoma20082008/httl-go/spi/loaders"
	"github.com/xiaoma20082008/httl-go/spi/resolvers"
	"github.com/xiaoma20082008/httl-go/utils"
)

const (
	DevelopConfigurationFile = "httl.properties"
	DefaultConfigurationFile = "httl-default.properties"
)

type defaultEngine struct {
	name            string
	defaultEncoding string
	defaultLocale   string
	version         string
	prefix          []string
	suffix          []string
	reloadable      bool
	preload         bool
	properties      map[string]any
	cache           spi.Cache
	compiler        spi.Compiler
	converter       spi.Converter
	filter          spi.Filter
	formatter       spi.Formatter
	interceptor     spi.Interceptor
	listener        spi.Listener
	loader          spi.Loader
	stringLoader    loaders.StringLoader
	logger          spi.Logger
	parser          spi.Parser
	resolver        spi.Resolver
	translator      spi.Translator

	httl.Engine
}

func (e *defaultEngine) GetName() string    { return e.name }
func (e *defaultEngine) GetVersion() string { return e.version }

func (e *defaultEngine) GetProperty(key string, dv string) string {
	v, ok := e.properties[key]
	if ok {
		return v.(string)
	}
	return dv
}
func (e *defaultEngine) GetAny(key string) any                { return nil }
func (e *defaultEngine) GetBool(key string, dv bool) bool     { return false }
func (e *defaultEngine) GetInt32(key string, dv string) int32 { return 0 }
func (e *defaultEngine) GetInt64(key string, dv string) int64 { return 0 }

func (e *defaultEngine) GetTemplate(name string, locale string, encoding string) (httl.Template, error) {
	if e.cache == nil {
		return e.parseTemplate(nil, name, locale, encoding)
	}
	var resource httl.Resource = nil
	if e.reloadable {
		resource, _ = e.GetResource(name, locale, encoding)
		return e.parseTemplate(&resource, name, locale, encoding)
	} else {
	}
	return nil, nil
}

func (e *defaultEngine) parseTemplate(r *httl.Resource, name string, locale string, encoding string) (httl.Template, error) {
	var resource = r
	if resource == nil {
		rr, _ := e.GetResource(name, locale, encoding)
		resource = &rr
	}
	source, _ := (*resource).Source()
	if e.filter != nil {
		source = e.filter.Filter(name, source)
	}
	node, _ := e.parser.Parse(source)
	return e.translator.Translate(resource, &node)
}

func (e *defaultEngine) GetResource(name string, locale string, encoding string) (httl.Resource, error) {
	if e.stringLoader.Exists(name, locale, encoding) {
		return e.stringLoader.Load(name, locale, encoding)
	}
	return e.loader.Load(name, locale, encoding)
}

func (e *defaultEngine) NewContext(parent map[string]any, current map[string]any) map[string]any {
	res := make(map[string]any, len(parent)+len(current))
	for k, v := range parent {
		res[k] = v
	}
	for k, v := range current {
		res[k] = v
	}
	return res
}

func (e *defaultEngine) initialize() error {
	e.resolver = resolvers.NewMultiResolver()
	e.resolver = resolvers.NewGlobalResolver()
	e.resolver = resolvers.NewSessionResolver()
	e.resolver = resolvers.NewEngineResolver()
	return nil
}

func mergeProperties(d, c map[string]string) map[string]string {
	if d == nil {
		return c
	}
	if c == nil {
		return d
	}
	if d == nil && c == nil {
		return map[string]string{}
	}

	res := make(map[string]string, len(d)+len(c))

	for k, v := range d {
		res[k] = v
	}
	for k, v := range c {
		res[k] = v
	}
	return res
}

func NewDefaultEngine() httl.Engine {
	defaultConfiguration, _ := utils.LoadProperties(DefaultConfigurationFile)
	currentConfiguration, _ := utils.LoadProperties(DevelopConfigurationFile)
	configuration := mergeProperties(defaultConfiguration, currentConfiguration)
	if len(configuration) == 0 {
		panic("")
	}
	engine := &defaultEngine{
		properties: map[string]any{},
	}
	err := engine.initialize()
	if err != nil {
		panic(err)
	}
	return engine
}
