/*
Copyright (c) 2018, Igor Lesik
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

The views and conclusions contained in the software and documentation are those
of the authors and should not be interpreted as representing official policies,
either expressed or implied, of the FreeBSD Project.
*/

/*
Package with functions implementing Functional idioms
and idioms from other languages.

Author: Igor Lesik 2018

To test run:
    go test -cover [-v] fun 
*/
package fun

import (
    "runtime"
    "fmt"
)

// Value of any type. TODO in go1.9 try type Any = interface{}
type Any interface{}

// Returns true if value:Any has type string.
func IsAnyString(val Any) bool {
    _, ok := val.(string)
    return ok
}

// Panics if cond is false.
func Assert(cond bool, msg string) {
    if cond == false {
        pc, fn, line, _ := runtime.Caller(1)
        panic(fmt.Sprintf("Assert in %s[%s:%d], \"%s\"",
            runtime.FuncForPC(pc).Name(), fn, line, msg))
    }
}

type Block func()
type ExprBlock func() Any

// "if" expression that returns a value depending on a condition.
func IfExpr(condition bool, ifBlock ExprBlock, elseBlock ExprBlock) Any {
    if condition == true {
        return ifBlock()
    }
    return elseBlock()
}

// Type Int extends basic type int.
type Int int

// Type UInt extends basic type uint.
type UInt uint

// Ruby-like n.times
func (times Int) Times(code func(i int)) {
    for index := 0; index < int(times); index++ {
        code(index)
    }
}

/*

Example:
    count := 0
    Times(7, func(i int) {
        count = i
    })
*/
func Times(times int, code func(i int)) {
    for i := 0; i < times; i++ {
       code(i)
    }
}
