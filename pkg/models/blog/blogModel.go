package blog

type BlogModel struct {
	BlogID      string `JSON:"blogID"       dynamodbav:"blogID"`
	BlogTitle   string `JSON:"blogTitle"    dynamodbav:"blogTitle"`
	BlogContent string `JSON:"blogContent"  dynamodbav:"blogContent"`
	CreatedBy   string `JSON:"createdBy"    dynamodbav:"createdBy"`
	Status      string `JSON:"status"       dynamodbav:"status"`
	CreatedAt   string `JSON:"createdAt"    dynamodbav:"createdAt"`
}

type BlogsWithPagination struct {
	Blogs             *[]BlogModel `JSON:"blogID"`
	Limit             string       `JSON:"limit"`
	FirstEvaluatedKey string       `JSON:"previousKey"`
	LastEvaluatedKey  string       `JSON:"nextKey"`
}
