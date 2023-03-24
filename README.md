# Pipe2GPT

Pipe2GPT 是一个使用 Go 编写的命令行工具，用于与 OpenAI GPT-3.5-turbo 模型进行交互。用户可以从 `conf.json` 文件中获取指定类型的 prompt，并将从标准输入获取的文本一起作为 GPT-3.5-turbo 模型的输入。程序通过 OpenAI API 发送请求并获取模型生成的回复，然后将回复内容输出到控制台。

Pipe2GPT is a command-line tool written in Go that interacts with the OpenAI GPT-3.5-turbo model. Users can obtain a specific type of prompt from the `conf.json` file and use it along with the text obtained from standard input as input to the GPT-3.5-turbo model. The program sends requests through the OpenAI API and receives the model-generated responses, which are then output to the console.

~~~command
$ kubectl describe po test-deployment-7fcc4f4fd-p8j27 | bin/pipe2gpt --type=k8s

这个输出展示了一个状态为“等待”的 Pod。原因是它无法从镜像仓库中拉取镜像，因此出现了“ImagePullBackOff”错误。可能的原因包括镜像仓库不可访问
  镜像名称错误等。对应的容器“test-container”状态为“等待（Waiting）”，它包含一个“while”循环，不断输出“Hello, Kubernetes”，每隔15秒退出一次
  此时容器还没有启动，因此“Ready”字段是“False”。同时，“ContainersReady”字段也为“False”，因为所有的容器都还没有就绪。最后，事件列表中显示
  这个 Pod 最近的事件，包括拉取镜像和拉取镜像失败的事件。

~~~

## 开始 / Getting Started

### 安装 / Installation

#### 中文
1. 确保您已安装 [Go](https://golang.org/doc/install) 。
2. 克隆此仓库到您的本地计算机。
3. 进入项目根目录，运行 `./build.sh` 以构建

项目。
4. 配置环境变量 `OPENAI_TOKEN`。

#### English

1. Ensure you have [Go](https://golang.org/doc/install) installed.
2. Clone this repository to your local machine.
3. In the project root directory, run `./build.sh` to build the project.
Set up the environment variable `OPENAI_TOKEN`.

### 使用 / Usage

#### 中文

修改 `conf.json` 文件，添加您需要的 Prompt 类型，用管道将命令行输出内容给 Pipe2GPT。

例如 `kubectl get pods | pipe2gpt --type=k8s`

#### English

To modify the conf.json file, add the prompt type(s) you need and pipe the command line output to pipe2gpt.

For example： `kubectl get pods | pipe2gpt --type=k8s`

## 注意事项 / Notes

请确保您已在系统中设置了 OpenAI API 密钥（环境变量 `OPENAI_TOKEN`）。

Make sure you have set the OpenAI API key (environment variable `OPENAI_TOKEN`) in your system.

## 许可 / License

用 MIT 许可。有关详细信息，请参阅 LICENSE 文件。

This project is licensed under the MIT License. For more information, see the LICENSE file.

