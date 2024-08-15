package middleware

// func RequireAuth(c *gin.Context) {
// 	fmt.Println("Authenticating...")
// 	tokenString, _ := c.Cookie("Authorization")
// 	fmt.Println(tokenString)

// 	tokenString = strings.ToLower(tokenString)

// 	// if tokenString != "admin" {
// 	// 	// c.JSON(401, gin.H{
// 	// 	// 	"message": "Unauthorized",
// 	// 	// })
// 	// 	c.AbortWithStatus(http.StatusUnauthorized)
// 	// 	return
// 	// }

// 	c.Next()

// }
