package main
 
import "testing"


func Dummy() int {
    return 1
}

func TestDummy(t *testing.T){
    got := Dummy()
    if got != 1 {
        t.Errorf("1 = %d (wtf?); want 1", got)
    }
}

