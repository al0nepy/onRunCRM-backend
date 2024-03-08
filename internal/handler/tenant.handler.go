package handler

import (
	"context"

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
	collTenant := mongodb.DB.Collection("tenant")
	tenantModel := model.TenantModel{
		Name: req.Name,
	}
	_, err = collTenant.InsertOne(ctx, tenantModel)

	return utils.HandleOperationResult(err)
}

func (tenantHandler *TenantHandlerServer) UpdateTenants(ctx context.Context, req *pb.UpdateTenantsRequest) (status *pb.Status, err error) {
	collTenant := mongodb.DB.Collection("tenant")
	filter := bson.D{{Key: "_id", Value: req.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: req.Name}}}}
	_, err = collTenant.UpdateOne(ctx, filter, update)

	return utils.HandleOperationResult(err)
}

func (tenantHandler *TenantHandlerServer) DeleteTenants(ctx context.Context, req *pb.IdTenantsRequest) (status *pb.Status, err error) {
	collTenant := mongodb.DB.Collection("tenant")
	filter := bson.D{{Key: "_id", Value: req.Id}}
	_, err = collTenant.DeleteOne(ctx, filter)

	return utils.HandleOperationResult(err)
}
