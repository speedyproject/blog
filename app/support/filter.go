package support

import "github.com/revel/revel"

//LoginFilter check login status
func LoginFilter(c *revel.Controller) revel.Result {

	uid := c.Session["UID"]

	revel.INFO.Printf("Login check UID: %s", uid)

	if uid == "" {
		// return c.Redirect(routes.Login.SignIn())
	}

	res, _ := Cache.Get(SPY_ADMIN_INFO + uid).Result()

	revel.INFO.Printf("Login check cache data: %v", res)

	if res == "" {
		// return c.Redirect(routes.Login.SignIn())
	}

	return nil
}
