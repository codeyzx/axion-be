package migration

import (
	"axion/database"
	"axion/model/entity"
	"fmt"
	"log"
)

func RunMigration() {

	err := database.DB.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Auction{}, &entity.AuctionHistory{}, &entity.History{})

	database.DB.Exec("CREATE OR REPLACE FUNCTION trigger_products_insert_history() RETURNS TRIGGER AS $$ BEGIN INSERT INTO histories VALUES(DEFAULT, CONCAT('Product with id ', NEW.id, ' (', NEW.name, ')', ' has been added'), NOW(), NOW(), NULL); RETURN NEW; END; $$ LANGUAGE plpgsql;")
	database.DB.Exec("CREATE TRIGGER trigger_products_insert_history AFTER INSERT ON products FOR EACH ROW EXECUTE FUNCTION trigger_products_insert_history();")
	database.DB.Exec("CREATE OR REPLACE FUNCTION trigger_products_update_history() RETURNS TRIGGER AS $$ BEGIN INSERT INTO histories VALUES(DEFAULT, CONCAT('Product with id ', OLD.id, ' (', OLD.name, ')', ' has been updated'), NOW(), NOW(), NULL); RETURN NEW; END; $$ LANGUAGE plpgsql;")
	database.DB.Exec("CREATE TRIGGER trigger_products_update_history AFTER UPDATE ON products FOR EACH ROW EXECUTE FUNCTION trigger_products_update_history();")
	database.DB.Exec("CREATE OR REPLACE FUNCTION trigger_products_delete_history() RETURNS TRIGGER AS $$ BEGIN INSERT INTO histories VALUES(DEFAULT, CONCAT('Product with id ', OLD.id, ' (', OLD.name, ')', ' has been removed'), NOW(), NOW(), NULL); RETURN NEW; END; $$ LANGUAGE plpgsql;")
	database.DB.Exec("CREATE TRIGGER trigger_products_delete_history AFTER DELETE ON products FOR EACH ROW EXECUTE FUNCTION trigger_products_delete_history();")
	database.DB.Exec("CREATE OR REPLACE FUNCTION trigger_auction() RETURNS TRIGGER AS $$ BEGIN UPDATE auctions SET last_price = NEW.price, bidder_id = NEW.user_id, bidders_count = bidders_count + 1 WHERE id = NEW.auction_id AND NEW.price > last_price AND status = 'open'; RETURN NEW; END; $$ LANGUAGE plpgsql;")
	database.DB.Exec("CREATE TRIGGER trigger_auction BEFORE INSERT ON auction_histories FOR EACH ROW EXECUTE FUNCTION trigger_auction();")
	database.DB.Exec("CREATE OR REPLACE FUNCTION update_product(p_product_id VARCHAR(255), p_user_id INT, p_name VARCHAR(255), p_description TEXT, p_price DECIMAL(10,2), p_image VARCHAR(255)) RETURNS VOID AS $$ BEGIN UPDATE products SET name = p_name, description = p_description, price = p_price, image = p_image WHERE products.id = (SELECT auctions.product_id FROM auctions WHERE auctions.product_id = p_product_id AND auctions.user_id = p_user_id); END; $$ LANGUAGE plpgsql;")
	database.DB.Exec("CREATE OR REPLACE FUNCTION get_product(user_id INT) RETURNS SETOF products AS $$ BEGIN RETURN QUERY SELECT products.* FROM products INNER JOIN auctions ON products.id = auctions.product_id WHERE auctions.user_id = user_id; END; $$ LANGUAGE plpgsql;")
	
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
