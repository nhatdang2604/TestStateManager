package helpers

import (
	"log"
	"strings"
)

func GetBatchInsertQuery(
	tableName string,
	columnNames []string,
	insertData [](map[string]interface{}),
) (sql string, bindingValues []interface{}) {

	var stringBuilder strings.Builder
	stringBuilder.WriteString("INSERT INTO " + tableName + "(" + strings.Join(columnNames, ", ") + ") VALUES ")

	var questionArray []string
	var questionString string
	var questionStringArray []string
	bindingValues = []interface{}{}

	//Build []string{"?", "?", "?", ....}
	for range columnNames {
		questionArray = append(questionArray, "?")
	}

	//Build "(?, ?, ?, ...., ?)"
	questionString = "(" + strings.Join(questionArray, ", ") + ")"

	//Build []string{"(?, ?, ?, ...., ?)", "(?, ?, ?, ...., ?)", "(?, ?, ?, ...., ?)", ..., "(?, ?, ?, ...., ?)"}
	for _, rowData := range insertData {
		questionStringArray = append(questionStringArray, questionString)

		//Build the bindingValues for "?"
		for _, columnName := range columnNames {
			bindingValues = append(bindingValues, rowData[columnName])
		}
	}

	//Build "(?, ?, ?, ...., ?), (?, ?, ?, ...., ?), (?, ?, ?, ...., ?), ... (?, ?, ?, ...., ?)"

	//Build "INSERT INTO tableName(col1, col2,...,coln) VALUES (?, ?, ?, ...., ?), (?, ?, ?, ...., ?), (?, ?, ?, ...., ?), ... (?, ?, ?, ...., ?)"
	stringBuilder.WriteString(strings.Join(questionStringArray, ", "))
	sql = stringBuilder.String()

	log.Printf("SQL: %v, with value: %v", sql, bindingValues)
	return sql, bindingValues
}
