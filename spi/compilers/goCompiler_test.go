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

package compilers

import (
	"github.com/xiaoma20082008/httl-go/utils"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func TestNewGoCompiler(t *testing.T) {
	compiler := NewGoCompiler()
	code, e := utils.ReadFile("./testdata/aboutCompiledTemplate.txt")
	if e != nil {
		panic("read file failed")
	}
	ast, e := compiler.Compile(code)
	if e != nil {
		panic("compile code failed")
	}
	e = printer.Fprint(os.Stdout, token.NewFileSet(), ast)
	if e != nil {
		panic("print ast failed")
	}
}
