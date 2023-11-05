# Kronos

Kronos是一个数据整合工具，用于整合多方数据到本地ES中，包括网页内容、API数据等，并提供自定义插件的功能。

## 功能

记录数据：通过插件推送数据，将数据保存到ES。

自定义插件：提供自定义插件的功能，用户可以根据自己的需求编写插件，并将其集成到Kronos中。插件分为脚本和服务两种模式。

> 脚本：可以手动执行，也可以定义多个`定时执行`任务，比如：A插件可以根据参数判断全量拉取数据，或增量拉取数据，那就可以定义每周增量更新，每月全量更新。

> 服务：长期运行，可以通过编写第三方服务或脚本调用。

## 目录描述

- controller 主要的控制服务
- plugins 插件
  - articles => 浏览器文章插件
  - github_stars => GitHub stars插件
  - rss_pull => rss拉取插件
- web 页面展示
  - simple-searcher => 极其简单且简陋的搜索页面
  - kronos-v1 => 新版搜索页面

## 部署

暂时没有处理docker部署，只能通过源码部署。

```bash
git clone https://github.com/liCells/kronos.git
# 编译主要控制服务
cd kronos/controller
# 会得到一个可执行文件
go build .
# 执行即可
./controller

# 处理config.json中的基本配置

# 切换到你想要用的插件目录
cd ../plugins/${plugin_name}
# 按照README.md中的说明编译插件
# 将最终编译好的文件放到config.json中指定的插件位置

# 处理config.json中的插件配置
```

## 配置项

> 默认配置文件为`./config.json`，也可以指定配置文件路径：`./controller ./config.json`

```json5
{
    // 服务端口
    "port": 18000,
    "es": {
        "host": "http://10.0.0.2",
        "port": 19200,
        // ES 分词器，可以不设置
        "analyzer": "ik_smart"
    },
    "control": {
        "renew": 30,
        "allowed_maximum_number_of_disconnections": 3
    },
    // 要激活的插件，和name对应
    "activate_extensions": [
        "Articles",
        "Rss_Pull"
    ],
    // 服务
    "services": [
        {
            // 注意要和activate_extensions对应
            "name": "Articles",
            "description": "...",
            "version": "0.0.1",
            "author": "LZ",
            // 可执行文件路径，也可以换为 java -jar 这样的命令
            "relative_path": "./plugins/article/articles",
            // 所需的ES设置
            "es_index_setting": {
                // 索引名称
                "name": "articles",
                // 创建索引的json
                "setting": "{\"mappings\":{\"properties\":{\"title\":{\"type\":\"text\"},\"content\":{\"type\":\"text\"},\"tag\":{\"type\":\"text\"},\"url\":{\"type\":\"text\"},\"recordTime\":{\"type\":\"date\"}}}}",
                // 映射用的字段，统一字段，用于页面展示
                "mapping": {
                    "id": "title",
                    "title": "title",
                    "content": "content",
                    "source": "url",
                    "tag": "tag"
                }
            },
            // 执行时的参数，会拼接到relative_path后
            "manual_execution_params": "http://localhost:18000 18001"
        }
    ],
    // 脚本，和服务基本一致，但是多了commands，用于定时执行
    "scripts": [
        {
            "name": "Github Stars",
            "description": "...",
            "version": "0.0.1",
            "author": "LZ",
            "relative_path": "./plugins/stars/github_stars",
            "es_index_setting": {
                "name": "github_stars",
                "setting": "{\"mappings\":{\"properties\":{\"name\":{\"type\":\"text\"},\"full_name\":{\"type\":\"text\"},\"html_url\":{\"type\":\"text\"},\"clone_url\":{\"type\":\"text\"},\"default_branch\":{\"type\":\"text\"},\"language\":{\"type\":\"keyword\"},\"readme_url\":{\"type\":\"text\"},\"readme_content\":{\"type\":\"text\"}}}}",
                "mapping": {
                    "id": "full_name",
                    "title": "name",
                    "content": "readme_content",
                    "source": "html_url",
                    "tag": "language"
                }
            },
            // 搜索模式，目前用于指定字段的权重
            "search_scheme": [
                {
                    "field": "title",
                    "boost": 3
                },
                {
                    "field": "content",
                    "boost": 1
                }
            ],
            // 用于定时执行
            "commands": [
                {
                    "params": "./plugins/stars/failedFile liCells http://localhost:18000",
                    "cron": "0 0 0 * * *"
                }
            ],
            // 用于手动执行脚本
            "manual_execution_params": "./plugins/stars/failedFile liCells http://localhost:18000"
        }
    ]
}
```

## web页面

目前只实现了非常简陋的一个搜索页面

![searcher](https://github.com/liCells/kronos/blob/main/web/simple-searcher/imgs/simple-searcher.png?raw=true)

当然你也可以自己实现页面，只要调用相应的接口就行。

## plugins

- [x] [Articles](https://github.com/liCells/kronos/tree/main/plugins/articles)
- [x] [GitHub Stars](https://github.com/liCells/kronos/tree/main/plugins/github_stars)
- [x] [Rss Pull](https://github.com/liCells/kronos/tree/main/plugins/rss)

## 待办事项

- [x] 路由
- [x] 数据处理
- [x] 插件变量处理
- [ ] 可配置项
  - [ ] sqlite 持久化插件数据
  - [x] ES
  - [x] service plugin
    - [ ] 通过controller获取到ES信息，直接连接
  - [x] script plugin
- [x] 搜索
  - [x] 指定目标索引
  - [x] 指定分词器
- [ ] web页面
- [ ] 心跳监听，自动重启
- [ ] 获取服务状态
- [ ] 日志

## 贡献

还处于项目初期，有些功能尚不完善。如果您有任何建议或者发现了bug，欢迎提出issue或者提交PR。
