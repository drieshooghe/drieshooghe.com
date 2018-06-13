// -------------- DEPENDENCIES --------------
var gulp = require('gulp');
var ts = require("gulp-typescript");

// -------------- TASKS --------------

// Compile css
gulp.task('css', function () {
    var postcss = require('gulp-postcss');
    var tailwindcss = require('tailwindcss');
  
    return gulp.src('resources/styles/main.css')
      .pipe(postcss([
        tailwindcss('./resources/styles/tailwind.js'),
        require('autoprefixer'),
      ]))
      .pipe(gulp.dest('static/'));
  });

// Compile javascript
gulp.task("ts", function () {
    var tsResult = gulp.src("resources/scripts/*.ts")
        .pipe(ts({
              noImplicitAny: true,
              out: "main.js"
        }));
    return tsResult.js.pipe(gulp.dest("static/"));
});