package data

import "gorm.io/gorm"

func ResetDB(db *gorm.DB) {
	_ = db.Migrator().DropTable(&Pricing{}, &Feature{})
	_ = db.AutoMigrate(&Pricing{}, &Feature{})
	SeedData(db)
}

func SeedData(db *gorm.DB) {
	features := []Feature{
		{FeatureName: "45 minutes of daily chat"},
		{FeatureName: "Unlimited messages"},
		{FeatureName: "Ask questions to our human therapists"},
		{FeatureName: "Weekly video sessions"},
		{FeatureName: "Personalized therapy plans"},
		{FeatureName: "24/7 support"},
		{FeatureName: "Monthly progress reports"},
	}

	for i := range features {
		db.Where(Feature{FeatureName: features[i].FeatureName}).FirstOrCreate(&features[i])
	}

	appPlans := []Pricing{
		{
			PlanName:     "Introductory Care",
			Description:  "Explore the basics of our therapy platform",
			PricePerWeek: 0,
			Features:     features[:3],
		},
		{
			PlanName:     "Standard Care",
			Description:  "Get started with essential therapy features",
			PricePerWeek: 10,
			Features:     features[:5],
		},
		{
			PlanName:     "Comprehensive Care",
			Description:  "Experience full therapy support with human involvement",
			PricePerWeek: 20,
			Features:     features,
		},
	}

	for i := range appPlans {
		db.Where(Pricing{PlanName: appPlans[i].PlanName}).FirstOrCreate(&appPlans[i])
	}
}
