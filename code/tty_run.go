f, err := utils.RecvFile(socket)
go func() {
  _, _ = io.Copy(epollConsole, os.Stdin)
}()
t.wg.Add(1)
go t.copyIO(os.Stdout, epollConsole)
