## 音素 单字 单词 的存储方式
目录
```bash
.
|--音素
|  |--a
|--单字
|  |--你
|--单词
|  |--你好
```
### 音素
```bash
a
|--a.mp4 # 音素的视频
|--desc.txt # 音素的发音描述
|--spell.jpg # 音素的发音图
|--voice.mp3 # 音素的发音
|--lip.jpg # 唇形图
```

### 单字
```bash
你
|--你.mp4 # 单字的视频
|--desc.txt # 单字的发音描述
|--spell.jpg # 单字的发音图
|--spell.txt # 单字的拼音, 里面写"nǐ"(不带引号)
|--voice.mp3 # 单字的发音
|--lip.jpg # 唇形图
|--video.mp4 # 单字的视频
|--meta.txt # 单字的元数据
```
meta.txt的内容为:
```bash
你
nǐ
```

### 单词
```bash
你好
|--你好.mp4 # 单词的视频
|--video.mp4 # 单词的视频
|--meta.txt # 单词的元数据
```
meta.txt的内容为:
```bash
你好
nǐ hǎo
```
