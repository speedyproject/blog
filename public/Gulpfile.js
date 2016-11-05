var gulp = require('gulp'),
    less = require('gulp-less'),
    minifyCSS = require('gulp-minify-css'),
    gulpUtil = require('gulp-util'),
    uglify = require('gulp-uglify');


gulp.task('less', function () {
    gulp.src('css/**/*.less')
        .pipe(less())
        .pipe(gulp.dest('build/css/'));
});

//// 压缩 css
//gulp.task('css',function(){
//    gulp.src('public/stylesheets/user/share/index.css')
//        .pipe(minifyCSS())
//        .pipe(gulp.dest('public/stylesheets/user/share/'))
//});

gulp.task("js", function(){
    gulp.src("js/**/*.js")
        .pipe(uglify().on("error",gulpUtil.log))
    .pipe(gulp.dest("build/js/"));

})

gulp.task("w", function () {
   gulp.watch("css/**/*.less", ['less']);
});
