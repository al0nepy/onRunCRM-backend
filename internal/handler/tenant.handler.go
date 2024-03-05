package handler

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.onRunCRM/internal/database/mongodb"
	"go.onRunCRM/internal/model"
	pb "go.onRunCRM/pkg/api"
	"go.onRunCRM/pkg/utils"
)

type TenantHandlerServer struct {
	pb.TenantsServiceServer
}

func (tenantHandler *TenantHandlerServer) CreateTenants(ctx context.Context, req *pb.CreateTenantsRequest) (status *pb.Status, err error) {
	coll := mongodb.DB.Collection("tenant")
	tenantModel := model.TenantModel{
		Name: req.Name,
	}

	if _, err := coll.InsertOne(ctx, tenantModel); err != nil {
		log.Printf(utils.ERROR_MESSAGE, err.Error())
		return &pb.Status{
			Code:    500,
			Message: err.Error(),
		}, err
	}

	return &pb.Status{
		Code:    200,
		Message: utils.SUCCESS_MESSAGE,
	}, nil
}

func (tenantHandler *TenantHandlerServer) UpdateTenants(ctx context.Context, req *pb.UpdateTenantsRequest) (status *pb.Status, err error) {
	coll := mongodb.DB.Collection("tenant")
	filter := bson.D{{Key: "_id", Value: req.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: req.Name}}}}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		log.Printf(utils.ERROR_MESSAGE, err.Error())
		return &pb.Status{
			Code:    500,
			Message: err.Error(),
		}, err
	}

	return &pb.Status{
		Code:    200,
		Message: utils.SUCCESS_MESSAGE,
	}, nil
}

func (tenantHandler *TenantHandlerServer) DeleteTenants(ctx context.Context, req *pb.IdTenantsRequest) (status *pb.Status, err error) {
	coll := mongodb.DB.Collection("tenant")
	filter := bson.D{{Key: "_id", Value: req.Id}}

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		log.Printf(utils.ERROR_MESSAGE, err.Error())
		return &pb.Status{
			Code:    500,
			Message: err.Error(),
		}, err
	}

	return &pb.Status{
		Code:    200,
		Message: utils.SUCCESS_MESSAGE,
	}, nil
}
