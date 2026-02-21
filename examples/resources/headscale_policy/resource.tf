# Basic ACL policy
resource "headscale_policy" "basic" {
  policy = jsonencode({
    "acls" : [{
      "action" : "accept"
    }]
  })
}

# Complex policy with groups and tag owners
resource "headscale_policy" "complex" {
  policy = jsonencode({
    "groups" : {
      "group:admin" : ["user1@", "user2@"],
      "group:dev" : ["user3@", "user4@"]
    },
    "tagOwners" : {
      "tag:prod" : ["group:admin"],
      "tag:dev" : ["group:dev"]
    },
    "acls" : [
      {
        "action" : "accept",
        "src" : ["group:admin"],
        "dst" : ["*:*"]
      },
      {
        "action" : "accept",
        "src" : ["group:dev"],
        "dst" : ["tag:dev:*"]
      }
    ],
    "ssh" : [
      {
        "action" : "accept",
        "src" : ["group:admin"],
        "dst" : ["autogroup:member"],
        "users" : ["root", "ubuntu"]
      }
    ]
  })
}
