在下面这些场景下，我们会使用go generate命令：
```txt
yacc：从 .y 文件生成 .go 文件；
protobufs：从 protocol buffer 定义文件（.proto）生成 .pb.go 文件；
Unicode：从 UnicodeData.txt 生成 Unicode 表；
HTML：将 HTML 文件嵌入到 go 源码；
bindata：将形如 JPEG 这样的文件转成 go 代码中的字节数组。
```

go generate命令格式如下所示：
```txt
go generate [-run regexp] [-n] [-v] [-x] [command] [build flags] [file.go... | packages]
参数说明如下：
-run 正则表达式匹配命令行，仅执行匹配的命令；
-v 输出被处理的包名和源文件名；
-n 显示不执行命令；
-x 显示并执行命令；
command 可以是在环境变量 PATH 中的任何命令。
```

执行go generate命令时，也可以使用一些环境变量，如下所示:
```txt
$GOARCH 体系架构（arm、amd64 等）；
$GOOS 当前的 OS 环境（linux、windows 等）；
$GOFILE 当前处理中的文件名；
$GOLINE 当前命令在文件中的行号；
$GOPACKAGE 当前处理文件的包名；
$DOLLAR 固定的$，不清楚具体用途。
```
