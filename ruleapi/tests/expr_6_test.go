package tests

import (
	"context"

	"github.com/project-flogo/rules/common/model"
)

//1 arithmetic operation

func checkC6(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	return true
}

func a6(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
}

// func BenchmarkTestExp(b *testing.B) {

// 	rs, _ := createRuleSession()
// 	r1 := ruleapi.NewRule("r1")
// 	r1.AddExprCondition("c1", "$.t2.p2 > $.t1.p1", nil)
// 	r1.SetAction(a6)
// 	r1.SetContext(nil)

// 	rs.AddRule(r1)

// 	rs.Start(nil)
// 	b.ResetTimer()
// 	var ctx context.Context
// 	for n := 0; n < b.N; n++ {
// 		s1 := rand.NewSource(time.Now().UnixNano())
// 		r1 := rand.New(s1)
// 		a := strconv.Itoa(r1.Intn(1000000000000))
// 		t1, _ := model.NewTupleWithKeyValues("t1", "t1"+a)
// 		t1.SetInt(nil, "p1", n)
// 		t1.SetDouble(nil, "p2", float64(n)+1.3)
// 		t1.SetString(nil, "p3", "t3")

// 		ctx = context.WithValue(context.TODO(), TestKey{}, b)
// 		err := rs.Assert(ctx, t1)
// 		if err != nil {
// 			b.Error(err)
// 			b.Fail()
// 		}
// 		t2, _ := model.NewTupleWithKeyValues("t2", "t2"+a)
// 		t2.SetInt(nil, "p1", n)
// 		t2.SetDouble(nil, "p2", float64(n)+1.0001)
// 		t2.SetString(nil, "p3", "t3")

// 		ctx = context.WithValue(context.TODO(), TestKey{}, b)
// 		err = rs.Assert(ctx, t2)
// 		if err != nil {
// 			b.Error(err)
// 			b.Fail()
// 		}
// 		// fmt.Println(t1, t2)
// 	}
// 	rs.Unregister()

// }

// func BenchMarkTestExp5(b *testing.B) {
// 	rs, _ := createRuleSession()
// 	r1 := ruleapi.NewRule("r1")
// 	r1.AddExprCondition("c1", "(( $.t2.p1 - $.t1.p1  == 3) && ($.t1.p2 > $.t2.p2) && ($.t1.p3 == $.t2.p3))", nil)
// 	r1.AddCondition("C2", []string{"t1"}, checkC6, nil)
// 	r1.SetAction(a6)
// 	r1.SetContext(nil)

// 	rs.AddRule(r1)

// 	rs.Start(nil)

// 	var ctx context.Context
// 	for n := 0; n < b.N; n++ {
// 		s1 := rand.NewSource(time.Now().UnixNano())
// 		r1 := rand.New(s1)
// 		a := strconv.Itoa(r1.Intn(1000000000000))
// 		t1, _ := model.NewTupleWithKeyValues("t1", "t1"+a)
// 		t1.SetInt(nil, "p1", n+1)
// 		t1.SetDouble(nil, "p2", float64(n)+1.3)
// 		t1.SetString(nil, "p3", "t3"+a)

// 		ctx = context.WithValue(context.TODO(), TestKey{}, b)
// 		err := rs.Assert(ctx, t1)
// 		if err != nil {
// 			b.Error(err)
// 			b.Fail()
// 		}

// 		t2, _ := model.NewTupleWithKeyValues("t2", "t2"+a)
// 		t2.SetInt(nil, "p1", n+4)
// 		t2.SetDouble(nil, "p2", float64(n)+1.1)
// 		t2.SetString(nil, "p3", "t3"+a)

// 		ctx = context.WithValue(context.TODO(), TestKey{}, b)
// 		rs.Assert(ctx, t2)
// 		if err != nil {
// 			b.Error(err)
// 			b.Fail()
// 		}
// 	}
// 	rs.Unregister()
// }
