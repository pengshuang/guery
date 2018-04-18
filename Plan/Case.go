package Plan

import (
	"github.com/xitongsys/guery/Context"
	"github.com/xitongsys/guery/DataSource"
	"github.com/xitongsys/guery/parser"
)

type CaseNode struct {
	Whens []*WhenClauseNode
	Else  *ExpressionNode
}

func NewCaseNode(ctx *Context.Context, whens []parser.IWhenClauseContext, el parser.IExpressionContext) *CaseNode {
	res := &CaseNode{
		Whens: []*WhenClauseNode{},
		Else:  NewExpressionNode(ctx, el),
	}
	for _, w := range whens {
		res.Whens = append(res.Whens, NewWhenClauseNode(ctx, w))
	}
	return res
}

func (self *CaseNode) Result(input *DataSource.DataSource) interface{} {
	input.Reset()
	var res interface{}
	for _, w := range self.Whens {
		res = w.Result(input)
		if res != nil {
			return res
		}
	}
	res = self.Else.Result(input)
	return res
}

////////
type WhenClauseNode struct {
	Condition *ExpressionNode
	Res       *ExpressionNode
}

func NewWhenClauseNode(ctx *Context.Context, wh parser.IWhenClauseContext) *WhenClauseNode {
	tt := wh.(*parser.WhenClauseContext)
	ct, rt := tt.GetCondition(), tt.GetResult()
	res := &WhenClauseNode{
		Condition: NewExpressionNode(ctx, ct),
		Res:       NewExpressionNode(ctx, rt),
	}
	return res
}

func (self *WhenClauseNode) Result(input *DataSource.DataSource) interface{} {
	input.Reset()
	var res interface{} = nil
	if self.Condition.Result(input).(bool) {
		res = self.Res.Result(input)
	}
	return res
}
