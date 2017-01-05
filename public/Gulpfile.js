var gulp = require("gulp")
var sass = require("gulp-sass")
var gulpUtil = require('gulp-util');
var uglify = require('gulp-uglify')
var inlinejs = require("gulp-inline-js")

gulp.task("sass", function(){
    gulp.src("css/**/*.scss")
        .pipe(sass())
        .pipe(gulp.dest("css"));
})

gulp.task("js", function(){
    gulp.src("js/**/*.js")
        .pipe(inlinejs())
        .pipe(uglify().on("error",gulpUtil.log))
        .pipe(gulp.dest("js/"))
})

gulp.task("watchjs", function(){
    var res = ['js'];
    gulp.watch("js/**/*.js",res)
})

gulp.task("watchcss", function(){
    var res = ['sass'];
    gulp.watch("css/**/*.scss", res)
})

gulp.task("default",function(){
    gulp.start("watchjs","watchcss","sass","js");
})