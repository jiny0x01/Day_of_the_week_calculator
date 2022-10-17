package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYear(t *testing.T) {
	case1 := calcYear(2022)
	require.Equal(t, 5, case1)
	// 윤년 5번, 21+5 26 %
	case2 := calcYear(400)
	require.Equal(t, 5, case2)

	case3 := calcYear(1)
	require.Equal(t, 0, case3)

	case4 := calcYear(1601)
	require.Equal(t, 0, case4)

	case5 := calcYear(1901)
	require.Equal(t, 1, case5)

	case6 := calcYear(1935)
	require.Equal(t, 1, case6)

	case7 := calcYear(622)
	require.Equal(t, 1, case7)
}

func TestMonth(t *testing.T) {
	case1 := calcMonth(1)
	require.Equal(t, 0, case1)

	case2 := calcMonth(2)
	require.Equal(t, 3, case2)

	case3 := calcMonth(3)
	require.Equal(t, 3, case3)

	case4 := calcMonth(4)
	require.Equal(t, 6, case4)

	case5 := calcMonth(5)
	require.Equal(t, 1, case5)

	case6 := calcMonth(6)
	require.Equal(t, 4, case6)

	case7 := calcMonth(7)
	require.Equal(t, 6, case7)

	case8 := calcMonth(8)
	require.Equal(t, 2, case8)

	case9 := calcMonth(9)
	require.Equal(t, 5, case9)

	case10 := calcMonth(10)
	require.Equal(t, 0, case10)

	case11 := calcMonth(11)
	require.Equal(t, 3, case11)

	case12 := calcMonth(12)
	require.Equal(t, 5, case12)
}

func TestDay(t *testing.T) {
	case1 := calcDay(27)
	require.Equal(t, 6, case1)

	case2 := calcDay(1)
	require.Equal(t, 1, case2)

	case3 := calcDay(30)
	require.Equal(t, 2, case3)

	case4 := calcDay(31)
	require.Equal(t, 3, case4)

	case5 := calcDay(10)
	require.Equal(t, 3, case5)

	case6 := calcDay(14)
	require.Equal(t, 0, case6)

	case7 := calcDay(15)
	require.Equal(t, 1, case7)
}
