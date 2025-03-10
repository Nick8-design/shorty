package handlels

import (
	"math/rand"
	//"os"
	"shorty/db"
	"shorty/models"

	"github.com/gofiber/fiber/v2"
)



func generateShortCode(length int) string {
    const charset = "hijkABCDEF3456789lmnopqrstGHIJKLMNOPQRSTUVWXYZ012345abcdefg6789lmnopqrstuvwxyz"
   

    shortCode := make([]byte, length)
    for i := range shortCode {
        shortCode[i] = charset[rand.Intn(len(charset))]
    }
    return string(shortCode)
}

func ShortenURL(c *fiber.Ctx) error {
   var longShort models.LongShort

   var request struct {
	URL string `json:"url"`
}

    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

	
    db.Db.Where("long_url = ?",request.URL).First(&longShort)




    if longShort.ShortUrl==""{
        shortCode := generateShortCode(6) 
        longShort.LongUrl=request.URL

        // data:=os.Getenv("Durl")
        // if data==""{
        //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        //         "error": "Base URL not found",
        //     })
        // }
    data:="https://shorty-xdb4.onrender.com/"
        
    
        longShort.ShortUrl = data + shortCode
    
    
        db.Db.Create(&longShort)
        // err := db.RedisClient.Set(db.Ctx, longShort.ShortUrl,longShort.LongUrl , 0).Err()
        // if err != nil {
        //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        //         "error": "Failed to cache in Redis",
        //     })
        // }

        return c.JSON(longShort.ShortUrl)


      } else{
        return c.JSON(longShort.ShortUrl)

      }


}



func RedirectUrl(c *fiber.Ctx)error{
    // data:=os.Getenv("Durl")
	// if data==""{
    //     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    //         "error": "Base URL not found",
    //     })
    // }

       data:="https://shorty-xdb4.onrender.com/"

    var shortCode =data+c.Params("short")

    // longURL, err := db.RedisClient.Get(db.Ctx, shortCode).Result()
    // if err == nil {
    //     return c.Redirect(longURL, fiber.StatusMovedPermanently)
    // }

    

    var longShort models.LongShort

  if err:=db.Db.Where("short_url = ?",shortCode).First(&longShort).Error;err!=nil{
    return c.Status(404).JSON(fiber.Map{
        "error":err})
  }

    // 🔹 Store in Redis for future requests
    // _ = db.RedisClient.Set(db.Ctx, shortCode, longShort.LongUrl, 0).Err()

  return c.Redirect(longShort.LongUrl,fiber.StatusMovedPermanently)



}
