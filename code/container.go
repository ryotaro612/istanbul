type Container struct {
  // コンテナのID
  id string
  stateDir string
  // spec.jsonの内容
  config *configs.Config
  # 中略
}
