# Rss Pull

拉取RSS源的插件。

> 配合[RssHub](https://docs.rsshub.app/)使用就可以实现拉取大量RSS源的功能。

## 使用

```bash
./rss_pull ./config.json normal
```

> 第二个参数为配置文件中的`resources`中的key，用于指定拉取哪些源。
> 将命令拆分配置到controller中的config.json中

### controller config.json script 配置示例

```json5
{
  "name": "Rss_Pull",
  "description": "...",
  "version": "0.0.1",
  "author": "LZ",
  "relative_path": "./plugins/rss_pull/rss_pull",
  "es_index_setting": {
    "name": "rss_pull",
    "setting": "{\"mappings\":{\"properties\":{\"title\":{\"type\":\"text\"},\"source_url\":{\"type\":\"text\"},\"tag\":{\"type\":\"text\"},\"content\":{\"type\":\"text\"}}}}",
    "mapping": {
      "id": "source_url",
      "title": "title",
      "content": "content",
      "source": "source_url",
      "tag": "tag"
    }
  },
  "commands": [
    {
      "params": "./plugins/rss_pull/config.json normal",
      "cron": "0 0 0 * * *"
    }
  ],
  "manual_execution_params": "./plugins/rss_pull/config.json normal"
}
```

## 配置项

```json5
{
  // 服务地址
  "serverUrl": "http://10.0.0.2:18000",
  // 接口路径
  "saveDataUrlSuffix": "/data/doc/save/data"
  // rss源
  "resources": {
    // 可以随意指定名称，用于标志拉取哪些源
    "normal": [
      // rss 源地址
      "https://xxxxxx/rss"
    ],
    "twitter": [
      // RssHub 源地址
      "https://xxxxxx/twitter/user/xxxx"
    ]
  }
}
```