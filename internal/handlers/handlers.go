package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_app_for_code_whisperer/internal/models"
	"go_app_for_code_whisperer/pkg/database"
)

// CreateUser handles the creation of a new user.
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := database.DB.Database("go_app").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, user)
}

// GetUser handles retrieving a user by ID.
func GetUser(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	collection := database.DB.Database("go_app").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser handles updating a user's details.
func UpdateUser(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := database.DB.Database("go_app").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": user,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return c.NoContent(http.StatusOK)
}

// DeleteUser handles deleting a user.
func DeleteUser(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	collection := database.DB.Database("go_app").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.NoContent(http.StatusNoContent)
}

// CreateProduct handles the creation of a new product.
func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := database.DB.Database("go_app").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	product.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, product)
}

// GetProduct handles retrieving a product by ID.
func GetProduct(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	var product models.Product
	collection := database.DB.Database("go_app").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}

// UpdateProduct handles updating a product's details.
func UpdateProduct(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := database.DB.Database("go_app").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": product,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update product")
	}

	return c.NoContent(http.StatusOK)
}

// DeleteProduct handles deleting a product.
func DeleteProduct(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	collection := database.DB.Database("go_app").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete product")
	}

	return c.NoContent(http.StatusNoContent)
}

// CreateOrder handles the creation of a new order.
func CreateOrder(c echo.Context) error {
	var order models.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := database.DB.Database("go_app").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	order.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, order)
}

// GetOrder handles retrieving an order by ID.
func GetOrder(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid order ID")
	}

	var order models.Order
	collection := database.DB.Database("go_app").Collection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Order not found")
	}

	return c.JSON(http.StatusOK, order)
}

// UpdateInventory handles updating inventory.
func UpdateInventory(c echo.Context) error {
	var inventory models.Inventory
	if err := c.Bind(&inventory); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	collection := database.DB.Database("go_app").Collection("inventory")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"product_id": inventory.ProductID}
	update := bson.M{
		"$set": bson.M{"quantity": inventory.Quantity},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update inventory")
	}

	return c.NoContent(http.StatusOK)
}
