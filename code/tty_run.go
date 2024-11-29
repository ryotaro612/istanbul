func (t *tty) recvtty(socket *os.File) (Err error) {
  // 中略
  f, err := utils.RecvFile(socket)
  go func() {_, _ = io.Copy(epollConsole, os.Stdin)}()
  t.wg.Add(1)
  go t.copyIO(os.Stdout, epollConsole)
  // 中略
}
