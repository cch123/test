# Go 1.13 study

## defer 在栈上分配

```go
	0x003a 00058 (deferstack.go:4)	PCDATA	$0, $1
	0x003a 00058 (deferstack.go:4)	LEAQ	""..autotmp_1+8(SP), AX
	0x003f 00063 (deferstack.go:4)	PCDATA	$0, $0
	0x003f 00063 (deferstack.go:4)	MOVQ	AX, (SP)
	0x0043 00067 (deferstack.go:4)	CALL	runtime.deferprocStack(SB)
	0x0048 00072 (deferstack.go:4)	TESTL	AX, AX
	0x004a 00074 (deferstack.go:4)	JNE	92
	0x004c 00076 (deferstack.go:5)	XCHGL	AX, AX
	0x004d 00077 (deferstack.go:5)	CALL	runtime.deferreturn(SB)
	0x0052 00082 (deferstack.go:5)	MOVQ	64(SP), BP
	0x0057 00087 (deferstack.go:5)	ADDQ	$72, SP
	0x005b 00091 (deferstack.go:5)	RET
```

deferproc 变成了 deferprocStack

什么条件下使用 deferproc，什么条件下使用 deferprocStack？

### 在 for 循环中的 defer 仍然是在堆上

```golang
	case ODEFER:
		if e.loopdepth == 1 { // top level
			n.Esc = EscNever // force stack allocation of defer record (see ssa.go)
			break
		}
```

```golang
if n.Op == OFOR || n.Op == OFORUNTIL || n.Op == ORANGE {
	e.loopdepth++
}
```

```golang
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer func() {
			for {
				var a = make([]int, 128)
				fmt.Println(a)
			}
		}()
	}
}

```

###  2

```golang
// augmentParamHole augments parameter holes as necessary for use in
// go/defer statements.
func (e *Escape) augmentParamHole(k EscHole, where *Node) EscHole {
	if where == nil {
		return k
	}

	// Top level defers arguments don't escape to heap, but they
	// do need to last until end of function. Tee with a
	// non-transient location to avoid arguments from being
	// transiently allocated.
	if where.Op == ODEFER && e.loopDepth == 1 {
		where.Esc = EscNever // force stack allocation of defer record (see ssa.go)
		// TODO(mdempsky): Eliminate redundant EscLocation allocs.
		return e.teeHole(k, e.newLoc(nil, false).asHole())
	}

	return e.heapHole()
}

```


### 3

```golang
// make a new Node off the books
func tempAt(pos src.XPos, curfn *Node, t *types.Type) *Node {
	if curfn == nil {
		Fatalf("no curfn for tempname")
	}
	if curfn.Func.Closure != nil && curfn.Op == OCLOSURE {
		Dump("tempname", curfn)
		Fatalf("adding tempname to wrong closure function")
	}
	if t == nil {
		Fatalf("tempname called with nil type")
	}

	s := &types.Sym{
		Name: autotmpname(len(curfn.Func.Dcl)),
		Pkg:  localpkg,
	}
	n := newnamel(pos, s)
	s.Def = asTypesNode(n)
	n.Type = t
	n.SetClass(PAUTO)
	n.Esc = EscNever
	n.Name.Curfn = curfn
	n.Name.SetUsed(true)
	n.Name.SetAutoTemp(true)
	curfn.Func.Dcl = append(curfn.Func.Dcl, n)

	dowidth(t)

	return n.Orig
}
```

### 4

```golang
// splitSlot returns a slot representing the data of parent starting at offset.
func (e *ssafn) splitSlot(parent *ssa.LocalSlot, suffix string, offset int64, t *types.Type) ssa.LocalSlot {
	s := &types.Sym{Name: parent.N.(*Node).Sym.Name + suffix, Pkg: localpkg}

	n := &Node{
		Name: new(Name),
		Op:   ONAME,
		Pos:  parent.N.(*Node).Pos,
	}
	n.Orig = n

	s.Def = asTypesNode(n)
	asNode(s.Def).Name.SetUsed(true)
	n.Sym = s
	n.Type = t
	n.SetClass(PAUTO)
	n.SetAddable(true)
	n.Esc = EscNever
	n.Name.Curfn = e.curfn
	e.curfn.Func.Dcl = append(e.curfn.Func.Dcl, n)
	dowidth(t)
	return ssa.LocalSlot{N: n, Type: t, Off: 0, SplitOf: parent, SplitOffset: offset}
}
```