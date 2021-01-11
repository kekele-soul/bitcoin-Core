package structToStr

import (
	"fmt"
	"reflect"
)

func ToStr(entity interface{}) (str string) {
	typeOf := reflect.TypeOf(entity)
	if typeOf.Kind() != reflect.Struct {
		return "Unstructured calls"
	}
	valueOf := reflect.ValueOf(entity)
	//str += "{\n"
	str += "{<br>"
	for i := 0; i < typeOf.NumField(); i++ {
		name := typeOf.Field(i).Name
		type_ := typeOf.Field(i).Type
		value := valueOf.FieldByIndex([]int{i})
		switch type_.Kind() {
		case reflect.Struct:
			structName := fmt.Sprintf("%q", name)
			str += structName + ":"
			str += ToStr(value.Interface())
		case reflect.Map:
			//还没想好怎么搞
			continue
		case reflect.Slice:
			//str += fmt.Sprintf("%q: [\n", name)
			str += fmt.Sprintf("%q: [<br>", name)
			for i := 0; i < value.Len(); i++ {
				if value.Index(i).Kind() == reflect.Struct {
					str += ToStr(value.Index(i).Interface())
				}else {
					str += fmt.Sprintf("%v\n", value.Index(i))
				}
			}
			//str += fmt.Sprintf("]\n")
			str += fmt.Sprintf("]<br>")
		default:
			value_ := fmt.Sprintf("%v", value.Interface())
			if type_.Kind() == reflect.String {
				//str += fmt.Sprintf("%q:%q,\n", name, value_)
				str += fmt.Sprintf("\"%s\": \"%s\",<br>", name, value_)
			} else {
				//str += fmt.Sprintf("%q:%s,\n", name, value_)
				str += fmt.Sprintf("\"%s\": %s,<br>", name, value_)
			}
		}
	}
	//str += "},\n"
	str += "},<br>"
	return
}
