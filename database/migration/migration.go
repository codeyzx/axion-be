package migration

import (
	"axion/database"
	"axion/model/entity"
	"fmt"
	"log"
)

func RunMigration() {

	err := database.DB.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Auction{}, &entity.AuctionHistory{}, &entity.History{})

	database.DB.Exec("CREATE OR REPLACE TRIGGER `trigger_products_insert_history` AFTER INSERT ON `products` FOR EACH ROW INSERT INTO histories VALUES(NULL, CONCAT('Product with id ', NEW.id, ' (', NEW.name, ')', ' has been added'),NOW(), NOW(), NULL);")
	database.DB.Exec("CREATE OR REPLACE TRIGGER `trigger_products_update_history` AFTER UPDATE ON `products` FOR EACH ROW INSERT INTO histories VALUES(NULL, CONCAT('Product with id ', OLD.id, ' (', OLD.name, ')', ' has been updated'),NOW(), NOW(), NULL);")
	database.DB.Exec("CREATE OR REPLACE TRIGGER `trigger_products_delete_history` AFTER DELETE ON `products` FOR EACH ROW INSERT INTO histories VALUES(NULL, CONCAT('Product with id ', OLD.id, ' (', OLD.name, ')', ' has been removed'),NOW(), NOW(), NULL);")
	database.DB.Exec("CREATE OR REPLACE TRIGGER `trigger_auction` BEFORE INSERT ON `auction_histories` FOR EACH ROW UPDATE `auctions` SET last_price = new.price, user_id = new.user_id, bidders_count = bidders_count + 1	WHERE id = new.auction_id AND new.price > last_price AND status = 'open';")
	database.DB.Exec("CREATE OR REPLACE PROCEDURE update_product(IN p_product_id VARCHAR(255), IN p_user_id INT, IN p_name VARCHAR(255), IN p_description TEXT, IN p_price DECIMAL(10,2), IN p_image VARCHAR(255)) BEGIN UPDATE products SET name = p_name, description = p_description, price = p_price, image = p_image WHERE products.id = (SELECT auctions.product_id FROM auctions WHERE auctions.product_id = p_product_id AND auctions.user_id = p_user_id); END;")
	database.DB.Exec("CREATE OR REPLACE PROCEDURE delete_product(IN p_product_id VARCHAR(255), IN p_user_id INT) BEGIN DELETE FROM products WHERE products.id = (SELECT auctions.product_id FROM auctions WHERE auctions.product_id = p_product_id AND auctions.user_id = p_user_id); END;")

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
