package consts

type MenuDish struct {
	// @tg desc=`айди блюда`
	// @tg example=10
	Id int64 `json:"id"`
	// @tg desc=`название блюда`
	// @tg example="бердра барбекю"
	Name string `json:"name"`
	// @tg desc=`категория блюда`
	// @tg example="Закуски"
	Category string `json:"category"`
	// @tg desc=`было ли выбрано блюдо`
	// @tg example="true"
	Choice bool `json:"choice"`
}
