package main

import (
	sp "github.com/scipipe/scipipe"
)

func main() {
	// Initialize processes
	foo := sp.NewFromShell("foowriter", "echo 'foo' > {o:foo}")
	f2b := sp.NewFromShell("foo2bar", "sed 's/foo/resultforreplacefoo/g' {i:foo} > {o:bar}")
	snk := sp.NewSink() // Will just receive file targets, doing nothing

	// Add output file path formatters for the components created above
	foo.SetPathStatic("foo", "foo.txt")
	// 在foo的path后面加后缀名，这函数名也是醉
	f2b.SetPathExtend("foo", "bar", ".bar")

	// Connect network
	f2b.In["foo"].Connect(foo.Out["foo"])
	snk.Connect(f2b.Out["bar"])

	// Add to a pipeline runner and run
	pl := sp.NewPipelineRunner()
	pl.AddProcesses(foo, f2b, snk)
	pl.Run()
}
