## syncパッケージ

非同期処理における排他制御や同期処理を実行するための機能を提供するパッケージ

1. [Mutexによる排他制御(Mutex.Lock、Mutex.Unlock)](./sync_sample_001.go)
2. [RWMutexによる排他制御(RWMutex.RLock、RWMutex.RUnlock)](./sync_sample_002.go)
3. [WaitGroupによるゴルーチンの完了待ち(WaitGroup.Add、WaitGroup.Done、WaitGroup.Wait)](./sync_sample_003.go)