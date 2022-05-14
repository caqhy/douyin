-- 放置 sql 语句的文件

-- 视频表（DB.AutoMigrate(&Video{}) 生成）
create table videos(
    id             bigint auto_increment primary key,
    created_at     datetime(3) null,
    updated_at     datetime(3) null,
    deleted_at     datetime(3) null,
    user_id        bigint      null,
    play_url       longtext    null,
    cover_url      longtext    null,
    tag            longtext    null,
    favorite_count bigint      null,
    comment_count  bigint      null
);