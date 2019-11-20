package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	i int
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	c.i += len(p)
	return len(p), nil
}

func (c *ByteCounter) String() string {
	return "ByteConter"
}

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
	//return "ce"
}

type ReaderT interface {
	Read(p []byte) (n int, err error)
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int)  {
	p[i], p[j] = p[j], p[i]
}

func main()  {
	// 7.1 接口约定
	//ce := Celsius(1.2)
	//fmt.Println(ce)
	//var c ByteCounter
	//fmt.Println(&c)
	//c.Write([]byte("hello"))
	//fmt.Println(c)
	//c.i = 0
	//var name = "Dolly"
	//fmt.Fprintf(&c, "hello, %s", name)
	//fmt.Println(c)


	// 7.2 接口类型
	// 将line25


	// 7.3 实现接口的条件
	//var w io.Writer
	//w = os.Stdout
	//w = new(bytes.Buffer)
	//
	//var rwc io.ReadWriteCloser
	//rwc = os.Stdout
	//
	//w = rwc
	//fmt.Printf("%T\n", w)
	//
	//// 空接口
	//var any interface{}
	//any = true
	//any = 12.34
	//fmt.Printf("%T\n", any)
	//
	//var _ io.Writer = (*bytes.Buffer)(nil)


	// 7.4 flag.Value接口
	//var period = flag.Duration("period", 1 * time.Second, "sleep period")
	//flag.Parse()
	//fmt.Printf("Sleeping for %v . . . ", *period)
	//time.Sleep(*period)
	//fmt.Println()

	// 7.5 接口值
	//var w io.Writer
	//fmt.Println(w)
	//w = os.Stdout
	//fmt.Println(w)
	//w = new(bytes.Buffer)
	//fmt.Printf("%T\n", w)

	// 7.5.1 一个包含nil指针的接口不是nil接口
	//var buf *bytes.Buffer
	//var out io.Writer
	//fmt.Printf("%v %v\n", out, out==nil)
	//out = buf
	//fmt.Printf("%v %v\n", out, out==nil)


	// 7.6 sort.Interface接口
	//strSlice := StringSlice([]string{"bc", "ab", "de"})
	//sort.Sort(strSlice)
	//fmt.Println(strSlice)


	// 7.8 error接口
	//fmt.Println(errors.New("124"))
	//var err error = syscall.Errno(2)
	//fmt.Println(err.Error())


	// 7.9 示例：表达式求值


	// 7.10 类型断言
	//var w io.Writer
	//w = os.Stdout
	//fmt.Printf("%T\n", w)
	//f := w.(*os.File)
	//f.WriteString("zhouhao\n")
	//fmt.Printf("%T\n", f)
	////c := w.(*bytes.Buffer)

	//var w io.Writer
	//w = os.Stdout
	//rw := w.(io.ReadWriter)
	//w = new(ByteCounter)
	////rw = w.(io.ReadWriter)
	//var rw1 io.ReadWriter
	//w = rw1.(io.Writer)
	//fmt.Printf("%T\n", rw)

	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok{
		fmt.Println(f)
	}










}
