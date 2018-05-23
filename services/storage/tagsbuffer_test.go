package storage

import (
	"reflect"
	"testing"

	"github.com/influxdata/influxdb/models"
)

func Test_tagsBuffer_copyTags(t *testing.T) {
	type fields struct {
		sz  int
		i   int
		buf models.Tags
	}
	type args struct {
		src models.Tags
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.Tags
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := &tagsBuffer{
				sz:  tt.fields.sz,
				i:   tt.fields.i,
				buf: tt.fields.buf,
			}
			if got := tb.copyTags(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagsBuffer.copyTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTagsBuffer_CopyTags(b *testing.B) {

}
