# GitHub Stars

拉取GitHub中指定用户的stars。

## 使用

```bash
# 路由地址可以不指定
# ./github_stars <failedFilePath> <username> <serverUrl default:http://localhost:18000> <proxy:http://127.0.0.1:7890>
./github_stars ./failedFile liCells http://localhost:18000 http://127.0.0.1:7890
```

### controller config.json script 配置示例

```json5
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
  "commands": [
    {
      "params": "./plugins/stars/failedFile liCells http://localhost:18000",
      "cron": "0 0 0 * * *"
    }
  ],
  "manual_execution_params": "./plugins/stars/failedFile liCells http://localhost:18000"
}
```