package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	sdm "github.com/strongdm/strongdm-sdk-go/v3"
)

//
// /*
// * need a fake client
// * thinking about it now; I think that whatever sdm.New returns should be mimicked
// * so instead of doing sdm.Client, what I have done is assigned the Client type to a variable
// * it should be sufficient to be passed to the test function
// ------ Executing now
// I believe when this will be executed we will get a nil pointer exception : and that's what happened when I ran this function.
// this happened because fakeClient.Resources in nil
// */
// // TestCreateResource : First Try
// func TestCreateResource_firstTry(t *testing.T) {
// 	// need a fake client
// 	// thinking about it now; I think that whatever sdm.New returns should be mimicked
// 	// so instead of doing sdm.Client, what I have done is assigned the Client type to a variable
// 	// it should be sufficient to be passed to the test function
// 	fakeClient := sdm.Client{}
// 	ctx := context.Background()
// 	r := sdm.ElasticacheRedis{Name: "test-redis"}
// 	fmt.Println("nil pointer exception will occur", fakeClient.Resources())
// 	err := CreateResourceV1(ctx, &fakeClient, &r)
// 	assert.NoError(t, err)
// }
//
// // type Resourcer interface {
// // 	Create(ctx context.Context, resource sdm.Resource) (*sdm.ResourceCreateResponse, error)
// // }

// your own type called FakeClient or mockClient - whatever you want to call it
type FakeClient struct{}

// this type implements the interface that you had defined where the implemenration was written
func (f *FakeClient) Create(ctx context.Context, resource sdm.Resource) (*sdm.ResourceCreateResponse, error) {
	fmt.Println("Create for fake client")
	return &sdm.ResourceCreateResponse{Resource: &sdm.ElasticacheRedis{Name: "manand-test"}}, nil
}

func (f *FakeClient) Delete(ctx context.Context, resourceID string) (*sdm.ResourceDeleteResponse, error) {
	fmt.Println("Delete for fake client")
	return &sdm.ResourceDeleteResponse{}, nil
}

func (f *FakeClient) Get(ctx context.Context, resourceID string) (*sdm.ResourceGetResponse, error) {
	fmt.Println("Delete for fake client")
	return &sdm.ResourceGetResponse{Resource: &sdm.ElasticacheRedis{Name: "manand-test", Tags: sdm.Tags{"Key": "Value"}}}, nil
}

func (f *FakeClient) Update(ctx context.Context, resource sdm.Resource) (*sdm.ResourceUpdateResponse, error) {
	fmt.Println("Update for fake client")
	return &sdm.ResourceUpdateResponse{Resource: &sdm.ElasticacheRedis{Name: "manand-test"}}, nil
}

func TestMyResource_GetResourceV2(t *testing.T) {
	fc := FakeClient{}
	mr := MyResource{Client: &fc}
	ctx := context.Background()
	err := mr.GetResourceV2(ctx, "someID")
	assert.NoError(t, err)

}

// TestCreateResource : Second Try
func TestCreateResourceV2(t *testing.T) {
	// fakeClient should provide me access to create function; all I need is access to Create function from the fake client
	// so can I have my own implementation of the Create function?
	// i.e. the object I create in test function will make a call to the create function I will write and when sdm.Client will make a call to the Create function sdm has in their package
	// what should be the interface?
	fc := FakeClient{}
	// set your client as a the fakeClient
	mr := MyResource{Client: &fc}
	ctx := context.Background()
	r := sdm.ElasticacheRedis{Name: "manand-test"}
	err := mr.CreateResourceV2(ctx, &r)
	assert.NoError(t, err)
	// assert.Equal(t)
}

func TestMyResource_DeleteResourceV2(t *testing.T) {
	fc := FakeClient{}
	mr := MyResource{Client: &fc}
	ctx := context.Background()
	err := mr.DeleteResourceV2(ctx, "someID")
	assert.NoError(t, err)
}

func TestMyResource_UpdateResourceV2(t *testing.T) {
	fc := FakeClient{}
	mr := MyResource{Client: &fc}
	ctx := context.Background()
	r := sdm.ElasticacheRedis{Name: "manand-test"}
	err := mr.UpdateResourceV2(ctx, &r)
	assert.NoError(t, err)
}
