if err := unix.IoctlSetInt(0, unix.TIOCSCTTY, 0); err != nil {
  return err
}
