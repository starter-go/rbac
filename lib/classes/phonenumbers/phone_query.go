package phonenumbers

// PhoneNumberQuery 查询参数
type PhoneNumberQuery struct {
	Pagination Pagination
	All        bool // 查询全部条目
	Want       *PhoneNumberDTO

	Q string // the query.text
	A []any  // the query.args

}
