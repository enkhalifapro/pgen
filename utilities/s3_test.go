package utilities

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"
)

func TestS3(t *testing.T) {
	t.Skip("test s3 working")
	ctx := context.Background()
	uv := "meta1value"
	m := map[string]*string{"meta1": &uv}
	s := S3{}
	if err := s.Put(ctx, "test/key/1", bytes.NewReader([]byte("Test value")), m); err != nil {
		t.Error(err)
		return
	}

	body, meta, err := s.Get(ctx, "test/key/1")
	if err != nil {
		t.Error(err)
		return
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("data was recieved", func(t *testing.T) {
		if "Test value" != string(data) {
			t.Fail()
			return
		}
	})

	t.Run("metadata was recieved", func(t *testing.T) {
		// s3 capitalize keys in map.
		if v, ok := meta["Meta1"]; !ok || *v != "meta1value" {
			t.Fail()
		}
	})
}
