// Copyright 2015 Joel Wu
// Copyright 2012 Gary Burd
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package redis

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ErrNil indicates that a reply value is nil.
var ErrNil = errors.New("nil reply")

func errNegativeInt(v int64) error {
	return fmt.Errorf("redigo: unexpected negative value %v for Uint64", v)
}

// Int is a helper that converts a command reply to an integer. If err is not
// equal to nil, then Int returns 0, err. Otherwise, Int converts the
// reply to an int as follows:
//
//  Reply type    Result
//  integer       int(reply), nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Int(reply interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case int64:
		x := int(reply)
		if int64(x) != reply {
			return 0, strconv.ErrRange
		}
		return x, nil
	case []byte:
		n, err := strconv.ParseInt(string(reply), 10, 0)
		return int(n), err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("unexpected type %T for Int", reply)
}

// Ints is a helper that converts an array command reply to a []int.
// If err is not equal to nil, then Ints returns nil, err.
func Ints(reply interface{}, err error) ([]int, error) {
	var result []int
	err = sliceHelper(reply, err, "Ints", func(n int) { result = make([]int, n) }, func(i int, v interface{}) error {
		switch v := v.(type) {
		case int64:
			n := int(v)
			if int64(n) != v {
				return strconv.ErrRange
			}
			result[i] = n
			return nil
		case []byte:
			n, err := strconv.Atoi(string(v))
			result[i] = n
			return err
		default:
			return fmt.Errorf("redigo: unexpected element type for Ints, got type %T", v)
		}
	})
	return result, err
}

// Int64 is a helper that converts a command reply to 64 bit integer. If err is
// not equal to nil, then Int returns 0, err. Otherwise, Int64 converts the
// reply to an int64 as follows:
//
//  Reply type    Result
//  integer       reply, nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Int64(reply interface{}, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case int64:
		return reply, nil
	case []byte:
		n, err := strconv.ParseInt(string(reply), 10, 64)
		return n, err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("unexpected type %T for Int64", reply)
}

// Int64s is a helper that converts an array command reply to a []int64.
// If err is not equal to nil, then Int64s returns nil, err. Nil array
// items are stay nil. Int64s returns an error if an array item is not a
// bulk string or nil.
func Int64s(reply interface{}, err error) ([]int64, error) {
	var result []int64
	err = sliceHelper(reply, err, "Int64s", func(n int) { result = make([]int64, n) }, func(i int, v interface{}) error {
		switch v := v.(type) {
		case int64:
			result[i] = v
			return nil
		case []byte:
			n, err := strconv.ParseInt(string(v), 10, 64)
			result[i] = n
			return err
		default:
			return fmt.Errorf("redigo: unexpected element type for Int64s, got type %T", v)
		}
	})
	return result, err
}

// Uint64 is a helper that converts a command reply to 64 bit unsigned integer.
// If err is not equal to nil, then Uint64 returns 0, err. Otherwise, Uint64 converts the
// reply to an uint64 as follows:
//
//  Reply type    Result
//  +integer      reply, nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Uint64(reply interface{}, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case int64:
		if reply < 0 {
			return 0, errNegativeInt(reply)
		}
		return uint64(reply), nil
	case []byte:
		n, err := strconv.ParseUint(string(reply), 10, 64)
		return n, err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("redigo: unexpected type for Uint64, got type %T", reply)
}

// Uint64s is a helper that converts an array command reply to a []uint64.
// If err is not equal to nil, then Uint64s returns nil, err. Nil array
// items are stay nil. Uint64s returns an error if an array item is not a
// bulk string or nil.
func Uint64s(reply interface{}, err error) ([]uint64, error) {
	var result []uint64
	err = sliceHelper(reply, err, "Uint64s", func(n int) { result = make([]uint64, n) }, func(i int, v interface{}) error {
		switch v := v.(type) {
		case uint64:
			result[i] = v
			return nil
		case []byte:
			n, err := strconv.ParseUint(string(v), 10, 64)
			result[i] = n
			return err
		default:
			return fmt.Errorf("redigo: unexpected element type for Uint64s, got type %T", v)
		}
	})
	return result, err
}

// Uint32 is a helper that converts a command reply to 64 bit unsigned integer.
// If err is not equal to nil, then Uint32 returns 0, err. Otherwise, Uint32 converts the
// reply to an uint64 as follows:
//
//  Reply type    Result
//  +integer      reply, nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Int32(reply interface{}, err error) (int32, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case int64:
		if reply < 0 {
			return 0, errNegativeInt(reply)
		}
		return int32(reply), nil
	case []byte:
		n, err := strconv.ParseUint(string(reply), 10, 64)
		return int32(n), err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("redigo: unexpected type for Uint64, got type %T", reply)
}

// Int32s is a helper that converts an array command reply to a []int64.
// If err is not equal to nil, then Int32s returns nil, err. Nil array
// items are stay nil. Int32s returns an error if an array item is not a
// bulk string or nil.
func Int32s(reply interface{}, err error) ([]int32, error) {
	var result []int32
	err = sliceHelper(reply, err, "Int32s", func(n int) { result = make([]int32, n) }, func(i int, v interface{}) error {
		switch v := v.(type) {
		case int64:
			result[i] = int32(v)
			return nil
		case []byte:
			n, err := strconv.ParseInt(string(v), 10, 64)
			result[i] = int32(n)
			return err
		default:
			return fmt.Errorf("redigo: unexpected element type for Int32s, got type %T", v)
		}
	})
	return result, err
}

// Uint32 is a helper that converts a command reply to 64 bit unsigned integer.
// If err is not equal to nil, then Uint32 returns 0, err. Otherwise, Uint32 converts the
// reply to an uint64 as follows:
//
//  Reply type    Result
//  +integer      reply, nil
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Uint32(reply interface{}, err error) (uint32, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case int64:
		if reply < 0 {
			return 0, errNegativeInt(reply)
		}
		return uint32(reply), nil
	case []byte:
		n, err := strconv.ParseUint(string(reply), 10, 64)
		return uint32(n), err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("redigo: unexpected type for Uint32, got type %T", reply)
}

// Uint32s is a helper that converts an array command reply to a []int64.
// If err is not equal to nil, then Uint32s returns nil, err. Nil array
// items are stay nil. Uint32s returns an error if an array item is not a
// bulk string or nil.
func Uint32s(reply interface{}, err error) ([]uint32, error) {
	var result []uint32
	err = sliceHelper(reply, err, "Uint32s", func(n int) { result = make([]uint32, n) }, func(i int, v interface{}) error {
		switch v := v.(type) {
		case int64:
			result[i] = uint32(v)
			return nil
		case []byte:
			n, err := strconv.ParseInt(string(v), 10, 64)
			result[i] = uint32(n)
			return err
		default:
			return fmt.Errorf("redigo: unexpected element type for Uint32s, got type %T", v)
		}
	})
	return result, err
}

// Float64 is a helper that converts a command reply to 64 bit float. If err is
// not equal to nil, then Float64 returns 0, err. Otherwise, Float64 converts
// the reply to an int as follows:
//
//  Reply type    Result
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Float64(reply interface{}, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case []byte:
		n, err := strconv.ParseFloat(string(reply), 64)
		return n, err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("unexpected type %T for Float64", reply)
}

// Float64s is a helper that converts an array command reply to a []float64. If
// err is not equal to nil, then Float64s returns nil, err. Nil array items are
// converted to 0 in the output slice. Floats64 returns an error if an array
// item is not a bulk string or nil.
func Float64s(reply interface{}, err error) ([]float64, error) {
	var result []float64
	err = sliceHelper(reply, err, "Float64s", func(n int) { result = make([]float64, n) }, func(i int, v interface{}) error {
		p, ok := v.([]byte)
		if !ok {
			return fmt.Errorf("redigo: unexpected element type for Floats64, got type %T", v)
		}
		f, err := strconv.ParseFloat(string(p), 64)
		result[i] = f
		return err
	})
	return result, err
}

// Float32 is a helper that converts a command reply to 64 bit float. If err is
// not equal to nil, then Float32 returns 0, err. Otherwise, Float32 converts
// the reply to an int as follows:
//
//  Reply type    Result
//  bulk string   parsed reply, nil
//  nil           0, ErrNil
//  other         0, error
func Float32(reply interface{}, err error) (float32, error) {
	if err != nil {
		return 0, err
	}
	switch reply := reply.(type) {
	case []byte:
		n, err := strconv.ParseFloat(string(reply), 64)
		return float32(n), err
	case nil:
		return 0, ErrNil
	case redisError:
		return 0, reply
	}
	return 0, fmt.Errorf("unexpected type %T for Float64", reply)
}

// Float32s is a helper that converts an array command reply to a []float64. If
// err is not equal to nil, then Float64s returns nil, err. Nil array items are
// converted to 0 in the output slice. Floats64 returns an error if an array
// item is not a bulk string or nil.
func Float32s(reply interface{}, err error) ([]float32, error) {
	var result []float32
	err = sliceHelper(reply, err, "Float32s", func(n int) { result = make([]float32, n) }, func(i int, v interface{}) error {
		p, ok := v.([]byte)
		if !ok {
			return fmt.Errorf("redigo: unexpected element type for Float32s, got type %T", v)
		}
		f, err := strconv.ParseFloat(string(p), 64)
		result[i] = float32(f)
		return err
	})
	return result, err
}

// String is a helper that converts a command reply to a string. If err is not
// equal to nil, then String returns "", err. Otherwise String converts the
// reply to a string as follows:
//
//  Reply type      Result
//  bulk string     string(reply), nil
//  simple string   reply, nil
//  nil             "",  ErrNil
//  other           "",  error
func String(reply interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	switch reply := reply.(type) {
	case []byte:
		return string(reply), nil
	case string:
		return reply, nil
	case nil:
		return "", ErrNil
	case redisError:
		return "", reply
	}
	return "", fmt.Errorf("unexpected type %T for String", reply)
}

// Strings is a helper that converts an array command reply to a []string. If
// err is not equal to nil, then Strings returns nil, err. Nil array items are
// converted to "" in the output slice. Strings returns an error if an array
// item is not a bulk string or nil.
func Strings(reply interface{}, err error) ([]string, error) {
	values, err := Values(reply, err)
	if err != nil {
		return nil, err
	}

	strings := make([]string, len(values))
	slice := make([]interface{}, len(values))
	for i, _ := range strings {
		slice[i] = &strings[i]
	}

	if _, err = Scan(values, slice...); err != nil {
		return nil, err
	}

	return strings, nil
}

// Bytes is a helper that converts a command reply to a slice of bytes. If err
// is not equal to nil, then Bytes returns nil, err. Otherwise Bytes converts
// the reply to a slice of bytes as follows:
//
//  Reply type      Result
//  bulk string     reply, nil
//  simple string   []byte(reply), nil
//  nil             nil, ErrNil
//  other           nil, error
func Bytes(reply interface{}, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	switch reply := reply.(type) {
	case []byte:
		return reply, nil
	case string:
		return []byte(reply), nil
	case nil:
		return nil, ErrNil
	case redisError:
		return nil, reply
	}
	return nil, fmt.Errorf("unexpected type %T for Bytes", reply)
}

// ByteSlices is a helper that converts an array command reply to a [][]byte.
// If err is not equal to nil, then ByteSlices returns nil, err. Nil array
// items are stay nil. ByteSlices returns an error if an array item is not a
// bulk string or nil.
func ByteSlices(reply interface{}, err error) ([][]byte, error) {
	var result [][]byte
	err = sliceHelper(reply, err, "ByteSlices", func(n int) { result = make([][]byte, n) }, func(i int, v interface{}) error {
		p, ok := v.([]byte)
		if !ok {
			return fmt.Errorf("redigo: unexpected element type for ByteSlices, got type %T", v)
		}
		result[i] = p
		return nil
	})
	return result, err
}

// Bool is a helper that converts a command reply to a boolean. If err is not
// equal to nil, then Bool returns false, err. Otherwise Bool converts the
// reply to boolean as follows:
//
//  Reply type      Result
//  integer         value != 0, nil
//  bulk string     strconv.ParseBool(reply)
//  nil             false, ErrNil
//  other           false, error
func Bool(reply interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	switch reply := reply.(type) {
	case int64:
		return reply != 0, nil
	case []byte:
		return strconv.ParseBool(string(reply))
	case nil:
		return false, ErrNil
	case redisError:
		return false, reply
	}
	return false, fmt.Errorf("unexpected type %T for Bool", reply)
}

// MultiBulk is a helper that converts an array command reply to a []interface{}.
//
// Deprecated: Use Values instead.
func MultiBulk(reply interface{}, err error) ([]interface{}, error) { return Values(reply, err) }

// Values is a helper that converts an array command reply to a []interface{}.
// If err is not equal to nil, then Values returns nil, err. Otherwise, Values
// converts the reply as follows:
//
//  Reply type      Result
//  array           reply, nil
//  nil             nil, ErrNil
//  other           nil, error
func Values(reply interface{}, err error) ([]interface{}, error) {
	if err != nil {
		return nil, err
	}
	switch reply := reply.(type) {
	case []interface{}:
		return reply, nil
	case nil:
		return nil, ErrNil
	case redisError:
		return nil, reply
	}
	return nil, fmt.Errorf("unexpected type %T for Values", reply)
}

func sliceHelper(reply interface{}, err error, name string, makeSlice func(int), assign func(int, interface{}) error) error {
	if err != nil {
		return err
	}
	switch reply := reply.(type) {
	case []interface{}:
		makeSlice(len(reply))
		for i := range reply {
			if reply[i] == nil {
				continue
			}
			if err := assign(i, reply[i]); err != nil {
				return err
			}
		}
		return nil
	case nil:
		return ErrNil
	case redisError:
		return reply
	}
	return fmt.Errorf("redigo: unexpected type for %s, got type %T", name, reply)
}

// StringStringMap is a helper that converts an array of strings (alternating key, value)
// into a map[string]string. The HGETALL and CONFIG GET commands return replies in this format.
// Requires an even number of values in result.
func StringStringMap(result interface{}, err error) (map[string]string, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("expect even number elements for StringStringMap")
	}

	m := make(map[string]string, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, okKey := values[i].([]byte)
		value, okValue := values[i+1].([]byte)
		if !okKey || !okValue {
			return nil, errors.New("expect bulk string for StringStringMap")
		}
		m[string(key)] = string(value)
	}

	return m, nil
}

// StringIntMap is a helper that converts an array of strings (alternating key, value)
// into a map[string]int. The HGETALL commands return replies in this format.
// Requires an even number of values in result.
func StringIntMap(result interface{}, err error) (map[string]int, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: StringIntMap expects even number of values result")
	}
	m := make(map[string]int, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: StringIntMap key not a bulk string value")
		}
		value, err := Int(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}

// StringInt64Map is a helper that converts an array of strings (alternating key, value)
// into a map[string]int64. The HGETALL commands return replies in this format.
// Requires an even number of values in result.
func StringInt64Map(result interface{}, err error) (map[string]int64, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: StringInt64Map expects even number of values result")
	}
	m := make(map[string]int64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: StringInt64Map key not a bulk string value")
		}
		value, err := Int64(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}

// StringUint64Map is a helper that converts an array of strings (alternating key, value)
// into a map[string]uint64. The HGETALL commands return replies in this format.
// Requires an even number of values in result.
func StringUint64Map(result interface{}, err error) (map[string]uint64, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: StringUint64Map expects even number of values result")
	}
	m := make(map[string]uint64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, errors.New("redigo: StringUint64Map key not a bulk string value")
		}
		value, err := Uint64(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[string(key)] = value
	}
	return m, nil
}

// Uint64Uint64Map is a helper that converts an array of strings (alternating key, value)
// into a map[uint64]uint64. The HGETALL commands return replies in this format.
// Requires an even number of values in result.
func Uint64Uint64Map(result interface{}, err error) (map[uint64]uint64, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("redigo: Uint64Uint64Map expects even number of values result")
	}
	m := make(map[uint64]uint64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, err := Uint64(values[i+1], nil)
		if err != nil {
			return nil, errors.New("redigo: Uint64Uint64Map key not a bulk string value")
		}
		value, err := Uint64(values[i+1], nil)
		if err != nil {
			return nil, err
		}
		m[key] = value
	}
	return m, nil
}

// Positions is a helper that converts an array of positions (lat, long)
// into a [][2]float64. The GEOPOS command returns replies in this format.
func Positions(result interface{}, err error) ([]*[2]float64, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	positions := make([]*[2]float64, len(values))
	for i := range values {
		if values[i] == nil {
			continue
		}
		p, ok := values[i].([]interface{})
		if !ok {
			return nil, fmt.Errorf("redigo: unexpected element type for interface slice, got type %T", values[i])
		}
		if len(p) != 2 {
			return nil, fmt.Errorf("redigo: unexpected number of values for a member position, got %d", len(p))
		}
		lat, err := Float64(p[0], nil)
		if err != nil {
			return nil, err
		}
		long, err := Float64(p[1], nil)
		if err != nil {
			return nil, err
		}
		positions[i] = &[2]float64{lat, long}
	}
	return positions, nil
}

// SlowLog represents a redis SlowLog
type SlowLog struct {
	// ID is a unique progressive identifier for every slow log entry.
	ID int64

	// Time is the unix timestamp at which the logged command was processed.
	Time time.Time

	// ExecutationTime is the amount of time needed for the command execution.
	ExecutionTime time.Duration

	// Args is the command name and arguments
	Args []string

	// ClientAddr is the client IP address (4.0 only).
	ClientAddr string

	// ClientName is the name set via the CLIENT SETNAME command (4.0 only).
	ClientName string
}

// SlowLogs is a helper that parse the SLOWLOG GET command output and
// return the array of SlowLog
func SlowLogs(result interface{}, err error) ([]SlowLog, error) {
	rawLogs, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	logs := make([]SlowLog, len(rawLogs))
	for i, rawLog := range rawLogs {
		rawLog, ok := rawLog.([]interface{})
		if !ok {
			return nil, errors.New("redigo: slowlog element is not an array")
		}

		var log SlowLog

		if len(rawLog) < 4 {
			return nil, errors.New("redigo: slowlog element has less than four elements")
		}
		log.ID, ok = rawLog[0].(int64)
		if !ok {
			return nil, errors.New("redigo: slowlog element[0] not an int64")
		}
		timestamp, ok := rawLog[1].(int64)
		if !ok {
			return nil, errors.New("redigo: slowlog element[1] not an int64")
		}
		log.Time = time.Unix(timestamp, 0)
		duration, ok := rawLog[2].(int64)
		if !ok {
			return nil, errors.New("redigo: slowlog element[2] not an int64")
		}
		log.ExecutionTime = time.Duration(duration) * time.Microsecond

		log.Args, err = Strings(rawLog[3], nil)
		if err != nil {
			return nil, fmt.Errorf("redigo: slowlog element[3] is not array of string. actual error is : %s", err.Error())
		}
		if len(rawLog) >= 6 {
			log.ClientAddr, err = String(rawLog[4], nil)
			if err != nil {
				return nil, fmt.Errorf("redigo: slowlog element[4] is not a string. actual error is : %s", err.Error())
			}
			log.ClientName, err = String(rawLog[5], nil)
			if err != nil {
				return nil, fmt.Errorf("redigo: slowlog element[5] is not a string. actual error is : %s", err.Error())
			}
		}
		logs[i] = log
	}
	return logs, nil
}

// Scan copies from src to the values pointed at by dest.
//
// The values pointed at by dest must be an integer, float, boolean, string,
// []byte, interface{} or slices of these types. Scan uses the standard strconv
// package to convert bulk strings to numeric and boolean types.
//
// If a dest value is nil, then the corresponding src value is skipped.
//
// If a src element is nil, then the corresponding dest value is not modified.
//
// To enable easy use of Scan in a loop, Scan returns the slice of src
// following the copied values.
/*func Scan(src []interface{}, dst ...interface{}) ([]interface{}, error) {
	if len(src) < len(dst) {
		return nil, errors.New("mismatch length of source and dest")
	}
	var err error
	for i, d := range dst {
		err = convertAssign(d, src[i])
		if err != nil {
			break
		}
	}
	return src[len(dst):], err
}*/

/*func ensureLen(d reflect.Value, n int) {
	if n > d.Cap() {
		d.Set(reflect.MakeSlice(d.Type(), n, n))
	} else {
		d.SetLen(n)
	}
}*/

/*func cannotConvert(d reflect.Value, s interface{}) error {
	return fmt.Errorf("redigo: Scan cannot convert from %s to %s",
		reflect.TypeOf(s), d.Type())
}

func convertAssignInt(d reflect.Value, s int64) (err error) {
	switch d.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		d.SetInt(s)
		if d.Int() != s {
			err = strconv.ErrRange
			d.SetInt(0)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if s < 0 {
			err = strconv.ErrRange
		} else {
			x := uint64(s)
			d.SetUint(x)
			if d.Uint() != x {
				err = strconv.ErrRange
				d.SetUint(0)
			}
		}
	case reflect.Bool:
		d.SetBool(s != 0)
	default:
		err = cannotConvert(d, s)
	}
	return
}

func convertAssignValue(d reflect.Value, s interface{}) (err error) {
	switch s := s.(type) {
	case []byte:
		err = convertAssignBytes(d, s)
	case int64:
		err = convertAssignInt(d, s)
	default:
		err = cannotConvert(d, s)
	}
	return err
}
*/

/*func convertAssign(d interface{}, s interface{}) (err error) {
	// Handle the most common destination types using type switches and
	// fall back to reflection for all other types.
	switch s := s.(type) {
	case nil:
		// ingore
	case []byte:
		switch d := d.(type) {
		case *string:
			*d = string(s)
		case *int:
			*d, err = strconv.Atoi(string(s))
		case *int64:
			*d, err = strconv.ParseInt(string(s), 10, 64)
		case *bool:
			*d, err = strconv.ParseBool(string(s))
		case *[]byte:
			*d = s
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			if d := reflect.ValueOf(d); d.Type().Kind() != reflect.Ptr {
				err = cannotConvert(d, s)
			} else {
				err = convertAssignBytes(d.Elem(), s)
			}
		}
	case int64:
		switch d := d.(type) {
		case *int:
			x := int(s)
			if int64(x) != s {
				err = strconv.ErrRange
				x = 0
			}
			*d = x
		case *int64:
			*d = s
		case *bool:
			*d = s != 0
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			if d := reflect.ValueOf(d); d.Type().Kind() != reflect.Ptr {
				err = cannotConvert(d, s)
			} else {
				err = convertAssignInt(d.Elem(), s)
			}
		}
	case []interface{}:
		switch d := d.(type) {
		case *[]interface{}:
			*d = s
		case *interface{}:
			*d = s
		case nil:
			// skip value
		default:
			if d := reflect.ValueOf(d); d.Type().Kind() != reflect.Ptr {
				err = cannotConvert(d, s)
			} else {
				err = convertAssignValues(d.Elem(), s)
			}
		}
	case redisError:
		err = s
	default:
		err = cannotConvert(reflect.ValueOf(d), s)
	}
	return
}
*/
