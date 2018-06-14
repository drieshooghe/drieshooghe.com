// -------------- DEPENDENCIES --------------
var gulp = require('gulp');
var ts = require("gulp-typescript");
var merge = require('event-stream').merge;
var concat = require("gulp-concat");
var uglify = require("gulp-uglify");
var sourcemaps = require('gulp-sourcemaps');
var postcss = require('gulp-postcss');
var tailwindcss = require('tailwindcss');
var atImport = require('postcss-import');

var jsVendors = [
    'node_modules/cookieconsent/build/cookieconsent.min.js'
]
var cssVendors = [
    'node_modules/cookieconsent/build/cookieconsent.min.css'
]
// -------------- TASKS --------------

// Compile css
gulp.task('styles', function () {
  
    var css =  gulp.src('resources/styles/main.css')
      .pipe(postcss([
        atImport(),
        tailwindcss('./resources/scripts/tailwind.js'),
        require('autoprefixer'),
      ]));
      
    var vendors = gulp.src(cssVendors);

      return merge(vendors, css)
        .pipe(concat('main.css'))
        .pipe(gulp.dest('static/'));
  });

// Compile javascript
gulp.task("scripts", function () {
    
    var vendors = gulp.src(jsVendors);

    var custom = gulp.src('resources/scripts/*.ts')
        .pipe(ts({
            noImplicitAny: true,
            out: 'custom.js',
            lib: ["es2015", "es2016", "es2017", "dom"]
        }));

    return merge(vendors, custom)
        .pipe(sourcemaps.init())
        .pipe(concat('main.js'))
        .pipe(uglify())
        .pipe(sourcemaps.write('/'))
        .pipe(gulp.dest('static/'));
});