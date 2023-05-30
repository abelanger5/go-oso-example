package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"

	"go-oso-example/prisma/db"

	"github.com/osohq/go-oso"
)

//go:embed main.polar
var MainPolicy string

func main() {
	oso, err := oso.NewOso()

	if err != nil {
		panic(err)
	}

	oso.RegisterClassWithName(reflect.TypeOf(map[string]interface{}{}), nil, "PlatformModel")
	oso.RegisterClassWithName(reflect.TypeOf(map[string]interface{}{}), nil, "UserModel")
	oso.RegisterClassWithName(reflect.TypeOf(map[string]interface{}{}), nil, "TenantModel")
	// oso.RegisterClass(reflect.TypeOf(db.TenantModel{}), nil)

	if err := oso.LoadString(MainPolicy); err != nil {
		fmt.Printf("Failed to start: %s", err)
		panic(err)
	}

	// create some test data
	user1 := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:   "123",
			Role: db.RoleCLIENT,
		},
		RelationsUser: db.RelationsUser{
			Memberships: []db.TenantMemberModel{
				{
					InnerTenantMember: db.InnerTenantMember{
						TenantID:   "123",
						TenantRole: db.TenantRoleADMIN,
					},
				},
			},
		},
	}

	user2 := &db.UserModel{
		InnerUser: db.InnerUser{
			ID:   "123",
			Role: db.RoleCLIENT,
		},
		RelationsUser: db.RelationsUser{
			Memberships: []db.TenantMemberModel{
				{
					InnerTenantMember: db.InnerTenantMember{
						TenantID:   "456",
						TenantRole: db.TenantRoleADMIN,
					},
				},
			},
		},
	}

	tenant := &db.TenantModel{
		InnerTenant: db.InnerTenant{
			ID: "123",
		},
	}

	user1Transformed, _ := transform(user1)
	user2Transformed, _ := transform(user2)
	tenantTransformed, _ := transform(tenant)

	err = oso.Authorize(user1Transformed, "update", tenantTransformed)

	fmt.Println("Is user 1 authorized to update the tenant: ", err == nil)

	err = oso.Authorize(user2Transformed, "update", tenantTransformed)

	fmt.Println("Is user 2 authorized to update the tenant: ", err == nil)
}

func transform(v interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	jsonBytes, err := json.Marshal(v)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
