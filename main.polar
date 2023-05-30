allow(actor, action, resource) if
  has_permission(actor, action, resource);

has_role(user: UserModel, role: String, tenant: TenantModel) if
  membership in user.memberships and
  membership.tenantId = tenant.id and
  membership.tenantRole = role;

has_relation(_parent: PlatformModel, "parent", _child: TenantModel) if
    true;

actor UserModel {}

resource PlatformModel {
  roles = [ "ADMIN", "CLIENT"];
  permissions = [ "get", "update", "delete", "list", "create" ];

  "get" if "CLIENT";
  "list" if "CLIENT";
  "create" if "ADMIN";
    "update" if "ADMIN";
    "delete" if "ADMIN";

  "CLIENT" if "ADMIN";
}

resource TenantModel {
  roles = [ "OWNER", "ADMIN", "MEMBER" ];
  permissions = [ "get", "update", "delete", "list", "create" ];
  relations = { parent: PlatformModel };

  "get" if "MEMBER";
  "update" if "ADMIN";
  "delete" if "OWNER";

  "MEMBER" if "ADMIN";
  "ADMIN" if "OWNER";

  "OWNER" if "ADMIN" on "parent";
}