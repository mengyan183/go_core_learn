package diff_pac

import (
	"flag"
	// 如果文件路径最后一个文件夹的名称相同,如果package名称不相同,则不会冲突,反之则会冲突
	"go_core_learn/lesson3/diff_pac/flag"
	// 引入包 为src或mod 文件绝对路径
	//lib "go_core_learn/lesson3/diff_pac/lib" // 当定义了别名不为 "_"时,则要求使用自定义的别名来调用引入包; 反之如果定义为 "_",则调用引入包下的代码不需要使用任何前缀,可以当做当前包直接使用
	. "go_core_learn/lesson3/diff_pac/lib" //作为当前包直接使用
	internal "go_core_learn/lesson3/diff_pac/lib/internal1"

	//"go_core_learn/lesson3/diff_pac/lib" //当未定义别名时,使用引入包下代码时要求使用引入包定义的package名称 lib_re
)

func Import() {
	//lib.Hello() // 使用自定义别名
	//Hello() // 将引入包作为当前包使用,别名为 "."
	Hello()
	// lib_re.Hello()  // 未对引入包设置别名
	internal.HelloReNameInternal()
}

func sameNamePackage()  {
	flag.Args()
	flag1.SamePackageName()
}
