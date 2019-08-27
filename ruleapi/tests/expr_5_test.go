package tests

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

//1 arithmetic operation
func Test_5_Expr(t *testing.T) {

	actionCount := map[string]int{"count": 0}
	rs, _ := createRuleSession()
	r1 := ruleapi.NewRule("r1")
	r1.AddExprCondition("c1", "(($.t1.p1 + $.t2.p1) == 5) && (($.t1.p2 > $.t2.p2) && ($.t1.p3 == $.t2.p3))", nil)
	r1.SetAction(a5)
	r1.SetContext(actionCount)

	rs.AddRule(r1)

	rs.Start(nil)

	var ctx context.Context

	t1, _ := model.NewTupleWithKeyValues("t1", "t1")
	t1.SetInt(nil, "p1", 1)
	t1.SetDouble(nil, "p2", 1.3)
	t1.SetString(nil, "p3", "t3")

	ctx = context.WithValue(context.TODO(), TestKey{}, t)
	rs.Assert(ctx, t1)

	t2, _ := model.NewTupleWithKeyValues("t2", "t2")
	t2.SetInt(nil, "p1", 4)
	t2.SetDouble(nil, "p2", 1.1)
	t2.SetString(nil, "p3", "t3")

	ctx = context.WithValue(context.TODO(), TestKey{}, t)
	rs.Assert(ctx, t2)
	rs.Unregister()
	count := actionCount["count"]
	if count != 1 {
		t.Errorf("expected [%d], got [%d]\n", 1, count)
	}
}

func BenchmarkTestExp5(b *testing.B) {
	rs, _ := createRuleSession()
	r1 := ruleapi.NewRule("r1")
	r1.AddExprCondition("c1", "(( $.t2.p1 - $.t1.p1  == 3) && ($.t1.p2 > $.t2.p2) && ($.t1.p3 == $.t2.p3))", nil)
	r1.AddCondition("C2", []string{"t1"}, checkC1, nil)
	r1.SetAction(b5)
	r1.SetContext(nil)

	rs.AddRule(r1)

	rs.Start(nil)

	var ctx context.Context
	for n := 0; n < b.N; n++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		a := strconv.Itoa(r1.Intn(1000000000000))
		t1, _ := model.NewTupleWithKeyValues("t1", "t1"+a)
		t1.SetInt(nil, "p1", n+1)
		t1.SetDouble(nil, "p2", float64(n)+1.3)
		t1.SetString(nil, "p3", "t3"+a)

		ctx = context.WithValue(context.TODO(), TestKey{}, b)
		err := rs.Assert(ctx, t1)
		if err != nil {
			b.Error(err)
			b.Fail()
		}

		t2, _ := model.NewTupleWithKeyValues("t2", "t2"+a)
		t2.SetInt(nil, "p1", n+4)
		t2.SetDouble(nil, "p2", float64(n)+1.1)
		t2.SetString(nil, "p3", "t3"+a)

		ctx = context.WithValue(context.TODO(), TestKey{}, b)
		rs.Assert(ctx, t2)
		if err != nil {
			b.Error(err)
			b.Fail()
		}
		fmt.Println(t1, t2)
	}
	rs.Unregister()
}

func a5(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	t := ctx.Value(TestKey{}).(*testing.T)
	t.Logf("Test_5_Expr executed!")
	actionCount := ruleCtx.(map[string]int)
	count := actionCount["count"]
	actionCount["count"] = count + 1
}

func b5(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	// b := ctx.Value(TestKey{}).(*testing.B)
	// b.Logf("Test_5_Expr executed!")
	// actionCount := ruleCtx.(map[string]int)
	// count := actionCount["count"]
	// actionCount["count"] = count + 1
	fmt.Println("Test")
}
