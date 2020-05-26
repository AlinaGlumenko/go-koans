package go_koans

import (
	"bytes"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		data := make([]byte, in.Len())
		r := io.Reader(in)
		r.Read(data)
		// out.Write(data)
		w := io.Writer(out)
		w.Write(data)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		data := make([]byte, in.Len())
		r := io.Reader(in)
		r.Read(data)
		// out.Write(data[:5])
		w := io.Writer(out)
		w.Write(data[:5])

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
