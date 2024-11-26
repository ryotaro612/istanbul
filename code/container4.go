type Container struct {	
  criuVersion int
  state       containerState
  created     time.Time
  // 名前つきパイプ
  fifo *os.File	
}
