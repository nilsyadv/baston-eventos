package repository

import "gorm.io/gorm"

type ConditionalClause func(db *gorm.DB, out interface{}) (*gorm.DB, error)

// PreloadAssociations specified associations to be preloaded
func PreloadAssociations(preloadAssociations []string) ConditionalClause {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if preloadAssociations != nil {
			for _, association := range preloadAssociations {
				db = db.Preload(association)
			}
		}
		return db, nil
	}
}

// Order will order the results
func Order(value interface{}, reorder bool) ConditionalClause {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Order(value)
		return db, nil
	}
}

// Filter will filter the results
func Filter(condition string, args ...interface{}) ConditionalClause {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		return db, nil
	}
}
