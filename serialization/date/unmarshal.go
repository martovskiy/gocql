package date

import (
	"fmt"
	"reflect"
	"time"
)

func Unmarshal(data []byte, value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil

	case *int32:
		return DecInt32(data, v)
	case *int64:
		return DecInt64(data, v)
	case *uint32:
		return DecUint32(data, v)
	case *string:
		return DecString(data, v)
	case *time.Time:
		return DecTime(data, v)

	case **int32:
		return DecInt32R(data, v)
	case **int64:
		return DecInt64R(data, v)
	case **uint32:
		return DecUint32R(data, v)
	case **string:
		return DecStringR(data, v)
	case **time.Time:
		return DecTimeR(data, v)
	default:

		// Custom types (type MyDate uint32) can be deserialized only via `reflect` package.
		// Later, when generic-based serialization is introduced we can do that via generics.
		rv := reflect.ValueOf(value)
		rt := rv.Type()
		if rt.Kind() != reflect.Ptr {
			return fmt.Errorf("failed to unmarshal date: unsupported value type (%T)(%[1]v), supported types: ~int32, ~int64, ~uint32,  ~string, time.Time", value)
		}
		if rt.Elem().Kind() != reflect.Ptr {
			return DecReflect(data, rv)
		}
		return DecReflectR(data, rv)
	}
}
