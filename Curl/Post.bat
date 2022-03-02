
:start
:: Post Add book Test
	curl -Lvso /dev/null -d  "@Wines.json" -X POST http://localhost:8585/books
	curl GET http://localhost:8585/books/1
:: Post Add Wine Test
curl -Lvso /dev/null -d  "@Wines.json" -X POST http://localhost:8585/wines
goto start