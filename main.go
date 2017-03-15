package main

import (
	"errors"
	"log"
	"reflect"

	"strconv"

	"fmt"

	"go/types"
	"unsafe"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/tools/go/pointer"
)

type ParentStruct struct {
	P_A string
	P_B int
	P_C bool
	P_D ChildStructA
	P_E *string
	P_F *int
	P_G *ChildStructB
	P_H *wrappers.UInt32Value
}

type ChildStructA struct {
	CA_A string
	CA_B int
	CA_C bool
	CA_D GrandsonStruct
}

type ChildStructB struct {
	CB_A string
}

type GrandsonStruct struct {
	G_A string
	G_B int
	G_C bool
}

var m = map[interface{}]interface{}{
	"P_A":  "ParentStruct P_A string",
	"P_B":  "1",
	"P_C":  "true",
	"CA_A": "ChildStruct CA_A string",
	"G_A":  "GrandsonStruct G_A string",
	"P_F":  "432",
	"P_H":  "56",
	"CB_A": "*ChildStructB CB_A string",
}

type wkt interface {
	XXX_WellKnownType() string
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.Lshortfile)
	log.Println("map is ", m)

	st := ParentStruct{}
	m["P_E"] = "ParentStruct P_E *string"

	log.Println("[before] struct is ", &st)
	if err := MapToStruct(m, &st); err != nil {
		log.Println("[err] struct is ", &st)
		log.Fatalln(err)
	}
	log.Println("[after] struct is ", &st)
	log.Println("pointer struct is ", *st.P_E, *st.P_F, *st.P_H)
	log.Println("[[        Success        ]]")
}

func protoMessageToValue(p proto.Message, str string) (reflect.Value, error) {

	if wkt, ok := p.(wkt); ok {
		switch wkt.XXX_WellKnownType() {
		case "DoubleValue":
			x, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.DoubleValue{Value: float64(x)}), nil
		case "FloatValue":
			x, err := strconv.ParseFloat(str, 32)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.FloatValue{Value: float32(x)}), nil
		case "Int32Value":
			x, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.Int32Value{Value: int32(x)}), nil
		case "Int64Value":
			x, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.Int64Value{Value: int64(x)}), nil
		case "UInt32Value":
			x, err := strconv.ParseUint(str, 10, 32)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.UInt32Value{Value: uint32(x)}), nil
		case "UInt64Value":
			x, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.UInt64Value{Value: uint64(x)}), nil
		case "BoolValue":
			x, err := strconv.ParseBool(str)
			if err != nil {
				return reflect.Value{}, err
			}
			return reflect.ValueOf(&wrappers.BoolValue{Value: bool(x)}), nil
		case "StringValue":
			return reflect.ValueOf(&wrappers.StringValue{Value: str}), nil
		case "BytesValue":
			return reflect.ValueOf(&wrappers.BytesValue{Value: []byte(str)}), nil
		}
	}
	return reflect.Value{}, errors.New("Not proto.Message")

}

func ptrToValue(typ reflect.Type, in interface{}) (reflect.Value, error) {

	i, err := strConvType(in, typ)
	if err != nil {
		return reflect.Value{}, err
	}
	switch typ.Kind() {
	case reflect.Invalid:
	case reflect.Bool:
		value := i.(bool)
		return reflect.ValueOf(&value), nil
	case reflect.Int:
		value := i.(int)
		return reflect.ValueOf(&value), nil
	case reflect.Int8:
		value := i.(int8)
		return reflect.ValueOf(&value), nil
	case reflect.Int16:
		value := i.(int16)
		return reflect.ValueOf(&value), nil
	case reflect.Int32:
		value := i.(int32)
		return reflect.ValueOf(&value), nil
	case reflect.Int64:
		value := i.(int64)
		return reflect.ValueOf(&value), nil
	case reflect.Uint:
		value := i.(uint)
		return reflect.ValueOf(&value), nil
	case reflect.Uint8:
		value := i.(uint8)
		return reflect.ValueOf(&value), nil
	case reflect.Uint16:
		value := i.(uint16)
		return reflect.ValueOf(&value), nil
	case reflect.Uint32:
		value := i.(uint32)
		return reflect.ValueOf(&value), nil
	case reflect.Uint64:
		value := i.(uint64)
		return reflect.ValueOf(&value), nil
	case reflect.Uintptr:
		value := i.(uintptr)
		return reflect.ValueOf(&value), nil
	case reflect.Float32:
		value := i.(float32)
		return reflect.ValueOf(&value), nil
	case reflect.Float64:
		value := i.(float64)
		return reflect.ValueOf(&value), nil
	case reflect.Complex64:
		value := i.(complex64)
		return reflect.ValueOf(&value), nil
	case reflect.Complex128:
		value := i.(complex128)
		return reflect.ValueOf(&value), nil
	case reflect.Array:
		value := i.(types.Array)
		return reflect.ValueOf(&value), nil
	case reflect.Chan:
		value := i.(types.Chan)
		return reflect.ValueOf(&value), nil
	case reflect.Func:
		value := i.(types.Func)
		return reflect.ValueOf(&value), nil
	case reflect.Interface:
		value := i.(types.Interface)
		return reflect.ValueOf(&value), nil
	case reflect.Map:
		value := i.(types.Map)
		return reflect.ValueOf(&value), nil
	case reflect.Ptr:
		value := i.(pointer.Pointer)
		return reflect.ValueOf(&value), nil
	case reflect.Slice:
		value := i.(types.Slice)
		return reflect.ValueOf(&value), nil
	case reflect.String:
		value := i.(string)
		return reflect.ValueOf(&value), nil
	case reflect.Struct:
		value := i.(types.Struct)
		return reflect.ValueOf(&value), nil
	case reflect.UnsafePointer:
		value := i.(unsafe.Pointer)
		return reflect.ValueOf(&value), nil
	default:

	}
	return reflect.Value{}, errors.New("Not")
}

func MapToStruct(m map[interface{}]interface{}, st interface{}) error {

	for mapKey, mapValue := range m {
		refValue, noFiled := findFieldFromInterface(mapKey.(string), reflect.Indirect(reflect.ValueOf(st)))
		if noFiled != nil {
			continue
		}

		if refValue.FieldByName(mapKey.(string)).Kind() == reflect.Ptr {
			// google protobuf
			if proMes, ok := refValue.FieldByName(mapKey.(string)).Interface().(proto.Message); ok {
				rv, err := protoMessageToValue(proMes, mapValue.(string))
				if err != nil || !rv.IsValid() {
					return err
				}
				refValue.FieldByName(mapKey.(string)).Set(rv)
				continue
			}

			rv, err := ptrToValue(refValue.FieldByName(mapKey.(string)).Type().Elem(), mapValue)
			if err != nil || !rv.IsValid() {
				return err
			}
			refValue.FieldByName(mapKey.(string)).Set(rv)
			continue
		}

		rv, err := strConvType(mapValue, refValue.FieldByName(mapKey.(string)).Type())
		if err != nil {
			return err
		}
		refValue.FieldByName(mapKey.(string)).Set(reflect.ValueOf(rv))
		continue
	}
	return nil
}

func findFieldFromInterface(k string, v reflect.Value) (reflect.Value, error) {
	if v.FieldByName(k).CanSet() {
		return v, nil
	}

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Ptr:
			if v.Field(i).Type().Elem().Kind() == reflect.Struct {
				v.Field(i).Set(reflect.New(v.Field(i).Type().Elem()))
				if v, err := findFieldFromInterface(k, v.Field(i).Elem()); err == nil {
					return v, nil
				}
			}
		case reflect.Struct:
			if v, err := findFieldFromInterface(k, v.Field(i)); err == nil {
				return v, nil
			}
		}
	}
	return v, errors.New("Nothing Field")
}

func strConvType(in interface{}, typ reflect.Type) (interface{}, error) {
	if s, ok := in.(string); ok {
		switch typ.Kind() {
		case reflect.String:
			return in, nil
		case reflect.Bool:
			x, err := strconv.ParseBool(s)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("[Error] %v", err))
			}
			return reflect.ValueOf(x).Convert(typ).Interface(), nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(s, 10, typ.Bits())
			if err != nil || reflect.Zero(typ).OverflowInt(x) {
				return nil, errors.New(fmt.Sprintf("[Error] %v", err))
			}
			return reflect.ValueOf(x).Convert(typ).Interface(), nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			x, err := strconv.ParseUint(s, 10, typ.Bits())
			if err != nil || reflect.Zero(typ).OverflowUint(x) {
				return nil, errors.New(fmt.Sprintf("[Error] %v", err))
			}
			return reflect.ValueOf(x).Convert(typ).Interface(), nil
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(s, typ.Bits())
			if err != nil || reflect.Zero(typ).OverflowFloat(x) {
				return nil, errors.New(fmt.Sprintf("[Error] %v", err))
			}
			return reflect.ValueOf(x).Convert(typ).Interface(), nil
		default:
			return nil, errors.New(fmt.Sprintf("[Error] %v", "Unknown reflect Type"))
		}
	}
	return nil, errors.New(fmt.Sprintf("[Error] %v", "in is not string"))
}
