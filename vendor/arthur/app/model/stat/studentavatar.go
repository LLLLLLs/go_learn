/*
Created on 2018-11-05 09:35:07
author: Auto Generate
*/
package stat

type StudentAvatar struct {
	Id        int16 `model:"pk"` //形象ID
	Talent    int16              //天赋ID
	IsMale    bool               //是否为男性
	Graduated bool               //是否已毕业
}
