package seeds

import (
	"codebase-app/internal/adapter"
	"context"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jmoiron/sqlx"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

// Seed struct.
type Seed struct {
	db *sqlx.DB
}

// NewSeed return a Seed with a pool of connection to a dabase.
func newSeed(db *sqlx.DB) Seed {
	return Seed{
		db: db,
	}
}

func Execute(db *sqlx.DB, table string, total int) {
	seed := newSeed(db)
	seed.run(table, total)
}

// Run seeds.
func (s *Seed) run(table string, total int) {

	switch table {
	case "roles":
		s.rolesSeed()
	case "users":
		s.usersSeed(total)
	case "category":
		s.categorySeed()
	case "shop":
		s.shopsSeed(total)
	case "product":
		s.productsSeed(total)
	case "all":
		s.rolesSeed()
		s.usersSeed(total)
		s.categorySeed()
	case "delete-all":
		s.deleteAll()
	default:
		log.Warn().Msg("No seed to run")
	}

	if table != "" {
		log.Info().Msg("Seed ran successfully")
		log.Info().Msg("Exiting ...")
		if err := adapter.Adapters.Unsync(); err != nil {
			log.Fatal().Err(err).Msg("Error while closing database connection")
		}
		os.Exit(0)
	}
}

func (s *Seed) deleteAll() {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		} else {
			err = tx.Commit()
			if err != nil {
				log.Error().Err(err).Msg("Error committing transaction")
			}
		}
	}()

	_, err = tx.Exec(`DELETE FROM users`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting users")
		return
	}
	log.Info().Msg("users table deleted successfully")

	_, err = tx.Exec(`DELETE FROM roles`)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting roles")
		return
	}
	log.Info().Msg("roles table deleted successfully")

	log.Info().Msg("=== All tables deleted successfully ===")
}

// rolesSeed seeds the roles table.
func (s *Seed) rolesSeed() {
	roleMaps := []map[string]any{
		{"name": "admin"},
		{"name": "end_user"},
	}

	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	_, err = tx.NamedExec(`
		INSERT INTO roles (name)
		VALUES (:name)
	`, roleMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating roles")
		return
	}

	log.Info().Msg("roles table seeded successfully")
}

// users
func (s *Seed) usersSeed(total int) {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	type generalData struct {
		Id   string `db:"id"`
		Name string `db:"name"`
	}

	var (
		roles    = make([]generalData, 0)
		userMaps = make([]map[string]any, 0)
	)

	err = s.db.Select(&roles, `SELECT id, name FROM roles`)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting roles")
		return
	}

	for i := 0; i < total; i++ {
		selectedRole := roles[gofakeit.Number(0, len(roles)-1)]

		dataUserToInsert := make(map[string]any)
		dataUserToInsert["id"] = ulid.Make().String()
		dataUserToInsert["role_id"] = selectedRole.Id
		dataUserToInsert["name"] = gofakeit.Name()
		dataUserToInsert["email"] = gofakeit.Email()
		dataUserToInsert["whatsapp_number"] = gofakeit.Phone()
		dataUserToInsert["password"] = "$2y$10$mVf4BKsfPSh/pjgHjvk.JOlGdkIYgBGyhaU9WQNMWpYskK9MZlb0G" // password

		userMaps = append(userMaps, dataUserToInsert)
	}

	var (
		endUserId   string
		adminUserId string
	)

	// iterate over roles to get service advisor id
	for _, role := range roles {
		if role.Name == "admin" {
			adminUserId = role.Id
			continue
		}
		if role.Name == "end_user" {
			endUserId = role.Id
			continue
		}
	}

	EndUser := map[string]any{
		"id":              ulid.Make().String(),
		"role_id":         endUserId,
		"name":            "Irham",
		"email":           "irham@fake.com",
		"whatsapp_number": gofakeit.Phone(),
		"password":        "$2y$10$mVf4BKsfPSh/pjgHjvk.JOlGdkIYgBGyhaU9WQNMWpYskK9MZlb0G", // password
	}

	AdminUser := map[string]any{
		"id":              ulid.Make().String(),
		"role_id":         adminUserId,
		"name":            "Fathan",
		"email":           "fathan@fake.com",
		"whatsapp_number": gofakeit.Phone(),
		"password":        "$2y$10$mVf4BKsfPSh/pjgHjvk.JOlGdkIYgBGyhaU9WQNMWpYskK9MZlb0G", // password
	}

	userMaps = append(userMaps, EndUser)
	userMaps = append(userMaps, AdminUser)

	_, err = tx.NamedExec(`
		INSERT INTO users (id, role_id, name, email, whatsapp_number, password)
		VALUES (:id, :role_id, :name, :email, :whatsapp_number, :password)
	`, userMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating users")
		return
	}

	log.Info().Msg("users table seeded successfully")
}

// categorySeed seeds the category table.
func (s *Seed) categorySeed() {
	categoryMaps := []map[string]any{
		{"name": "Elektronik"},
		{"name": "Rumah Tangga"},
		{"name": "Kecantikan & Perawatan"},
		{"name": "Kesehatan"},
		{"name": "Otomotif"},
		{"name": "Olahraga"},
		{"name": "Makanan & Minuman"},
		{"name": "Hobi"},
	}

	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	_, err = tx.NamedExec(`
		INSERT INTO category (name)
		VALUES (:name)
	`, categoryMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating category")
		return
	}

	log.Info().Msg("category table seeded successfully")
}

func (s *Seed) shopsSeed(total int) {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	var (
		shopMaps = make([]map[string]any, 0)
		// users   = make([]string, 0)
	)

	// err = s.db.Select(&users, `SELECT id FROM users`)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Error selecting users")
	// 	return
	// }
	user_id := "e7544cc7-f561-4fcc-8226-f1ce0b3edf0f"

	for i := 0; i < total; i++ {
		dataShopToInsert := make(map[string]any)
		dataShopToInsert["name"] = gofakeit.Company()
		dataShopToInsert["description"] = gofakeit.Sentence(10)
		dataShopToInsert["terms"] = gofakeit.Sentence(10)
		// dataShopToInsert["user_id"] = users[gofakeit.Number(0, len(users)-1)]
		dataShopToInsert["user_id"] = user_id

		shopMaps = append(shopMaps, dataShopToInsert)
	}

	_, err = tx.NamedExec(`
		INSERT INTO shops (name, description, terms, user_id)
		VALUES (:name, :description, :terms, :user_id)
	`, shopMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating shops")
		return
	}

	log.Info().Msg("shops table seeded successfully")
}

func (s *Seed) productsSeed(total int) {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
			log.Error().Err(err).Msg("Error rolling back transaction")
			return
		}

		err = tx.Commit()
		if err != nil {
			log.Error().Err(err).Msg("Error committing transaction")
		}
	}()

	type generalData struct {
		Id       string `db:"id"`
	}

	var (
		categories = make([]generalData, 0)
		productMaps = make([]map[string]any, 0)
		shops = make([]generalData, 0)
	)

	err = s.db.Select(&categories, `SELECT id FROM category`)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting categories")
		return
	}

	err = s.db.Select(&shops, `SELECT id FROM shops`)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting shops")
		return
	}

	for i := 0; i < total; i++ {
		selectedCategory := categories[gofakeit.Number(0, len(categories)-1)]
		selectedShop := shops[gofakeit.Number(0, len(shops)-1)]

		dataProductToInsert := make(map[string]any)
		dataProductToInsert["category_id"] = selectedCategory.Id
		dataProductToInsert["name"] = gofakeit.ProductName()
		dataProductToInsert["brand"] = gofakeit.ProductName()
		dataProductToInsert["price"] = gofakeit.Price(10000, 500000)
		dataProductToInsert["stock"] = gofakeit.Number(10, 500)
		dataProductToInsert["shop_id"] = selectedShop.Id

		productMaps = append(productMaps, dataProductToInsert)
	}

	_, err = tx.NamedExec(`
		INSERT INTO product (category_id, name, brand, price, stock, shop_id)
		VALUES (:category_id, :name, :brand, :price, :stock, :shop_id)	
	`, productMaps)
	if err != nil {
		log.Error().Err(err).Msg("Error creating products")
		return
	}

	log.Info().Msg("products table seeded successfully")
}