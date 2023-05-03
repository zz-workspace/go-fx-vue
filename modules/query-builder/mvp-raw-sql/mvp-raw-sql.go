package mvprawsql

import (
	"fmt"
	"regexp"
	"strings"

	queryBuilder "fast-api.io/modules/query-builder"
	"github.com/tidwall/gjson"
)

func CreateQueryBuilder(
	input gjson.Result,
	function gjson.Result,
	bodyInput queryBuilder.TestEndpointRequest,
) string {
	functionContext := function.Map()["context"]
	query := functionContext.Map()["dbo"].Map()["raw"].Str
	re := regexp.MustCompile(`\{\{\s*\$(\w+)\.(\w+)\s*\}\}`)
	query = re.ReplaceAllStringFunc(query, func(s string) string {
		return "$" + re.ReplaceAllString(s, "$1.$2")
	})
	for _, input := range input.Array() {
		inputName := input.Map()["name"].Str
		value := gjson.Get(string(bodyInput.InputData), inputName)
		query = strings.ReplaceAll(query, fmt.Sprintf("$input.%s", inputName), value.String())
	}

	return query
}
