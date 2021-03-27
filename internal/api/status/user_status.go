package status

type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	// success status code
	UserCreateCode = 201001
	UserLoginCode  = 201002
	UserTokenCode  = 201003

	// error status code
	UserParamErrorCode  = 201011
	UserCreateErrorCode = 201012
	UserLoginErrorCode  = 201013
	UserTokenErrorCode  = 201014
	UserOtherErrorCode  = 201010
)

var codeText = map[int]string{
	UserCreateCode: "The user to create a successfully",
	UserLoginCode:  "User login successfully",
	UserTokenCode:  "Token obtained successfully",

	UserParamErrorCode:  "User request parameter is wrong, please try again",
	UserCreateErrorCode: "User creation failed. Please try again",
	UserLoginErrorCode:  "User login failed, please login again",
	UserTokenErrorCode:  "Your account is private. Access token been expired",
	UserOtherErrorCode:  "An unknown error has occurred",
}
