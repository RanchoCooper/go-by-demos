package main

import (
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/6
 */

func BenchmarkWithoutPool(b *testing.B) {
    var p *Person
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < 10000; j++ {
            p = new(Person)
            p.Age = 23
        }
    }
}

func BenchmarkWithPool(b *testing.B) {
    var p *Person
    b.ReportAllocs()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        for j := 0; j < 10000; j++ {
            p = personPool.Get().(*Person)
            p.Age = 23
            personPool.Put(p)
        }
    }
}
