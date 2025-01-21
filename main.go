package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Webhook handler function
func webhookHandler(c *fiber.Ctx) error {
	// ตรวจสอบว่าเป็นการร้องขอ Validation หรือไม่
	validationToken := c.Query("validationToken")
	if validationToken != "" {
		// ถ้าเป็นการร้องขอ Validation ส่ง validationToken กลับไป
		return c.SendString(validationToken)
	}

	// ถ้าไม่ใช่ ValidationToken ปกติให้แสดงข้อมูลของ notification
	var payload map[string]interface{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to decode request body")
	}

	// แสดงข้อมูลจาก payload สำหรับการตรวจสอบ
	log.Println("Received notification:", payload)

	// ส่ง Status 200 OK กลับไป
	return c.SendStatus(fiber.StatusOK)
}

func main() {
	// สร้าง app ด้วย Fiber
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ค่าเริ่มต้น (สำหรับการทดสอบใน local)
	}
	// กำหนด route สำหรับ Webhook endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Render!")
	})
	app.Post("/webhook", webhookHandler)

	// ตั้งค่า HTTP server
	log.Println("Server started on :8080")
	log.Fatal(app.Listen(":"+port))
}
