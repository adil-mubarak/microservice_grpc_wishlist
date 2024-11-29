package services

import (
	"context"
	"fmt"
	"microservice_grpc_wishlist/models"
	"microservice_grpc_wishlist/pb/wishlist"

	"gorm.io/gorm"
)

type WishlistService struct {
	wishlist.UnimplementedWishlistServiceServer
	DB *gorm.DB
}

func (w *WishlistService) AddToWishlist(ctx context.Context, req *wishlist.AddToWishlistRequest) (*wishlist.AddToWishlistResponse, error) {
	wishlists := models.Wishlist{
		UserID:    int(req.UserID),
		ProductID: int(req.ProductID),
	}

	if err := w.DB.Create(&wishlists).Error; err != nil {
		return nil, fmt.Errorf("failed to create wishlist %v", err)
	}

	return &wishlist.AddToWishlistResponse{
		Id:     uint32(wishlists.ID),
		Status: "item added to the wishlist",
	}, nil
}

func (w *WishlistService) RemoveFromWishlist(ctx context.Context, req *wishlist.RemoveFromWishlistRequest) (*wishlist.RemoveFromeWishlistResponse, error) {
	var wishlists models.Wishlist

	if err := w.DB.Where("user_id = ? AND id = ?", req.UserID, req.Id).First(&wishlists).Delete(&wishlists).Error; err != nil {
		return nil, fmt.Errorf("failed to remove the item %v", err)
	}

	return &wishlist.RemoveFromeWishlistResponse{
		Id:     uint32(wishlists.ID),
		Status: "item removed successfully",
	}, nil
}

func (w *WishlistService) ViewWishlist(ctx context.Context, req *wishlist.ViewWishlistRequset) (*wishlist.ViewWishlistResponse, error) {
	var wishlists []models.Wishlist

	if err := w.DB.Where("user_id = ?", req.UserID).Find(&wishlists).Error; err != nil {
		return nil, fmt.Errorf("did't find the wishlsit")
	}

	var wishlistResponse []*wishlist.Wishlist
	for _, wsh := range wishlists {
		wishlistResponse = append(wishlistResponse, &wishlist.Wishlist{
			Id:        uint32(wsh.ID),
			UserID:    uint32(wsh.UserID),
			ProductID: uint32(wsh.ProductID),
		})
	}
	return &wishlist.ViewWishlistResponse{
		Wishlists: wishlistResponse,
	}, nil

}
