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

func GetUpdateQuery(
	tableName string,
	idColumnName string,
	updateData map[string]interface{},
) (sql string, bindingValues []interface{}) {

	bindingValues = make([]interface{}, len(updateData))

	//Get array of column's name with " = ?" postfix
	var columnNameSetters []string = make([]string, len(updateData)-1)
	var idx int = 0
	for columnName, columnValue := range updateData {

		//Ignore the id column name, because we do not update this column
		if idColumnName != columnName {
			columnNameSetters[idx] = columnName + " = ?"
			bindingValues[idx] = columnValue

			idx++
		}
	}

	//Get id set mask with "?"

	sql = strings.Join([]string{
		"UPDATE", tableName,
		"SET", strings.Join(columnNameSetters, ", "),
		"WHERE", idColumnName, "= ?",
	}, " ")

	//Append the primary key in the last, for the WHERE statement
	bindingValues[idx] = updateData[idColumnName]

	//Logging
	log.Printf("SQL: %v, with value: %v", sql, bindingValues)

	return sql, bindingValues
}

func GetFindByIdQuery(
	tableName string,
	columnNames []string,
	searchIds []interface{},
	idColumnName string,
	isLimit bool,
) (sql string, bindingValues []interface{}) {

	var limit string = ""
	if isLimit {
		limit = "LIMIT 1"
	}

	var idMaskArray []string = make([]string, len(searchIds))
	for idx, _ := range searchIds {
		idMaskArray[idx] = "?"
	}
	var idMaskSet string = "(" + strings.Join(idMaskArray, ", ") + ")"

	sql = strings.Join([]string{
		"SELECT", strings.Join(columnNames, ", "),
		"FROM", tableName,
		"WHERE", idColumnName, "IN", idMaskSet,
		limit,
	}, " ")

	bindingValues = searchIds

	log.Printf("SQL: %v, with value: %v", sql, bindingValues)

	return sql, bindingValues
}
