# fclone_api

Golang API for NyaMisty/fclone

## Example usage

- First, initialize rclone context
```
g_rclone := &fclone_api.RcloneUtil{
    RcloneMode:   "rc",
    RcloneRcAddr: RunArg.RcAddr,
    RcloneRcAuth: RunArg.RcAuth,
    MaxTransfer:  int64(RunArg.MaxTransfer),
}
g_rclone.Init()
log.Infof("Initialized Rclone with arg: %v", g_rclone)
```

- Then, upload files
```
g_rclone.RcatSize(RunArg.UploadPath+"/"+path, size, modTime, RunArg.UploadBuffer, func(resp interface{}, err error) {
    if err != nil {
        log.Errorf("StreamFactory handler rcatSize(%s, %d) failed, resp %v, err: %v", path, size, resp, err)
        failedItems = append(failedItems, FailItemInfo{archiveIndex, path})
    }
    log.Infof("StreamFactory handler rcatSize(%s, %d) finished, resp %v err %v", path, size, resp, err)
})
```

- Finally, **BESURE** to wait for async reqs to allow fclone_api to commit transfers
```
g_rclone.WaitAllAsyncReq()
```