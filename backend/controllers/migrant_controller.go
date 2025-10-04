package controllers

import (
	"time"
	"visa-tracker/database"
	"visa-tracker/models"

	"github.com/gofiber/fiber/v2"
)

// Структура для входящего запроса
type CreateMigrantRequest struct {
	FullName    string  `json:"full_name"`
	Passport    string  `json:"passport"`
	Nationality string  `json:"nationality"`
	StayType    string  `json:"stay_type"`
	EntryDate   string  `json:"entry_date"` // Принимаем как строку
	VisaExpiry  *string `json:"visa_expiry,omitempty"`
	AllowedDays *int    `json:"allowed_days,omitempty"`
}

func GetAllMigrants(c *fiber.Ctx) error {
	var migrants []models.Migrant
	database.DB.Find(&migrants)
	return c.JSON(migrants)
}

func GetExpiredMigrants(c *fiber.Ctx) error {
	var migrants []models.Migrant
	database.DB.Find(&migrants)

	expired := []models.Migrant{}
	now := time.Now()

	for _, m := range migrants {
		if m.StayType == "visa" && m.VisaExpiry != nil && now.After(*m.VisaExpiry) {
			expired = append(expired, m)
		}
		if m.StayType == "visa-free" && m.AllowedDays != nil {
			deadline := m.EntryDate.AddDate(0, 0, *m.AllowedDays)
			if now.After(deadline) {
				expired = append(expired, m)
			}
		}
	}

	return c.JSON(expired)
}

func CreateMigrant(c *fiber.Ctx) error {
	var req CreateMigrantRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON: " + err.Error()})
	}

	// Парсим даты
	entryDate, err := time.Parse("2006-01-02", req.EntryDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid entry_date format"})
	}

	var visaExpiry *time.Time
	if req.VisaExpiry != nil && *req.VisaExpiry != "" {
		ve, err := time.Parse("2006-01-02", *req.VisaExpiry)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid visa_expiry format"})
		}
		visaExpiry = &ve
	}

	// Валидация
	if req.FullName == "" || req.Passport == "" || req.Nationality == "" ||
		req.StayType == "" || req.EntryDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing required fields"})
	}

	if req.StayType == "visa" && visaExpiry == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "visa_expiry is required for visa type"})
	}

	if req.StayType == "visa-free" && req.AllowedDays == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "allowed_days is required for visa-free type"})
	}

	// Создаем модель для БД
	migrant := models.Migrant{
		FullName:    req.FullName,
		Passport:    req.Passport,
		Nationality: req.Nationality,
		StayType:    req.StayType,
		EntryDate:   entryDate,
		VisaExpiry:  visaExpiry,
		AllowedDays: req.AllowedDays,
	}

	// Сохраняем в БД
	result := database.DB.Create(&migrant)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create migrant: " + result.Error.Error(),
		})
	}

	return c.JSON(migrant)
}

// Остальные функции остаются без изменений
func UpdateMigrant(c *fiber.Ctx) error {
	id := c.Params("id")
	var migrant models.Migrant
	if err := database.DB.First(&migrant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "migrant not found"})
	}

	var req struct {
		FullName    *string `json:"full_name,omitempty"`
		Passport    *string `json:"passport,omitempty"`
		Nationality *string `json:"nationality,omitempty"`
		StayType    *string `json:"stay_type,omitempty"`
		EntryDate   *string `json:"entry_date,omitempty"`
		VisaExpiry  *string `json:"visa_expiry,omitempty"`
		AllowedDays *int    `json:"allowed_days,omitempty"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if req.FullName != nil {
		migrant.FullName = *req.FullName
	}
	if req.Passport != nil {
		migrant.Passport = *req.Passport
	}
	if req.Nationality != nil {
		migrant.Nationality = *req.Nationality
	}
	if req.StayType != nil {
		migrant.StayType = *req.StayType
	}
	if req.EntryDate != nil {
		if date, err := time.Parse("2006-01-02", *req.EntryDate); err == nil {
			migrant.EntryDate = date
		}
	}
	if req.VisaExpiry != nil {
		if date, err := time.Parse("2006-01-02", *req.VisaExpiry); err == nil {
			migrant.VisaExpiry = &date
		}
	}
	if req.AllowedDays != nil {
		migrant.AllowedDays = req.AllowedDays
	}

	database.DB.Save(&migrant)
	return c.JSON(migrant)
}

func DeleteMigrant(c *fiber.Ctx) error {
	id := c.Params("id")
	var migrant models.Migrant
	if err := database.DB.First(&migrant, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "migrant not found"})
	}
	database.DB.Delete(&migrant)
	return c.SendStatus(fiber.StatusNoContent)
}
