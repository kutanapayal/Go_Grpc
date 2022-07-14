
#with absolute path
protoc -Igreet/proto --go_out=. --go_opt=module=go_Grpc --go-grpc_opt=module=go_Grpc --go-grpc_out=. greet/proto/dummy.proto

#updated cmd with relative path 
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  greet/proto/dummy.proto

#Using Makefile       
make greet

all                            Generate Pbs and build
greet                          Generate Pbs and build for greet
calculator                     Generate Pbs and build for calculator
blog                           Generate Pbs and build for blog
test                           Launch tests
clean                          Clean generated files
clean_greet                    Clean generated files for greet
clean_calculator               Clean generated files for calculator
clean_blog                     Clean generated files for blog
rebuild                        Rebuild the whole project
about                          Display info related to the build
help                           Show this help
