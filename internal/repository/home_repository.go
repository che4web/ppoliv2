package repository

import (
	"fmt"

	"gorm.io/gorm"

	"ppoliv2/internal/models"
)

type HomePageData struct {
	HeroCards []models.HeroCard
	Features  []models.Feature
	Steps     []models.Step
	Products  []models.Product
	Reviews   []models.Review
}

type HomeRepository struct {
	db *gorm.DB
}

func NewHomeRepository(db *gorm.DB) *HomeRepository {
	return &HomeRepository{db: db}
}

func (r *HomeRepository) MigrateAndSeed() error {
	if err := r.db.AutoMigrate(&models.HeroCard{}, &models.Feature{}, &models.Step{}, &models.Product{}, &models.Review{}); err != nil {
		return fmt.Errorf("automigrate: %w", err)
	}

	if err := r.seedHeroCards(); err != nil {
		return err
	}
	if err := r.seedFeatures(); err != nil {
		return err
	}
	if err := r.seedSteps(); err != nil {
		return err
	}
	if err := r.seedProducts(); err != nil {
		return err
	}
	if err := r.seedReviews(); err != nil {
		return err
	}

	return nil
}

func (r *HomeRepository) HomeData() (HomePageData, error) {
	var data HomePageData

	if err := r.db.Order("id asc").Find(&data.HeroCards).Error; err != nil {
		return data, err
	}
	if err := r.db.Order("id asc").Find(&data.Features).Error; err != nil {
		return data, err
	}
	if err := r.db.Order("number asc").Find(&data.Steps).Error; err != nil {
		return data, err
	}
	if err := r.db.Order("id asc").Find(&data.Products).Error; err != nil {
		return data, err
	}
	if err := r.db.Order("id asc").Find(&data.Reviews).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *HomeRepository) seedHeroCards() error {
	items := []models.HeroCard{
		{Title: "Капельный полив", Description: "Экономичный и точечный полив", LinkText: "Перейти"},
		{Title: "Дождеватели", Description: "Равномерный полив газона", LinkText: "Перейти"},
		{Title: "Контроллеры", Description: "Управление поливом по расписанию", LinkText: "Перейти"},
		{Title: "Трубы и фитинги", Description: "Надёжные комплектующие", LinkText: "Перейти"},
	}
	for _, item := range items {
		if err := r.db.FirstOrCreate(&models.HeroCard{}, models.HeroCard{Title: item.Title}).Error; err != nil {
			return err
		}
		if err := r.db.Model(&models.HeroCard{}).Where("title = ?", item.Title).Updates(map[string]any{"description": item.Description, "link_text": item.LinkText}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *HomeRepository) seedFeatures() error {
	items := []models.Feature{{Icon: "💧", Title: "Экономия воды до 50%"}, {Icon: "📶", Title: "Автоматизация и Wi‑Fi управление"}, {Icon: "🔧", Title: "Простая установка своими руками"}, {Icon: "🌿", Title: "Подходит для любого участка"}}
	for _, item := range items {
		if err := r.db.FirstOrCreate(&models.Feature{}, models.Feature{Title: item.Title}).Error; err != nil {
			return err
		}
		if err := r.db.Model(&models.Feature{}).Where("title = ?", item.Title).Update("icon", item.Icon).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *HomeRepository) seedSteps() error {
	items := []models.Step{{Number: 1, Title: "Вы выбираете систему полива"}, {Number: 2, Title: "Мы рассчитываем комплектацию"}, {Number: 3, Title: "Вы устанавливаете или заказываете монтаж"}, {Number: 4, Title: "Система поливает автоматически"}}
	for _, item := range items {
		if err := r.db.FirstOrCreate(&models.Step{}, models.Step{Number: item.Number}).Error; err != nil {
			return err
		}
		if err := r.db.Model(&models.Step{}).Where("number = ?", item.Number).Update("title", item.Title).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *HomeRepository) seedProducts() error {
	items := []models.Product{
		{Name: "Комплект капельного полива AquaGreen AG‑100", Price: "4 990 ₽", Meta: "до 100 м²"},
		{Name: "Дождеватель роторный Hunter SRM‑04", Price: "1 890 ₽", Meta: "до 350 м²"},
		{Name: "Контроллер полива Rain Bird ESP‑RZXe", Price: "12 990 ₽", Meta: "Wi‑Fi, 4 зоны"},
		{Name: "Труба ПНД для полива Ø 32 мм", Price: "2 750 ₽", Meta: "Бухта 25 м"},
	}
	for _, item := range items {
		if err := r.db.FirstOrCreate(&models.Product{}, models.Product{Name: item.Name}).Error; err != nil {
			return err
		}
		if err := r.db.Model(&models.Product{}).Where("name = ?", item.Name).Updates(map[string]any{"price": item.Price, "meta": item.Meta}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *HomeRepository) seedReviews() error {
	items := []models.Review{{Author: "Александр П.", Text: "Подобрали систему для теплицы. Всё работает идеально!", Date: "15 мая 2024"}, {Author: "Мария С.", Text: "Установили автополив газона. Очень удобно и экономит время.", Date: "28 апреля 2024"}}
	for _, item := range items {
		if err := r.db.FirstOrCreate(&models.Review{}, models.Review{Author: item.Author}).Error; err != nil {
			return err
		}
		if err := r.db.Model(&models.Review{}).Where("author = ?", item.Author).Updates(map[string]any{"text": item.Text, "date": item.Date}).Error; err != nil {
			return err
		}
	}
	return nil
}
