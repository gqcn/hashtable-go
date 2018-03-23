// go test *.go -bench=".*" -benchmem

package htable_test

import (
    "testing"
    "gitee.com/johng/hashtable-go/htable"
)

var hta  = htable.New(8)
var gom  = make(map[int]interface{})

func BenchmarkHtable_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        hta.Set(i, i)
    }
}

func BenchmarkHtable_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        hta.Get(i)
    }
}
//
//func BenchmarkDrh_Remove(b *testing.B) {
//    for i := 0; i < b.N; i++ {
//        hta.Remove(i)
//    }
//}

func BenchmarkGoMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gom[i] = i
    }
}

func BenchmarkGoMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        if _, ok := gom[i]; ok {}
    }
}

//func BenchmarkGoMap_Remove(b *testing.B) {
//    for i := 0; i < b.N; i++ {
//        delete(gom, i)
//    }
//}