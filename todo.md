### 最终使用
``` golang
var document Document
document = zdocx.read("xxx.docx")
// ...
```


1. 解压缩word文档，使之变成一个个xml文件 `box`
    - `box.pack`
    - `box.unpack`

2. 解析xml，变成一个个对象 `parse`