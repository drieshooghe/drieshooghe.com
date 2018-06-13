// -------------- DEPENDENCIES --------------
var gulp = require('gulp');
var vendorPaths = [
    'node_modules/cookieconsent/build/cookieconsent.min.js'
]
// -------------- TASKS --------------

// Compile css
gulp.task('styles', function () {
    var postcss = require('gulp-postcss');
    var tailwindcss = require('tailwindcss');
    var atImport = require('postcss-import');
  
    return gulp.src('resources/styles/main.css')
      .pipe(postcss([
        atImport(),
        tailwindcss('./resources/scripts/tailwind.js'),
        require('autoprefixer'),
      ]))
      .pipe(gulp.dest('static/'));
  });

// Compile javascript
gulp.task("scripts", function () {
    var ts = require("gulp-typescript");
    var merge = require('event-stream').merge;
    var concat = require("gulp-concat");
    var uglify = require("gulp-uglify");
    var sourcemaps = require('gulp-sourcemaps');

    var vendors = gulp.src(vendorPaths);

    var custom = gulp.src('resources/scripts/*.ts')
        .pipe(ts({
            noImplicitAny: true,
            out: 'custom.js'
        }));

    return merge(vendors, custom)
        .pipe(sourcemaps.init())
        .pipe(concat('main.js'))
        .pipe(uglify())
        .pipe(sourcemaps.write('/'))
        .pipe(gulp.dest('static/'));
});