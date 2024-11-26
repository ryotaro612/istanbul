type Container struct {
 // 続き
  initProcess          parentProcess
  initProcessStartTime uint64
  // 起動の排他制御
  m           sync.Mutex	
  // 中略
}
