package union

import "github.com/iGoogle-ink/gopay"

type UnionPayClient struct {
	MerId  string
	IsProd bool
}

func (this *UnionPayClient) Consume(bm gopay.BodyMap) {

}
