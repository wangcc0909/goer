package db

func Insert(bean interface{}) error {
	_, err := database.Insert(bean)
	return err
}
