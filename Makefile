test:
	DATABASE_URL=mysql://root:password@localhost:3306                     go test
	DATABASE_URL=mysql://root:password@localhost:3306/                    go test
	DATABASE_URL=mysql://root:password@localhost:3306/test_database_name  go test
