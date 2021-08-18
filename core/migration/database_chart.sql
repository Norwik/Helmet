-- https://dbdiagram.io/

Table option {
  id int [pk, increment] // auto-increment
  key varchar
  value mediumtext
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table endpoint {
  id int [pk, increment] // auto-increment
  status varchar
  listen_path varchar
  token varchar
  name varchar
  upstreams mediumtext
  balancing varchar
  authorization varchar
  authentication varchar
  rate_limit varchar
  circuit_breaker varchar
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table auth_method {
  id int [pk, increment] // auto-increment
  name varchar
  description varchar
  type varchar
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table endpoint_auth_method {
  id int [pk, increment]
  auth_method_id integer
  endpoint_id integer
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table key_based_auth_data {
  id int [pk, increment]
  name varchar
  api_key varchar
  meta varchar
  auth_method_id integer
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table basic_auth_data {
  id int [pk, increment]
  name varchar
  username varchar
  password varchar
  meta varchar
  auth_method_id integer
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table oauth_data {
  id int [pk, increment]
  name varchar
  client_id varchar
  client_secret varchar
  meta varchar
  auth_method_id integer
  created_at datetime [default: `now()`]
  updated_at datetime
}

Table oauth_access_data {
  id int [pk, increment]
  access_token varchar
  meta varchar
  expire_at datetime
  oauth_data_id integer
  created_at datetime [default: `now()`]
  updated_at datetime
}

Ref: auth_method.id > endpoint_auth_method.auth_method_id  [delete: cascade, update: cascade]
Ref: endpoint.id > endpoint_auth_method.endpoint_id [delete: cascade, update: cascade]
Ref: auth_method.id > key_based_auth_data.auth_method_id [delete: cascade, update: cascade]
Ref: auth_method.id > basic_auth_data.auth_method_id [delete: cascade, update: cascade]
Ref: auth_method.id > oauth_data.auth_method_id [delete: cascade, update: cascade]
Ref: oauth_data.id > oauth_access_data.oauth_data_id [delete: cascade, update: cascade]
