package handler

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.onRunCRM/internal/database/mongodb"
	"go.onRunCRM/internal/model"
	pb "go.onRunCRM/pkg/api"
	"go.onRunCRM/pkg/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StaffHandlerServer struct {
	pb.StaffServiceServer
}

func (staffHandler *StaffHandlerServer) GetAllStaff(req *pb.IdTenantRequest, stream pb.StaffService_GetAllStaffServer) error {
	data := &model.StaffModel{}
	collStaff := mongodb.DB.Collection("staff")
	filter := bson.M{
		"tenantId": req.TenantId,
	}

	cursor, err := collStaff.Find(context.TODO(), filter)

	if err != nil {
		log.Panic(err.Error())
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)

		if err != nil {
			return err
		}

		stream.Send(&pb.Staff{
			Id:         data.ID.Hex(),
			Name:       data.Name,
			Email:      data.Email,
			Position:   data.Position,
			Department: data.Department,
			TenantId:   data.TenantID.Hex(),
			CreatedAt:  &timestamppb.Timestamp{},
			UpdatedAt:  &timestamppb.Timestamp{}},
		)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

func (staffHandler *StaffHandlerServer) GetOneStaff(ctx context.Context, req *pb.IdsStaffRequest) (staff *pb.Staff, err error) {
	collStaff := mongodb.DB.Collection("staff")
	filter := bson.M{"_id": req.Id, "tenantId": req.TenantId}

	err = collStaff.FindOne(ctx, filter).Decode(&staff)

	if err != nil {
		return nil, err
	}

	return staff, nil
}

func (staffHandler *StaffHandlerServer) CreateStaff(ctx context.Context, req *pb.CUStaffRequest) (status *pb.Status, err error) {
	collStaff := mongodb.DB.Collection("staff")
	tenantId, _ := primitive.ObjectIDFromHex(req.TenantId)
	staffModel := model.StaffModel{
		Name:       req.Name,
		Email:      req.Email,
		Position:   req.Position,
		Department: req.Department,
		TenantID:   tenantId,
	}

	_, err = collStaff.InsertOne(ctx, staffModel)

	return utils.HandleOperationResult(err)
}

func (staffHandler *StaffHandlerServer) UpdateStaff(ctx context.Context, req *pb.CUStaffRequest) (status *pb.Status, err error) {
	collStaff := mongodb.DB.Collection("staff")
	filter := bson.D{{Key: "_id", Value: req.Id}}
	update := bson.D{{Key: "$set", Value: req}}
	_, err = collStaff.UpdateOne(ctx, filter, update)

	return utils.HandleOperationResult(err)
}

func (staffHandler *StaffHandlerServer) DeleteStaff(ctx context.Context, req *pb.IdsStaffRequest) (status *pb.Status, err error) {
	collStaff := mongodb.DB.Collection("staff")
	filter := bson.D{{Key: "_id", Value: req.Id}}
	_, err = collStaff.DeleteOne(ctx, filter)

	return utils.HandleOperationResult(err)
}
