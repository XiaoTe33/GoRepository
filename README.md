# 苏宁易购实战项目

[![](https://img.shields.io/badge/docs-接_口_文_档-green.svg)](https://console-docs.apipost.cn/preview/964bf1a9921afa9b/b2c4952490094e3e) [![GoDoc](https://pkg.go.dev/badge/github.com/XiaeTe33)](https://github.com/XiaoTe33/winter-examination)[![](https://img.shields.io/badge/FE-柳亦钦_(PIPI--1021)-#c586c0.svg)](https://github.com/PIPI-1021)[![](https://img.shields.io/badge/BE-刘力延_ _(XiaoTe33)-#43a1de.svg)](https://github.com/XiaoTe33)

##  实现的功能

1. 用户

   - 基本功能

     - 用户注册

     - 登录（用户名和密码，token，二维码登录）

     - token保持登陆状态

     - 退出登录
     - 上传头像
     - 修改个人信息（头像，用户名等）
     - 充值余额，查看余额
     - 忘记密码

   - 项目功能

     - 搜索商品（支持按店铺查询、分类查询、模糊查询，可按销量、评分、价格排序）
     - 加入购物车，我的购物车，移除购物车
     - 收藏商品，查看我的收藏，取消收藏
     - 新增收货地址，我的收货地址，删除收货地址
     - 抢购优惠券（可处理高并发），我的优惠券
     - 下订单（可处理高并发），我的订单（可状态分类），订单状态（已付款，取消订单，确认收货），变更订单收货地址，订单号查询订单
     - 评价商品（文字，图片，评分），删除评价

2. 商家

   - 基本功能

     - 创建店铺，修改店铺信息，店铺公告
     - 上架商品，下架商品，进货，卸货，修改商品信息，查看本店商品

   - 其他功能

     - 发放优惠券（现金券，折扣券，满减券），优惠券有效期设置

     - 本店订单

3. 其他

   - 其他功能

     - 加载前端静态资源

     - 提供了直接获取数据库信息的接口，便于前端查看数据，~~让队友帮忙一起找bug~~
     - 一键导入原始数据
     - 重开（清空上传的图片，数据库信息），~~删库跑路~~



## 亮点

1. 加密存储

   - 密码加密存储

     ```go
     package service
     
     func Register(u model.UserRegisterReq) {
     	dao.AddUser(model.User{
     		Username: u.Username,
     		Password: utils.SHA256Secret(u.Password),//sha256加密
     		Phone:    u.Phone,
     		Email:    u.Email,
     	})
     }
     ```

     ```go
     package utils
     
     func SHA256Secret(str string) string {
     	h := sha256.New()
     	h.Write([]byte(str + "secret"))
     	return fmt.Sprintf("%x", h.Sum(nil))
     }
     
     func Md5Encoded(pre string) string {
     	return fmt.Sprintf("%x", md5.Sum([]byte(pre)))
     }
     
     // Md5EncodedWithTime 防止文件名重复
     func Md5EncodedWithTime(pre string) string {
     	return fmt.Sprintf("%x", md5.Sum([]byte(pre+fmt.Sprintf("%b", time2.Now().Nanosecond()))))
     }
     ```

     

2. 防止sql注入

   - 使用通配符防止sql注入

     ```go
     package dao
     
     func AddUser(user model.User) {
     	sqlStr := "insert into `users` ( `username`, `password`, `phone`, `email`) values ( ?, ?, ?, ?)"
     	_, err := Db.Exec(sqlStr, user.Username, user.Password, user.Phone, user.Email)
     	if err != nil {
     		fmt.Println("add user fail ...")
     		return
     	}
     }
     ```

     

3. 链接生成（头像上传，商品图片）

   - 上传文件

     ```go
     package app
     
     func AddUserPhoto(c *gin.Context) {
     	photo, err := c.FormFile("photo")
     	if handleError(c, err) {
     		return
     	}
     	ok, style := utils.IsValidPictureFile(photo.Filename)
     	if !ok {
     		jsonError(c, "仅支持jpg,png,jfif格式的图片")
     		return
     	}
     	err = c.SaveUploadedFile(photo, conf.LocalSavePathOfUserPhotos+c.GetString("userId")+style)//保存到本地
     	if handleError(c, err) {
     		return
     	}
     	service.AddUserPhoto(c.GetString("userId"))
     	jsonSuccess(c)
     }
     ```

     

   - 加载静态资源

     ```go
     const (
     GinPathOfUserPhoto          = "user/photo/" //路由地址
     GinPathOfGoodsPicture       = "goods/picture/"
     GinPathOfEvaluationPictures = "evaluation/picture/"
     GinPathOfQR                 = "qr/"
     )
     
     //加载静态资源
     r.Static(conf.GinPathOfQR, "./src/static/qr")
     r.Static(conf.GinPathOfUserPhoto, "./src/static/user/photos")
     r.Static(conf.GinPathOfGoodsPicture, "./src/static/goods/pictures")
     r.Static(conf.GinPathOfEvaluationPictures, "./src/static/evaluation/pictures")
     ```

     

   - 生成链接

     ```go
     package conf
     
     const(
     
         IP = "http://39.101.72.18"
         
         FrontEndPort = ":9091"
     	BackEndPort  = ":9090"
         
     	WebLinkPathOfUserPhoto          = IP + FrontEndPort + "/" + GinPathOfUserPhoto //链接地址
     	WebLinkPathOfGoodsPicture       = IP + FrontEndPort + "/" + GinPathOfGoodsPicture
     	WebLinkPathOfEvaluationPictures = IP + FrontEndPort + "/" + GinPathOfEvaluationPictures
     	WebLinkPathOfQR                 = IP + FrontEndPort + "/" + GinPathOfQR
     )
     ```

4. 校验

   - gin绑定器校验和自定义错误信息

     - 绑定器实现

     ```go
     package app
     
     func Register(c *gin.Context) {
     	var user model.UserRegisterReq
     	err := c.ShouldBind(&user)//gin绑定
     	if handleBindingError(c, err, &user) {//处理绑定时错误
     		return
     	}
     	service.Register(user)
     	jsonSuccess(c)
     }
     ```

     ```go
     package model
     
     type UserRegisterReq struct {
     	Username   string `json:"username"   binding:"required,max=20,IsValidUsername"       form:"username"   err:"用户名过长或已被注册"`
     	Password   string `json:"password"   binding:"required,min=6,max=20,IsValidPassword" form:"password"   err:"密码长度应为6~20"`
     	RePassword string `json:"rePassword" binding:"required,eqfield=Password"             form:"rePassword" err:"两次密码不一致"`
     	Email      string `json:"email"      binding:"required,IsValidEmail"                 form:"email"      err:"邮箱格式不正确或已被注册"`
     	Phone      string `json:"phone"      binding:"required,IsValidPhone"                 form:"phone"      err:"手机号格式不正确或已被注册"`
     }
     ```

     ```go
     package utils
     
     
     func InitUserVal() {
         //注册绑定器
         
     	addValidator("IsValidUsername", IsValidUsername)
     	addValidator("IsValidEmail", IsValidEmail)
     	addValidator("IsValidPhone", IsValidPhone)
     	addValidator("IsValidPassword", IsValidPassword)
     	addValidator("IsRegisteredUsername", IsRegisteredUsername)
     }
     
     
     func addValidator(tag string, fun validator.Func) {
         //注册
     	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
     		_ = v.RegisterValidation(tag, fun)
     	}
     }
     
     //以下为校验函数
     
     func IsRegisteredUsername(fl validator.FieldLevel) bool {
     	username, _ := fl.Field().Interface().(string)
     	return dao.QueryUserByUsername(username) != model.User{}
     }
     
     func IsValidPassword(fl validator.FieldLevel) bool {
     	password, _ := fl.Field().Interface().(string)
     	return regexp.MustCompile(`^[!-~]{6,20}$`).MatchString(password)
     }
     
     func IsValidUsername(fl validator.FieldLevel) bool {
     	username, _ := fl.Field().Interface().(string)
     	return dao.QueryUserByUsername(username) == model.User{}
     }
     
     func IsValidEmail(fl validator.FieldLevel) bool {
     	email, _ := fl.Field().Interface().(string)
     	return regexp.MustCompile(`^(\w+(.\w+)*)*@(\w+(.\w+)+)+$`).MatchString(email) &&
     		dao.QueryUserByEmail(email) == model.User{}
     }
     
     func IsValidPhone(fl validator.FieldLevel) bool {
     	phone, _ := fl.Field().Interface().(string)
     	return regexp.MustCompile("^1[3-9][0-9]{9}$").MatchString(phone) &&
     		dao.QueryUserByPhone(phone) == model.User{}
     }
     
     ```

   - 自定义错误信息的实现

     ```go 
     package app
     
     // 绑定器处理错误
     func handleBindingError(c *gin.Context, err error, obj any) bool {
     	msg := getErrTag(err, obj)
     	if msg != "" {
     		jsonError(c, msg)
     		return true
     	}
     	return false
     }
     
     func jsonError(c *gin.Context, err string) {
     	c.AbortWithStatusJSON(400, gin.H{
     		"status": http.StatusBadRequest,
     		"msg":    err,
     	})
     }
     
     // getErrTag 返回结构体中的msg参数
     // 此函数模仿自(http://docs.fengfengzhidao.com)
     func getErrTag(err error, obj any) string {
     	// 使用的时候，需要传obj的指针
     	getObj := reflect.TypeOf(obj)
     	// 将err接口断言为具体类型
     	if errs, ok := err.(validator.ValidationErrors); ok {
     		// 断言成功
     		for _, e := range errs {
     			// 循环每一个错误信息
     			// 根据报错字段名，获取结构体的具体字段
     			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
     				msg := f.Tag.Get("err")
     				return msg
     			}
     		}
     	}
     	if err != nil {
     		return err.Error()
     	} else {
     		return ""
     	}
     }
     ```

   - 定义接口处理错误

     ```go
     package app
     
     // 接口处理错误
     func handleValidReq(c *gin.Context, req model.Req) bool {
     	if handleError(c, req.Valid()/*调用接口的校验方法*/) {
     		return true
     	}
     	return false
     }
     ```

     ```go
     package model
     
     
     //校验接口
     type Req interface {
     	Valid() error
     }
     
     type AddCouponReq struct {
     	//...
     }
     
     //实现接口
     func (req AddCouponReq) Valid() error {
         //优惠券的入参校验
     	if req.Kind == "0" {
     		if !isValidDiscount0(req.Discount) {
     			return errors.New("折扣格式不正确")
     		}
     	}
     	if req.Kind == "1" {
     		if !isValidDiscount1(req.Discount) {
     			return errors.New("折扣格式不正确")
     		}
     	}
     	if req.Kind == "2" {
     		if !isValidDiscount2(req.Discount) {
     			return errors.New("折扣格式不正确")
     		}
     	}
     	if !isValidTime(req.BeginAt) {
     		return errors.New("开始时间格式不正确")
     	}
     	if !isValidTime(req.EndAt) {
     		return errors.New("结束时间格式不正确")
     	}
     	begin, _ := time.Parse(`2006-01-02 15:04:05`, req.BeginAt)
     	end, _ := time.Parse(`2006-01-02 15:04:05`, req.EndAt)
     	if begin.Unix() > end.Unix() {
     		return errors.New("开始时间应早于结束时间")
     	}
     	return nil
     }
     ```

5. 优惠券

   - 优惠功能的实现

     ```go
     package model
     
     // Coupon 优惠券
     type Coupon struct {
     	Id       string `json:"id"`
     	ShopId   string `json:"shopId"`
     	Name     string `json:"name"`
     	Kind     string `json:"kind"` //优惠券种类 0.现金券  1.折扣券  2.满减券
     	Amount   string `json:"amount"`
     	Discount string `json:"discount"` //现金券格式:一位或两位小数 , 折扣券格式: 小于1的一位或两位小数 , 满减券格式: "a,b"(满a减b)
     	BeginAt  string `json:"beginAt"`
     	EndAt    string `json:"endAt"`
     }
     
     // Discounts 优惠券的打折方法
     func (c Coupon) Discounts(pre string) (cur string) {
     	if c.Kind == "0" {
     		discount, _ := strconv.ParseFloat(c.Discount, 64)
     		prePrice, _ := strconv.ParseFloat(pre, 64)
     		if discount >= prePrice {
     			return "0.00"
     		}
     		return fmt.Sprintf("%.2f", prePrice-discount)
     	}
     	if c.Kind == "1" {
     		discount, _ := strconv.ParseFloat(c.Discount, 64)
     		prePrice, _ := strconv.ParseFloat(pre, 64)
     		return fmt.Sprintf("%.2f", discount*prePrice)
     	}
     	if c.Kind == "2" {
     		split := strings.Split(c.Discount, ",")
     		if len(split) != 2 {
     			return pre
     		}
     		prePrice, _ := strconv.ParseFloat(pre, 64)
     		base, _ := strconv.ParseFloat(split[0], 64)
     		cutDown, _ := strconv.ParseFloat(split[1], 64)
     		return fmt.Sprintf("%.2f", prePrice-float64(int(prePrice)/int(base))*cutDown)
     	}
     	return ""
     }
     ```

6. 全局唯一id生成（生成订单号，商品id）

   - 订单号

     ```go
     package utils
     
     var orderIdLock = make(chan struct{}, 1)
     var a = map[string]int64{}
     
     func GetOrderId() string {
     	orderIdLock <- struct{}{}
        
     	DurationTimeStamp := time.Now().Unix() - conf.OrderIdBaseTimeStamp //2023-1-1 00:00:00
     	a[time.Now().Format("20060102")]++
     	num := a[time.Now().Format("20060102")]
         //雪花算法
     	sprintf := fmt.Sprintf("%d", DurationTimeStamp<<conf.OrderIdLeftShiftNumber|num)
     	<-orderIdLock
     	return sprintf
     }
     ```

   - 商品id类似，只换了个左移位数

7. 高并发处理（下订单、秒杀优惠券）

   - 下订单

     ```go
     package service
     
     func AddOrder2(req model.AddOrderReq, userId string) error {
     	var solderId string
     	//秒杀的原子性问题处理部分
     	for {
     		goods := dao.QueryGoodsById(req.GoodsId)
     		if goods != (model.Goods{}) {
     			solderId = goods.ShopId
     			goodsAmt, _ := strconv.Atoi(goods.Amount)
     			amt, err := strconv.Atoi(req.Amount)
     			if err != nil {
     				return errors.New("请输入整数")
     			}
     			//只有一个人抢到锁
     			dao.GoodsChan <- struct{}{}
     			//没库存就释放锁
     			if goodsAmt < amt {
     				<-dao.GoodsChan
     				return errors.New("库存不足")
     			}
     			//把修改前的数量goods.Amount也带上
     			//这样在修改数据的时候，一旦数据被已经修改过，就不能再修改第二次
     			if dao.DownGoodsAmountByOrder(req.GoodsId, req.Amount, goods.Amount) == 0 {
     				//修改失败再去枪锁
     				continue
     			} else {
     				//修改成功新增订单
     				price, _ := strconv.ParseFloat(goods.Price, 64)
     				originPrice := fmt.Sprintf("%.2f", price*float64(amt))
     				discount := "无优惠"
     				actualPrice := originPrice
     				if req.CouponId != "" {
     					coupon := dao.QueryCouponById(req.CouponId)
     
     					discount = coupon.GetDiscountString()
     					actualPrice = coupon.Discounts(originPrice)
     
     				}
     				dao.AddOrder(model.Order{
     					Id:          utils.GetOrderId(),
     					BuyerId:     userId,
     					SolderId:    solderId,
     					GoodsId:     req.GoodsId,
     					Address:     req.Address,
     					Amount:      req.Amount,
     					Style:       req.Style,
     					Discount:    discount,
     					OriginPrice: originPrice,
     					ActualPrice: actualPrice,
     					Time:        time.Now().Format("2006-01-02 15:04:05"),
     				})
     			}
     			return nil
     		}
     		return errors.New("没找到id为" + req.GoodsId + "的商品")
     	}
     }
     
     ```

     ```go
     package dao
     
     func DownGoodsAmountByOrder(goodsId string, num string, preAmount string) int {
     	sqlStr := "update goods set goods_amount = goods_amount - ? ,goods_sold_amount = goods_sold_amount + ? where goods_id = ? and goods_amount = ?"
     	n, err := Db.Exec(sqlStr, num, num, goodsId, preAmount)
     	if err != nil {
     		fmt.Println("DownGoodsAmountByOrder failed ...")
     	}
     	affected, err := n.RowsAffected()
     	if err != nil {
     		fmt.Println("DownGoodsAmountByOrder RowsAffected err ...")
     	}
     	<-GoodsChan
     	return int(affected)
     }
     ```

   - 秒杀优惠券相似

     ```go
     package service
     
     func FetchCoupon(userId string, couponId string) error {
     	u := dao.QueryUserById(userId)
     	var cp []string
     	err := json.Unmarshal([]byte(u.Coupon), &cp)
     	if err != nil {
     		fmt.Println("FetchCoupon json.Unmarshal failed ...")
     		return err
     	}
     	for i := 0; i < len(cp); i++ {
     		if cp[i] == couponId {
     			return errors.New("您已经有该优惠券了，不可再领取")
     		}
     	}
     	//秒杀抢券，原子性问题思路同下订单的高并发
     	for {
     		coupon := dao.QueryCouponById(couponId)
     		if coupon == (model.Coupon{}) {
     			return errors.New("优惠券不存在")
     		}
     		amt, _ := strconv.Atoi(coupon.Amount)
     		dao.CouponChan <- struct{}{}
     		if amt <= 0 {
     			<-dao.CouponChan
     			return errors.New("你来晚啦,优惠券已经被抢光啦")
     		}
     		if dao.DownCouponAmount(couponId, coupon.Amount) == 0 {
     			continue
     		} else {
     			cp = append(cp, couponId)
     			bytes, err := json.Marshal(cp)
     			if err != nil {
     				fmt.Println("FetchCoupon json.Marshal failed ...")
     				return err
     			}
     			u.Coupon = string(bytes)
     			dao.UpdateUser(u)
     			return nil
     		}
     	}
     }
     ```

     

8. 扫码登陆

   - 实现思路

     1. 用户点击验证码登录，前端发送请求生成一个五分钟有效期的二维码(里面是一个url)
     2. 后端返回前端一个二维码和轮询地址
     3. 前端轮询
     4. 用户扫码（相当于发送一个get请求）
     5. 后端接收到请求，获取参数中的token
     6. 如果token在有效期内，则将二维码状态设为已被扫，并获取用户信息（因为没有app信息给我获取，所以这里只是生成了一个临时用户信息 用户名[扫码用户xxxxxxx] 密码[888888]）
     7. 由于二维码状态发生改变，前端收到的状态码会变成200，同时收到该用户的token
     8. 配合刷新token就可以保持登陆状态了

   - 用到的第三方库

     ```go
         //"github.com/boombuler/barcode"
     	//"github.com/boombuler/barcode/qr"
     ```

   - 代码

     ```go
     package app
     
     var keys = map[string]bool{}
     
     //二维码路由
     func GetQR(c *gin.Context) {
     	var req struct {
     		Phone string `json:"phone" form:"phone" binding:"IsValidPhone"`
     	}
     	if handleBindingError(c, c.ShouldBindQuery(&req), &req) {
     		return
     	}
     	jwt := utils.CreateJWTWithDuration("扫码用户"+utils.GetGoodsId(), 300) //5分钟有效期
     	name := utils.Md5EncodedWithTime(jwt)
     	link := utils.GenerateQR(conf.IP+conf.BackEndPort+"/qr/key/"+jwt, name)
     	jsonData(c, gin.H{
     		"qrLink":      link,                                             //图片链接
     		"inquireLink": conf.IP + conf.BackEndPort + "/qr/status/" + jwt, //轮询地址
     	})
     }
     
     func GetQRStatus(c *gin.Context) {
     	token := c.Param("key")//获取到token
     	if !utils.IsValidJWT(token) {//若五分钟过去token失效
     		jsonError(c, "二维码失效")
     		return
     	}
     	if keys[token] {//判断二维码是否被扫过
     		id := utils.GetGoodsId()//用雪花生成防止名字重复
     		service.Register(model.UserRegisterReq{
     			Username:   "扫码用户" + id,
     			Password:   "888888",//默认密码
     			RePassword: "888888",
     			Email:      "",
     			Phone:      "",
     		})
     		jsonToken(c, "扫码用户"+id)
     		delete(keys, token)//删除此key
     		return
     	}
     	c.JSON(400, gin.H{
     		"status": http.StatusContinue,
     		"msg":    "未识别到用户扫码",
     	})
     }
     
     func JudgeQR(c *gin.Context) {
     	token := c.Param("key")
     	if !utils.IsValidJWT(token) {//若五分钟过去token失效
     		c.String(400, "二维码已过期")
     		return
     	}
     	keys[token] = true//将状态设为ture
     	c.String(200, "扫码成功,请等待网页跳转!")
     }
     ```

     ```go
     package utils
     
     //生成二维码的函数
     func GenerateQR(content string, name string) (link string) {
     	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)
     
     	qrCode, _ = barcode.Scale(qrCode, 200, 200)//200x200像素
     
     	file, _ := os.Create(conf.LocalSavePathOfQR + name + ".jpg")
     	defer file.Close()
     
     	png.Encode(file, qrCode)
     
     	return conf.WebLinkPathOfQR + name + ".jpg"//生成链接
     }
     ```

9. 部署云服务器

   - 接口：http://39.101.72.18:9090
   - 静态资源：http://39.101.72.18:9091

10. 清空数据，一键导入原始数据

    - 清除

      ```scss
      http://39.101.72.18:9090/RESTART/:key
      ```

    - 导入数据

      ```scss
      go run ./src/utils/put/put.go
      ```

      
