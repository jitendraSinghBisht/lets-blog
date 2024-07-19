package user

type UserModel struct {
	UserId    string `JSON:"userId"       dynamodbav:"userId"`
	UserName  string `JSON:"userName"     dynamodbav:"userName"`
	UserEmail string `JSON:"userEmail"    dynamodbav:"userEmail"`
	Password  string `JSON:"password"     dynamodbav:"password"`
	CreatedAt string `JSON:"createdAt"    dynamodbav:"createdAt"`
}

type UserReqModel struct {
	UserName  string `JSON:"userName"`
	UserEmail string `JSON:"userEmail"`
	Password  string `JSON:"password"`
}
