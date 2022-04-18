package service

import (
	"context"
	"github.com/Ruadgedy/pcbook/pb"
	"github.com/Ruadgedy/pcbook/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "asdasdad"

	laptopDuplicateId := sample.NewLaptop()
	storeDuplicateID := NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateId)
	assert.NoError(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInvalidId,
			store:  NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateId,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{Laptop: tc.laptop}

			server := NewLaptopServer(tc.store)
			resp, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.NotEmpty(t, resp.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, resp.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, resp)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, st.Code(), tc.code)
			}
		})
	}
}
