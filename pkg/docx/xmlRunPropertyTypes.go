package docx

/*Color 字体颜色*/
type Color struct {
	/* 字体颜色值，如 D4F4F2，前面不带#号 */
	Value string
	/* 字体的主题颜色，如运用了主题，以主题为主 */
	Theme string
}

/*
Language 字体语言

This element specifies the languages which shall be used to check spelling and grammar (if requested) when
processing the contents of this run.

例如: <w:lang w:val="en-US" w:eastAsia="zh-CN" w:bidi="he-IL"/>
*/
type Language struct {
	/* 指定在处理使用拉丁字符的运行内容时(由运行内容的Unicode字符值决定)应用于检查拼写和语法(如果请求)的语言 */
	Value string
	/* 指定在处理使用复杂脚本字符的运行内容时应使用的语言，由运行内容的Unicode字符值决定。 */
	Bidi string
	/* 指定在处理使用东亚字符的运行内容时应使用的语言 */
	EastAsian string
}

// RunFonts 最多有4种字体槽
type RunFonts struct {
	// 默认提示所用的子图
	Hint string
	// 处理Ascii字符时所使用的字体
	Ascii string
	// 处理 High ANSI 字符时所使用的字体
	HAnsi string
	// 处理东南亚 East Asian 文字所使用的字体，包括中文
	EastAsia string
	// 处理 Complex Script 字符时所使用的字体
	Cs string
	// Ascii字符所使用的主题
	AsciiTheme string
	// High ANSI字符所使用的主题
	HAnsiTheme string
	// 东南亚文字所使用的主题
	EastAsiaTheme string
}
