package Plan

import (
	"context"
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/parser"
)

type PlanOrderByNode struct {
	Input    PlanNode
	Output   PlanNode
	Metadata *Util.Metadata
}

func NewPlanOrderByNode(input, output PlanNode, items []parser.ISortItemContext) *PlanOrderByNode {
	return &PlanOrderByNode{
		Input:    input,
		Output:   output,
		Metadata: Util.NewDefaultMetadata(),
	}
}

func (self *PlanOrderByNode) GetNodeType() PlanNodeType {
	return ORDERBYNODE
}

func (self *PlanOrderByNode) String() string {
	res := "PlanOrderByNode {\n"
	res += "Input: " + self.Input.String() + "\n"
	res += "}\n"
	return res
}

func (self *PlanOrderByNode) GetMetadata() *Util.Metadata {
	return self.Metadata
}

func (self *PlanOrderByNode) SetMetadata() error {
	err := self.Input.SetMetadata()
	if err != nil {
		return err
	}
	self.Metadata.Copy(self.Input.GetMetadata())

}
