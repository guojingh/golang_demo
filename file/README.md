## golang 提供的对文件的操作
### 简单
- 打开文件
  - file, err := os.Open("文件路径")
  - file, err := os.OpenFile("文件路径","读写权限","文件权限")
- 关闭文件
  - file.Close()
- 读文件
  - content, err := ioutil.ReadFile("文件路径") 一次性读取文件全部内容，适合小文件读取，直接使用
  - 有缓存读取，大文件读取，需要配合 os.Open() 进行使用
    - reader := bufio.NewReader("文件路径")
    - str, err := reader.ReadString("每段结束标识符")
- 写文件需要配合 os.Open("文件路径进行使用")
  - writer, err := bufio.NewWriter(文件路径) 创建 Writer
  - n, err := writer.WriterString(str) 写入文件字符串
  - writer.Flush() 对流中的内容进行 Flush