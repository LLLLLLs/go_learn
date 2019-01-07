/*
Author : Haoyuan Liu
Time   : 2018/6/27
*/
package dcapi

import "errors"

const ACT = "act"

func Recovery() HandlerFunc {
	return func(c Context) {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
				c.Abort()
			}
		}()
		Next(c)
	}
}

func DoAction() HandlerFunc {
	return func(c Context) {
		act, ok := c.Value(ACT).(Action)
		if !ok {
			c.AbortWithError(errors.New("cannot find action"))
			return
		}
		act.Do(c)

		Next(c)
	}
}
