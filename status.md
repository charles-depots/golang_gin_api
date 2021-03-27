## Status Code And Message

HTTP status codes as registered with IANA.
See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml

### 规则:

* 遵循 RFC 7231 规范
* 前三位为 RFC 状态码，第四位代表模块，第五位代表成功或失败，第六位为编号

### 实例：

> 203101

203: RFC 状态码

1: 基础模块

0：错误， 1: 成功

0：状态码编号, 自增, 0代表未知错误

### 基础模块 1

|  Field  |  Code   | Message |Create_time |Update_time|
| ---   |  ---  | ---  |  ---  |  ---  |
|  StatusOK  |  200111  |  OK  |  2021-03-27  |  ---  |
|  BadRequest  |  404102  |  Request was missing the 'redirect_uri' parameter.   |  2021-03-27  |  ---  |

### 用户模块 2

|  Field  |  Code   | Message |Create_time |Update_time|
| ---   |  ---  | ---  |  ---  |  ---  |
|  OAuthorTokenExpire  |  203201  | Invalid JWT token. The token is expired.   |  2021-03-27  |  ---  |
|  OAuthorInvalidToken  |  401201  |  Invalid token (token not found)  |  2021-03-27  |  ---  |
|  OAuth2AccessDenied  |  401202  |  Invalid token does not contain user information.  |  2021-03-27  |  ---  |
|  ...  |    |    |    |    |