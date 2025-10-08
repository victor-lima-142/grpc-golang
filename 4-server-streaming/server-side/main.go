package main

import (
	"bufio"
	"department/src/pb/department"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	department.DepartmentServiceServer
}

func (s *server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersonServer) error {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ";")

		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.Atoi(data[3])
		departmentId, _ := strconv.Atoi(data[4])

		if int32(departmentId) == req.GetDepartmentId() {
			if err = srv.Send(&department.ListPersonResponse{
				Id:     int32(id),
				Name:   name,
				Email:  email,
				Income: int32(income),
			}); err != nil {
				return fmt.Errorf("Error sending request: %v", err)
			}
		}
	}

	return nil
}

func main() {
	fmt.Println("Starting grpc server")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}
	s := grpc.NewServer()
	department.RegisterDepartmentServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
