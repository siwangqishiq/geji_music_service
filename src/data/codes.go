package data

const HTTP_CODE_SUCCESS int = 200
const HTTP_CODE_SERVER_ERR int = 400
const HTTP_CODE_CLIENT_ERR int = 500

const ERR_CODE_NOMID int = 50001          //无mid参数
const ERR_CODE_NOTFOUND_MUSIC int = 50002 //未发现音乐

const ERR_CODE_DATABASE_ERROR = 50003 //数据库错误

const ERR_CODE_ACCOUNT_NOT_EXIST = 60003      //账号不存在
const ERR_CODE_ACCOUNT_REPEAT = 60004         //账号重复
const ERR_CODE_LOGIN_PWDERROR = 60005         //登录密码错误
const ERR_CODE_LOGIN_GEN_TOKEN_FAILED = 60006 //登录token生成错误
