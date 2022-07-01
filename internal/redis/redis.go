package redis

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

// Redis is a simple redis client based on RESP protocol.
type Redis struct {
	err    error
	buffer *bufio.ReadWriter
	conn   net.Conn
}

// New creates a new redis client.
func New(addr string) (*Redis, error) {
	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err != nil {
		return nil, err
	}
	return &Redis{conn: conn, buffer: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))}, nil
}

const (
	bulkString   = '$'
	simpleString = '+'
	array        = '*'
	errorString  = '-'
	integer      = ':'
	separator    = "\r\n"
)

// Do executes a command.
func (r *Redis) Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	t, ok := ctx.Deadline()
	if ok {
		_ = r.conn.SetDeadline(t)
	}
	r.array(len(args) + 1)
	r.bulkString(cmd)
	for _, v := range args {
		switch vv := v.(type) {
		case string:
			r.bulkString(vv)
		case []byte:
			r.bulkBytes(vv)
		default:
			r.bulkString(fmt.Sprintf("%v", vv))
		}
	}
	if r.err != nil {
		return nil, r.err
	}
	if err := r.buffer.Flush(); err != nil {
		return nil, err
	}
	return r.readPart()
}

func (r *Redis) writeString(s string) {
	if r.err != nil {
		return
	}
	if _, err := r.buffer.WriteString(s); err != nil {
		r.err = err
	}
}

func (r *Redis) write(b []byte) {
	if r.err != nil {
		return
	}
	if _, err := r.buffer.Write(b); err != nil {
		r.err = err
	}
}

func (r *Redis) writeByte(b byte) {
	if r.err != nil {
		return
	}
	if err := r.buffer.WriteByte(b); err != nil {
		r.err = err
	}
}

func (r *Redis) array(size int) {
	r.writeByte(array)
	r.writeString(strconv.Itoa(size))
	r.writeString(separator)
}

func (r *Redis) bulkString(s string) {
	r.writeByte(bulkString)
	r.writeString(strconv.Itoa(len(s)))
	r.writeString(separator)
	r.writeString(s)
	r.writeString(separator)
}

func (r *Redis) bulkBytes(b []byte) {
	r.writeByte(bulkString)
	r.writeString(strconv.Itoa(len(b)))
	r.writeString(separator)
	r.write(b)
	r.writeString(separator)
}

func (r *Redis) readPart() (interface{}, error) {
	b, err := r.buffer.ReadByte()
	if err != nil {
		return nil, err
	}
	switch b {
	case simpleString:
		return r.readLine()
	case bulkString:
		return r.readBulkString()
	case integer:
		return r.readLine()
	case errorString:
		s, err := r.readLine()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("redis error: %s", string(s))
	case array:
		size, err := r.readInteger()
		if err != nil {
			return nil, err
		}
		results := make([]interface{}, size)
		for i := range results {
			a, err := r.readPart()
			if err != nil {
				return nil, err
			}
			results[i] = a
		}
		return results, nil
	default:
		return nil, fmt.Errorf("unknown type: %c", b)
	}
}

func (r *Redis) readLine() ([]byte, error) {
	var (
		s []byte
	)
	for {
		b, err := r.buffer.ReadByte()
		if err != nil {
			return nil, err
		}
		if b == '\n' && len(s) > 0 && s[len(s)-1] == '\r' {
			return s[:len(s)-1], nil
		}
		s = append(s, b)
	}
}

func (r *Redis) readBulkString() ([]byte, error) {
	size, err := r.readInteger()
	if err != nil {
		return nil, err
	}
	if size < 0 {
		return nil, nil
	}
	data := make([]byte, size)
	if _, err := io.ReadFull(r.buffer, data); err != nil {
		return nil, err
	}
	if err := r.readSeparator(); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Redis) readInteger() (int64, error) {
	line, err := r.readLine()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(string(line), 10, 64)
}

func (r *Redis) readSeparator() error {
	b, err := r.buffer.ReadByte()
	if err != nil {
		return err
	}
	if b != '\r' {
		return fmt.Errorf("expect \\r\\n, got %c", b)
	}
	b, err = r.buffer.ReadByte()
	if err != nil {
		return err
	}
	if b != '\n' {
		return fmt.Errorf("expect \\r\\n, got %c", b)
	}
	return nil
}

// String gets an string response from redis reply.
func String(o interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	s, ok := o.([]byte)
	if !ok {
		return "", fmt.Errorf("bad type of redis result: %v", o)
	}
	return string(s), nil
}

// Bytes gets an bytes response from redis reply.
func Bytes(o interface{}, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	s, ok := o.([]byte)
	if !ok {
		return nil, fmt.Errorf("bad type of redis result: %v", o)
	}
	return s, nil
}

// Int gets an integer response from redis reply.
func Int(o interface{}, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	s, ok := o.([]byte)
	if !ok {
		return 0, fmt.Errorf("bad type of redis result: %v", o)
	}
	return strconv.Atoi(string(s))
}

// Int64 gets an 64 bits integer response from redis reply.
func Int64(o interface{}, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	s, ok := o.([]byte)
	if !ok {
		return 0, fmt.Errorf("bad type of redis result: %v", o)
	}
	return strconv.ParseInt(string(s), 10, 64)
}

func sliceHelper(o interface{}, err error, construct func(int), setter func(int, []byte) error) error {
	if err != nil {
		return err
	}
	s, ok := o.([]interface{})
	if !ok {
		return fmt.Errorf("bad type of redis result: %v", o)
	}
	construct(len(s))
	for i := range s {
		if vv, ok := s[i].([]byte); ok {
			if err := setter(i, vv); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("bad type of redis result: %v", s)
		}
	}
	return nil
}

// Strings gets a slice of string response from redis reply.
func Strings(o interface{}, err error) ([]string, error) {
	var ss []string
	err = sliceHelper(o, err,
		func(i int) {
			ss = make([]string, i)
		},
		func(i int, b []byte) error {
			ss[i] = string(b)
			return nil
		})
	if err != nil {
		return nil, err
	}
	return ss, nil
}

// Ints gets a slice of integer response from redis reply.
func Ints(o interface{}, err error) ([]int, error) {
	var ss []int
	err = sliceHelper(o, err,
		func(i int) {
			ss = make([]int, i)
		},
		func(i int, b []byte) error {
			parses, err1 := strconv.Atoi(string(b))
			if err != nil {
				return err1
			}
			ss[i] = parses
			return nil
		})
	if err != nil {
		return nil, err
	}
	return ss, nil
}

// Int64s gets a slice of 64 bits integer response from redis reply.
func Int64s(o interface{}, err error) ([]int64, error) {
	var ss []int64
	err = sliceHelper(o, err,
		func(i int) {
			ss = make([]int64, i)
		},
		func(i int, b []byte) error {
			parses, err1 := strconv.ParseInt(string(b), 10, 64)
			if err != nil {
				return err1
			}
			ss[i] = parses
			return nil
		})
	if err != nil {
		return nil, err
	}
	return ss, nil
}
