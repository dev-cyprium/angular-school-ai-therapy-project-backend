package data

import (
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	appPlans := []Pricing{
		{
			PlanName:     "Free",
			Description:  " Make sure the platform works for you ",
			PricePerWeek: 0,
			Features: []Feature{
				{
					FeatureName: "45 minutes of daily chat",
				},
				{
					FeatureName: "Unlimited messages",
				},
				{
					FeatureName: "Ask questions to our human therapists",
				},
			},
		},
	}

	for i := range appPlans {
		db.Where(Pricing{PlanName: appPlans[i].PlanName}).FirstOrCreate(&appPlans[i])
	}
}
