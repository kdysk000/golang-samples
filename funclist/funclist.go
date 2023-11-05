package funclist

import (
	myarray "go-samples/samples/basic/array"
	myfor "go-samples/samples/basic/for"
	myfunc "go-samples/samples/basic/func"
	mygoroutine "go-samples/samples/basic/goroutine"
	myif "go-samples/samples/basic/if"
	myinterface "go-samples/samples/basic/interface"
	mymap "go-samples/samples/basic/map"
	myslice "go-samples/samples/basic/slice"
	mystruct "go-samples/samples/basic/struct"
	myswitch "go-samples/samples/basic/switch"
	mycsv "go-samples/samples/csv"
	myerrors "go-samples/samples/errors"
	myfilepath "go-samples/samples/filepath"
	myjson "go-samples/samples/json"
	mylog "go-samples/samples/log"
	myos "go-samples/samples/os"
	myslices "go-samples/samples/slices"
	mystrconv "go-samples/samples/strconv"
	mystrings "go-samples/samples/strings"
	mytime "go-samples/samples/time"
)

func GetFunc(name string) []func() {
	arrayfunc := []func(){
		myarray.ArraySample001,
		myarray.ArraySample002,
		myarray.ArraySample003,
		myarray.ArraySample004,
		myarray.ArraySample005,
		myarray.ArraySample006,
		myarray.ArraySample007,
		myarray.ArraySample008,
		myarray.ArraySample009,
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
	forfunc := []func(){
		myfor.ForSample001,
		myfor.ForSample002,
		myfor.ForSample003,
		myfor.ForSample004,
		myfor.ForSample005,
		myfor.ForSample006,
	}
	funcfunc := []func(){
		myfunc.FuncSample001,
		myfunc.FuncSample002,
		myfunc.FuncSample003,
		myfunc.FuncSample004,
		myfunc.FuncSample005,
		myfunc.FuncSample006,
		myfunc.FuncSample007,
		myfunc.FuncSample008,
		myfunc.FuncSample009,
		myfunc.FuncSample010,
		myfunc.FuncSample011,
		myfunc.FuncSample012,
		myfunc.FuncSample013,
	}
	goroutinefunc := []func(){
		mygoroutine.GoroutineSample001,
		mygoroutine.GoroutineSample002,
		mygoroutine.GoroutineSample003,
		mygoroutine.GoroutineSample004,
	}
	iffunc := []func(){
		myif.IfSample001,
		myif.IfSample002,
	}
	interfacefunc := []func(){
		myinterface.InterfaceSample001,
		myinterface.InterfaceSample002,
		myinterface.InterfaceSample003,
		myinterface.InterfaceSample004,
		myinterface.InterfaceSample005,
	}
	jsonfunc := []func(){
		myjson.JsonSample001,
		myjson.JsonSample002,
		myjson.JsonSample003,
		myjson.JsonSample004,
	}
	logfunc := []func(){
		mylog.LogSample001,
		mylog.LogSample002,
		mylog.LogSample003,
		mylog.LogSample004,
		mylog.LogSample005,
	}
	mapfunc := []func(){
		mymap.MapSample001,
		mymap.MapSample002,
		mymap.MapSample003,
		mymap.MapSample004,
		mymap.MapSample005,
		mymap.MapSample006,
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
		myos.OsSample013,
		myos.OsSample014,
		myos.OsSample015,
		myos.OsSample016,
		myos.OsSample017,
		myos.OsSample018,
		myos.OsSample019,
		myos.OsSample020,
		myos.OsSample021,
		myos.OsSample022,
	}
	slicefunc := []func(){
		myslice.SliceSample001,
		myslice.SliceSample002,
		myslice.SliceSample003,
		myslice.SliceSample004,
		myslice.SliceSample005,
		myslice.SliceSample006,
		myslice.SliceSample007,
		myslice.SliceSample008,
		myslice.SliceSample009,
	}
	slicesfunc := []func(){
		myslices.SlicesSample001,
		myslices.SlicesSample002,
		myslices.SlicesSample003,
		myslices.SlicesSample004,
		myslices.SlicesSample005,
		myslices.SlicesSample006,
		myslices.SlicesSample007,
		myslices.SlicesSample008,
		myslices.SlicesSample009,
		myslices.SlicesSample010,
		myslices.SlicesSample011,
		myslices.SlicesSample012,
		myslices.SlicesSample013,
		myslices.SlicesSample014,
		myslices.SlicesSample015,
		myslices.SlicesSample016,
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
	structfunc := []func(){
		mystruct.StructSample001,
		mystruct.StructSample002,
		mystruct.StructSample003,
		mystruct.StructSample004,
		mystruct.StructSample005,
		mystruct.StructSample006,
	}
	switchfunc := []func(){
		myswitch.SwitchSample001,
		myswitch.SwitchSample002,
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
		mytime.TimeSample016,
		mytime.TimeSample017,
		mytime.TimeSample018,
		mytime.TimeSample019,
		mytime.TimeSample020,
		mytime.TimeSample021,
		mytime.TimeSample022,
	}

	switch name {
	case "array":
		return arrayfunc
	case "csv":
		return csvfunc
	case "errors":
		return errorsfunc
	case "filepath":
		return filepathfunc
	case "for":
		return forfunc
	case "func":
		return funcfunc
	case "goroutine":
		return goroutinefunc
	case "if":
		return iffunc
	case "interface":
		return interfacefunc
	case "json":
		return jsonfunc
	case "log":
		return logfunc
	case "map":
		return mapfunc
	case "os":
		return osfunc
	case "slice":
		return slicefunc
	case "slices":
		return slicesfunc
	case "strconv":
		return strconvfunc
	case "strings":
		return stringsfunc
	case "struct":
		return structfunc
	case "switch":
		return switchfunc
	case "time":
		return timefunc
	default:
		return nil
	}
}
