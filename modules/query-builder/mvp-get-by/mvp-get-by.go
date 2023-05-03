package mvpgetby

import (
	"fmt"
	"strings"

	queryBuilder "fast-api.io/modules/query-builder"
	"github.com/tidwall/gjson"
)

func CreateQueryBuilder(
	input gjson.Result,
	function gjson.Result,
	bodyInput queryBuilder.TestEndpointRequest,
	tableName string,
) string {
	functionInput := function.Map()["input"]
	var whereConditions []string
	for _, condition := range functionInput.Array() {
		conditionMap := condition.Map()
		fieldNameMap := conditionMap["fieldName"].Map()
		fieldValueMap := conditionMap["fieldValue"].Map()
		conditionName := fieldNameMap["value"]
		var conditionValue string
		conditionOp := conditionMap["op"]
		// var fieldValue interface{}
		if fieldValueMap["type"].Str == "input" {
			// str, _ := lo.Find(endpointInput.Array(), func(input gjson.Result) bool {
			// 	return input.Map()["name"].Str == fieldValueMap["name"].Str
			// })
			conditionValue = gjson.Get(string(bodyInput.InputData), fieldValueMap["value"].Str).String()
		} else {
			conditionValue = fieldValueMap["value"].Str
		}
		fmt.Println(fieldValueMap["type"])
		whereCondition := fmt.Sprintf(
			"%s %s '%s'",
			conditionName,
			conditionOp,
			conditionValue,
		)
		whereConditions = append(whereConditions, whereCondition)
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT 1", tableName, strings.Join(whereConditions, " AND "))

	return query
}
