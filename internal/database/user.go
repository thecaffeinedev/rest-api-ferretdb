package database

import (
	"context"
	"errors"

	"github.com/thecaffeinedev/rest-api-ferretdb/internal/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type user struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func fromModel(in users.User) user {
	return user{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func toModel(in user) users.User {
	return users.User{
		ID:       in.ID.String(),
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func (d *Database) GetUser(ctx context.Context, email string) (users.User, error) {
	var out user
	err := d.
		Collection.
		FindOne(ctx, bson.M{"email": email}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return users.User{}, ErrUserNotFound
		}
		return users.User{}, err
	}
	return toModel(out), nil
}

func (d *Database) CreateUser(ctx context.Context, user users.User) (users.User, error) {
	out, err := d.
		Collection.
		InsertOne(ctx, fromModel(user))
	if err != nil {
		return users.User{}, err
	}
	user.ID = out.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (d *Database) UpdateUser(ctx context.Context, user users.User) (users.User, error) {
	in := bson.M{}
	if user.Name != "" {
		in["name"] = user.Name
	}
	if user.Password != "" {
		in["password"] = user.Password
	}
	out, err := d.
		Collection.
		UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": in})
	if err != nil {
		return users.User{}, err
	}
	if out.MatchedCount == 0 {
		return users.User{}, ErrUserNotFound
	}
	return user, nil
}

func (d *Database) DeleteUser(ctx context.Context, email string) error {
	out, err := d.
		Collection.
		DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
