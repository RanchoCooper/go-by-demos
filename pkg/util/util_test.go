package util

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

/**
 * @author Rancho
 * @date 2021/12/15
 */

func TestRandPickWithExcept(t *testing.T) {
    r1 := RandPickWithExcept(10, 5, nil)
    assert.Equal(t, 5, len(*r1))

    r2 := RandPickWithExcept(10, 5, []int{1, 2, 3, 4, 5})
    assert.Equal(t, []int{0, 6, 7, 8, 9}, *r2)

    r3 := RandPickWithExcept(5, 5, []int{1, 2, 3, 4, 5})
    assert.Nil(t, r3)
}

func TestHasOverlap(t *testing.T) {
    assert.False(t, HasOverlap(nil, nil))
    assert.False(t, HasOverlap([]int{}, []int{}))
    assert.False(t, HasOverlap([]int{1, 1, 2, 2, 3, 3}, []int{4, 4, 5, 5, 6, 6}))
    assert.True(t, HasOverlap([]int{1, 1, 2, 2, 3, 3}, []int{4, 4, 5, 5, 6, 6, 3}))
}
