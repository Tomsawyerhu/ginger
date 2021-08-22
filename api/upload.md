### 通用参数
|参数名|参数格式|参数说明|
|-----|------|-------|
|SessionId|string|会话id|

# API
### 文件上传
API名称:upload  
API请求类型:POST  
API参数:

|参数名|参数格式|参数说明|
|-----|------|-------|
|FileName|string|文件名|
|FileSize|int|文件大小|
|FileFormat|string|文件格式|
|FileAuth|int|文件权限|
|NeedBak|int|是否需要备份|
|SegmentId|int|块号|
|SegmentSize|int|块大小|
|IsLastSegment|int|是否是最后一块|
|NeedCheck|int|是否需要校验|
|CheckSum|string|校验和(只有最后一个文件块的请求携带)|
|CheckMethod|string|校验算法|

