/*
   package httpapi implements the HTTP API for CRUD operations on the database
   and membership management.

   The key-value operations are available through:
      - /kvstore/create?key=abc&value=123 to create an entry
      - /kvstore/read?key=abc to read an entry
      - /kvstore/update?key=abc&value=456 to update an entry
      - /kvstore/delete?key=abc to delete an entry

   Membership management is available through:
      - /nodes/add?node=123.45.67.89 to add a node
      - /nodes/remove?node=123.45.67.89 to remove a node

   By default, all of these can be used via HTTP GET requests, so for a program
   running locally, any web browser can be used to send requests, e.g:
   http://localhost:8080/kvstore/read?key=42

*/

package httpapi
