package funclist

import (
	mybasic "go-samples/samples/basic"
	mycsv "go-samples/samples/csv"
	myerrors "go-samples/samples/errors"
	myfilepath "go-samples/samples/filepath"
	myjson "go-samples/samples/json"
	mylog "go-samples/samples/log"
	myos "go-samples/samples/os"
	mystrconv "go-samples/samples/strconv"
	mystrings "go-samples/samples/strings"
	mytime "go-samples/samples/time"
)

func GetFunc(name string) []func() {
	basicfunc := []func(){
		mybasic.BasicSample001,
		mybasic.BasicSample002,
		mybasic.BasicSample003,
		mybasic.BasicSample004,
		mybasic.BasicSample005,
		mybasic.BasicSample006,
		mybasic.BasicSample007,
		mybasic.BasicSample008,
		mybasic.BasicSample009,
		mybasic.BasicSample010,
		mybasic.BasicSample011,
		mybasic.BasicSample012,
		mybasic.BasicSample013,
		mybasic.BasicSample014,
		mybasic.BasicSample015,
		mybasic.BasicSample016,
		mybasic.BasicSample017,
		mybasic.BasicSample018,
		mybasic.BasicSample019,
		mybasic.BasicSample020,
		mybasic.BasicSample021,
		mybasic.BasicSample022,
		mybasic.BasicSample023,
		mybasic.BasicSample024,
		mybasic.BasicSample025,
		mybasic.BasicSample026,
		mybasic.BasicSample027,
		mybasic.BasicSample028,
		mybasic.BasicSample029,
		mybasic.BasicSample030,
	}
	csvfunc := []func(){
		mycsv.CsvSample001,
		mycsv.CsvSample002,
		mycsv.CsvSample003,
		mycsv.CsvSample004,
	}
	errorsfunc := []func(){
		myerrors.ErrorsSample001,
		myerrors.ErrorsSample002,
		myerrors.ErrorsSample003,
	}
	filepathfunc := []func(){
		myfilepath.FilepathSample001,
		myfilepath.FilepathSample002,
		myfilepath.FilepathSample003,
		myfilepath.FilepathSample004,
		myfilepath.FilepathSample005,
		myfilepath.FilepathSample006,
		myfilepath.FilepathSample007,
		myfilepath.FilepathSample008,
		myfilepath.FilepathSample009,
		myfilepath.FilepathSample010,
	}
	jsonfunc := []func(){
		myjson.JsonSample001,
		myjson.JsonSample002,
		myjson.JsonSample003,
		myjson.JsonSample004,
	}
	logfunc := []func(){
		mylog.LogSample001,
	}
	osfunc := []func(){
		myos.OsSample001,
		myos.OsSample002,
		myos.OsSample003,
		myos.OsSample004,
		myos.OsSample005,
		myos.OsSample006,
		myos.OsSample007,
		myos.OsSample008,
		myos.OsSample009,
		myos.OsSample010,
		myos.OsSample011,
		myos.OsSample012,
	}
	strconvfunc := []func(){
		mystrconv.StrconvSample001,
		mystrconv.StrconvSample002,
		mystrconv.StrconvSample003,
		mystrconv.StrconvSample004,
		mystrconv.StrconvSample005,
		mystrconv.StrconvSample006,
	}
	stringsfunc := []func(){
		mystrings.StringsSample001,
		mystrings.StringsSample002,
		mystrings.StringsSample003,
		mystrings.StringsSample004,
		mystrings.StringsSample005,
		mystrings.StringsSample006,
		mystrings.StringsSample007,
		mystrings.StringsSample008,
		mystrings.StringsSample009,
		mystrings.StringsSample010,
		mystrings.StringsSample011,
		mystrings.StringsSample012,
		mystrings.StringsSample013,
		mystrings.StringsSample014,
		mystrings.StringsSample015,
		mystrings.StringsSample016,
		mystrings.StringsSample017,
		mystrings.StringsSample018,
		mystrings.StringsSample019,
		mystrings.StringsSample020,
		mystrings.StringsSample021,
		mystrings.StringsSample022,
		mystrings.StringsSample023,
		mystrings.StringsSample024,
		mystrings.StringsSample025,
		mystrings.StringsSample026,
		mystrings.StringsSample027,
		mystrings.StringsSample028,
		mystrings.StringsSample029,
	}
	timefunc := []func(){
		mytime.TimeSample001,
		mytime.TimeSample002,
		mytime.TimeSample003,
		mytime.TimeSample004,
		mytime.TimeSample005,
		mytime.TimeSample006,
		mytime.TimeSample007,
		mytime.TimeSample008,
		mytime.TimeSample009,
		mytime.TimeSample010,
		mytime.TimeSample011,
		mytime.TimeSample012,
		mytime.TimeSample013,
		mytime.TimeSample014,
		mytime.TimeSample015,
	}

	switch name {
	case "basic":
		return basicfunc
	case "csv":
		return csvfunc
	case "errors":
		return errorsfunc
	case "filepath":
		return filepathfunc
	case "json":
		return jsonfunc
	case "log":
		return logfunc
	case "os":
		return osfunc
	case "strconv":
		return strconvfunc
	case "strings":
		return stringsfunc
	case "time":
		return timefunc
	default:
		return nil
	}
}
