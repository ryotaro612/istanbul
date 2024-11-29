pty, slavePath, err := console.NewPty()
// 中略
if err := utils.SendRawFd(
  socket, pty.Name(), pty.Fd()); err != nil {
  return err
}
