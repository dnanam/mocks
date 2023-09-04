package redis

import (
	"context"
	"fmt"

	sdm "github.com/strongdm/strongdm-sdk-go/v3"
)

// // CreateResourceV1 : first attempt for creating a resource
// /*
// Questions when I started to think about how to test this
// * how do I pass in a fake client
// * what will happen when CreateResource is called and eventually client.Resources().Create(ctx, resource) is executed?
// */
// func CreateResourceV1(ctx context.Context, client *sdm.Client, resource sdm.Resource) error {
// 	createResp, err := client.Resources().Create(ctx, resource)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("created successfully", createResp.Resource.GetName())
// 	return nil
// }
//
// // type MyResource struct {
// // 	Client sdm.Client
// // }

type Resourcer interface {
	//	Resources() *sdm.Resources
	Create(ctx context.Context, resource sdm.Resource) (*sdm.ResourceCreateResponse, error)
	Delete(ctx context.Context, resourceID string) (*sdm.ResourceDeleteResponse, error)
	Get(ctx context.Context, resourceID string) (*sdm.ResourceGetResponse, error)
	Update(ctx context.Context, resource sdm.Resource) (*sdm.ResourceUpdateResponse, error)
}

// elasticRedis implements the
type elasticRedis struct {
	Client sdm.Client
}

func (f *elasticRedis) Get(ctx context.Context, resourceID string) (*sdm.ResourceGetResponse, error) {
	fmt.Println("Create")
	return f.Client.Resources().Get(ctx, resourceID)
}

// Create this function is more like a wrapper on top of the sdm library Create function
func (f *elasticRedis) Create(ctx context.Context, resource sdm.Resource) (*sdm.ResourceCreateResponse, error) {
	fmt.Println("Create")
	return f.Client.Resources().Create(ctx, resource)
}

func (f *elasticRedis) Update(ctx context.Context, resource sdm.Resource) (*sdm.ResourceUpdateResponse, error) {
	return f.Client.Resources().Update(ctx, resource)
}

func (f *elasticRedis) Delete(ctx context.Context, resourceID string) (*sdm.ResourceDeleteResponse, error) {
	return f.Client.Resources().Delete(ctx, resourceID)
}

// MyResource implements the Resourcer interface and eventually provides access to functions defined in the interface
type MyResource struct {
	Client Resourcer
}

func (m *MyResource) GetResourceV2(ctx context.Context, resouceID string) error {
	resp, err := m.Client.Get(ctx, resouceID)
	if err != nil {
		return err
	}
	fmt.Println(resp.Resource.GetTags())
	return nil
}

func (m *MyResource) CreateResourceV2(ctx context.Context, resource sdm.Resource) error {
	// sdm.Client.Resources()
	createResp, err := m.Client.Create(ctx, resource)
	if err != nil {
		return err
	}
	fmt.Println("created successfully", createResp.Resource.GetName())
	return nil
}

func (m *MyResource) DeleteResourceV2(ctx context.Context, resourceID string) error {
	resp, err := m.Client.Delete(ctx, resourceID)
	if err != nil {
		return err
	}
	fmt.Println("resource deleted", resp.Meta)
	return nil
}

func (m *MyResource) UpdateResourceV2(ctx context.Context, resource sdm.Resource) error {
	resp, err := m.Client.Update(ctx, resource)
	if err != nil {
		return err
	}
	fmt.Println("resource updated", resp.Resource.GetName())
	return nil
}
