cleanup() {
  rm -rf p1 p2
}

MY_GO_PATH="/mnt/c/Program\ Files/Go/bin/go.exe"

cleanup
mkfifo p1 p2
$MY_GO_PATH run other/other.go > p1 < p2 &
$MY_GO_PATH run main.go < p1 > p2
cleanup
