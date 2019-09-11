package tests

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

func checkC1(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	// fmt.Println("In Condition C1")
	return true
}

func actionA1(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	// fmt.Println("In Action A1 Start")

	// change t1 field
	t1 := tuples["t1"].(model.MutableTuple)
	t1.SetString(ctx, "p3", "somethingnew")

	// fmt.Println("In Action A1 End")
	firedMap := ruleCtx.(map[string]string)
	firedMap["A1"] = "Fired"
}

func checkC2(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	fmt.Println("In Condition C2")
	return true
}

func actionA2(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Println("In Action A2 Start")
	t1 := tuples["t1"]
	val, _ := t1.GetString("p3")
	fmt.Println("In Action A2 End ", val)
	firedMap := ruleCtx.(map[string]string)
	firedMap["A2"] = "Fired"
}

func checkC3(ruleName string, condName string, tuples map[model.TupleType]model.Tuple, ctx model.RuleContext) bool {
	fmt.Println("In Condition C3")
	return true
}

func actionA3(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	fmt.Println("In Action A3 Start")
	t1 := tuples["t1"]
	val, _ := t1.GetString("p3")
	fmt.Println("In Action A3 End ", val)
	firedMap := ruleCtx.(map[string]string)
	firedMap["A3"] = "Fired"
}

func BenchmarkTestApiOne(b *testing.B) {
	rs, _ := createRuleSession()

	actionMap := make(map[string]string)

	// rule 1
	r1 := ruleapi.NewRule("R1")
	r1.AddCondition("C1", []string{"t1"}, checkC1, nil)
	r1.SetAction(actionA1)
	r1.SetPriority(1)
	r1.SetContext(actionMap)
	rs.AddRule(r1)

	rs.Start(nil)
	for n := 0; n < b.N; n++ {
		// s1 := rand.NewSource(time.Now().UnixNano())
		// r1 := rand.New(s1)

		// t1, _ := model.NewTupleWithKeyValues("t1", "Tom"+strconv.Itoa(r1.Intn(1000000000000000000)))
		t1, _ := model.NewTupleWithKeyValues("t1", "Tom"+randomString(randomBytes(16)))
		t1.SetString(nil, "p3", "test")
		err := rs.Assert(nil, t1)
		if err != nil {
			fmt.Println("Tom")
			b.Error(err)
			b.Fail()
		}
	}
	rs.Unregister()
}

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/"

func randomBytes(n int) []byte {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}

func randomString(bytes []byte) string {
	for i, b := range bytes {
		bytes[i] = letters[b%64]
	}

	return string(bytes)
}
