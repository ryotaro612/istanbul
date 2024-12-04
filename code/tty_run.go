f, err := utils.RecvFile(socket)
// 中略 socketからマスタの記述子を取得
go func() {
  _, _ = io.Copy(epollConsole, os.Stdin)
}()
t.wg.Add(1)
go t.copyIO(os.Stdout, epollConsole)
