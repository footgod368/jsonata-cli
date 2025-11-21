package jsonata

import (
	"github.com/bytedance/mockey"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

const jsonString = `
    {
        "orders": [
            {"price": 10, "quantity": 3},
            {"price": 0.5, "quantity": 10},
            {"price": 100, "quantity": 1}
        ]
    }
`

const expr = `$sum(orders.(price*quantity)) = 135 ? 135 : -1`

func TestEval(t *testing.T) {
	mockey.PatchConvey("success", t, func() {
		result, err := Eval(jsonString, expr)
		convey.So(err, convey.ShouldBeNil)
		convey.So(result, convey.ShouldEqual, "135")
	})
}
