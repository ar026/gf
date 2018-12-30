// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// Package gtest provides simple and useful test utils/测试模块.
package gtest

import (
    "fmt"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/util/gconv"
    "os"
    "reflect"
)

// 封装一个测试用例
func Case(f func()) {
    defer func() {
        if err := recover(); err != nil {
            glog.To(os.Stderr).Println(err)
            glog.Header(false).PrintBacktrace(4)
        }
    }()
    f()
}

// 断言判断, 相等
func Assert(value, expect interface{}) {
    rv := reflect.ValueOf(value)
    if rv.Kind() == reflect.Ptr {
        if rv.IsNil() {
            value = nil
        }
    }
    if value != expect {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, value, expect))
    }
}

// 断言判断, 相等
func AssertEQ(value, expect interface{}) {
    rv := reflect.ValueOf(value)
    if rv.Kind() == reflect.Ptr {
        if rv.IsNil() {
            value = nil
        }
    }
    if value != expect {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v == %v`, value, expect))
    }
}

// 断言判断, 不相等
func AssertNE(value, expect interface{}) {
    rv := reflect.ValueOf(value)
    if rv.Kind() == reflect.Ptr {
        if rv.IsNil() {
            value = nil
        }
    }
    if value == expect {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v != %v`, value, expect))
    }
}

// 断言判断, value > expect; 注意: 仅有字符串、整形、浮点型才可以比较
func AssertGT(value, expect interface{}) {
    passed := false
    switch reflect.ValueOf(expect).Kind() {
        case reflect.String:
            passed = gconv.String(value) > gconv.String(expect)

        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            passed = gconv.Int(value) > gconv.Int(expect)

        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            passed = gconv.Uint(value) > gconv.Uint(expect)

        case reflect.Float32, reflect.Float64:
            passed = gconv.Float64(value) > gconv.Float64(expect)
    }
    if !passed {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v > %v`, value, expect))
    }
}

// 断言判断, value >= expect; 注意: 仅有字符串、整形、浮点型才可以比较
func AssertGTE(value, expect interface{}) {
    passed := false
    switch reflect.ValueOf(expect).Kind() {
        case reflect.String:
            passed = gconv.String(value) >= gconv.String(expect)

        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            passed = gconv.Int(value) >= gconv.Int(expect)

        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            passed = gconv.Uint(value) >= gconv.Uint(expect)

        case reflect.Float32, reflect.Float64:
            passed = gconv.Float64(value) >= gconv.Float64(expect)
    }
    if !passed {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v >= %v`, value, expect))
    }
}

// 断言判断, value < expect; 注意: 仅有字符串、整形、浮点型才可以比较
func AssertLT(value, expect interface{}) {
    passed := false
    switch reflect.ValueOf(expect).Kind() {
        case reflect.String:
            passed = gconv.String(value) < gconv.String(expect)

        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            passed = gconv.Int(value) < gconv.Int(expect)

        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            passed = gconv.Uint(value) < gconv.Uint(expect)

        case reflect.Float32, reflect.Float64:
            passed = gconv.Float64(value) < gconv.Float64(expect)
    }
    if !passed {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v < %v`, value, expect))
    }
}

// 断言判断, value <= expect; 注意: 仅有字符串、整形、浮点型才可以比较
func AssertLTE(value, expect interface{}) {
    passed := false
    switch reflect.ValueOf(expect).Kind() {
        case reflect.String:
            passed = gconv.String(value) <= gconv.String(expect)

        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            passed = gconv.Int(value) <= gconv.Int(expect)

        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            passed = gconv.Uint(value) <= gconv.Uint(expect)

        case reflect.Float32, reflect.Float64:
            passed = gconv.Float64(value) <= gconv.Float64(expect)
    }
    if !passed {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v <= %v`, value, expect))
    }
}


// 断言判断, value IN expect; 注意: expect必须为slice类型
func AssertIN(value, expect interface{}) {
    passed := false
    switch reflect.ValueOf(expect).Kind() {
        case reflect.Slice, reflect.Array:
            for _, v := range gconv.Interfaces(expect) {
                if v == value {
                    passed = true
                    break
                }
            }
    }
    if !passed {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v IN %v`, value, expect))
    }
}

// 断言判断, value NOT IN expect; 注意: expect必须为slice类型
func AssertNI(value, expect interface{}) {
    passed := false
    switch reflect.ValueOf(expect).Kind() {
        case reflect.Slice, reflect.Array:
            for _, v := range gconv.Interfaces(expect) {
                if v == value {
                    passed = true
                    break
                }
            }
    }
    if passed {
        panic(fmt.Sprintf(`[ASSERT] EXPECT %v NOT IN %v`, value, expect))
    }
}

// 提示错误不退出
func Error(message...interface{}) {
    glog.To(os.Stderr).Println(`[ERROR]`, fmt.Sprint(message...))
    glog.Header(false).PrintBacktrace(1)
}

// 提示错误并退出
func Fatal(message...interface{}) {
    glog.To(os.Stderr).Println(`[FATAL]`, fmt.Sprint(message...))
    glog.Header(false).PrintBacktrace(1)
    os.Exit(1)
}