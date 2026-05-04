package tui

import (
	terminal "github.com/terminaldotshop/terminal-sdk-go"
	"github.com/terminaldotshop/terminal/go/pkg/api"
)

// IsDemo reports whether the app is running in demo mode.
func IsDemo() bool {
	return api.IsDemo()
}

// MockViewData returns a realistic ViewInitResponseData for demo mode.
func MockViewData() terminal.ViewInitResponseData {
	products := []terminal.Product{
		{
			ID:          "prd_demo_coffee",
			Name:        "coffee",
			Description: "A rich, dark roast single-origin coffee from the highlands of Ethiopia. Notes of blueberry, jasmine, and dark chocolate.",
			Subscription: terminal.ProductSubscriptionAllowed,
			Tags: terminal.ProductTags{
				Color:    "#6F4E37",
				Featured: true,
				MarketNa: true,
			},
			Variants: []terminal.ProductVariant{
				{
					ID:    "var_demo_coffee_12oz",
					Name:  "12oz",
					Price: 1800,
				},
				{
					ID:    "var_demo_coffee_1lb",
					Name:  "1lb",
					Price: 2800,
				},
			},
		},
		{
			ID:          "prd_demo_tshirt",
			Name:        "tee",
			Description: "The official terminal.shop t-shirt. 100% organic cotton, screen printed with our logo. Made to last as long as your longest running process.",
			Tags: terminal.ProductTags{
				Color:    "#1a1a2e",
				Featured: true,
				MarketNa: true,
			},
			Variants: []terminal.ProductVariant{
				{
					ID:    "var_demo_tshirt_sm",
					Name:  "sm",
					Price: 2500,
				},
				{
					ID:    "var_demo_tshirt_md",
					Name:  "md",
					Price: 2500,
				},
				{
					ID:    "var_demo_tshirt_lg",
					Name:  "lg",
					Price: 2500,
				},
			},
		},
		{
			ID:          "prd_demo_mug",
			Name:        "mug",
			Description: "A ceramic mug with the terminal.shop logo. Perfect for your morning coffee, your afternoon coffee, and your evening coffee.",
			Tags: terminal.ProductTags{
				Color:    "#2d2d2d",
				MarketNa: true,
			},
			Variants: []terminal.ProductVariant{
				{
					ID:    "var_demo_mug_std",
					Name:  "standard",
					Price: 1500,
				},
			},
		},
		{
			ID:          "prd_demo_cron",
			Name:        "cron",
			Description: "Subscribe to a monthly delivery of freshly roasted coffee. We source the best beans from around the world and ship them directly to your door.",
			Subscription: terminal.ProductSubscriptionRequired,
			Tags: terminal.ProductTags{
				Color:    "#FFD700",
				MarketNa: true,
			},
			Variants: []terminal.ProductVariant{
				{
					ID:    "var_demo_cron_monthly",
					Name:  "monthly",
					Price: 2200,
				},
			},
		},
	}

	addresses := []terminal.Address{
		{
			ID:       "adr_demo_1",
			Name:     "Demo User",
			Street1:  "123 Terminal Ave",
			Street2:  "Suite 404",
			City:     "San Francisco",
			Province: "CA",
			Country:  "US",
			Zip:      "94105",
		},
	}

	cards := []terminal.Card{
		{
			ID:    "crd_demo_1",
			Brand: "visa",
			Last4: "4242",
			Expiration: terminal.CardExpiration{
				Month: 12,
				Year:  2027,
			},
		},
	}

	cart := terminal.Cart{
		AddressID: "adr_demo_1",
		CardID:    "crd_demo_1",
		Items:     []terminal.CartItem{},
		Subtotal:  0,
		Amount: terminal.CartAmount{
			Subtotal: 0,
			Shipping: 0,
		},
	}

	subscriptions := []terminal.Subscription{
		{
			ID:               "sub_demo_1",
			ProductVariantID: "var_demo_cron_monthly",
			AddressID:        "adr_demo_1",
			CardID:           "crd_demo_1",
			Quantity:         1,
			Price:            2200,
			Next:             "2026-06-01",
			Created:          "2026-01-15",
			Schedule: terminal.SubscriptionSchedule{
				Type:     terminal.SubscriptionScheduleTypeFixed,
				Interval: 1,
			},
		},
	}

	tokens := []terminal.Token{
		{
			ID:      "tok_demo_abc123",
			Token:   "trm_demo_****",
			Created: "2026-03-10",
		},
	}

	apps := []terminal.App{
		{
			ID:          "app_demo_1",
			Name:        "demo-integration",
			RedirectUri: "https://example.com/callback",
			Secret:      "sec_demo_****",
		},
	}

	orders := []terminal.Order{
		{
			ID:      "ord_demo_1",
			Created: "2026-04-20",
			Amount: terminal.OrderAmount{
				Subtotal: 4300,
				Shipping: 599,
			},
			Items: []terminal.OrderItem{
				{
					ID:               "itm_demo_1",
					Amount:           1800,
					Quantity:         1,
					ProductVariantID: "var_demo_coffee_12oz",
					Description:      "coffee 12oz",
				},
				{
					ID:               "itm_demo_2",
					Amount:           2500,
					Quantity:         1,
					ProductVariantID: "var_demo_tshirt_md",
					Description:      "tee md",
				},
			},
			Shipping: terminal.OrderShipping{
				Name:     "Demo User",
				Street1:  "123 Terminal Ave",
				City:     "San Francisco",
				Province: "CA",
				Country:  "US",
				Zip:      "94105",
			},
			Tracking: terminal.OrderTracking{
				Service: "USPS Priority Mail",
				Number:  "9400111899223850777117",
				Status:  terminal.OrderTrackingStatusDelivered,
			},
		},
	}

	profile := terminal.Profile{
		User: terminal.ProfileUser{
			ID:          "usr_demo_1",
			Name:        "Demo User",
			Email:       "demo@terminal.shop",
			Fingerprint: "aa:bb:cc:dd:ee:ff",
		},
	}

	region := terminal.RegionNa

	return terminal.ViewInitResponseData{
		Products:      products,
		Cart:          cart,
		Cards:         cards,
		Addresses:     addresses,
		Subscriptions: subscriptions,
		Tokens:        tokens,
		Apps:          apps,
		Orders:        orders,
		Profile:       profile,
		Region:        region,
	}
}

// mockCartAfterUpdate returns an updated cart after a SetItem operation in demo mode.
func mockCartAfterUpdate(cart terminal.Cart) terminal.Cart {
	cart.Subtotal = 0
	for _, item := range cart.Items {
		cart.Subtotal += item.Subtotal
	}
	cart.Amount.Subtotal = cart.Subtotal
	return cart
}
