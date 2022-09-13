# 使用指南

## 1. 启动服务

注意：远端部署时，YourIP=虚机ip，本机部署不需要--ip

## 1.1 Win

1. 使用默认参数[127.0.0.1:9090]

   双击打开convertor_win.exe文件

2. 使用自定义参数

   使用终端进入convertor_win.exe所在目录，输入以下命令

   ```shell
   convertor_win --port [YourPort] --ip [YourIP]
   ```

## 1.2 Mac

注意：arm平台未测试

1. 使用默认参数[127.0.0.1:9090]

   双击打开convertor_mac文件

2. 使用自定义参数

   使用终端进入converter_mac所在目录，输入以下命令
   
   ```shell
   ./converter_mac --port [YourPort] --ip [YourIP]
   ```

## 1.3 Linux

1. 使用默认参数[127.0.0.1:9090]

   双击打开convertor_linux文件

2. 使用自定义参数

   使用终端进入convertor_linux所在目录，输入以下命令
   
   ```shell
   ./convertor_linux --port [YourPort] --ip [YourIP]
   ```

## 2. 生成邮箱地址

目前支持的几种转换方法：

- 张三丰 --> sanfeng.zhang
- 张三丰 --> zhangsanfeng
- 张三丰 --> zhangsf
- 张三丰 --> zsf

按钮说明：

- `转换`：将姓名文本框中的内容进行转换，转换结果到右边文本框
- `清空`：清空左右两边文本框中的内容
- `选择文件`：选择本地文件识别其中内容，目前仅支持txt、xlsx格式文件
- `提交`：点击提交，系统自动识别文件中的姓名到姓名输入框
- `下载`：将转换结果下载到本地，目前只支持txt格式

### 2.1 根据输入生成

注意：请以如下格式输入

每行一个中文名字

<img src="https://img-kay.oss-cn-shanghai.aliyuncs.com/images/image-20220913135059528.png" alt="image-20220913135059528" style="zoom:67%;" />

### 2.2 根据本地文件生成

#### 2.2.1 根据txt文件

注意：txt请遵循如下格式：

每行一个中文名字

<img src="https://img-kay.oss-cn-shanghai.aliyuncs.com/images/image-20220913135126040.png" alt="image-20220913135126040" style="zoom:67%;" />

#### 2.2.2 根据xlsx文件

注意：txt请遵循如下格式：

sheet1中的某一列为姓名列，且第一个cell内容为文本“姓名”

<img src="https://img-kay.oss-cn-shanghai.aliyuncs.com/images/image-20220913135221091.png" alt="image-20220913135221091" style="zoom:67%;" />

## 3. 下载文件到本地

### 3.1 txt下载

点击下载即可

### 3.2 xlsx下载[暂不支持]