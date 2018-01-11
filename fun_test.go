package fun

import (
    "testing"
    //. "fun"
)

func TestIsAnyString(t *testing.T) {
    var strVal Any = "I am a string"
    if IsAnyString(strVal) != true {
        t.Errorf("IsAnyString must be true for \"%v\"", strVal)
    }

    var intVal Any = 7
    if IsAnyString(intVal) == true {
        t.Errorf("IsAnyString must be false for \"%v\"", intVal)
    }
}

func TestAssert(t *testing.T) {
    Assert(true == true, "true is true")
    defer func() {
        if r := recover(); r != nil {
            t.Log("Recovered ", r)
        }
    }()
    Assert(true == false, "true is false ... really?")
}

func TestIfExpr(t *testing.T) {

    // To make sure our closures capture outside variables.
    const num7 int = 7

    test_if := IfExpr(1 == 1,
    func() Any {
        return num7+1
    },func() Any {
        return num7+2
    })

    if test_if != num7+1 {
        t.Error ("IfExpr true branch error when condition is 1 == 1")
    }

    test_if = IfExpr(1 == 2,
    func() Any {
        return num7+1
    },func() Any {
        return num7+2
    })

    if test_if != num7+2 {
        t.Error ("IfExpr false branch error when condition is 1 == 2")
    }

}

func TestInt(t *testing.T) {
    if (Int(2) + 2) != 4 {
        t.Error("Int(2) + 2 != 4, Int does not support int builtin operators")
    }
}

func TestTimes(t *testing.T) {
    count := 0
    Int(5).Times(func(i int) {
        count++
    })

    if count != 5 {
        t.Error ("Int(n).Times does execute correct number of times")
    }

    count = 0
    Times(7, func(i int) {
        count = i
    })

    if count != 6 {
        t.Error ("Times(n,block) does execute correct number of times")
    }


}
