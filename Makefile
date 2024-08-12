tailwind:
	tailwindcss build -i tailwind.css -o assets/styles.css --minify
tailwind_watch:
	tailwindcss build -i tailwind.css -o assets/styles.css --watch
build:
	make tailwind && go build -o dist/calendar
