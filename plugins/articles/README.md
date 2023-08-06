# Article

通过浏览器插件的形式，来获取页面的内容，主要用于保存某些文章。

## 使用

浏览器导入frontend文件夹下的插件。

> 浏览器插件通过control+/来触发保存面板

```bash
# ./articles <serverUrl> <port>
./articles http://localhost:18000 18001
```

> 注意如果修改了port或者浏览器端和articles服务不在同一台电脑，需要同步修改`frontend/config.js`中的路径，重新导入插件
> 例如：`var remote_path = 'http://10.0.0.2:18001/';`

### controller config.json service 配置示例

```json5
{
  "name": "Articles",
  "description": "...",
  "version": "0.0.1",
  "author": "LZ",
  "relative_path": "./plugins/article/articles",
  "es_index_setting": {
    "name": "articles",
    "setting": "{\"mappings\":{\"properties\":{\"title\":{\"type\":\"text\"},\"content\":{\"type\":\"text\"},\"tag\":{\"type\":\"text\"},\"url\":{\"type\":\"text\"},\"recordTime\":{\"type\":\"date\"}}}}",
    "mapping": {
      "id": "title",
      "title": "title",
      "content": "content",
      "source": "url",
      "tag": "tag"
    }
  },
  "manual_execution_params": "http://localhost:18000 18001"
}
```