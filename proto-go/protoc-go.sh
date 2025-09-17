protoc --go_out=. \
       --go_opt=paths=source_relative \
       --go-triple_out=. \
       --go-triple_opt=paths=source_relative \
  $1/$2
