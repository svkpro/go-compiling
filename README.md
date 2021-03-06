### Preface

Working as a PHP developer for a long time, I ignored a lot of things from a programming perspective. PHP is an interpreted scripting language with dynamic typization, and you don't have to control strict data types, memory allocation or building process. You can, but you don’t have to. But recently I’ve discovered  GoLang. And I’ve started learning from basics, like compiling. So I believe the article will be interesting for beginners or developers who’s traveling from interpreted languages to GoLang.
To demonstrate the compiling process we will use following program code:

```go
package main

import (
  "fmt"
)

type Greeter struct {
  helloPhrase string
}

func (g Greeter) Hello() {
  fmt.Println(g.helloPhrase)
}

func main() {
  g := Greeter{helloPhrase: "Hey everyone!"}
  g.Hello()
}
```

### Go compiling process

First of all it should be clarified that the name "gc" stands for "Go compiler", and has little to do with uppercase "GC", which stands for garbage collection. According documentation the compiler may be logically split in four phases:

#### 1. Parsing

[cmd/compile/internal/syntax](https://github.com/golang/go/tree/master/src/cmd/compile/internal/syntax) (lexer, parser, syntax tree)
In the first phase of compilation, source code is tokenized (lexical analysis), parsed (syntax analysis), and a syntax tree is constructed for each source file. Each syntax tree is an exact representation of the respective source file, with nodes corresponding to the various elements of the source such as expressions, declarations, and statements. The syntax tree also includes position information which is used for error reporting and the creation of debugging information.

```text
main.go:1:1package "package"
main.go:1:9IDENT   "main"
main.go:1:13;       "\n"
main.go:3:1import  "import"
main.go:3:8(       ""
main.go:4:2STRING  "\"fmt\""
main.go:4:7;       "\n"
main.go:5:1)       ""
main.go:5:2;       "\n"
main.go:7:1type    "type"
main.go:7:6IDENT   "Greeter"
main.go:7:14struct  "struct"
main.go:7:21{       ""
main.go:8:2IDENT   "helloPhrase"
main.go:8:14IDENT   "string"
main.go:8:20;       "\n"
main.go:9:1}       ""
main.go:9:2;       "\n"
main.go:11:1func    "func"
main.go:11:6(       ""
main.go:11:7IDENT   "g"
main.go:11:9IDENT   "Greeter"
main.go:11:16)       ""
main.go:11:18IDENT   "Hello"
main.go:11:23(       ""
main.go:11:24)       ""
main.go:11:26{       ""
main.go:12:2IDENT   "fmt"
main.go:12:5.       ""
main.go:12:6IDENT   "Println"
main.go:12:13(       ""
main.go:12:14IDENT   "g"
main.go:12:15.       ""
main.go:12:16IDENT   "helloPhrase"
main.go:12:27)       ""
main.go:12:28;       "\n"
main.go:13:1}       ""
main.go:13:2;       "\n"
main.go:15:1func    "func"
main.go:15:6IDENT   "main"
main.go:15:10(       ""
main.go:15:11)       ""
main.go:15:13{       ""
main.go:16:2IDENT   "g"
main.go:16:4:=      ""
main.go:16:7IDENT   "Greeter"
main.go:16:14{       ""
main.go:16:15IDENT   "helloPhrase"
main.go:16:26:       ""
main.go:16:28STRING  "\"Hey everyone!\""
main.go:16:43}       ""
main.go:16:44;       "\n"
main.go:17:2IDENT   "g"
main.go:17:3.       ""
main.go:17:4IDENT   "Hello"
main.go:17:9(       ""
main.go:17:10)       ""
main.go:17:11;       "\n"
main.go:18:1}       ""
main.go:18:2;       "\n"
main.go:18:2EOF     ""
```
#### 2. Type-checking and AST (Abstract Syntax Tree) transformations

[cmd/compile/internal/gc](https://github.com/golang/go/tree/master/src/cmd/compile/internal/gc) (create compiler AST, type checking, AST transformations)
The gc package includes an AST definition carried over from when it was written in C. All of its code is written in terms of it, so the first thing that the gc package must do is convert the syntax package's syntax tree to the compiler's AST representation. The AST is then type-checked. Type-checking includes certain extra checks, such as "declared and not used" as well as determining whether or not a function terminates.
Certain transformations are also done on the AST. Some nodes are refined based on type information, such as string additions being split from the arithmetic addition node type. Some other examples are dead code elimination, function call inlining, and escape analysis.

```text
0  *ast.File {
     1  .  Package: main.go:1:1
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: main.go:1:9
     4  .  .  Name: "main"
     5  .  }
     6  .  Decls: []ast.Decl (len = 4) {
     7  .  .  0: *ast.GenDecl {
     8  .  .  .  TokPos: main.go:3:1
     9  .  .  .  Tok: import
    10  .  .  .  Lparen: main.go:3:8
    11  .  .  .  Specs: []ast.Spec (len = 1) {
    12  .  .  .  .  0: *ast.ImportSpec {
    13  .  .  .  .  .  Path: *ast.BasicLit {
    14  .  .  .  .  .  .  ValuePos: main.go:4:2
    15  .  .  .  .  .  .  Kind: STRING
    16  .  .  .  .  .  .  Value: "\"fmt\""
    17  .  .  .  .  .  }
    18  .  .  .  .  .  EndPos: -
    19  .  .  .  .  }
    20  .  .  .  }
    21  .  .  .  Rparen: main.go:5:1
    22  .  .  }
    23  .  .  1: *ast.GenDecl {
    24  .  .  .  TokPos: main.go:7:1
    25  .  .  .  Tok: type
    26  .  .  .  Lparen: -
    27  .  .  .  Specs: []ast.Spec (len = 1) {
    28  .  .  .  .  0: *ast.TypeSpec {
    29  .  .  .  .  .  Name: *ast.Ident {
    30  .  .  .  .  .  .  NamePos: main.go:7:6
    31  .  .  .  .  .  .  Name: "Greeter"
    32  .  .  .  .  .  .  Obj: *ast.Object {
    33  .  .  .  .  .  .  .  Kind: type
    34  .  .  .  .  .  .  .  Name: "Greeter"
    35  .  .  .  .  .  .  .  Decl: *(obj @ 28)
    36  .  .  .  .  .  .  }
    37  .  .  .  .  .  }
    38  .  .  .  .  .  Assign: -
    39  .  .  .  .  .  Type: *ast.StructType {
    40  .  .  .  .  .  .  Struct: main.go:7:14
    41  .  .  .  .  .  .  Fields: *ast.FieldList {
    42  .  .  .  .  .  .  .  Opening: main.go:7:21
    43  .  .  .  .  .  .  .  List: []*ast.Field (len = 1) {
    44  .  .  .  .  .  .  .  .  0: *ast.Field {
    45  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
    46  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
    47  .  .  .  .  .  .  .  .  .  .  .  NamePos: main.go:8:2
    48  .  .  .  .  .  .  .  .  .  .  .  Name: "helloPhrase"
    49  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
    50  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
    51  .  .  .  .  .  .  .  .  .  .  .  .  Name: "helloPhrase"
    52  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 44)
    53  .  .  .  .  .  .  .  .  .  .  .  }
    54  .  .  .  .  .  .  .  .  .  .  }
    55  .  .  .  .  .  .  .  .  .  }
    56  .  .  .  .  .  .  .  .  .  Type: *ast.Ident {
    57  .  .  .  .  .  .  .  .  .  .  NamePos: main.go:8:14
    58  .  .  .  .  .  .  .  .  .  .  Name: "string"
    59  .  .  .  .  .  .  .  .  .  }
    60  .  .  .  .  .  .  .  .  }
    61  .  .  .  .  .  .  .  }
    62  .  .  .  .  .  .  .  Closing: main.go:9:1
    63  .  .  .  .  .  .  }
    64  .  .  .  .  .  .  Incomplete: false
    65  .  .  .  .  .  }
    66  .  .  .  .  }
    67  .  .  .  }
    68  .  .  .  Rparen: -
    69  .  .  }
    70  .  .  2: *ast.FuncDecl {
    71  .  .  .  Recv: *ast.FieldList {
    72  .  .  .  .  Opening: main.go:11:6
    73  .  .  .  .  List: []*ast.Field (len = 1) {
    74  .  .  .  .  .  0: *ast.Field {
    75  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
    76  .  .  .  .  .  .  .  0: *ast.Ident {
    77  .  .  .  .  .  .  .  .  NamePos: main.go:11:7
    78  .  .  .  .  .  .  .  .  Name: "g"
    79  .  .  .  .  .  .  .  .  Obj: *ast.Object {
    80  .  .  .  .  .  .  .  .  .  Kind: var
    81  .  .  .  .  .  .  .  .  .  Name: "g"
    82  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 74)
    83  .  .  .  .  .  .  .  .  }
    84  .  .  .  .  .  .  .  }
    85  .  .  .  .  .  .  }
    86  .  .  .  .  .  .  Type: *ast.Ident {
    87  .  .  .  .  .  .  .  NamePos: main.go:11:9
    88  .  .  .  .  .  .  .  Name: "Greeter"
    89  .  .  .  .  .  .  .  Obj: *(obj @ 32)
    90  .  .  .  .  .  .  }
    91  .  .  .  .  .  }
    92  .  .  .  .  }
    93  .  .  .  .  Closing: main.go:11:16
    94  .  .  .  }
    95  .  .  .  Name: *ast.Ident {
    96  .  .  .  .  NamePos: main.go:11:18
    97  .  .  .  .  Name: "Hello"
    98  .  .  .  }
    99  .  .  .  Type: *ast.FuncType {
   100  .  .  .  .  Func: main.go:11:1
   101  .  .  .  .  Params: *ast.FieldList {
   102  .  .  .  .  .  Opening: main.go:11:23
   103  .  .  .  .  .  Closing: main.go:11:24
   104  .  .  .  .  }
   105  .  .  .  }
   106  .  .  .  Body: *ast.BlockStmt {
   107  .  .  .  .  Lbrace: main.go:11:26
   108  .  .  .  .  List: []ast.Stmt (len = 1) {
   109  .  .  .  .  .  0: *ast.ExprStmt {
   110  .  .  .  .  .  .  X: *ast.CallExpr {
   111  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
   112  .  .  .  .  .  .  .  .  X: *ast.Ident {
   113  .  .  .  .  .  .  .  .  .  NamePos: main.go:12:2
   114  .  .  .  .  .  .  .  .  .  Name: "fmt"
   115  .  .  .  .  .  .  .  .  }
   116  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
   117  .  .  .  .  .  .  .  .  .  NamePos: main.go:12:6
   118  .  .  .  .  .  .  .  .  .  Name: "Println"
   119  .  .  .  .  .  .  .  .  }
   120  .  .  .  .  .  .  .  }
   121  .  .  .  .  .  .  .  Lparen: main.go:12:13
   122  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
   123  .  .  .  .  .  .  .  .  0: *ast.SelectorExpr {
   124  .  .  .  .  .  .  .  .  .  X: *ast.Ident {
   125  .  .  .  .  .  .  .  .  .  .  NamePos: main.go:12:14
   126  .  .  .  .  .  .  .  .  .  .  Name: "g"
   127  .  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 79)
   128  .  .  .  .  .  .  .  .  .  }
   129  .  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
   130  .  .  .  .  .  .  .  .  .  .  NamePos: main.go:12:16
   131  .  .  .  .  .  .  .  .  .  .  Name: "helloPhrase"
   132  .  .  .  .  .  .  .  .  .  }
   133  .  .  .  .  .  .  .  .  }
   134  .  .  .  .  .  .  .  }
   135  .  .  .  .  .  .  .  Ellipsis: -
   136  .  .  .  .  .  .  .  Rparen: main.go:12:27
   137  .  .  .  .  .  .  }
   138  .  .  .  .  .  }
   139  .  .  .  .  }
   140  .  .  .  .  Rbrace: main.go:13:1
   141  .  .  .  }
   142  .  .  }
   143  .  .  3: *ast.FuncDecl {
   144  .  .  .  Name: *ast.Ident {
   145  .  .  .  .  NamePos: main.go:15:6
   146  .  .  .  .  Name: "main"
   147  .  .  .  .  Obj: *ast.Object {
   148  .  .  .  .  .  Kind: func
   149  .  .  .  .  .  Name: "main"
   150  .  .  .  .  .  Decl: *(obj @ 143)
   151  .  .  .  .  }
   152  .  .  .  }
   153  .  .  .  Type: *ast.FuncType {
   154  .  .  .  .  Func: main.go:15:1
   155  .  .  .  .  Params: *ast.FieldList {
   156  .  .  .  .  .  Opening: main.go:15:10
   157  .  .  .  .  .  Closing: main.go:15:11
   158  .  .  .  .  }
   159  .  .  .  }
   160  .  .  .  Body: *ast.BlockStmt {
   161  .  .  .  .  Lbrace: main.go:15:13
   162  .  .  .  .  List: []ast.Stmt (len = 2) {
   163  .  .  .  .  .  0: *ast.AssignStmt {
   164  .  .  .  .  .  .  Lhs: []ast.Expr (len = 1) {
   165  .  .  .  .  .  .  .  0: *ast.Ident {
   166  .  .  .  .  .  .  .  .  NamePos: main.go:16:2
   167  .  .  .  .  .  .  .  .  Name: "g"
   168  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   169  .  .  .  .  .  .  .  .  .  Kind: var
   170  .  .  .  .  .  .  .  .  .  Name: "g"
   171  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 163)
   172  .  .  .  .  .  .  .  .  }
   173  .  .  .  .  .  .  .  }
   174  .  .  .  .  .  .  }
   175  .  .  .  .  .  .  TokPos: main.go:16:4
   176  .  .  .  .  .  .  Tok: :=
   177  .  .  .  .  .  .  Rhs: []ast.Expr (len = 1) {
   178  .  .  .  .  .  .  .  0: *ast.CompositeLit {
   179  .  .  .  .  .  .  .  .  Type: *ast.Ident {
   180  .  .  .  .  .  .  .  .  .  NamePos: main.go:16:7
   181  .  .  .  .  .  .  .  .  .  Name: "Greeter"
   182  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 32)
   183  .  .  .  .  .  .  .  .  }
   184  .  .  .  .  .  .  .  .  Lbrace: main.go:16:14
   185  .  .  .  .  .  .  .  .  Elts: []ast.Expr (len = 1) {
   186  .  .  .  .  .  .  .  .  .  0: *ast.KeyValueExpr {
   187  .  .  .  .  .  .  .  .  .  .  Key: *ast.Ident {
   188  .  .  .  .  .  .  .  .  .  .  .  NamePos: main.go:16:15
   189  .  .  .  .  .  .  .  .  .  .  .  Name: "helloPhrase"
   190  .  .  .  .  .  .  .  .  .  .  }
   191  .  .  .  .  .  .  .  .  .  .  Colon: main.go:16:26
   192  .  .  .  .  .  .  .  .  .  .  Value: *ast.BasicLit {
   193  .  .  .  .  .  .  .  .  .  .  .  ValuePos: main.go:16:28
   194  .  .  .  .  .  .  .  .  .  .  .  Kind: STRING
   195  .  .  .  .  .  .  .  .  .  .  .  Value: "\"Hey everyone!\""
   196  .  .  .  .  .  .  .  .  .  .  }
   197  .  .  .  .  .  .  .  .  .  }
   198  .  .  .  .  .  .  .  .  }
   199  .  .  .  .  .  .  .  .  Rbrace: main.go:16:43
   200  .  .  .  .  .  .  .  .  Incomplete: false
   201  .  .  .  .  .  .  .  }
   202  .  .  .  .  .  .  }
   203  .  .  .  .  .  }
   204  .  .  .  .  .  1: *ast.ExprStmt {
   205  .  .  .  .  .  .  X: *ast.CallExpr {
   206  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
   207  .  .  .  .  .  .  .  .  X: *ast.Ident {
   208  .  .  .  .  .  .  .  .  .  NamePos: main.go:17:2
   209  .  .  .  .  .  .  .  .  .  Name: "g"
   210  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 168)
   211  .  .  .  .  .  .  .  .  }
   212  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
   213  .  .  .  .  .  .  .  .  .  NamePos: main.go:17:4
   214  .  .  .  .  .  .  .  .  .  Name: "Hello"
   215  .  .  .  .  .  .  .  .  }
   216  .  .  .  .  .  .  .  }
   217  .  .  .  .  .  .  .  Lparen: main.go:17:9
   218  .  .  .  .  .  .  .  Ellipsis: -
   219  .  .  .  .  .  .  .  Rparen: main.go:17:10
   220  .  .  .  .  .  .  }
   221  .  .  .  .  .  }
   222  .  .  .  .  }
   223  .  .  .  .  Rbrace: main.go:18:1
   224  .  .  .  }
   225  .  .  }
   226  .  }
   227  .  Scope: *ast.Scope {
   228  .  .  Objects: map[string]*ast.Object (len = 2) {
   229  .  .  .  "Greeter": *(obj @ 32)
   230  .  .  .  "main": *(obj @ 147)
   231  .  .  }
   232  .  }
   233  .  Imports: []*ast.ImportSpec (len = 1) {
   234  .  .  0: *(obj @ 12)
   235  .  }
   236  .  Unresolved: []*ast.Ident (len = 2) {
   237  .  .  0: *(obj @ 56)
   238  .  .  1: *(obj @ 112)
   239  .  }
   240  }
```

#### 3. Generic SSA

[cmd/compile/internal/gc](https://github.com/golang/go/tree/master/src/cmd/compile/internal/gc) (converting to SSA)
[cmd/compile/internal/ssa](https://github.com/golang/go/tree/master/src/cmd/compile/internal/ssa) (SSA passes and rules)
In this phase, the AST is converted into Static Single Assignment (SSA) form, a lower-level intermediate representation with specific properties that make it easier to implement optimizations and to eventually generate machine code from it. During this conversion, function intrinsics are applied. These are special functions that the compiler has been taught to replace with heavily optimized code on a case-by-case basis. Certain nodes are also lowered into simpler components during the AST to SSA conversion, so that the rest of the compiler can work with them. For instance, the copy builtin is replaced by memory moves, and range loops are rewritten into for loops. Some of these currently happen before the conversion to SSA due to historical reasons, but the long-term plan is to move all of them here.
Then, a series of machine-independent passes and rules are applied. These do not concern any single computer architecture, and thus run on all GOARCH variants.

```text
package main:
  type  Greeter    struct{helloPhrase string}
    method (Greeter) Hello()
  func  init       func()
  var   init$guard bool
  func  main       func()

# Name: main.init
# Package: main
# Synthetic: package initializer
func init():
0:                                                                entry P:0 S:2
        t0 = *init$guard                                                   bool
        if t0 goto 2 else 1
1:                                                           init.start P:1 S:1
        *init$guard = true:bool
        t1 = fmt.init()                                                      ()
        jump 2
2:                                                            init.done P:2 S:0
        return

# Name: main.main
# Package: main
# Location: main.go:15:6
# Locals:
#   0:  t0 Greeter
func main():
0:                                                                entry P:0 S:0
        t0 = local Greeter (g)                                         *Greeter
        t1 = &t0.helloPhrase [#0]                                       *string
        *t1 = "Hey everyone!":string
        t2 = *t0                                                        Greeter
        t3 = (Greeter).Hello(t2)                                             ()
        return
```
#### 4. Generating machine code

[cmd/compile/internal/ssa](https://github.com/golang/go/tree/master/src/cmd/compile/internal/ssa) (SSA lowering and arch-specific passes)
[cmd/internal/obj](https://github.com/golang/go/tree/master/src/cmd/internal/obj) (machine code generation)
The machine-dependent phase of the compiler begins with the "lower" pass, which rewrites generic values into their machine-specific variants. For example, on amd64 memory operands are possible, so many load-store operations may be combined.
Note that the lower pass runs all machine-specific rewrite rules, and thus it currently applies lots of optimizations too.
Once the SSA has been "lowered" and is more specific to the target architecture, the final code optimization passes are run. This includes yet another dead code elimination pass, moving values closer to their uses, the removal of local variables that are never read from, and register allocation.
Other important pieces of work done as part of this step include stack frame layout, which assigns stack offsets to local variables, and pointer liveness analysis, which computes which on-stack pointers are live at each GC safe point.
At the end of the SSA generation phase, Go functions have been transformed into a series of obj.Prog instructions. These are passed to the assembler (cmd/internal/obj), which turns them into machine code and writes out the final object file. The object file will also contain  reflect data, export data, and debugging information.

```text
os.(*File).close STEXT dupok nosplit size=26 args=0x18 locals=0x0
        0x0000 00000 (<autogenerated>:1)        TEXT    os.(*File).close(SB), DUPOK|NOSPLIT|ABIInternal, $0-24
        0x0000 00000 (<autogenerated>:1)        PCDATA  $0, $-2
        0x0000 00000 (<autogenerated>:1)        PCDATA  $1, $-2
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0000 00000 (<autogenerated>:1)        FUNCDATA        $2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0000 00000 (<autogenerated>:1)        PCDATA  $0, $1
        0x0000 00000 (<autogenerated>:1)        PCDATA  $1, $1
        0x0000 00000 (<autogenerated>:1)        MOVQ    ""..this+8(SP), AX
        0x0005 00005 (<autogenerated>:1)        MOVQ    (AX), AX
        0x0008 00008 (<autogenerated>:1)        PCDATA  $0, $0
        0x0008 00008 (<autogenerated>:1)        PCDATA  $1, $0
        0x0008 00008 (<autogenerated>:1)        MOVQ    AX, ""..this+8(SP)
        0x000d 00013 (<autogenerated>:1)        XORPS   X0, X0
        0x0010 00016 (<autogenerated>:1)        MOVUPS  X0, "".~r0+16(SP)
        0x0015 00021 (<autogenerated>:1)        JMP     os.(*file).close(SB)
        0x0000 48 8b 44 24 08 48 8b 00 48 89 44 24 08 0f 57 c0  H.D$.H..H.D$..W.
        0x0010 0f 11 44 24 10 e9 00 00 00 00                    ..D$......
        rel 22+4 t=8 os.(*file).close+0
"".Greeter.Hello STEXT size=163 args=0x10 locals=0x58
        0x0000 00000 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        TEXT    "".Greeter.Hello(SB), ABIInternal, $88-16
        0x0000 00000 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        MOVQ    (TLS), CX
        0x0009 00009 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        CMPQ    SP, 16(CX)
        0x000d 00013 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $0, $-2
        0x000d 00013 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        JLS     153
        0x0013 00019 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $0, $-1
        0x0013 00019 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        SUBQ    $88, SP
        0x0017 00023 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        MOVQ    BP, 80(SP)
        0x001c 00028 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        LEAQ    80(SP), BP
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $0, $-2
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $1, $-2
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        FUNCDATA        $0, gclocals·2d7c1615616d4cf40d01b3385155ed6e(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        FUNCDATA        $1, gclocals·ffd148479e14c29ee3c68361945c5d25(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        FUNCDATA        $2, gclocals·bfec7e55b3f043d1941c093912808913(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        FUNCDATA        $3, "".Greeter.Hello.stkobj(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $1, $0
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    "".g+96(SP), AX
        0x0026 00038 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $0
        0x0026 00038 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, (SP)
        0x002a 00042 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $1, $1
        0x002a 00042 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    "".g+104(SP), AX
        0x002f 00047 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, 8(SP)
        0x0034 00052 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        CALL    runtime.convTstring(SB)
        0x0039 00057 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x0039 00057 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    16(SP), AX
        0x003e 00062 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $1, $2
        0x003e 00062 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        XORPS   X0, X0
        0x0041 00065 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVUPS  X0, ""..autotmp_13+64(SP)
        0x0046 00070 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $2
        0x0046 00070 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        LEAQ    type.string(SB), CX
        0x004d 00077 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x004d 00077 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    CX, ""..autotmp_13+64(SP)
        0x0052 00082 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $0
        0x0052 00082 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, ""..autotmp_13+72(SP)
        0x0057 00087 (<unknown line number>)    NOP
        0x0057 00087 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x0057 00087 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x005e 00094 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $2
        0x005e 00094 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
        0x0065 00101 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x0065 00101 ($GOROOT/src/fmt/print.go:274)     MOVQ    CX, (SP)
        0x0069 00105 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $0
        0x0069 00105 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x006e 00110 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x006e 00110 ($GOROOT/src/fmt/print.go:274)     PCDATA  $1, $1
        0x006e 00110 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_13+64(SP), AX
        0x0073 00115 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $0
        0x0073 00115 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x0078 00120 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x0081 00129 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x008a 00138 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x008f 00143 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    80(SP), BP
        0x0094 00148 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        ADDQ    $88, SP
        0x0098 00152 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        RET
        0x0099 00153 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        NOP
        0x0099 00153 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $1, $-1
        0x0099 00153 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $0, $-2
        0x0099 00153 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        CALL    runtime.morestack_noctxt(SB)
        0x009e 00158 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        PCDATA  $0, $-1
        0x009e 00158 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:11)        JMP     0
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 86  dH..%....H;a....
        0x0010 00 00 00 48 83 ec 58 48 89 6c 24 50 48 8d 6c 24  ...H..XH.l$PH.l$
        0x0020 50 48 8b 44 24 60 48 89 04 24 48 8b 44 24 68 48  PH.D$`H..$H.D$hH
        0x0030 89 44 24 08 e8 00 00 00 00 48 8b 44 24 10 0f 57  .D$......H.D$..W
        0x0040 c0 0f 11 44 24 40 48 8d 0d 00 00 00 00 48 89 4c  ...D$@H......H.L
        0x0050 24 40 48 89 44 24 48 48 8b 05 00 00 00 00 48 8d  $@H.D$HH......H.
        0x0060 0d 00 00 00 00 48 89 0c 24 48 89 44 24 08 48 8d  .....H..$H.D$.H.
        0x0070 44 24 40 48 89 44 24 10 48 c7 44 24 18 01 00 00  D$@H.D$.H.D$....
        0x0080 00 48 c7 44 24 20 01 00 00 00 e8 00 00 00 00 48  .H.D$ .........H
        0x0090 8b 6c 24 50 48 83 c4 58 c3 e8 00 00 00 00 e9 5d  .l$PH..X.......]
        0x00a0 ff ff ff                                         ...
        rel 5+4 t=17 TLS+0
        rel 53+4 t=8 runtime.convTstring+0
        rel 73+4 t=16 type.string+0
        rel 90+4 t=16 os.Stdout+0
        rel 97+4 t=16 go.itab.*os.File,io.Writer+0
        rel 139+4 t=8 fmt.Fprintln+0
        rel 154+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=164 args=0x0 locals=0x58
        0x0000 00000 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        TEXT    "".main(SB), ABIInternal, $88-0
        0x0000 00000 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        MOVQ    (TLS), CX
        0x0009 00009 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        CMPQ    SP, 16(CX)
        0x000d 00013 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $0, $-2
        0x000d 00013 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        JLS     154
        0x0013 00019 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $0, $-1
        0x0013 00019 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        SUBQ    $88, SP
        0x0017 00023 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        MOVQ    BP, 80(SP)
        0x001c 00028 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        LEAQ    80(SP), BP
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $0, $-2
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $1, $-2
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        FUNCDATA        $1, gclocals·568470801006e5c0dc3947ea998fe279(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        FUNCDATA        $2, gclocals·bfec7e55b3f043d1941c093912808913(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        FUNCDATA        $3, "".main.stkobj(SB)
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:17)        PCDATA  $0, $-1
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:17)        PCDATA  $1, $-1
        0x0021 00033 (<unknown line number>)    NOP
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:17)        PCDATA  $0, $1
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:17)        PCDATA  $1, $0
        0x0021 00033 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:17)        LEAQ    go.string."Hey everyone!"(SB), AX
        0x0028 00040 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $0
        0x0028 00040 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, (SP)
        0x002c 00044 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    $13, 8(SP)
        0x0035 00053 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        CALL    runtime.convTstring(SB)
        0x003a 00058 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x003a 00058 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    16(SP), AX
        0x003f 00063 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $1, $1
        0x003f 00063 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        XORPS   X0, X0
        0x0042 00066 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVUPS  X0, ""..autotmp_14+64(SP)
        0x0047 00071 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $2
        0x0047 00071 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        LEAQ    type.string(SB), CX
        0x004e 00078 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x004e 00078 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    CX, ""..autotmp_14+64(SP)
        0x0053 00083 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $0
        0x0053 00083 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, ""..autotmp_14+72(SP)
        0x0058 00088 (<unknown line number>)    NOP
        0x0058 00088 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x0058 00088 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x005f 00095 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $2
        0x005f 00095 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
        0x0066 00102 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x0066 00102 ($GOROOT/src/fmt/print.go:274)     MOVQ    CX, (SP)
        0x006a 00106 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $0
        0x006a 00106 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x006f 00111 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x006f 00111 ($GOROOT/src/fmt/print.go:274)     PCDATA  $1, $0
        0x006f 00111 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_14+64(SP), AX
        0x0074 00116 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $0
        0x0074 00116 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x0079 00121 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x0082 00130 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x008b 00139 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x0090 00144 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    80(SP), BP
        0x0095 00149 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        ADDQ    $88, SP
        0x0099 00153 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        RET
        0x009a 00154 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        NOP
        0x009a 00154 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $1, $-1
        0x009a 00154 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $0, $-2
        0x009a 00154 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        CALL    runtime.morestack_noctxt(SB)
        0x009f 00159 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        PCDATA  $0, $-1
        0x009f 00159 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:15)        JMP     0
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 87  dH..%....H;a....
        0x0010 00 00 00 48 83 ec 58 48 89 6c 24 50 48 8d 6c 24  ...H..XH.l$PH.l$
        0x0020 50 48 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24  PH......H..$H.D$
        0x0030 08 0d 00 00 00 e8 00 00 00 00 48 8b 44 24 10 0f  ..........H.D$..
        0x0040 57 c0 0f 11 44 24 40 48 8d 0d 00 00 00 00 48 89  W...D$@H......H.
        0x0050 4c 24 40 48 89 44 24 48 48 8b 05 00 00 00 00 48  L$@H.D$HH......H
        0x0060 8d 0d 00 00 00 00 48 89 0c 24 48 89 44 24 08 48  ......H..$H.D$.H
        0x0070 8d 44 24 40 48 89 44 24 10 48 c7 44 24 18 01 00  .D$@H.D$.H.D$...
        0x0080 00 00 48 c7 44 24 20 01 00 00 00 e8 00 00 00 00  ..H.D$ .........
        0x0090 48 8b 6c 24 50 48 83 c4 58 c3 e8 00 00 00 00 e9  H.l$PH..X.......
        0x00a0 5c ff ff ff                                      \...
        rel 5+4 t=17 TLS+0
        rel 36+4 t=16 go.string."Hey everyone!"+0
        rel 54+4 t=8 runtime.convTstring+0
        rel 74+4 t=16 type.string+0
        rel 91+4 t=16 os.Stdout+0
        rel 98+4 t=16 go.itab.*os.File,io.Writer+0
        rel 140+4 t=8 fmt.Fprintln+0
        rel 155+4 t=8 runtime.morestack_noctxt+0
"".(*Greeter).Hello STEXT dupok size=211 args=0x8 locals=0x58
        0x0000 00000 (<autogenerated>:1)        TEXT    "".(*Greeter).Hello(SB), DUPOK|WRAPPER|ABIInternal, $88-8
        0x0000 00000 (<autogenerated>:1)        MOVQ    (TLS), CX
        0x0009 00009 (<autogenerated>:1)        CMPQ    SP, 16(CX)
        0x000d 00013 (<autogenerated>:1)        PCDATA  $0, $-2
        0x000d 00013 (<autogenerated>:1)        JLS     179
        0x0013 00019 (<autogenerated>:1)        PCDATA  $0, $-1
        0x0013 00019 (<autogenerated>:1)        SUBQ    $88, SP
        0x0017 00023 (<autogenerated>:1)        MOVQ    BP, 80(SP)
        0x001c 00028 (<autogenerated>:1)        LEAQ    80(SP), BP
        0x0021 00033 (<autogenerated>:1)        MOVQ    32(CX), BX
        0x0025 00037 (<autogenerated>:1)        TESTQ   BX, BX
        0x0028 00040 (<autogenerated>:1)        JNE     189
        0x002e 00046 (<autogenerated>:1)        NOP
        0x002e 00046 (<autogenerated>:1)        PCDATA  $0, $-2
        0x002e 00046 (<autogenerated>:1)        PCDATA  $1, $-2
        0x002e 00046 (<autogenerated>:1)        FUNCDATA        $0, gclocals·2d7c1615616d4cf40d01b3385155ed6e(SB)
        0x002e 00046 (<autogenerated>:1)        FUNCDATA        $1, gclocals·ffd148479e14c29ee3c68361945c5d25(SB)
        0x002e 00046 (<autogenerated>:1)        FUNCDATA        $2, gclocals·bfec7e55b3f043d1941c093912808913(SB)
        0x002e 00046 (<autogenerated>:1)        FUNCDATA        $3, "".(*Greeter).Hello.stkobj(SB)
        0x002e 00046 (<autogenerated>:1)        PCDATA  $0, $1
        0x002e 00046 (<autogenerated>:1)        PCDATA  $1, $1
        0x002e 00046 (<autogenerated>:1)        MOVQ    ""..this+96(SP), AX
        0x0033 00051 (<autogenerated>:1)        TESTQ   AX, AX
        0x0036 00054 (<autogenerated>:1)        JEQ     173
        0x0038 00056 (<autogenerated>:1)        MOVQ    8(AX), CX
        0x003c 00060 (<autogenerated>:1)        MOVQ    (AX), AX
        0x003f 00063 (<unknown line number>)    NOP
        0x003f 00063 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $0
        0x003f 00063 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, (SP)
        0x0043 00067 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    CX, 8(SP)
        0x0048 00072 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        CALL    runtime.convTstring(SB)
        0x004d 00077 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x004d 00077 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    16(SP), AX
        0x0052 00082 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $1, $2
        0x0052 00082 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        XORPS   X0, X0
        0x0055 00085 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVUPS  X0, ""..autotmp_14+64(SP)
        0x005a 00090 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $2
        0x005a 00090 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        LEAQ    type.string(SB), CX
        0x0061 00097 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $1
        0x0061 00097 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    CX, ""..autotmp_14+64(SP)
        0x0066 00102 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        PCDATA  $0, $0
        0x0066 00102 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    AX, ""..autotmp_14+72(SP)
        0x006b 00107 (<unknown line number>)    NOP
        0x006b 00107 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x006b 00107 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x0072 00114 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $2
        0x0072 00114 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
        0x0079 00121 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x0079 00121 ($GOROOT/src/fmt/print.go:274)     MOVQ    CX, (SP)
        0x007d 00125 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $0
        0x007d 00125 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x0082 00130 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $1
        0x0082 00130 ($GOROOT/src/fmt/print.go:274)     PCDATA  $1, $1
        0x0082 00130 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_14+64(SP), AX
        0x0087 00135 ($GOROOT/src/fmt/print.go:274)     PCDATA  $0, $0
        0x0087 00135 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x008c 00140 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x0095 00149 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x009e 00158 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x00a3 00163 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        MOVQ    80(SP), BP
        0x00a8 00168 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        ADDQ    $88, SP
        0x00ac 00172 (/home/vasilisalavei/GoProgramming/go-compiling/main.go:12)        RET
        0x00ad 00173 (<autogenerated>:1)        CALL    runtime.panicwrap(SB)
        0x00b2 00178 (<autogenerated>:1)        XCHGL   AX, AX
        0x00b3 00179 (<autogenerated>:1)        NOP
        0x00b3 00179 (<autogenerated>:1)        PCDATA  $1, $-1
        0x00b3 00179 (<autogenerated>:1)        PCDATA  $0, $-2
        0x00b3 00179 (<autogenerated>:1)        CALL    runtime.morestack_noctxt(SB)
        0x00b8 00184 (<autogenerated>:1)        PCDATA  $0, $-1
        0x00b8 00184 (<autogenerated>:1)        JMP     0
        0x00bd 00189 (<autogenerated>:1)        LEAQ    96(SP), DI
        0x00c2 00194 (<autogenerated>:1)        CMPQ    (BX), DI
        0x00c5 00197 (<autogenerated>:1)        JNE     46
        0x00cb 00203 (<autogenerated>:1)        MOVQ    SP, (BX)
        0x00ce 00206 (<autogenerated>:1)        JMP     46
        0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a0  dH..%....H;a....
        0x0010 00 00 00 48 83 ec 58 48 89 6c 24 50 48 8d 6c 24  ...H..XH.l$PH.l$
        0x0020 50 48 8b 59 20 48 85 db 0f 85 8f 00 00 00 48 8b  PH.Y H........H.
        0x0030 44 24 60 48 85 c0 74 75 48 8b 48 08 48 8b 00 48  D$`H..tuH.H.H..H
        0x0040 89 04 24 48 89 4c 24 08 e8 00 00 00 00 48 8b 44  ..$H.L$......H.D
        0x0050 24 10 0f 57 c0 0f 11 44 24 40 48 8d 0d 00 00 00  $..W...D$@H.....
        0x0060 00 48 89 4c 24 40 48 89 44 24 48 48 8b 05 00 00  .H.L$@H.D$HH....
        0x0070 00 00 48 8d 0d 00 00 00 00 48 89 0c 24 48 89 44  ..H......H..$H.D
        0x0080 24 08 48 8d 44 24 40 48 89 44 24 10 48 c7 44 24  $.H.D$@H.D$.H.D$
        0x0090 18 01 00 00 00 48 c7 44 24 20 01 00 00 00 e8 00  .....H.D$ ......
        0x00a0 00 00 00 48 8b 6c 24 50 48 83 c4 58 c3 e8 00 00  ...H.l$PH..X....
        0x00b0 00 00 90 e8 00 00 00 00 e9 43 ff ff ff 48 8d 7c  .........C...H.|
        0x00c0 24 60 48 39 3b 0f 85 63 ff ff ff 48 89 23 e9 5b  $`H9;..c...H.#.[
        0x00d0 ff ff ff                                         ...
        rel 5+4 t=17 TLS+0
        rel 73+4 t=8 runtime.convTstring+0
        rel 93+4 t=16 type.string+0
        rel 110+4 t=16 os.Stdout+0
        rel 117+4 t=16 go.itab.*os.File,io.Writer+0
        rel 159+4 t=8 fmt.Fprintln+0
        rel 174+4 t=8 runtime.panicwrap+0
        rel 180+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename.main SDWARFINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFINFO dupok size=42
        0x0000 04 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 11  .fmt.Println....
        0x0010 61 00 00 00 00 00 00 11 6e 00 01 00 00 00 00 11  a.......n.......
        0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
        rel 0+0 t=24 type.[]interface {}+0
        rel 0+0 t=24 type.error+0
        rel 0+0 t=24 type.int+0
        rel 19+4 t=29 go.info.[]interface {}+0
        rel 27+4 t=29 go.info.int+0
        rel 37+4 t=29 go.info.error+0
go.info."".Greeter.Hello$abstract SDWARFINFO dupok size=31
        0x0000 04 6d 61 69 6e 2e 47 72 65 65 74 65 72 2e 48 65  .main.Greeter.He
        0x0010 6c 6c 6f 00 01 01 11 67 00 00 00 00 00 00 00     llo....g.......
        rel 26+4 t=29 go.info."".Greeter+0
go.loc.os.(*File).close SDWARFLOC dupok size=0
go.info.os.(*File).close SDWARFINFO dupok size=55
        0x0000 03 6f 73 2e 28 2a 46 69 6c 65 29 2e 63 6c 6f 73  .os.(*File).clos
        0x0010 65 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  e...............
        0x0020 00 00 01 9c 00 00 00 00 01 0f 7e 72 30 00 01 f0  ..........~r0...
        0x0030 01 00 00 00 00 00 00                             .......
        rel 0+0 t=24 type.*os.File+0
        rel 0+0 t=24 type.error+0
        rel 18+8 t=1 os.(*File).close+0
        rel 26+8 t=1 os.(*File).close+26
        rel 36+4 t=30 gofile..<autogenerated>+0
        rel 49+4 t=29 go.info.error+0
go.range.os.(*File).close SDWARFRANGE dupok size=0
go.debuglines.os.(*File).close SDWARFMISC dupok size=12
        0x0000 04 01 0f 06 41 06 af 04 01 03 00 01              ....A.......
go.loc."".Greeter.Hello SDWARFLOC size=57
        0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 a3 00 00 00 00 00 00 00  ................
        0x0020 07 00 9c 93 08 91 08 93 08 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00                       .........
        rel 8+8 t=1 "".Greeter.Hello+0
go.info."".Greeter.Hello SDWARFINFO size=66
        0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 01 9c 13 00 00 00 00 00 00 00 00  ................
        0x0020 06 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 0c 12 00 00 00 00 00  ................
        0x0040 00 00                                            ..
        rel 0+0 t=24 type."".Greeter+0
        rel 0+0 t=24 type.[1]interface {}+0
        rel 1+4 t=29 go.info."".Greeter.Hello$abstract+0
        rel 5+8 t=1 "".Greeter.Hello+0
        rel 13+8 t=1 "".Greeter.Hello+163
        rel 24+4 t=29 go.info."".Greeter.Hello$abstract+22
        rel 28+4 t=29 go.loc."".Greeter.Hello+0
        rel 33+4 t=29 go.info.fmt.Println$abstract+0
        rel 37+8 t=1 "".Greeter.Hello+87
        rel 45+8 t=1 "".Greeter.Hello+143
        rel 53+4 t=30 gofile../home/vasilisalavei/GoProgramming/go-compiling/main.go+0
        rel 59+4 t=29 go.info.fmt.Println$abstract+15
go.range."".Greeter.Hello SDWARFRANGE size=0
go.debuglines."".Greeter.Hello SDWARFMISC size=36
        0x0000 04 02 03 05 14 0a cd 9c 06 41 04 03 06 02 1a 03  .........A......
        0x0010 81 02 fa 06 55 04 02 06 02 19 03 fe 7d fb 72 04  ....U.......}.r.
        0x0020 01 03 76 01                                      ..v.
runtime.nilinterequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
        0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
        0x0000 02                                               .
type.interface {} SRODATA dupok size=80
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.nilinterequal·f+0
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 44+4 t=6 type.*interface {}+0
        rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
        0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface
        0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]interface {}-+0
        rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
        0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[]interface {}-+0
        rel 44+4 t=6 type.*[]interface {}+0
        rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
        0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
        0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
        rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 01 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.nilinterequal·f+0
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*[1]interface {}-+0
        rel 44+4 t=6 type.*[1]interface {}+0
        rel 48+8 t=1 type.interface {}+0
        rel 56+8 t=1 type.[]interface {}+0
go.string."Hey everyone!" SRODATA dupok size=13
        0x0000 48 65 79 20 65 76 65 72 79 6f 6e 65 21           Hey everyone!
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=98
        0x0000 03 6d 61 69 6e 2e 6d 61 69 6e 00 00 00 00 00 00  .main.main......
        0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
        0x0020 00 01 0a 67 00 10 00 00 00 00 00 07 00 00 00 00  ...g............
        0x0030 00 00 00 00 00 00 00 00 11 12 00 00 00 00 00 06  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 0c 12 00 00 00 00 00 00  ................
        0x0060 00 00                                            ..
        rel 0+0 t=24 type.[1]interface {}+0
        rel 11+8 t=1 "".main+0
        rel 19+8 t=1 "".main+164
        rel 29+4 t=30 gofile../home/vasilisalavei/GoProgramming/go-compiling/main.go+0
        rel 38+4 t=29 go.info."".Greeter+0
        rel 44+4 t=29 go.info."".Greeter.Hello$abstract+0
        rel 48+4 t=29 go.range."".main+0
        rel 52+4 t=30 gofile../home/vasilisalavei/GoProgramming/go-compiling/main.go+0
        rel 58+4 t=29 go.info."".Greeter.Hello$abstract+22
        rel 64+4 t=29 go.info.fmt.Println$abstract+0
        rel 68+8 t=1 "".main+88
        rel 76+8 t=1 "".main+144
        rel 84+4 t=30 gofile../home/vasilisalavei/GoProgramming/go-compiling/main.go+0
        rel 90+4 t=29 go.info.fmt.Println$abstract+15
go.range."".main SDWARFRANGE size=64
        0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
        0x0010 28 00 00 00 00 00 00 00 58 00 00 00 00 00 00 00  (.......X.......
        0x0020 90 00 00 00 00 00 00 00 9a 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 8+8 t=1 "".main+0
go.debuglines."".main SDWARFMISC size=42
        0x0000 04 02 03 09 14 0a cd 9d 03 7f 51 06 37 06 69 06  ..........Q.7.i.
        0x0010 41 04 03 06 08 03 81 02 50 06 55 04 02 06 02 19  A.......P.U.....
        0x0020 03 fe 7d fb 76 04 01 03 72 01                    ..}.v...r.
""..inittask SNOPTRDATA size=32
        0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 fmt..inittask+0
runtime.strequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.strequal+0
go.loc."".(*Greeter).Hello SDWARFLOC dupok size=0
go.info."".(*Greeter).Hello SDWARFINFO dupok size=56
        0x0000 03 6d 61 69 6e 2e 28 2a 47 72 65 65 74 65 72 29  .main.(*Greeter)
        0x0010 2e 48 65 6c 6c 6f 00 00 00 00 00 00 00 00 00 00  .Hello..........
        0x0020 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0a 61  ...............a
        0x0030 00 0c 00 00 00 00 00 00                          ........
        rel 0+0 t=24 type.*"".Greeter+0
        rel 0+0 t=24 type.[1]interface {}+0
        rel 23+8 t=1 "".(*Greeter).Hello+0
        rel 31+8 t=1 "".(*Greeter).Hello+211
        rel 41+4 t=30 gofile..<autogenerated>+0
        rel 50+4 t=29 go.info.[]interface {}+0
go.range."".(*Greeter).Hello SDWARFRANGE dupok size=0
go.debuglines."".(*Greeter).Hello SDWARFMISC dupok size=52
        0x0000 04 01 0f 0a cd 06 08 5f 04 02 06 03 06 8c 06 37  ......._.......7
        0x0010 06 41 06 41 04 03 06 08 03 81 02 50 06 55 04 02  .A.A.......P.U..
        0x0020 06 02 19 03 fe 7d fb 04 01 06 03 79 6f 06 41 04  .....}.....yo.A.
        0x0030 01 03 00 01                                      ....
type..namedata.*main.Greeter. SRODATA dupok size=16
        0x0000 01 00 0d 2a 6d 61 69 6e 2e 47 72 65 65 74 65 72  ...*main.Greeter
type..namedata.*func(*main.Greeter)- SRODATA dupok size=23
        0x0000 00 00 14 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 47  ...*func(*main.G
        0x0010 72 65 65 74 65 72 29                             reeter)
type.*func(*"".Greeter) SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f2 14 aa 81 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func(*main.Greeter)-+0
        rel 48+8 t=1 type.func(*"".Greeter)+0
type.func(*"".Greeter) SRODATA dupok size=64
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 6f 29 c3 1a 02 08 08 33 00 00 00 00 00 00 00 00  o).....3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func(*main.Greeter)-+0
        rel 44+4 t=6 type.*func(*"".Greeter)+0
        rel 56+8 t=1 type.*"".Greeter+0
type..importpath."". SRODATA dupok size=7
        0x0000 00 00 04 6d 61 69 6e                             ...main
type..namedata.Hello. SRODATA dupok size=8
        0x0000 01 00 05 48 65 6c 6c 6f                          ...Hello
type..namedata.*func()- SRODATA dupok size=10
        0x0000 00 00 07 2a 66 75 6e 63 28 29                    ...*func()
type.*func() SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 9b 90 75 1b 08 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func()-+0
        rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00                                      ....
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func()-+0
        rel 44+4 t=6 type.*func()+0
type.*"".Greeter SRODATA size=88
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 54 d0 0d d5 09 08 08 36 00 00 00 00 00 00 00 00  T......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 01 00 01 00  ................
        0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*main.Greeter.+0
        rel 48+8 t=1 type."".Greeter+0
        rel 56+4 t=5 type..importpath."".+0
        rel 72+4 t=5 type..namedata.Hello.+0
        rel 76+4 t=25 type.func()+0
        rel 80+4 t=25 "".(*Greeter).Hello+0
        rel 84+4 t=25 "".(*Greeter).Hello+0
type..namedata.*func(main.Greeter)- SRODATA dupok size=22
        0x0000 00 00 13 2a 66 75 6e 63 28 6d 61 69 6e 2e 47 72  ...*func(main.Gr
        0x0010 65 65 74 65 72 29                                eeter)
type.*func("".Greeter) SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 68 f7 b8 ee 08 08 08 36 00 00 00 00 00 00 00 00  h......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func(main.Greeter)-+0
        rel 48+8 t=1 type.func("".Greeter)+0
type.func("".Greeter) SRODATA dupok size=64
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 32 78 a0 13 02 08 08 33 00 00 00 00 00 00 00 00  2x.....3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func(main.Greeter)-+0
        rel 44+4 t=6 type.*func("".Greeter)+0
        rel 56+8 t=1 type."".Greeter+0
type..namedata.helloPhrase- SRODATA dupok size=14
        0x0000 00 00 0b 68 65 6c 6c 6f 50 68 72 61 73 65        ...helloPhrase
type."".Greeter SRODATA size=136
        0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 a9 28 0e b2 07 08 08 19 00 00 00 00 00 00 00 00  .(..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 01 00 01 00 28 00 00 00 00 00 00 00  ........(.......
        0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0080 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.strequal·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*main.Greeter.+0
        rel 44+4 t=5 type.*"".Greeter+0
        rel 48+8 t=1 type..importpath."".+0
        rel 56+8 t=1 type."".Greeter+96
        rel 80+4 t=5 type..importpath."".+0
        rel 96+8 t=1 type..namedata.helloPhrase-+0
        rel 104+8 t=1 type.string+0
        rel 120+4 t=5 type..namedata.Hello.+0
        rel 124+4 t=25 type.func()+0
        rel 128+4 t=25 "".(*Greeter).Hello+0
        rel 132+4 t=25 "".Greeter.Hello+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
        rel 0+8 t=1 type.io.Writer+0
        rel 8+8 t=1 type.*os.File+0
        rel 24+8 t=1 os.(*File).Write+0
go.itablink.*os.File,io.Writer SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 go.itab.*os.File,io.Writer+0
type..importpath.fmt. SRODATA dupok size=6
        0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
        0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·2d7c1615616d4cf40d01b3385155ed6e SRODATA dupok size=11
        0x0000 03 00 00 00 01 00 00 00 01 00 00                 ...........
gclocals·ffd148479e14c29ee3c68361945c5d25 SRODATA dupok size=11
        0x0000 03 00 00 00 02 00 00 00 00 00 02                 ...........
gclocals·bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
        0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
"".Greeter.Hello.stkobj SRODATA size=24
        0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
        rel 16+8 t=1 type.[1]interface {}+0
gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
        0x0000 02 00 00 00 02 00 00 00 00 02                    ..........
"".main.stkobj SRODATA size=24
        0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
        rel 16+8 t=1 type.[1]interface {}+0
"".(*Greeter).Hello.stkobj SRODATA dupok size=24
        0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff ff ff ff ff  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
        rel 16+8 t=1 type.[1]interface {}+0
```