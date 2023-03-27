package learn

import (
	"testing"
)

func TestPostVideoPath(t *testing.T) {
	resp, err, ok := PostVideoPath("/root/src/user/standard/1.mp4")
	if err != nil {
		t.Errorf("err = %v", err)
		t.Errorf("ok = %v", ok)
	}
	if ok {
		t.Errorf("ok = %v", ok)
	}
	t.Logf("resp = %v", resp)
}
