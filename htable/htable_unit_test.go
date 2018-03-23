
package htable_test

import (
    "testing"
    "gitee.com/johng/hashtable-go/htable"
)

var table = htable.New(2)

func Test_Set(t *testing.T) {
    for i := 0; i < 10; i++ {
        table.Set(i, i*10)
    }
}

func Test_Get(t *testing.T) {
    for i := 0; i < 10; i++ {
        if r := table.Get(i); r != i*10 {
            t.Errorf("Invalid result:%v, for key:%v", r, i)
        }
    }
}
