package search

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"server/pkg/api"
	"testing"
)

//реализация интерефейса store.UserRepo from storerepo.go
type MockTestDB struct{}

func (mDB *MockTestDB) FindByWriter(writer string) ([]string, error) {

	switch writer {
	case "testAWriter":
		return []string{"testABook1", "testABook2"}, nil
	case "testBWriter":
		return []string{"testBBook1", "testBBook2"}, nil
	case "testCWriter":
		return []string{"testCBook1", "testCBook2", "testCBook3"}, nil
	case "testEmptyWriter":
		return nil, nil
	default:
		return nil, nil
	}
}

func (mDB *MockTestDB) FindByBook(book string) (string, error) {
	switch book {
	case "testABook":
		return "testAWriter", nil
	case "testBBook":
		return "testBWriter", nil
	case "testCBook":
		return "testCWriter", nil
	case "testEmptybook":
		return "", nil
	default:
		return "", nil
	}
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	//init test struct for mock...
	rep := &MockTestDB{}

	//init new logrus with panic level, for ignore logger...
	lg := logrus.New()
	lg.SetLevel(logrus.PanicLevel)

	api.RegisterSearchingServer(server, &GRPCServer{Logger: lg, Store: rep})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestGRPCServer_Search(t *testing.T) {
	testsBook := []struct {
		name   string
		book   string
		writer string
	}{
		{
			name:   "test1book",
			book:   "testABook",
			writer: "testAWriter",
		}, {
			name:   "test2book",
			book:   "testBBook",
			writer: "testBWriter",
		}, {
			name:   "test3book",
			book:   "testCBook",
			writer: "testCWriter",
		}, {
			name:   "test4book",
			book:   "testEmptybook",
			writer: "",
		},
	}

	testsWriter := []struct {
		name   string
		book   []string
		writer string
	}{
		{
			name:   "test1book",
			book:   []string{"testABook1", "testABook2"},
			writer: "testAWriter",
		}, {
			name:   "test2book",
			book:   []string{"testBBook1", "testBBook2"},
			writer: "testBWriter",
		}, {
			name:   "test3book",
			book:   []string{"testCBook1", "testCBook2", "testCBook3"},
			writer: "testCWriter",
		}, {
			name:   "test4book",
			book:   nil,
			writer: "testEmptyWriter",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := api.NewSearchingClient(conn)

	for _, tv := range testsBook {
		t.Run(tv.name, func(t *testing.T) {
			request := &api.SearchRequest{Book: tv.book}
			response, err := client.Search(ctx, request)
			if response.GetWriter() != tv.writer {
				t.Errorf("Внимание! Ошибка! Параметр из ответа %s не соответствует требуемому %s!\n", response.GetWriter(), tv.writer)
			}
			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, tv := range testsWriter {
		t.Run(tv.name, func(t *testing.T) {
			request := &api.SearchRequest{Writer: tv.writer}
			response, err := client.Search(ctx, request)
			for i, v := range response.GetBook() {
				if v != tv.book[i] {
					t.Errorf("Внимание! Ошибка! Параметр из ответа %s не соответствует требуемому %s!\n", response.Book, tv.book)
				}
			}
			if err != nil {
				t.Error(err)
			}
		})
	}
}
