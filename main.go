// Context

package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// 定义和生成Key, 根据官方建议不使用内置类型作为key.
type Key struct {
	name string
}

func NewKey() Key {
	return Key{name: "trace_id"}
}

var testKey Key = NewKey()

// 生成唯一Value
func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// 生成上下文
func ContextWithTraceID() context.Context {
	ctx := context.WithValue(context.Background(), testKey, NewRequestID())
	return ctx
}

// 查找上下文, 返回Value
func GetContextValue(ctx context.Context, k Key) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return "Can not get context value from key: " + k.name
	}
	return v
}

// 格式化打印Log
func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, testKey), message)
}

// 在生成Value Context后再加上一个timeout, 返回ctx和cancelFunc
func ContextWithTimeOut(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

func HandleTimeFromValueContext(ctx context.Context) {
	ctx, concel := ContextWithTimeOut(ctx, 3*time.Second)
	defer concel()
	deal(ctx)
}

func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			
		}

	}
}

// 入口
func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "Key-Value context\n")
	HandleTimeFromValueContext(ctx)
}

func main() {
	ProcessEnter(ContextWithTraceID())
}
