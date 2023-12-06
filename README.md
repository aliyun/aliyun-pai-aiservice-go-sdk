# aliyun-pai-aiservice-go-sdk

## 安装
```
go get github.com/aliyun/aliyun-pai-aiservice-go-sdk
```
## 使用说明
### AI 图片检测代码示例

```go
import sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"

host  := "http://ai-service.ce8cc13b6421545749e7b4605f3d02607.cn-hangzhou.alicontainer.com"
appId := os.Getenv("APPID")
token := os.Getenv("TOKEN")

client := sdk.NewClient(host, appId, token)

// 检查图片列表
images := []string{"https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/1.jpg", "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/0.jpg"}
response, err := client.AIGCApi.AigcImageCheck(images, "", nil)

if err != nil {
	log.Fatal(err)
    return
}

if response.Code != "OK" {
    log.Fatal("response error:%v", response)
    return
}

// 遍历图片检测结果
for _, result := range response.Data.CheckResults {
    fmt.Printf("code=%d\tfrontal=%v\turl=%s\n", result.Code, result.Frontal, result.Url)
}

```

### AI图片，训练，写真完整示例参考

```go
import sdk "github.com/aliyun/aliyun-pai-aiservice-go-sdk"

host  := "http://ai-service.ce8cc13b6421545749e7b4605f3d02607.cn-hangzhou.alicontainer.com"
appId := os.Getenv("APPID")
token := os.Getenv("TOKEN")

client := sdk.NewClient(host, appId, token)
images := []string{"https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/1.jpg", "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_images/foto/train/0.jpg"}
response, err := client.AIGCApi.AigcImagesTrain(images, "", nil)
if err != nil {
    log.Fatal(err)
}

// 任务ID
jobId := response.Data.JobId
modelId := response.Data.ModelId

for {
    jobResponse, err := client.JobApi.GetAsyncJobWithId(jobId)
    if err != nil {
        log.Fatal("get job fail", err)
    }

    if jobResponse.Data.Job.State == sdk.JOB_STATE_WAIT {
        fmt.Println("job running")
    } else if jobResponse.Data.Job.State == sdk.JOB_STATE_SUCCESS { // job success
        fmt.Println(jobResponse)
        fmt.Println("job success")
        break
    } else if jobResponse.Data.Job.State == sdk.JOB_STATE_FAILED {
        log.Fatal("job fail", err)
        return
    }

    time.Sleep(30 * time.Second)
}

createResponse, err := client.AIGCApi.AigcImagesCreate(modelId, "https://pai-aigc-photog-bj.oss-cn-beijing.aliyuncs.com/photog/user_weights/foto10/validation/global_step_Blue_1_100_0.jpg", "", nil)
if err != nil {
    log.Fatal(err)
}

// 生成的图片 base64 编码
fmt.Println(createResponse.Data.Image)

```
