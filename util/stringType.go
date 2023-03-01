package util

import (
	"fmt"
	"strconv"
	"strings"
)

// 除法计算，保留两位小数
func StrconvFloat(claNum1 int64, claNum2 int64) (value float64) {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(claNum1)/float64(claNum2)), 64)
	return value
}

type StringTypeInt struct {
	Int64   int64
	Int32   int32
	Float64 float64
	Float32 float32
	String  string
	Bool    bool
}

func (s *StringTypeInt) StringRepliceRes(srcText string, dstType string, fmtType string) error {

	switch {
	case dstType == "string" && fmtType == "/": //  StringReplice
		dstText := srcText[strings.LastIndex(srcText, "/")+1:]
		s.String = dstText

	case dstType == "float64" && fmtType == "(": // StringRepliceRes01
		resInt, err := strconv.ParseFloat(srcText[:strings.Index(srcText, "(")], 64)
		if err != nil {
			fmt.Println(err)
		}
		s.Float64 = resInt
	case dstType == "float64" && fmtType == "()": // StringRepliceRes00
		srcindextext := strings.Index(srcText, "(") + 1
		dstindextext := strings.Index(srcText, ")")
		resInt, err := strconv.ParseFloat(srcText[srcindextext:dstindextext], 64)
		if err != nil {
			fmt.Println(err)
		}
		s.Float64 = resInt
	case dstType == "int64" && fmtType == "": // StringRepliceRes03
		resInt, err := strconv.ParseInt(srcText, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		s.Int64 = resInt
	case dstType == "int64" && fmtType == "(": // StringRepliceRes02
		dstindextext := strings.Index(srcText, "(")
		if dstindextext < 1 {
			return nil
		}
		resInt, err := strconv.ParseInt(srcText[:dstindextext], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		s.Int64 = resInt
	}
	return nil
}

// string replice
func StringReplice(srcText string, Stype string) (dstText string) {
	// xxx/A
	if Stype == "A" {
		indextext := strings.LastIndex(srcText, "/")
		dstText = srcText[:indextext]
		// xxx/B#xxx
	} else if Stype == "B" {
		startindextext := strings.LastIndex(srcText, "/") + 1
		endindextext := strings.LastIndex(srcText, "#")
		dstText = srcText[startindextext:endindextext]
		// xxx#C
	} else if Stype == "C" {
		indextext := strings.LastIndex(srcText, "#") + 1
		dstText = srcText[indextext:]
		// D/xxx
	} else if Stype == "D" {
		indextext := strings.LastIndex(srcText, "/") + 1
		dstText = srcText[indextext:]
	}
	return
}

func StringRepliceRes00(srcText string) (resInt float64) {
	srcindextext := strings.Index(srcText, "(") + 1
	dstindextext := strings.Index(srcText, ")")
	restext := srcText[srcindextext:dstindextext]
	resInt, err := strconv.ParseFloat(restext, 64)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func StringRepliceRes01(srcText string) (resInt float64) {
	dstindextext := strings.Index(srcText, "(")
	restext := srcText[:dstindextext]
	resInt, err := strconv.ParseFloat(restext, 64)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func StringRepliceRes02(srcText string) (resInt int64) {
	dstindextext := strings.Index(srcText, "(")
	if dstindextext < 1 {
		return
	}
	restext := srcText[:dstindextext]
	resInt, err := strconv.ParseInt(restext, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func StringRepliceRes03(srcText string) (resInt int64) {
	resInt, err := strconv.ParseInt(srcText, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return
}
