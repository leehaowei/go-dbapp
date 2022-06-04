cat <<EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("First line of the program")
}
EOF
mkdir src

mkdir src/config
touch src/config/app.go

mkdir src/controllers
touch src/controllers/book.go

mkdir src/models
touch src/models/book.go

mkdir src/routes
touch src/routes/routes.go

mkdir src/utils
touch src/utils/utils.go