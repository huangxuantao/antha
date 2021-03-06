package data

// Code generated by gen.py. DO NOT EDIT.

import "github.com/pkg/errors"

// newNativeCompareFunc creates a function to compare elements of the given native Series.
func newNativeCompareFunc(nativeSeries *Series, asc bool) (compareFunc, error) {
	meta, ok := nativeSeries.meta.(*nativeSeriesMeta)
	if !ok {
		panic(errors.Errorf("series %+v is not native", nativeSeries))
	}

	// TODO: more optimal comparators for non-nullable columns
	switch nativeSeries.typ {
	case typeFloat64:
		return newNativeCompareFuncFloat64(meta, asc), nil
	case typeInt64:
		return newNativeCompareFuncInt64(meta, asc), nil
	case typeInt:
		return newNativeCompareFuncInt(meta, asc), nil
	case typeString:
		return newNativeCompareFuncString(meta, asc), nil
	case typeBool:
		return newNativeCompareFuncBool(meta, asc), nil
	case typeTimestampMillis:
		return newNativeCompareFuncTimestampMillis(meta, asc), nil
	case typeTimestampMicros:
		return newNativeCompareFuncTimestampMicros(meta, asc), nil
	default:
		// Currently we don't have a generic native series compare function.
		// However, it is possible to write a reflective one - at least for certain Kinds.
		return nil, errors.Errorf("The data type %+v is not supported, expecting a series of some supported primitive type", nativeSeries.typ)
	}
}

// newNativeSwapFunc creates a function to swap elements of the given native Series.
func newNativeSwapFunc(nativeSeries *Series) swapFunc {
	meta, ok := nativeSeries.meta.(*nativeSeriesMeta)
	if !ok {
		panic(errors.Errorf("series %+v is not native", nativeSeries))
	}

	switch nativeSeries.typ {
	case typeFloat64:
		return newNativeSwapFuncFloat64(meta)
	case typeInt64:
		return newNativeSwapFuncInt64(meta)
	case typeInt:
		return newNativeSwapFuncInt(meta)
	case typeString:
		return newNativeSwapFuncString(meta)
	case typeBool:
		return newNativeSwapFuncBool(meta)
	case typeTimestampMillis:
		return newNativeSwapFuncTimestampMillis(meta)
	case typeTimestampMicros:
		return newNativeSwapFuncTimestampMicros(meta)
	default:
		// a fallback swap func generator (very slow!)
		return newNativeSwapFuncGeneric(meta)
	}
}

// float64

func newNativeCompareFuncFloat64(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]float64)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareFloat64(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareFloat64(val1 float64, notNull1 bool, val2 float64, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareFloat64(val1, val2)
	}
	return applyAsc(result, asc)
}

func rawCompareFloat64(val1, val2 float64) int {
	switch {
	case val1 < val2:
		return -1
	case val1 > val2:
		return 1
	default:
		return 0
	}
}

func newNativeSwapFuncFloat64(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]float64)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}

// int64

func newNativeCompareFuncInt64(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]int64)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareInt64(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareInt64(val1 int64, notNull1 bool, val2 int64, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareInt64(val1, val2)
	}
	return applyAsc(result, asc)
}

func rawCompareInt64(val1, val2 int64) int {
	switch {
	case val1 < val2:
		return -1
	case val1 > val2:
		return 1
	default:
		return 0
	}
}

func newNativeSwapFuncInt64(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]int64)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}

// int

func newNativeCompareFuncInt(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]int)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareInt(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareInt(val1 int, notNull1 bool, val2 int, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareInt(val1, val2)
	}
	return applyAsc(result, asc)
}

func rawCompareInt(val1, val2 int) int {
	switch {
	case val1 < val2:
		return -1
	case val1 > val2:
		return 1
	default:
		return 0
	}
}

func newNativeSwapFuncInt(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]int)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}

// string

func newNativeCompareFuncString(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]string)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareString(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareString(val1 string, notNull1 bool, val2 string, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareString(val1, val2)
	}
	return applyAsc(result, asc)
}

func newNativeSwapFuncString(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]string)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}

// bool

func newNativeCompareFuncBool(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]bool)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareBool(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareBool(val1 bool, notNull1 bool, val2 bool, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareBool(val1, val2)
	}
	return applyAsc(result, asc)
}

func newNativeSwapFuncBool(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]bool)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}

// TimestampMillis

func newNativeCompareFuncTimestampMillis(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]TimestampMillis)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareTimestampMillis(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareTimestampMillis(val1 TimestampMillis, notNull1 bool, val2 TimestampMillis, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareTimestampMillis(val1, val2)
	}
	return applyAsc(result, asc)
}

func rawCompareTimestampMillis(val1, val2 TimestampMillis) int {
	switch {
	case val1 < val2:
		return -1
	case val1 > val2:
		return 1
	default:
		return 0
	}
}

func newNativeSwapFuncTimestampMillis(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]TimestampMillis)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}

// TimestampMicros

func newNativeCompareFuncTimestampMicros(nativeMeta *nativeSeriesMeta, asc bool) compareFunc {
	data := nativeMeta.rValue.Interface().([]TimestampMicros)
	notNull := nativeMeta.notNull

	return func(i, j int) int {
		return compareTimestampMicros(data[i], notNull[i], data[j], notNull[j], asc)
	}
}

func compareTimestampMicros(val1 TimestampMicros, notNull1 bool, val2 TimestampMicros, notNull2 bool, asc bool) int {
	result, ok := compareNulls(notNull1, notNull2)
	if !ok {
		result = rawCompareTimestampMicros(val1, val2)
	}
	return applyAsc(result, asc)
}

func rawCompareTimestampMicros(val1, val2 TimestampMicros) int {
	switch {
	case val1 < val2:
		return -1
	case val1 > val2:
		return 1
	default:
		return 0
	}
}

func newNativeSwapFuncTimestampMicros(nativeMeta *nativeSeriesMeta) swapFunc {
	data := nativeMeta.rValue.Interface().([]TimestampMicros)
	notNull := nativeMeta.notNull
	return func(i, j int) {
		data[i], data[j] = data[j], data[i]
		notNull[i], notNull[j] = notNull[j], notNull[i]
	}
}
