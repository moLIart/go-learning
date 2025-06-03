package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/model/room"
)

var (
	mongoClient *mongo.Client
	playersCol  *mongo.Collection
	roomsCol    *mongo.Collection
	boardsCol   *mongo.Collection
	gamesCol    *mongo.Collection
	redisClient *redis.Client
)

func Startup(mongoDataSource, redisDataSource string) {
	var err error
	clientOptions := options.Client().ApplyURI(mongoDataSource)
	mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	playersCol = mongoClient.Database("game_db").Collection("players")
	roomsCol = mongoClient.Database("game_db").Collection("rooms")
	boardsCol = mongoClient.Database("game_db").Collection("boards")
	gamesCol = mongoClient.Database("game_db").Collection("games")

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisDataSource,
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to MongoDB and Redis successfully")
}

func logActionToRedis(action, entityType string, entityID interface{}) {
	ctx := context.Background()
	key := "log:" + action + ":" + entityType + ":" + fmt.Sprint(entityID)
	value := time.Now().Format(time.RFC3339)
	redisClient.Set(ctx, key, value, time.Minute)
}

func AddEntity(entity interface{}) error {
	switch e := entity.(type) {
	case *room.Player:
		res, err := playersCol.InsertOne(context.TODO(), e)
		if err == nil {
			logActionToRedis("create", "player", res.InsertedID)
		}
		return err
	case *room.Room:
		res, err := roomsCol.InsertOne(context.TODO(), e)
		if err == nil {
			logActionToRedis("create", "room", res.InsertedID)
		}
		return err
	case *game.Board:
		res, err := boardsCol.InsertOne(context.TODO(), e)
		if err == nil {
			logActionToRedis("create", "board", res.InsertedID)
		}
		return err
	case *game.Game:
		res, err := gamesCol.InsertOne(context.TODO(), e)
		if err == nil {
			logActionToRedis("create", "game", res.InsertedID)
		}
		return err
	default:
		return nil
	}
}

func GetPlayersCount() (int64, error) {
	return playersCol.CountDocuments(context.TODO(), bson.D{{}})
}

func GetRoomsCount() (int64, error) {
	return roomsCol.CountDocuments(context.TODO(), bson.D{{}})
}

func GetBoardsCount() (int64, error) {
	return boardsCol.CountDocuments(context.TODO(), bson.D{{}})
}

func GetGamesCount() (int64, error) {
	return gamesCol.CountDocuments(context.TODO(), bson.D{{}})
}

func GetPlayers() ([]*room.Player, error) {
	var players []*room.Player
	cursor, err := playersCol.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var player room.Player
		if err := cursor.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, &player)
	}
	return players, nil
}

func GetRooms() ([]*room.Room, error) {
	var rooms []*room.Room
	cursor, err := roomsCol.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var room room.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func GetBoards() ([]*game.Board, error) {
	var boards []*game.Board
	cursor, err := boardsCol.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var board game.Board
		if err := cursor.Decode(&board); err != nil {
			return nil, err
		}
		boards = append(boards, &board)
	}
	return boards, nil
}

func GetGames() ([]*game.Game, error) {
	var games []*game.Game
	cursor, err := gamesCol.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var game game.Game
		if err := cursor.Decode(&game); err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	return games, nil
}

func GetPlayerByID(id int) (*room.Player, error) {
	var player room.Player
	err := playersCol.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&player)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func GetRoomByID(id int) (*room.Room, error) {
	var room room.Room
	err := roomsCol.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func GetBoardByID(id int) (*game.Board, error) {
	var board game.Board
	err := boardsCol.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&board)
	if err != nil {
		return nil, err
	}
	return &board, nil
}

func GetGameByID(id int) (*game.Game, error) {
	var game game.Game
	err := gamesCol.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func DeletePlayerByID(id int) (bool, error) {
	result, err := playersCol.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err == nil && result.DeletedCount > 0 {
		logActionToRedis("delete", "player", id)
	}
	return result.DeletedCount > 0, err
}

func DeleteRoomByID(id int) (bool, error) {
	result, err := roomsCol.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err == nil && result.DeletedCount > 0 {
		logActionToRedis("delete", "room", id)
	}
	return result.DeletedCount > 0, err
}

func DeleteBoardByID(id int) (bool, error) {
	result, err := boardsCol.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err == nil && result.DeletedCount > 0 {
		logActionToRedis("delete", "board", id)
	}
	return result.DeletedCount > 0, err
}

func DeleteGameByID(id int) (bool, error) {
	result, err := gamesCol.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err == nil && result.DeletedCount > 0 {
		logActionToRedis("delete", "game", id)
	}
	return result.DeletedCount > 0, err
}

func UpdatePlayerByID(id int, name string) (bool, error) {
	result, err := playersCol.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"name": name}})
	if err == nil && result.ModifiedCount > 0 {
		logActionToRedis("update", "player", id)
	}
	return result.ModifiedCount > 0, err
}

func UpdateRoomByID(id int, code string) (bool, error) {
	result, err := roomsCol.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"code": code}})
	if err == nil && result.ModifiedCount > 0 {
		logActionToRedis("update", "room", id)
	}
	return result.ModifiedCount > 0, err
}

func UpdateBoardByID(id int, size int) (bool, error) {
	result, err := boardsCol.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"size": size}})
	if err == nil && result.ModifiedCount > 0 {
		logActionToRedis("update", "board", id)
	}
	return result.ModifiedCount > 0, err
}
