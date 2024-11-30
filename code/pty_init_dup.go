fd, err := unix.Open(slavePath, unix.O_RDWR, 0)
// 中略
for _, i := range []int{0, 1, 2} {
  if err := unix.Dup3(fd, i, 0); err != nil {
	  return err
